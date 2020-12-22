package jsoniter

import (
	"github.com/json-iterator/go"
)

func Fuzz(data []byte) int {
	var stuff []string
	jsoniter.Unmarshal(data, &stuff)
	jsoniter.Marshal(&stuff)
	return 0
}
