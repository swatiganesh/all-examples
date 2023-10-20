package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TokenizerConfig struct {
	TokenChars []string `json:"token_chars"`
	MinGram    string   `json:"min_gram"`
	Type       string   `json:"type"`
	MaxGram    string   `json:"max_gram"`
}

type AnalyzerConfig struct {
	Filter    []string        `json:"filter"`
	Tokenizer TokenizerConfig `json:"tokenizer"`
}

type NormalizerConfig struct {
	CaseInsensitive struct {
		Filter     []string      `json:"filter"`
		Type       string        `json:"type"`
		CharFilter []interface{} `json:"char_filter"`
	} `json:"case_insensitive"`
}

type AnalysisConfig struct {
	Normalizer NormalizerConfig `json:"normalizer"`
	Analyzer   AnalyzerConfig   `json:"analyzer"`
	Tokenizer  TokenizerConfig  `json:"tokenizer"`
}

type IndexConfig struct {
	RefreshInterval  string         `json:"refresh_interval"`
	NumberOfShards   string         `json:"number_of_shards"`
	ProvidedName     string         `json:"provided_name"`
	CreationDate     string         `json:"creation_date"`
	Analysis         AnalysisConfig `json:"analysis"`
	NumberOfReplicas string         `json:"number_of_replicas"`
	UUID             string         `json:"uuid"`
	Version          struct {
		Created string `json:"created"`
	} `json:"version"`
}

type SettingsConfig struct {
	Index IndexConfig `json:"index"`
}

type Assignment1 struct {
	Comp7S20211122 struct {
		Settings SettingsConfig `json:"settings"`
	} `json:"comp-7-s-2021.11.22"`
	Comp7S20211123 struct {
		Settings SettingsConfig `json:"settings"`
	} `json:"comp-7-s-2021.11.23"`
}

func main() {

	//reading the json data from data.json file
	// content, err := ioutil.ReadFile("data.json")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	content, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	defer content.Close()

	//unmarshel the data structure into go struct
	var data Assignment1
	decoder := json.NewDecoder(content)
	err2 := decoder.Decode(&data)
	if err2 != nil {
		fmt.Println("error json unmarshelling")
		fmt.Println(err2.Error())
	}

	//for _, x := range data {

	// fmt.Println("Comp7S20211122 Refresh Interval:", data.Comp7S20211122.Settings.Index.RefreshInterval)

	// fmt.Println("Comp7S20211123 ProvidedName:", data.Comp7S20211123.Settings.Index.ProvidedName)

	// fmt.Println("Comp7S20211122 Analyzer Tokenizer:", data.Comp7S20211122.Settings.Index.Analysis.Tokenizer.TokenChars)
	// fmt.Println("Comp7S20211122 UUID:", data.Comp7S20211122.Settings.Index.UUID)

	fmt.Println("Comp7S20211123 Analyzer Tokenizer:", data.Comp7S20211123.Settings.Index.Analysis.Tokenizer)
	// fmt.Println("Comp7S20211123 UUID:", data.Comp7S20211123.Settings.Index.UUID)

	fmt.Println("Comp7S20211122 Analyzer auto complete version filter:", data.Comp7S20211122.Settings.Index.Analysis.Analyzer.Filter)

	fmt.Println("Normalizer Type:", data.Comp7S20211122.Settings.Index.Analysis.Normalizer.CaseInsensitive.Type)
	fmt.Println("Analyzer Tokenizer Type:", data.Comp7S20211122.Settings.Index.Analysis.Analyzer.Tokenizer.Type)

	// Accessing values from Comp7S20211123
	fmt.Println("Normalizer Type (Comp7S20211123):", data.Comp7S20211123.Settings.Index.Analysis.Normalizer.CaseInsensitive.Type)
	fmt.Println("Analyzer Tokenizer Type (Comp7S20211123):", data.Comp7S20211123.Settings.Index.Analysis.Analyzer.Tokenizer.Type)

}
