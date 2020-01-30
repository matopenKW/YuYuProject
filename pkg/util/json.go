package util

import (
	"encoding/json"
)

func GetList(path string, list []map[string]interface{}) ([]map[string]interface{}, error) {
	bytes, err := ReadFile(path + ".json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
