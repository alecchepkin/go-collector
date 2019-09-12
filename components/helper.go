package components

import "encoding/json"

func ToJson(i interface{}) (s string, err error) {
	d, err := json.Marshal(i)
	s = string(d)
	return
}
