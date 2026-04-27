package strutil

import (
	"encoding/json"
	"strings"
)

func ContainsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func ToJSON(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
