package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Integer struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewInteger(line int, col int, value interfaces.Expression) Integer {
	return Integer{Line: line, Col: col, Value: value}
}

func (i Integer) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
