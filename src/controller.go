package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unicode"
)

type retTuple struct {
	respCode int
	retStr   string
}

// Convert cognate list into a map
func getData(language string) map[string]string {
	path := "../data/" + language + ".json"
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteVal, _ := io.ReadAll(jsonFile)

	var cognates map[string]string
	json.Unmarshal(byteVal, &cognates)
	return cognates

}

func wordProcessor(st string) []string {
	words := []string{}
	placeholder := ""
	for _, element := range st {
		if element == ' ' {
			if placeholder != "" {
				words = append(words, placeholder)
				placeholder = ""
			}
		} else if unicode.IsPunct(element) {
			words = append(words, placeholder, string(element))
			placeholder = ""
		} else {
			placeholder += string(element)
		}

	}
	words = append(words, placeholder)
	return words

}

func handleTranslation(st sourceText) retTuple {
	// var cognates map[string]string
	retString := ""
	stString := st.Source[:2] + "-" + st.Target[:2]
	cognatesAddr, ok := languageMap[stString]

	if !ok {
		return retTuple{400, ""}
	}
	cognates := *cognatesAddr

	words := wordProcessor(st.PostContent)
	for _, word := range words {
		if val, ok := cognates[word]; ok {
			retString += val + " "
		} else {
			retString += word + " "
		}
	}
	return retTuple{200, retString[:len(retString)-1]}
}
