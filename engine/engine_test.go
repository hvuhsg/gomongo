package engine_test

import (
	"testing"

	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/engine/validation"
	"github.com/hvuhsg/gomongo/indexing"
	storage "github.com/hvuhsg/gomongo/storage/ram"
)

func createEngine() engine.IEngine {
	validator := validation.New()
	indexor := indexing.New()
	storage := storage.New()
	return engine.New(validator, indexor, storage)
}

func TestCreateDatabase(t *testing.T) {
	e := createEngine()

	err := e.CreateDatabase("##invalid")
	if err == nil {
		t.Errorf("'##invalid' is invalid name for database")
	}

	err = e.CreateDatabase("valid_name")
	if err != nil {
		t.Errorf("'valid_name' is a valid name for a database")
	}
}

func TestCraeteCollection(t *testing.T) {
	e := createEngine()

	err := e.CreateCollection("non exiting database", "col_name")
	if err == nil {
		t.Errorf("cant create collection on non-existing database")
	}

	e.CreateDatabase("db")

	err = e.CreateCollection("db", "##invalid")
	if err == nil {
		t.Errorf("'##invalid' is invalid name for collection")
	}

	err = e.CreateCollection("db", "valid_name")
	if err != nil {
		t.Errorf("'valid_name' is a valid name for a collection")
	}
}
