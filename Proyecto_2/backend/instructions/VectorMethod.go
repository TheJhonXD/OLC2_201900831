package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type VectorMethod struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Method     string
}

func NewVectorMethod(line int, col int, ide string, expression interfaces.Expression, method string) VectorMethod {
	return VectorMethod{Line: line, Col: col, Id: ide, Expression: expression, Method: method}
}

func (v VectorMethod) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
