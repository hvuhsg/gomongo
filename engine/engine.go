package engine

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

type IValidator interface {
	ValidateFilter(filter map[string]interface{}) error
	ValidateMutation(mutation map[string]interface{}) error
	ValidateName(name string) error
	ValidateDocument(document map[string]interface{}) error
}

type IIndexor interface {
	CreateIndex(database_name string, collection_name string, index map[string]interface{}) (string, error)
	DropIndex(index_id string) error
	QueryIndex(database_name string, collection_name string, filter map[string]interface{}) (IReadInstructions, error)
	GetDatabaseIndexes(database_name string) map[string]map[string]interface{}
	GetCollectionIndexes(database_name string, collection_name string) map[string]interface{}
}

type IStorage interface {
	CreateDatabase(database_name string) error
	DropDatabase(database_name string) error
	CreateCollection(database_name string, collection_name string) error
	DropCollection(database_name string, collection_name string) error
	Insert(database_name string, collection_name string, documents []map[string]interface{}) error
	Delete(database_name string, collection_name string, read_instructions IReadInstructions) error
	Find(database_name string, collection_name string, read_instructions IReadInstructions) ([]map[string]interface{}, error)
}

type IReadInstructions interface {
	And(*IReadInstructions) IReadInstructions
	Or(*IReadInstructions) IReadInstructions
	Not() IReadInstructions
	GetLookupKeys() []interface{}
	ReadAll() bool
}

type Engine struct {
	validator *IValidator
	indexor   *IIndexor
	storage   *IStorage
}

func NewEngine(validator IValidator, indexor IIndexor, storage IStorage) IEngine {
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
	return nil
}

func (engine Engine) Delete(database_name string, collection_name string, filter map[string]interface{}) error {
	return nil
}

func (engine Engine) Update(database_name string, collection_name string, filter map[string]interface{}, mutation map[string]interface{}) error {
	return nil
}

func (engine Engine) Replace(database_name string, collection_name string, filter map[string]interface{}, replacement map[string]interface{}) error {
	return nil
}

func (engine Engine) Find(database_name string, collection_name string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	return nil, nil
}
