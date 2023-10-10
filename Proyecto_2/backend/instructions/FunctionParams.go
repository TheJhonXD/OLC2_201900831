package instructions

import (
	"Server/environment"
	"Server/generator"
)

type Param struct {
	Line       int
	Col        int
	ExternalId string
	InternalId string
	Reference  bool
	Type       environment.TipoExpresion
}

func NewParam(line int, col int, externalId string, internalId string, reference bool, t environment.TipoExpresion) Param {
	return Param{Line: line, Col: col, ExternalId: externalId, InternalId: internalId, Reference: reference, Type: t}
}

func (p Param) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
