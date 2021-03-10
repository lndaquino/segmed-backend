package utils

import "encoding/json"

// ParseStruct handles structs format convertions
func ParseStruct(origin interface{}, destiny interface{}) (err error) {
	temp, err := json.Marshal(origin)
	if err != nil {
		return
	}
	if err = json.Unmarshal(temp, &destiny); err != nil {
		return
	}
	return
}
