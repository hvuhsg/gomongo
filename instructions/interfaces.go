package instructions

import "github.com/fatih/set"

type IReadInstructions interface {
	And(*IReadInstructions) IReadInstructions
	Or(*IReadInstructions) IReadInstructions
	Not() IReadInstructions
	GetLookupKeys() set.Interface
	IsExcluded(lookupKey interface{}) bool
	ReadAll() bool
	AddLookupKey(lookupKey interface{})
	AddExcludedlookupKey(lookupKey interface{})
}
