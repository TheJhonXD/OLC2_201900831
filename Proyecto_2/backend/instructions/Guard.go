package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Guard struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Block      []interface{}
}

func NewGuard(line int, col int, expression interfaces.Expression, block []interface{}) Guard {
	return Guard{line, col, expression, block}
}

func (g Guard) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
