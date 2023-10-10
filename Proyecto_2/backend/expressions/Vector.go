package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Vector struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Method     string
}

func NewVector(line int, col int, ide string, expression interfaces.Expression, method string) Vector {
	return Vector{Line: line, Col: col, Id: ide, Expression: expression, Method: method}
}

func (v Vector) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
