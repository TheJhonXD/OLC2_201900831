package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
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

func (c CallParams) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	/* var result interface{}
	if c.Expression != nil {
		val := c.Expression.Ejecutar(ast, env)
		result = val
	} else {
		result = nil
	}
	return result */
	fmt.Println("Entre CallParams")
	// fmt.Println("IDEEEEE: ", c.Expression)
	return environment.Symbol{Value: c.Expression.(Variable).Id}
}
