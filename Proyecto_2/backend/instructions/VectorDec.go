package instructions

import (
	"Server/environment"
	"Server/generator"
)

type VectorDec struct {
	Line    int
	Col     int
	Id      string
	Type    environment.TipoExpresion
	ExpList []interface{}
}

func NewVectorStmt(line int, col int, id string, tipo environment.TipoExpresion, expList []interface{}) VectorDec {
	return VectorDec{line, col, id, tipo, expList}
}

func (v VectorDec) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
