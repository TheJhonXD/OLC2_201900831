package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type CallParams struct {
	Line       int
	Col        int
	Id         string
	Amp        bool
	Expression interfaces.Expression
}

func NewCallParams(line int, col int, ide string, amp bool, exp interfaces.Expression) CallParams {
	return CallParams{Line: line, Col: col, Id: ide, Amp: amp, Expression: exp}
}

func (c CallParams) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
