package helper

import (
	"fmt"
	"testing"
)

func TestStringHandler(t *testing.T) {
	tt := []struct {
		value string
	}{
		{
			value: "		共有 281 页",
		},
	}
	for _, t := range tt {
		fmt.Println(StringHandler(t.value))
	}
}

func TestRegexHandler(t *testing.T) {
	tt := []struct {
		value string
	}{
		{
			value: "index_198.htm",
		},
	}
	for _, t := range tt {
		fmt.Println(RegexHandler(t.value))
	}
}
