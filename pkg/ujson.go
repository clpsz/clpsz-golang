package pkg

import (
	"encoding/json"
)

func MustGet(jsonStr string, keys []interface{}) interface{} {
	var curObject map[string]interface{}
	var curArray []interface{}

	var next interface{}
	var ok bool

	keyLen := len(keys)
	for i, key := range keys {
		switch key.(type) {
		case string:
			if i == 0 {
				json.Unmarshal([]byte(jsonStr), &curObject)
			} else {
				curObject, ok = next.(map[string]interface{})
				if !ok {
					panic("convert error")
				}
			}

			next = curObject[key.(string)]
			break
		case int:
			if i == 0 {
				json.Unmarshal([]byte(jsonStr), &curArray)
			} else {
				curArray, ok = next.([]interface{})
				if !ok {
					panic("convert error")
				}
			}
			next = curArray[key.(int)]

			break
		}
		if i == keyLen - 1 {
			return next
		}
	}

	panic("should never reach here")
}
