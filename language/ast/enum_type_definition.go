package ast

import (
	"github.com/chris-ramon/graphql-go/language/kinds"
)

// EnumTypeDefinition implements Node, TypeDefinition
type EnumTypeDefinition struct {
	Kind   string
	Loc    Location
	Name   *Name
	Values []interface{}
}

func NewEnumTypeDefinition(def *EnumTypeDefinition) *EnumTypeDefinition {
	if def == nil {
		def = &EnumTypeDefinition{}
	}
	return &EnumTypeDefinition{
		Kind:   kinds.EnumTypeDefinition,
		Loc:    def.Loc,
		Name:   def.Name,
		Values: def.Values,
	}
}

func (def *EnumTypeDefinition) GetKind() string {
	return def.Kind
}

func (def *EnumTypeDefinition) GetLoc() Location {
	return def.Loc
}

func (def *EnumTypeDefinition) GetName() *Name {
	return def.Name
}

func (def *EnumTypeDefinition) GetVariableDefinitions() []*VariableDefinition {
	return []*VariableDefinition{}
}

func (def *EnumTypeDefinition) GetSelectionSet() SelectionSet {
	return SelectionSet{}
}

func (def *EnumTypeDefinition) GetOperation() string {
	return ""
}
