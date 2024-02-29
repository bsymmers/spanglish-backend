package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Convert cognate list into a map
func getData() map[string]string {
	jsonFile, err := os.Open("../data/span.json")

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

func handleTranslation(st sourceText) string {
	var cognates map[string]string
	retString := ""
	if st.Source == "Spanish" && st.Target == "English" {
		cognates = sp_cognates
	} else {
		cognates = nil
	}

	words := wordProcessor(st.PostContent)
	for _, word := range words {
		if val, ok := cognates[word]; ok {
			retString += val + " "
		} else {
			retString += word + " "
		}
	}
	return retString[:len(retString)-1]
}
