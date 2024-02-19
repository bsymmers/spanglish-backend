package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	var cognateDict = getData()
}

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
