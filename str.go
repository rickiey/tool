package tool

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
)

const (
	empty = ""
	tab   = "\t"
)

// Stringify returns a string representation
func Stringify(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return empty, err
	}
	return string(b), nil
}

// Structify returns the original representation
func Structify(data string, value interface{}) error {
	return json.Unmarshal([]byte(data), value)
}

// PrettyJson returns a pretty json string
func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// IsNumeric 判断锁给字符串是否为数字
func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < byte('0') || s[i] > byte('9') {
			return false
		}
	}
	return true
}

func JsonString(v interface{}) string {
	b, _ := json.Marshal(v)
	return strings.TrimSpace(string(b))
}
