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

type AutocompleteAnalyzerConfig struct {
	Filter    []string `json:"filter"`
	Tokenizer string   `json:"tokenizer"`
}

type AnalyzerConfig struct {
	Autocomplete               AutocompleteAnalyzerConfig `json:"autocomplete"`
	AutocompleteVersionNumbers AutocompleteAnalyzerConfig `json:"autocomplete_version_numbers"`
}

type CaseInsensitiveNormalizerConfig struct {
	Filter     []string   `json:"filter"`
	Type       string     `json:"type"`
	CharFilter []struct{} `json:"char_filter"`
}

type NormalizerConfig struct {
	CaseInsensitive CaseInsensitiveNormalizerConfig `json:"case_insensitive"`
}

type AnalysisConfig struct {
	Normalizer NormalizerConfig `json:"normalizer"`
	Analyzer   AnalyzerConfig   `json:"analyzer"`
	Tokenizer  TokenizerConfig  `json:"tokenizer"`
}

type VersionConfig struct {
	Created string `json:"created"`
}

type IndexConfig struct {
	RefreshInterval  string         `json:"refresh_interval"`
	NumberOfShards   string         `json:"number_of_shards"`
	ProvidedName     string         `json:"provided_name"`
	CreationDate     string         `json:"creation_date"`
	Analysis         AnalysisConfig `json:"analysis"`
	NumberOfReplicas string         `json:"number_of_replicas"`
	UUID             string         `json:"uuid"`
	Version          VersionConfig  `json:"version"`
}

type SettingsConfig struct {
	Index IndexConfig `json:"index"`
}

type example struct {
	Settings SettingsConfig `json:"settings"`
}

type Assignment1 struct {
	Comp7S20211122 example `json:"comp-7-s-2021.11.22"`
	Comp7S20211123 example `json:"comp-7-s-2021.11.23"`
}

func main() {
	content, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	defer content.Close()

	var data Assignment1
	decoder := json.NewDecoder(content)
	err2 := decoder.Decode(&data)
	if err2 != nil {
		fmt.Println("error json unmarshelling")
		fmt.Println(err2.Error())
	}
	// Access and print values (you can add more print statements as needed)
	fmt.Println("Comp7S20211122:")
	fmt.Println("  Refresh Interval:", data.Comp7S20211122.Settings.Index.RefreshInterval)
	fmt.Println("  Number of Shards:", data.Comp7S20211122.Settings.Index.NumberOfShards)
	fmt.Println("  Case Insensitive Filter:", data.Comp7S20211122.Settings.Index.Analysis.Normalizer.CaseInsensitive.Filter[0])

	fmt.Println("Comp7S20211123:")
	fmt.Println("  Refresh Interval:", data.Comp7S20211123.Settings.Index.RefreshInterval)
	fmt.Println("  Number of Shards:", data.Comp7S20211123.Settings.Index.NumberOfShards)
	fmt.Println("  Case Insensitive Filter:", data.Comp7S20211123.Settings.Index.Analysis.Normalizer.CaseInsensitive.Filter[0])
}
