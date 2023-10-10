package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type String struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewString(line int, col int, value interfaces.Expression) String {
	return String{Line: line, Col: col, Value: value}
}

func (s String) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
