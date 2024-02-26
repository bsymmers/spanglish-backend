package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Convert cognate list into a map
func getData() map[string]string {
	jsonFile, err := os.Open("data/span.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteVal, _ := io.ReadAll(jsonFile)

	var cognates map[string]string
	json.Unmarshal(byteVal, &cognates)
	fmt.Println(cognates["athletic"])
	return cognates

}

// func translateCognates(cognates map[string]string, input string) {
// 	output := input
// 	for i := 0; i < len(input); i++ {
// 		elem, ok := cognates[input[i]]

// 		if ok {

// 		}

// 	}
// }
