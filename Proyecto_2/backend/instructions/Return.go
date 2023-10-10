package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Return struct {
	Line       int
	Col        int
	Expression interfaces.Expression
}

func NewReturn(line int, col int, exp interfaces.Expression) Return {
	return Return{line, col, exp}
}

func (r Return) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
