package instructions

import (
	"github.com/fatih/set"
	"github.com/hvuhsg/gomongo/engine"
)

type readInstructions struct {
	readAll           bool
	lookupKeys        set.Interface
	excludeLookUpKeys set.Interface
}

func New(readAll bool) engine.IReadInstructions {
	lookupKeys := set.New(set.NonThreadSafe)
	r := new(readInstructions)
	r.lookupKeys = lookupKeys
	r.readAll = readAll
	return *r
}

func (r readInstructions) And(other *engine.IReadInstructions) engine.IReadInstructions {
	resultReadInstructions := new(readInstructions)
	resultReadInstructions.readAll = r.readAll && (*other).ReadAll()
	resultReadInstructions.lookupKeys = set.Intersection(r.lookupKeys, (*other).GetLookupKeys())

	return *resultReadInstructions
}

func (r readInstructions) Or(other *engine.IReadInstructions) engine.IReadInstructions {
	resultReadInstructions := new(readInstructions)
	resultReadInstructions.readAll = r.readAll || (*other).ReadAll()
	resultReadInstructions.lookupKeys = set.Union(r.lookupKeys, (*other).GetLookupKeys())

	return *resultReadInstructions
}

func (r readInstructions) Not() engine.IReadInstructions {
	resultReadInstructions := new(readInstructions)
	resultReadInstructions.readAll = !r.readAll
	resultReadInstructions.excludeLookUpKeys = r.lookupKeys
	resultReadInstructions.lookupKeys = r.excludeLookUpKeys

	return *resultReadInstructions
}

func (r readInstructions) GetLookupKeys() set.Interface {
	return r.lookupKeys
}

func (r readInstructions) ReadAll() bool {
	return r.readAll
}

func (r readInstructions) IsExcluded(lookupKey interface{}) bool {
	return r.excludeLookUpKeys.Has(lookupKey)
}

func (r readInstructions) AddLookupKey(lookupKey interface{}) {
	r.lookupKeys.Add(lookupKey)
}

func (r readInstructions) AddExcludedlookupKey(lookupKey interface{}) {
	r.excludeLookUpKeys.Add(lookupKey)
}
