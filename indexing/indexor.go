package indexing

import (
	"github.com/hvuhsg/gomongo/engine"
	"github.com/hvuhsg/gomongo/instructions"
)

type Indexor struct{}

func New() engine.IIndexor {
	indexor := new(Indexor)
	return indexor
}

func (indexor Indexor) CreateIndex(database_name string, collection_name string, index map[string]interface{}) (string, error) {
	return "index_id", nil
}

func (indexor Indexor) DropIndex(index_id string) error {
	return nil
}

func (indexor Indexor) GetDatabaseIndexes(database_name string) map[string]map[string]interface{} {
	return nil
}

func (indexor Indexor) GetCollectionIndexes(database_name string, collection_name string) map[string]interface{} {
	return nil
}

func (indexor Indexor) QueryIndex(database_name string, collection_name string, filter map[string]interface{}) (engine.IReadInstructions, error) {
	read_instructions := instructions.New(true)
	return read_instructions, nil
}
