//go:build !jsoniter
// +build !jsoniter

package json

import (
	std "encoding/json"
)

// ToJSON transform struct to json
func ToJSON(v interface{}) string {
	b, err := std.Marshal(v)
	if err != nil {
		return ""
	}

	return string(b)
}

func ToJSONb(v interface{}) []byte {
	b, err := std.Marshal(v)
	if err != nil {
		return nil
	}
	return b
}

// ToJSONf transform struct to json and text format
func ToJSONf(v interface{}) string {
	b, err := std.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(b)
}

// ToJSONs ...strcut to jsonstr
func ToJSONs(v interface{}) string {
	b, err := std.MarshalIndent(v, "", "")
	if err != nil {
		return ""
	}
	return string(b)
}

// ToStruct ...json string to struct
func ToStruct(s string, v interface{}) error {
	return std.Unmarshal([]byte(s), v)
}

var (
	Unmarshal     = ToStruct
	Marshal       = ToJSONb
	MarshalString = ToJSON
)
