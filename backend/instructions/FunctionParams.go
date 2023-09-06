package instructions

import "Server/environment"

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

func (p Param) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	if p.ExternalId == "" {
		p.ExternalId = p.InternalId
	}
	if p.Type == environment.NULL {
		ast.SetError("Error en la definición de parámetros")
	}

	return environment.SymbolFuncParam{Line: p.Line, Col: p.Col, ExternalId: p.ExternalId, InternalId: p.InternalId, Reference: p.Reference, Type: p.Type}
}
