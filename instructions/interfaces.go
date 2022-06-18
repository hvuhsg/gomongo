package instructions

import "github.com/fatih/set"

type IReadInstructions interface {
	And(*IReadInstructions) IReadInstructions
	Or(*IReadInstructions) IReadInstructions
	Not() IReadInstructions
	GetInculdeKeys() set.Interface
	GetExcludeKeys() set.Interface
	IsExcluded(lookupKey any) bool
	ReadAll() bool
	AddInculdeKey(lookupKey any)
	AddExcludeKey(lookupKey any)
}
