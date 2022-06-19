package main

import (
	"fmt"

	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/engine/validation"
	"github.com/hvuhsg/gomongo/indexing"
	ram_storage "github.com/hvuhsg/gomongo/storage/ram"
)

func main() {
	storage := ram_storage.New()
	validator := validation.New()
	indexor := indexing.New()
	db_engine := engine.New(validator, indexor, storage)

	// Create database and collection
	db_engine.CreateDatabase("db_name")
	db_engine.CreateCollection("db_name", "col_name")

	// Insert document to collection
	document := map[string]interface{}{
		"hello":  "string",
		"number": 5,
		"float":  5.4,
		"bool":   true,
		"__arr":  []int{5, 8, -9},
	}
	db_engine.Insert("db_name", "col_name", []map[string]any{document})

	// Filter docuemnts from collection
	var filter = map[string]any{"number": 5.0}
	documents, err := db_engine.Find("db_name", "col_name", filter)

	fmt.Println(err)
	fmt.Println(documents)
}
