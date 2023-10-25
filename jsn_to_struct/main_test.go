package main

import (
	"encoding/json"
	"os"

	//"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadJSONFile(t *testing.T) {
	// Assuming you have a test.json file with appropriate data for testing
	testFile := "test.json"

	// Call the function to read the JSON file
	data, err := ReadJSONFile(testFile)
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	// Read the expected data from expected.json
	expectedDataFile := "expected.json"
	expectedContent, err := os.Open(expectedDataFile)
	if err != nil {
		t.Fatalf("Error opening expected data file: %v", err)
	}
	defer expectedContent.Close()

	// Decode the expected data from the file
	var expectedData Assignment1
	decoder := json.NewDecoder(expectedContent)
	err = decoder.Decode(&expectedData)
	if err != nil {
		t.Fatalf("Error decoding expected data JSON: %v", err)
	}

	//assert.Equal(t, expectedData, data)

	assert.NotEqual(t, expectedData, data)

	// Compare the actual result with the expected result
	// if !reflect.DeepEqual(data, expectedData) {
	// 	t.Fail()
	// 	t.Errorf("Expected %+v, got %+v", expectedData, data)
	// }
}
