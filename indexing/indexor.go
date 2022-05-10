package indexing

import "github.com/hvuhsg/gomongo/engine"

type readAllInstructions struct{}

func (instructions readAllInstructions) And(other *engine.IReadInstructions) engine.IReadInstructions {
	return instructions
}

func (instructions readAllInstructions) Or(other *engine.IReadInstructions) engine.IReadInstructions {
	return instructions
}

func (instructions readAllInstructions) Not() engine.IReadInstructions {
	return instructions
}

func (instructions readAllInstructions) GetLookupKeys() []interface{} {
	return nil
}

func (instructions readAllInstructions) ReadAll() bool {
	return true
}

type fakeIndexor struct{}

func NewFakeIndexor() engine.IIndexor {
	indexor := new(fakeIndexor)
	return indexor
}

func (indexor fakeIndexor) CreateIndex(database_name string, collection_name string, index map[string]interface{}) (string, error) {
	return "index_id", nil
}

func (indexor fakeIndexor) DropIndex(index_id string) error {
	return nil
}

func (indexor fakeIndexor) GetDatabaseIndexes(database_name string) map[string]map[string]interface{} {
	return nil
}

func (indexor fakeIndexor) GetCollectionIndexes(database_name string, collection_name string) map[string]interface{} {
	return nil
}

func (indexor fakeIndexor) QueryIndex(database_name string, collection_name string, filter map[string]interface{}) (engine.IReadInstructions, error) {
	read_instructions := readAllInstructions{}
	return read_instructions, nil
}
