package storage

import "github.com/hvuhsg/gomongo/instructions"

type IStorageDocument interface {
	GetData() map[string]any
	GetLookupKey() any
}

type IStorage interface {
	CreateDatabase(database_name string) error
	DropDatabase(database_name string) error
	CreateCollection(database_name string, collection_name string) error
	DropCollection(database_name string, collection_name string) error
	Insert(database_name string, collection_name string, documents []map[string]interface{}) error
	Delete(database_name string, collection_name string, read_instructions instructions.IReadInstructions) error
	Find(database_name string, collection_name string, read_instructions instructions.IReadInstructions) ([]IStorageDocument, error)
}
