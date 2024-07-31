package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/joho/godotenv"
)

type retTuple struct {
	respCode int
	retStr   string
}

type DeeplBody struct {
	Text        []string `json:"text"`
	Source_lang string   `json:"source_lang"`
	Target_lang string   `json:"target_lang"`
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

// Process given input string into words so that we can identify cognates
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

// Check if a given word is in the cognates list and if it is add it to the return string
func handleTranslation(st sourceText) retTuple {
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

func deeplHandler(st sourceText) retTuple {
	err := godotenv.Load()
	if err == nil {
		return retTuple{400, ""}
	}

	apiKey := "DeepL-Auth-Key " + os.Getenv("DEEPL_API_KEY")

	pc := []string{st.PostContent}

	deeplbody := DeeplBody{
		Text:        pc,
		Source_lang: deeplMap[st.Source],
		Target_lang: deeplMap[st.Target],
	}

	deeplJson, err := json.Marshal(deeplbody)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
		return retTuple{400, ""}
	}

	request, err := http.NewRequest("POST", "https://api-free.deepl.com/v2/translate", bytes.NewBuffer(deeplJson))
	if err != nil {
		log.Fatalf("impossible to build request: %s", err)
		return retTuple{400, ""}
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Content-Length", strconv.Itoa(int(request.ContentLength)))
	request.Header.Set("Authorization", apiKey)
	request.Header.Set("User-Agent", "Cognate-Translator/1.2.3")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
		return retTuple{400, ""}
	}
	log.Printf("status Code: %d", res.StatusCode)
	// read response body
	body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println(error)
		return retTuple{400, ""}
	}
	// close response body
	res.Body.Close()
	// print response body
	log.Printf("res body: %s", string(body))

	return retTuple{res.StatusCode, string(body)}

}
