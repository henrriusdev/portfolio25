package utils

import (
	"encoding/json"
)

// parseTemplJsonString converts any value to a JSON string that can be safely used in templ templates
// and with Alpine.js data initialization. It handles the JSON marshaling and escaping.
func ParseTemplJsonString(v interface{}) string {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		// Return empty object if marshaling fails
		return "{}"
	}
	// Convert to string and return
	return string(jsonBytes)
}
