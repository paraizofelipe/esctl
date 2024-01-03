package output

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/TylerBrock/colorjson"
)

func PrintPrettyJSON(jsonBytes []byte) {
	var obj map[string]interface{}
	err := json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		log.Fatal(err)
	}

	f := colorjson.NewFormatter()
	f.Indent = 2

	coloredJson, err := f.Marshal(obj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(coloredJson))
}
