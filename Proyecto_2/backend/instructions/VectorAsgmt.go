package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type VectorAsgmt struct {
	Line        int
	Col         int
	Id          string
	Position    interfaces.Expression
	Expression  interfaces.Expression
	Id2ndVector string
}

func NewVectorAsgmt(line int, col int, ide string, position interfaces.Expression, expression interfaces.Expression, idvector string) VectorAsgmt {
	return VectorAsgmt{Line: line, Col: col, Id: ide, Position: position, Expression: expression, Id2ndVector: idvector}
}

func (v VectorAsgmt) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
