package ram

import (
	"errors"

	"github.com/hvuhsg/gomongo/engine"
)

type ramStorage struct {
	databases map[string]map[string][]map[string]interface{}
}

func New() engine.IStorage {
	databases := make(map[string]map[string][]map[string]interface{}, 100)
	storage := ramStorage{databases: databases}
	return storage
}

func (storage *ramStorage) ensureDatabase(database_name string) error {
	_, ok := storage.databases[database_name]
	if !ok {
		return errors.New("database not found")
	}

	return nil
}

func (storage *ramStorage) ensureCollection(database_name string, collection_name string) error {
	err := storage.ensureDatabase(database_name)
	if err != nil {
		return err
	}

	_, ok := storage.databases[database_name][collection_name]
	if !ok {
		return errors.New("collection not found")
	}

	return nil
}

func (storage ramStorage) CreateDatabase(database_name string) error {
	if _, ok := storage.databases[database_name]; ok {
		return errors.New("database all ready exists")
	}

	database := make(map[string][]map[string]interface{}, 100)
	storage.databases[database_name] = database

	return nil
}

func (storage ramStorage) DropDatabase(database_name string) error {
	err := storage.ensureDatabase(database_name)
	if err != nil {
		return err
	}

	delete(storage.databases, database_name)

	return nil
}

func (storage ramStorage) CreateCollection(database_name string, collection_name string) error {
	err := storage.ensureDatabase(database_name)
	if err != nil {
		return err
	}

	_, ok := storage.databases[database_name][collection_name]
	if ok {
		return errors.New("collection already exists")
	}

	collection := make([]map[string]interface{}, 100)
	storage.databases[database_name][collection_name] = collection

	return nil
}

func (storage ramStorage) DropCollection(database_name string, collection_name string) error {
	err := storage.ensureCollection(database_name, collection_name)
	if err != nil {
		return err
	}

	delete(storage.databases, database_name)

	return nil
}

func (storage ramStorage) Insert(database_name string, collection_name string, documents []map[string]interface{}) error {
	err := storage.ensureCollection(database_name, collection_name)
	if err != nil {
		return err
	}

	collection := storage.databases[database_name][collection_name]
	collection = append(collection, documents...)
	storage.databases[database_name][collection_name] = collection

	return nil
}

func (storage ramStorage) Delete(database_name string, collection_name string, read_instructions engine.IReadInstructions) error {
	err := storage.ensureCollection(database_name, collection_name)
	if err != nil {
		return err
	}

	collection := storage.databases[database_name][collection_name]

	for lookup_key := range read_instructions.GetLookupKeys().List() {
		collection[lookup_key] = nil
	}

	return nil
}

func (storage ramStorage) Find(database_name string, collection_name string, read_instructions engine.IReadInstructions) ([]map[string]interface{}, error) {
	documents := make([]map[string]interface{}, 5000)

	collection := storage.databases[database_name][collection_name]

	for lookup_key := range read_instructions.GetLookupKeys().List() {
		documents = append(documents, collection[lookup_key])
	}

	return documents, nil
}
