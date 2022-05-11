package indexing

import "github.com/hvuhsg/gomongo/engine"

type readAllInstructions struct{}

func (instructions readAllInstructions) And(other *engine.IReadInstructions) engine.IReadInstructions {
	return instructions
}

func (instructions readAllInstructions) Or(other *engine.IReadInstructions) engine.IReadInstructions {
	return instructions
}

func (instructions readAllInstructions) Not() engine.IReadInstructions {
	return instructions
}

func (instructions readAllInstructions) GetLookupKeys() []interface{} {
	return nil
}

func (instructions readAllInstructions) ReadAll() bool {
	return true
}
