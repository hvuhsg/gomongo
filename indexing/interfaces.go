package indexing

import "github.com/hvuhsg/gomongo/instructions"

type IIndexor interface {
	CreateIndex(database_name string, collection_name string, index map[string]interface{}) (string, error)
	DropIndex(index_id string) error
	QueryIndex(database_name string, collection_name string, filter map[string]interface{}) (instructions.IReadInstructions, error)
	GetDatabaseIndexes(database_name string) map[string]map[string]interface{}
	GetCollectionIndexes(database_name string, collection_name string) map[string]interface{}
}
