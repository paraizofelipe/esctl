package file

import (
	"encoding/json"
	"os"
)

func ReadJSONFile(filePath string) (jsonStr string, err error) {

	var content []byte

	content, err = os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	jsonStr = string(content)

	return
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func IsContentValid(jsonStr string) bool {
	var jsonMap map[string]interface{}
	return json.Unmarshal([]byte(jsonStr), &jsonMap) == nil
}
