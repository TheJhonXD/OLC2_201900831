package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type If struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Block      []interface{}
	ElseIf     []interface{}
	ElseInstr  []interface{}
}

func NewIf(line int, col int, expression interfaces.Expression, block []interface{}, elseif []interface{}, elseinstr []interface{}) If {
	return If{line, col, expression, block, elseif, elseinstr}
}

func (i If) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
