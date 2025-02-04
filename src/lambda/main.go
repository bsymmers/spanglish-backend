package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type sourceText struct {
	Source      string `json:"Source"`
	Target      string `json:"Target"`
	PostContent string `json:"postContent"`
	Deepl       bool   `json:"Deepl"`
}

type ResponseObj struct {
	CognateResponse string
	DeeplResponse   string
}

var ensp_cognates map[string]string
var enfe_cognates map[string]string
var iten_cognates map[string]string
var spen_cognates map[string]string
var fren_cognates map[string]string
var enit_cognates map[string]string

var languageMap = map[string]*map[string]string{
	"En-Sp": &ensp_cognates,
	"En-Fr": &enfe_cognates,
	"It-En": &iten_cognates,
	"Sp-En": &spen_cognates,
	"Fr-En": &fren_cognates,
	"En-It": &enit_cognates,
}

var deeplMap = map[string]string{
	"English": "EN",
	"Spanish": "ES",
	"French":  "FR",
	"Italian": "IT",
}

func handleRequest(ctx context.Context, event events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// Parse the input event
	body := []byte(event.Body)
	var newSourceText sourceText
	log.Print(event)

	if err := json.Unmarshal(body, &newSourceText); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return events.LambdaFunctionURLResponse{StatusCode: 400}, err
	}
	event_log, _ := json.Marshal(newSourceText)
	log.Print(string(event_log))

	response := handleTranslation(newSourceText)

	var responseObj ResponseObj
	var respCode int
	responseObj.CognateResponse = response.retStr

	if (newSourceText.Deepl) && (response.respCode == 200) {
		deeplResponse := deeplHandler(newSourceText)
		responseObj.DeeplResponse = deeplResponse.retStr

		if deeplResponse.respCode == 400 {
			respCode = 400
		} else {
			respCode = 200
		}

	} else {
		responseObj.DeeplResponse = ""
		respCode = 200
	}

	respString, err := json.Marshal(responseObj)
	if err != nil {
		respCode = 400
		respString = nil
	}
	log.Print(string(respString))

	return events.LambdaFunctionURLResponse{StatusCode: respCode, Body: string(respString)}, nil
}

func main() {
	ensp_cognates = getData("EN-SP")
	enfe_cognates = getData("EN-FR")
	iten_cognates = getData("IT-EN")
	spen_cognates = getData("SP-EN")
	fren_cognates = getData("FR-EN")
	enit_cognates = getData("EN-IT")
	lambda.Start(handleRequest)
}
