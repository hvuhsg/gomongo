package main

import (
	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/indexing"
	"github.com/hvuhsg/gomongo/storage"
	"github.com/hvuhsg/gomongo/validation"
)

func main() {
	storage := storage.NewRamStorage()
	validator := validation.NewFakeValidator()
	indexor := indexing.NewFakeIndexor()
	db_engine := engine.NewEngine(validator, indexor, storage)

	db_engine.CreateDatabase("db_name")
	db_engine.CreateCollection("db_name", "col_name")
}
