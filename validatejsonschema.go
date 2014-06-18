package main

import (
	"flag"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

func die(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		die("Must specify schema filename")
	}
	schema := args[0]
	json := ""

	if len(args) >= 2 {
		json = args[1]
	}

	schemaJson, err := gojsonschema.GetFileJson(schema)
	if err != nil {
		die(fmt.Sprintf("Failed to read %s: %s", schema, err.Error()))
	}

	schemaDoc, err := gojsonschema.NewJsonSchemaDocument(schemaJson)
	if err != nil {
		die(fmt.Sprintf("JSON schema %s is invalid: %s", schema, err.Error()))
	}

	if json != "" {
		jsonDoc, err := gojsonschema.GetFileJson(json)
		if err != nil {
			die(fmt.Sprintf("JSON file %s is invalid: %s", json, err.Error()))
		}
		res := schemaDoc.Validate(jsonDoc)
		if res.Valid() {
			fmt.Printf("Valid JSON: %s matches schema %s\n", json, schema)
		} else {
			for _, desc := range res.Errors() {
				fmt.Printf("- %s\n", desc)
			}
		}
	}
}
