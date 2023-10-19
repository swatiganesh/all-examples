package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// type Assignment1 struct {
// 	Comp7S20211122 struct {
// 		Settings struct {
// 			Index struct {
// 				RefreshInterval string `json:"refresh_interval"`
// 				NumberOfShards  string `json:"number_of_shards"`
// 				ProvidedName    string `json:"provided_name"`
// 				CreationDate    string `json:"creation_date"`
// 				Analysis        struct {
// 					Normalizer struct {
// 						CaseInsensitive struct {
// 							Filter     []string      `json:"filter"`
// 							Type       string        `json:"type"`
// 							CharFilter []interface{} `json:"char_filter"`
// 						} `json:"case_insensitive"`
// 					} `json:"normalizer"`
// 					Analyzer struct {
// 						Autocomplete struct {
// 							Filter    []string `json:"filter"`
// 							Tokenizer string   `json:"tokenizer"`
// 						} `json:"autocomplete"`
// 						AutocompleteVersionNumbers struct {
// 							Filter    []string `json:"filter"`
// 							Tokenizer string   `json:"tokenizer"`
// 						} `json:"autocomplete_version_numbers"`
// 					} `json:"analyzer"`
// 					Tokenizer struct {
// 						AutocompleteVersionNumberTokenizer struct {
// 							TokenChars []string `json:"token_chars"`
// 							MinGram    string   `json:"min_gram"`
// 							Type       string   `json:"type"`
// 							MaxGram    string   `json:"max_gram"`
// 						} `json:"autocomplete_version_number_tokenizer"`
// 						AutocompleteTokenizer struct {
// 							TokenChars []string `json:"token_chars"`
// 							MinGram    string   `json:"min_gram"`
// 							Type       string   `json:"type"`
// 							MaxGram    string   `json:"max_gram"`
// 						} `json:"autocomplete_tokenizer"`
// 					} `json:"tokenizer"`
// 				} `json:"analysis"`
// 				NumberOfReplicas string `json:"number_of_replicas"`
// 				UUID             string `json:"uuid"`
// 				Version          struct {
// 					Created string `json:"created"`
// 				} `json:"version"`
// 			} `json:"index"`
// 		} `json:"settings"`
// 	} `json:"comp-7-s-2021.11.22"`
// 	Comp7S20211123 struct {
// 		Settings struct {
// 			Index struct {
// 				RefreshInterval string `json:"refresh_interval"`
// 				NumberOfShards  string `json:"number_of_shards"`
// 				ProvidedName    string `json:"provided_name"`
// 				CreationDate    string `json:"creation_date"`
// 				Analysis        struct {
// 					Normalizer struct {
// 						CaseInsensitive struct {
// 							Filter     []string      `json:"filter"`
// 							Type       string        `json:"type"`
// 							CharFilter []interface{} `json:"char_filter"`
// 						} `json:"case_insensitive"`
// 					} `json:"normalizer"`
// 					Analyzer struct {
// 						Autocomplete struct {
// 							Filter    []string `json:"filter"`
// 							Tokenizer string   `json:"tokenizer"`
// 						} `json:"autocomplete"`
// 						AutocompleteVersionNumbers struct {
// 							Filter    []string `json:"filter"`
// 							Tokenizer string   `json:"tokenizer"`
// 						} `json:"autocomplete_version_numbers"`
// 					} `json:"analyzer"`
// 					Tokenizer struct {
// 						AutocompleteVersionNumberTokenizer struct {
// 							TokenChars []string `json:"token_chars"`
// 							MinGram    string   `json:"min_gram"`
// 							Type       string   `json:"type"`
// 							MaxGram    string   `json:"max_gram"`
// 						} `json:"autocomplete_version_number_tokenizer"`
// 						AutocompleteTokenizer struct {
// 							TokenChars []string `json:"token_chars"`
// 							MinGram    string   `json:"min_gram"`
// 							Type       string   `json:"type"`
// 							MaxGram    string   `json:"max_gram"`
// 						} `json:"autocomplete_tokenizer"`
// 					} `json:"tokenizer"`
// 				} `json:"analysis"`
// 				NumberOfReplicas string `json:"number_of_replicas"`
// 				UUID             string `json:"uuid"`
// 				Version          struct {
// 					Created string `json:"created"`
// 				} `json:"version"`
// 			} `json:"index"`
// 		} `json:"settings"`
// 	} `json:"comp-7-s-2021.11.23"`
// }

