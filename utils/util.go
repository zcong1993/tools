package utils

import (
	"encoding/json"
	"os"
	"strconv"
)

// GetEnvOrDefault is helper func for getting env or default
func GetEnvOrDefault(key, d string) string {
	v := os.Getenv(key)
	if v == "" {
		v = d
	}
	return v
}

// Str2int64 is helper func for converting string to int64
func Str2int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// OutPut is helper func for writing pretty json data to std
func Output(val interface{}, err error) {
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Stdout.Write([]byte("\n"))
		return
	}
	b, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Stdout.Write([]byte("\n"))
		return
	}
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n"))
}
