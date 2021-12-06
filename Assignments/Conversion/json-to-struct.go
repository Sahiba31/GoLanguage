package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Version struct {
	Created string `json:"created"`
}

type AutoCompleteTokenizer struct {
	TokenChars []string `json:"token_chars"`
	MinGram    string   `json:"min_gram"`
	Type       string   `json:"type"`
	MaxGram    string   `json:"max_gram"`
}

type Tokenizer struct {
	TokenizerAutoCompleteVersionNo AutoCompleteTokenizer `json:"autocomplete_version_number_tokenizer"`
	TokenizerAutoComplete          AutoCompleteTokenizer `json:"autocomplete_tokenizer"`
}

type AutoComplete struct {
	Filter    []string `json:"filter"`
	Tokenizer string   `json:"tokenizer"`
}

type Analyzer struct {
	AnalyzerAutocomplete          AutoComplete `json:"autocomplete"`
	AnalyzerAutocompleteVersionNo AutoComplete `json:"autocomplete_version_numbers"`
}

type CaseInsensitive struct {
	Filter     []string `json:"filter"`
	Type       string   `json:"type"`
	CharFilter []string `json:"char_filter"`
}

type Normalizer struct {
	NormCaseInsensitive CaseInsensitive `json:"case_insensitive"`
}

type Analysis struct {
	AnalysisNormalizer Normalizer `json:"normalizer"`
	AnalysisAnalyzer   Analyzer   `json:"analyzer"`
	AnalysisTokenizer  Tokenizer  `json:"tokenizer"`
}

type IndexData struct {
	RefreshInterval string   `json:"refresh_interval"`
	NoOfShards      string   `json:"number_of_shards"`
	ProvidedName    string   `json:"provided_name"`
	CreationDate    string   `json:"creation_date"`
	IndexAnalysis   Analysis `json:"analysis"`
	NoOfReplicas    string   `json:"number_of_replicas"`
	UUID            string   `json:"uuid"`
	IndexVersion    Version  `json:"version"`
}

type Index struct {
	IndexNum IndexData `json:"index"`
}
type Setting struct {
	CompSetting Index `json:"settings"`
}
type Computer struct {
	CompName1 Setting `json:"comp-7-s-2021.11.22"`
	CompName2 Setting `json:"comp-7-s-2021.11.23"`
}

func main() {

	jsonFile, err := os.Open("string.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened string.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var comp Computer

	json.Unmarshal(byteValue, &comp)
	//fmt.Println(comp)
	fmt.Println(comp.CompName1.CompSetting.IndexNum.ProvidedName)
	fmt.Println(comp.CompName2.CompSetting.IndexNum.ProvidedName)
}
