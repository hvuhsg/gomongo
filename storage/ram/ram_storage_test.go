package ram_test

import (
	"testing"

	storage "github.com/hvuhsg/gomongo/storage/ram"
)

func TestCreateDatabase(t *testing.T) {
	ramStorage := storage.New()
	err := ramStorage.CreateDatabase("some_name")
	if err != nil {
		t.Fatalf("database not created")
	}

	err = ramStorage.CreateDatabase("some_name")
	if err == nil {
		t.Errorf("created database twice")
	}
}

func TestCreateCollection(t *testing.T) {
	ramStorage := storage.New()
	err := ramStorage.CreateCollection("db_name", "col_name")
	if err == nil {
		t.Errorf("created collection inside database that was not created")
	}

	ramStorage.CreateDatabase("db_name")

	err = ramStorage.CreateCollection("db_name", "collection")
	if err != nil {
		t.Fatalf("collection not created")
	}

	err = ramStorage.CreateCollection("db_name", "collection")
	if err == nil {
		t.Errorf("collection was created twice")
	}
}

func TestInsertDocuments(t *testing.T) {
	ramStorage := storage.New()
	ramStorage.CreateDatabase("db")
	ramStorage.CreateCollection("db", "col")
	var documents = [2]map[string]interface{}{
		{"a": 1},
		{"b": 2},
	}

	err := ramStorage.Insert("x", "col", documents[:])
	if err == nil {
		t.Errorf("can't insert documents into non existing db")
	}

	err = ramStorage.Insert("db", "x", documents[:])
	if err == nil {
		t.Errorf("can't insert documents into non existing collection")
	}

	err = ramStorage.Insert("db", "col", documents[:])
	if err != nil {
		t.Errorf("documents not inserted to collection")
	}
}

func TestDeleteDocuments(t *testing.T) {
	ramStorage := storage.New()
	ramStorage.CreateDatabase("db")
	ramStorage.CreateCollection("db", "col")

	var documents = [2]map[string]interface{}{
		{"a": 1},
		{"b": 2},
	}
	ramStorage.Insert("db", "col", documents[:])
}
