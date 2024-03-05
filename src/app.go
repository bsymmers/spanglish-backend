package main

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

func main() {
	ensp_cognates = getData("En-Sp")
	enfe_cognates = getData("En-Fr")
	iten_cognates = getData("It-En")
	spen_cognates = getData("Sp-En")
	fren_cognates = getData("Fr-En")
	enit_cognates = getData("En-It")

	routing()
}
