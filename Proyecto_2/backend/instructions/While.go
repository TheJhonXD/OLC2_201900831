package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type While struct {
	Line         int
	Col          int
	Expression   interfaces.Expression
	Instructions []interface{}
}

func NewWhile(line int, col int, expression interfaces.Expression, instructions []interface{}) While {
	return While{line, col, expression, instructions}
}

func (w While) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
