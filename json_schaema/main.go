package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
)

func main() {
	schemaContent, err := ioutil.ReadFile("schema.json")
	if err != nil {
		panic(err.Error())
	}
	jsonContent, err := ioutil.ReadFile("document.json")
	if err != nil {
		panic(err.Error())
	}

	loader1 := gojsonschema.NewStringLoader(string(schemaContent))
	schema, err := gojsonschema.NewSchema(loader1)
	if err != nil {
		panic(err.Error())
	}

	documentLoader := gojsonschema.NewStringLoader(string(jsonContent))
	result, err := schema.Validate(documentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}