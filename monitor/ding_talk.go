package monitor

import (
	"fmt"
	"utils/http_client"
)

type (
	DingTalk struct {
		Url       string `json:"url"`
		atMobiles []string
		isAtAll   bool
	}

	TextMsg struct {
		Msgtype string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
		At struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		} `json:"at"`
	}
)

func NewDingTalk(url string) *DingTalk {
	return &DingTalk{
		Url:       url,
		atMobiles: []string{},
		isAtAll:   false,
	}
}

func (dingTalk *DingTalk) WithAtMobiles(atMobiles []string) *DingTalk {
	if atMobiles != nil {
		dingTalk.atMobiles = atMobiles
	}
	return dingTalk
}

func (dingTalk *DingTalk) WithIsAtAll(isAtAll bool) *DingTalk {
	dingTalk.isAtAll = isAtAll
	return dingTalk
}

func (dingTalk *DingTalk) SendMsg(body interface{}) ([]byte, error) {
	client := http_client.GetHTTPClient()
	var resp interface{}
	data, err := client.Post(dingTalk.Url, nil, body, &resp)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (dingTalk *DingTalk) SendTextMsg(content string) error {
	if dingTalk.Url == "" {
		return fmt.Errorf("报警地址为空")
	}
	msg := TextMsg{
		Msgtype: "text",
	}
	msg.Text.Content = content
	msg.At.IsAtAll = dingTalk.isAtAll
	msg.At.AtMobiles = dingTalk.atMobiles
	_, err := dingTalk.SendMsg(msg)
	return err
}
