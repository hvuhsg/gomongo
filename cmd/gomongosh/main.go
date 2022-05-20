package main

import (
	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/indexing"
	ram_storage "github.com/hvuhsg/gomongo/storage/ram"
	"github.com/hvuhsg/gomongo/validation"
)

func main() {
	storage := ram_storage.New()
	validator := validation.New()
	indexor := indexing.New()
	db_engine := engine.New(validator, indexor, storage)

	document := map[string]interface{}{
		"hello":  "string",
		"number": 5,
		"float":  5.4,
		"bool":   true,
		"__arr":  []int{5, 8, -9},
	}

	validator.ValidateDocument(document)

	db_engine.CreateDatabase("db_name")
	db_engine.CreateCollection("db_name", "col_name")
}
