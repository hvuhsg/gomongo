package main

import (
	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/indexing"
	"github.com/hvuhsg/gomongo/storage"
	"github.com/hvuhsg/gomongo/validation"
)

func main() {
	storage := storage.NewRamStorage()
	validator := validation.New()
	indexor := indexing.NewFakeIndexor()
	db_engine := engine.NewEngine(validator, indexor, storage)

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