type Analyzer struct {
	Autocomplete struct {
		Filter    []string `json:"filter"`
		Tokenizer string   `json:"tokenizer"`
	} `json:"autocomplete"`
	AutocompleteVersionNumbers struct {
		Filter    []string `json:"filter"`
		Tokenizer string   `json:"tokenizer"`
	} `json:"autocomplete_version_numbers"`
}

type Tokenizer struct {
	AutocompleteVersionNumberTokenizer struct {
		TokenChars []string `json:"token_chars"`
		MinGram    string   `json:"min_gram"`
		Type       string   `json:"type"`
		MaxGram    string   `json:"max_gram"`
	} `json:"autocomplete_version_number_tokenizer"`
	AutocompleteTokenizer struct {
		TokenChars []string `json:"token_chars"`
		MinGram    string   `json:"min_gram"`
		Type       string   `json:"type"`
		MaxGram    string   `json:"max_gram"`
	} `json:"autocomplete_tokenizer"`
}

type Index struct {
	RefreshInterval string `json:"refresh_interval"`
	NumberOfShards  string `json:"number_of_shards"`
	ProvidedName    string `json:"provided_name"`
	CreationDate    string `json:"creation_date"`
	Analysis        struct {
		Normalizer struct {
			CaseInsensitive struct {
				Filter     []string      `json:"filter"`
				Type       string        `json:"type"`
				CharFilter []interface{} `json:"char_filter"`
			} `json:"case_insensitive"`
		} `json:"normalizer"`
		Analyzer  Analyzer  `json:"analyzer"`
		Tokenizer Tokenizer `json:"tokenizer"`
	} `json:"analysis"`
	NumberOfReplicas string `json:"number_of_replicas"`
	UUID             string `json:"uuid"`
	Version          struct {
		Created string `json:"created"`
	} `json:"version"`
}

type Settings struct {
	Index Index `json:"index"`
}

type Comp7S20211122 struct {
	Settings Settings `json:"settings"`
}

type Comp7S20211123 struct {
	Settings Settings `json:"settings"`
}

type Assignment1 struct {
	Comp7S20211122 Comp7S20211122 `json:"comp-7-s-2021.11.22"`
	Comp7S20211123 Comp7S20211123 `json:"comp-7-s-2021.11.23"`
}

func main() {

	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var data Assignment1
	err2 := json.Unmarshal(content, &data)
	if err2 != nil {
		fmt.Println("error json unmarshelling")
		fmt.Println(err2.Error())
	}

	//for _, x := range data {

	fmt.Println("Comp7S20211122 Refresh Interval:", data.Comp7S20211122.Settings.Index.RefreshInterval)

	fmt.Println("Comp7S20211123 ProvidedName:", data.Comp7S20211123.Settings.Index.ProvidedName)

	//fmt.Println("Comp7S20211123 Analyzer Tokenizer:", data.Comp7S20211123.Settings.Index.Analysis.Tokenizer.AutocompleteTokenizer.TokenChars)
	//fmt.Println("Comp7S20211123 UUID:", data.Comp7S20211123.Settings.Index.UUID)

	fmt.Println("Comp7S20211123 Analyzer Tokenizer:", data.Comp7S20211123.Settings.Index.Analysis.Tokenizer)
	fmt.Println("Comp7S20211123 UUID:", data.Comp7S20211123.Settings.Index.UUID)

}
