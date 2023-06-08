package gjson

import "encoding/json"

type JsonConvertInter interface {
	Encode(v interface{}) string
	Decode(data []byte, val interface{}) error
}

var JsonConvertUtil = _jsonConvertUtil{}

type _jsonConvertUtil struct {
}

func (_jsonConvertUtil) Encode(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (_jsonConvertUtil) Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}
