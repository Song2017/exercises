package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	strDemo := `{"test":123, "test2":"sadf"}`
	dictDemo := make(map[string]interface{})
	err := json.Unmarshal([]byte(strDemo), &dictDemo)
	if err != nil {
		panic("Youzan Config Error: " + err.Error())
	}
	fmt.Println(dictDemo["test"])
	fmt.Print(dictDemo["test2"], "\n")
}
