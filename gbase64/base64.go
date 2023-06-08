package gbase64

import "encoding/base64"

type Base64ConvertInter interface {
	Encode(s string) string
	Decode(s string) (string, error)
}

var Base64ConvertUtil = _base64ConvertUtil{}

type _base64ConvertUtil struct {
}

func (_base64ConvertUtil) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (_base64ConvertUtil) Decode(s string) (string, error) {
	ds, err := base64.StdEncoding.DecodeString(s)
	return string(ds), err
}
