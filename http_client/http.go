package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	DefaultClient *http.Client

	maxRetries = 3
	baseTime   = 50
	RetryCode  = map[int]struct{}{
		502: {},
		503: {},
		504: {},
	}
)

type Client struct {
	*http.Client
}

func GetHTTPClient() *Client {
	DefaultClient = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          20,
			MaxConnsPerHost:       100,
			MaxIdleConnsPerHost:   3,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   2 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	client := &Client{
		Client: DefaultClient,
	}
	return client
}

// Get http get
func (c *Client) Get(api string, params url.Values, body, response interface{}, headerKV ...string) ([]byte, error) {
	return c.Request(http.MethodGet, api, params, body, response, headerKV...)
}

// Put http put
func (c *Client) Put(api string, body, response interface{}, headerKV ...string) ([]byte, error) {
	return c.Request(http.MethodPut, api, nil, body, response, headerKV...)
}

// Del http delete
func (c *Client) Del(api string, response interface{}, headerKV ...string) ([]byte, error) {
	body, err := c.Request(http.MethodDelete, api, nil, nil, response, headerKV...)
	if err != nil {
		return nil, nil
	}
	return body, err
}

// Post http post
func (c *Client) Post(api string, params url.Values, body, response interface{}, headerKV ...string) ([]byte, error) {
	return c.Request(http.MethodPost, api, params, body, response, headerKV...)
}

// Upload http post file
func (c *Client) Upload(api, fieldname, filename string, params map[string]string, response interface{}) ([]byte, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for key, val := range params {
		if err := bodyWriter.WriteField(key, val); err != nil {
			return nil, err
		}
	}

	fileWriter, err := bodyWriter.CreateFormFile(fieldname, filepath.Base(filename))
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(fileWriter, f)
	if err != nil {
		return nil, err
	}
	defer bodyWriter.Close()

	return c.Request(http.MethodPost, api, nil, bodyBuf, response,
		"Content-Type", bodyWriter.FormDataContentType())
}

// Request http request
func (c *Client) Request(method, api string, params url.Values, bodyParams, response interface{}, headerKV ...string) ([]byte, error) {
	var (
		statusCode int
		err        error
		body       []byte
		request    *http.Request
	)

	for try := 1; try <= maxRetries; try++ {
		request, err = c.newRequest(method, api, params, bodyParams, headerKV...)
		if err != nil {
			return nil, err
		}

		now := time.Now()
		statusCode, body, err = c.do(request, response)
		latency := time.Since(now)
		// An error is returned caused by client policy.
		if err != nil {
			// obfuscate api sensitive information
			// we can use for range to handle this
			if strings.Contains(err.Error(), "connection reset by peer") ||
				strings.Contains(err.Error(), "handshake failure") {
				time.Sleep(time.Millisecond * time.Duration(try) * time.Duration(baseTime))
				continue
			}
			log.Printf("http request failed  {statusCode: %v, try: %v, method: %v params: %v, bodyParams: "+
				"%v latency: %v}", statusCode, try, method, params, bodyParams, latency.String())
			break
		}
		// An error from server.
		if statusCode != http.StatusOK {
			// obfuscate api sensitive information
			// we can use for range to handle this
			_, ok := RetryCode[statusCode]
			if ok {
				// retry
				time.Sleep(time.Millisecond * time.Duration(try) * time.Duration(baseTime))
				if try == maxRetries {
					log.Printf("http request failed  {statusCode: %v, try: %v, method: %v, params: %v, bodyParams: %v, "+
						"latency: %v}", statusCode, try, method, params, bodyParams, latency.String())
				}
				continue
			}

			if statusCode == http.StatusBadRequest {
				// not need to alarm
				log.Printf("http request failed  {statusCode: %v, try: %v, method: %v, params: %v, bodyParams: %v, "+
					"latency: %v}", statusCode, try, method, params, bodyParams, latency.String())
			} else if statusCode == http.StatusNotFound {
				log.Printf("http response code %v failed", statusCode)
				break
			} else {
				log.Printf("http request failed  {statusCode: %v, try: %v, method: %v,params: %v, bodyParams: %v, "+
					"latency: %v}", statusCode, try, method, params, bodyParams, latency.String())
			}
			err = fmt.Errorf("http response code %v failed", statusCode)
			break
		}
		// resp success
		return body, nil
	}
	// resp fail
	return nil, err
}

// NOTE: don't use params & bodyParams at the same time
func (c *Client) newRequest(method, reqURL string, params url.Values, bodyParams interface{},
	headerKV ...string) (*http.Request, error) {
	body, err := c.trimBody(bodyParams)
	if err != nil {
		return nil, err
	}

	pBody, contentType, rawQuery := c.trimParams(method, params)
	if body == nil && pBody != nil {
		body = pBody
	}

	request, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return nil, err
	}

	if request.URL.RawQuery != "" && rawQuery != "" {
		rawQuery = "&" + rawQuery
	}
	request.URL.RawQuery += rawQuery
	request.Header.Set("Content-Type", contentType)
	for i := 0; i < len(headerKV); i += 2 {
		request.Header.Set(headerKV[i], headerKV[i+1])
	}

	return request, nil
}

func (c *Client) do(request *http.Request, response interface{}) (statusCode int, body []byte, err error) {
	statusCode = -1
	resp, err := c.Client.Do(request)
	if resp != nil {
		statusCode = resp.StatusCode
		defer resp.Body.Close()
	}
	if err != nil {
		return statusCode, nil, err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return statusCode, nil, err
	}

	if statusCode != http.StatusOK {
		return statusCode, body, nil
	}

	if response == nil {
		return statusCode, body, nil
	}
	if err := json.Unmarshal(body, response); err != nil {
		return statusCode, nil, err
	}

	return statusCode, body, nil
}

func (c *Client) trimParams(method string, params url.Values) (body io.Reader, contentType, rawQuery string) {
	contentType = "application/json"
	switch method {
	case http.MethodPost, http.MethodPut:
		if len(params) > 0 {
			contentType = "application/x-www-form-urlencoded"
			body = strings.NewReader(params.Encode())
		}
	case http.MethodGet, http.MethodDelete:
		if len(params) > 0 {
			rawQuery = params.Encode()
		}
	}

	return
}

func (c *Client) trimBody(params interface{}) (io.Reader, error) {
	if params == nil {
		return nil, nil
	}

	if body, ok := params.(io.Reader); ok {
		return body, nil
	}

	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}
