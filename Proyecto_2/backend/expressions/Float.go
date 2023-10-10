package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Float struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewFloat(line int, col int, value interfaces.Expression) Float {
	return Float{Line: line, Col: col, Value: value}
}

func (f Float) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
