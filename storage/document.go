package storage

type storageDocument struct {
	data      map[string]any
	lookupKey any
}

func NewDocument(data map[string]any, lookupKey any) IStorageDocument {
	storageDoc := new(storageDocument)
	storageDoc.data = data
	storageDoc.lookupKey = lookupKey
	return storageDoc
}

func (doc storageDocument) GetData() map[string]any {
	return doc.data
}

func (doc storageDocument) GetLookupKey() any {
	return doc.lookupKey
}
