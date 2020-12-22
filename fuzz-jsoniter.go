package jsoniter

import (
	"github.com/json-iterator/go"
)

func Fuzz(data []byte) int {
	var stuff []byte
	jsoniter.Unmarshal(data, &stuff)
	jsoniter.Marshal(&data)
	return 0
}
