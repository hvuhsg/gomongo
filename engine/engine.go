package engine

import (
	"github.com/hvuhsg/gomongo/engine/filtering"
	"github.com/hvuhsg/gomongo/engine/validation"
	"github.com/hvuhsg/gomongo/indexing"
	"github.com/hvuhsg/gomongo/instructions"
	"github.com/hvuhsg/gomongo/storage"
)

type IEngine interface {
	CreateDatabase(database_name string) error
	DropDatabase(database_name string) error
	CreateCollection(database_name string, collection_name string) error
	DropCollection(database_name string, collection_name string) error
	Insert(database_name string, collection_name string, documents []map[string]interface{}) error
	Update(database_name string, collection_name string, filter map[string]interface{}, mutation map[string]interface{}) error
	Replace(database_name string, collection_name string, filter map[string]interface{}, replacement map[string]interface{}) error
	Delete(database_name string, collection_name string, filter map[string]interface{}) error
	Find(database_name string, collection_name string, filter map[string]interface{}) ([]map[string]interface{}, error)
}

type Engine struct {
	validator *validation.IValidator
	indexor   *indexing.IIndexor
	storage   *storage.IStorage
}

func New(validator validation.IValidator, indexor indexing.IIndexor, storage storage.IStorage) IEngine {
	engine := new(Engine)
	engine.validator = &validator
	engine.indexor = &indexor
	engine.storage = &storage
	return engine
}

func (engine Engine) CreateDatabase(database_name string) error {
	err := (*engine.validator).ValidateName(database_name)
	if err != nil {
		return err
	}

	err = (*engine.storage).CreateDatabase(database_name)
	if err != nil {
		return err
	}

	return nil
}

func (engine Engine) DropDatabase(database_name string) error {
	err := (*engine.storage).DropDatabase(database_name)
	if err != nil {
		return err
	}

	collections := (*engine.indexor).GetDatabaseIndexes(database_name)

	for _, indexes := range collections {
		for index_id := range indexes {
			err := (*engine.indexor).DropIndex(index_id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (engine Engine) CreateCollection(database_name string, collection_name string) error {
	err := (*engine.validator).ValidateName(database_name)
	if err != nil {
		return err
	}

	err = (*engine.validator).ValidateName(collection_name)
	if err != nil {
		return err
	}

	err = (*engine.storage).CreateCollection(database_name, collection_name)
	if err != nil {
		return err
	}

	return nil
}

func (engine Engine) DropCollection(database_name string, collection_name string) error {
	err := (*engine.storage).DropCollection(database_name, collection_name)
	if err != nil {
		return err
	}

	indexes := (*engine.indexor).GetCollectionIndexes(database_name, collection_name)

	for index_id := range indexes {
		err := (*engine.indexor).DropIndex(index_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (engine Engine) Insert(database_name string, collection_name string, documents []map[string]interface{}) error {
	err := (*engine.storage).Insert(database_name, collection_name, documents)
	if err != nil {
		return err
	}

	return nil
}

func (engine Engine) Delete(database_name string, collection_name string, filter map[string]interface{}) error {
	readInstructions, err := (*engine.indexor).QueryIndex(database_name, collection_name, filter)
	if err != nil {
		return err
	}

	storageDocuments, err := (*engine.storage).Find(database_name, collection_name, readInstructions)
	if err != nil {
		return err
	}

	deleteInstruactions := instructions.New(false)
	for _, storageDocument := range storageDocuments {
		deleteInstruactions.AddInculdeKey(storageDocument.GetLookupKey())
	}

	(*engine.storage).Delete(database_name, collection_name, deleteInstruactions)

	return nil
}

func (engine Engine) Update(database_name string, collection_name string, filter map[string]interface{}, mutation map[string]interface{}) error {
	return nil
}

func (engine Engine) Replace(database_name string, collection_name string, filter map[string]interface{}, replacement map[string]interface{}) error {
	return nil
}

func (engine Engine) Find(database_name string, collection_name string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	readInstructions, err := (*engine.indexor).QueryIndex(database_name, collection_name, filter)
	if err != nil {
		return nil, err
	}

	storageDocuments, err := (*engine.storage).Find(database_name, collection_name, readInstructions)
	if err != nil {
		return nil, err
	}

	documents := make([]map[string]any, 0, len(storageDocuments))
	for _, storageDocument := range storageDocuments {
		documents = append(documents, storageDocument.GetData())
	}

	var filteredDocuments []map[string]any
	if readInstructions.ReadAll() {
		filteredDocuments = filtering.FilterList(filter, documents)
	} else {
		filteredDocuments = documents
	}

	return filteredDocuments, err
}
