package json

import jsoniter "github.com/json-iterator/go"

var (
	Marshal   = jsoniter.ConfigCompatibleWithStandardLibrary.Marshal
	Unmarshal = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal
)
