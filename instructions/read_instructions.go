package instructions

import (
	"github.com/fatih/set"
)

type readInstructions struct {
	readAll bool
	include set.Interface
	exclude set.Interface
}

func New(readAll bool) IReadInstructions {
	inculdeKeys := set.New(set.NonThreadSafe)
	excludeKeys := set.New(set.NonThreadSafe)
	r := new(readInstructions)
	r.include = inculdeKeys
	r.exclude = excludeKeys
	r.readAll = readAll
	return *r
}

func (r readInstructions) And(other *IReadInstructions) IReadInstructions {
	resultReadInstructions := new(readInstructions)
	resultReadInstructions.readAll = r.readAll && (*other).ReadAll()
	resultReadInstructions.include = set.Intersection(r.include, (*other).GetInculdeKeys())
	resultReadInstructions.exclude = set.Difference(r.exclude, (*other).GetExcludeKeys())
	return *resultReadInstructions
}

func (r readInstructions) Or(other *IReadInstructions) IReadInstructions {
	resultReadInstructions := new(readInstructions)
	resultReadInstructions.readAll = r.readAll || (*other).ReadAll()
	resultReadInstructions.include = set.Difference(r.include, (*other).GetInculdeKeys())
	resultReadInstructions.exclude = set.Intersection(r.exclude, (*other).GetExcludeKeys())
	return *resultReadInstructions
}

func (r readInstructions) Not() IReadInstructions {
	resultReadInstructions := new(readInstructions)
	resultReadInstructions.readAll = !r.readAll
	resultReadInstructions.exclude = r.include
	resultReadInstructions.include = r.exclude

	return *resultReadInstructions
}

func (r readInstructions) GetInculdeKeys() set.Interface {
	return r.include
}

func (r readInstructions) GetExcludeKeys() set.Interface {
	return r.exclude
}

func (r readInstructions) ReadAll() bool {
	return r.readAll
}

func (r readInstructions) IsExcluded(lookupKey interface{}) bool {
	return r.exclude.Has(lookupKey)
}

func (r readInstructions) AddInculdeKey(lookupKey interface{}) {
	if r.exclude.Has(lookupKey) {
		panic(lookupKey)
	}
	r.include.Add(lookupKey)
}

func (r readInstructions) AddExcludeKey(lookupKey interface{}) {
	if r.include.Has(lookupKey) {
		panic(lookupKey)
	}
	r.exclude.Add(lookupKey)
}
