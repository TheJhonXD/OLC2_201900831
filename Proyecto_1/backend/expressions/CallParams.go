package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
	"strings"
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
	fmt.Println("Entre CallParams EXP")
	// fmt.Println("IDEEEEE: ", c.Expression)
	fmt.Println(fmt.Sprintf("%T", c.Expression))
	if strings.Contains(fmt.Sprintf("%T", c.Expression), "Primitive") {
		// fmt.Println("Entre CallParams EXP PRIMITIVE")
		return environment.Symbol{Value: c.Expression.Ejecutar(ast, env)}
	} else if strings.Contains(fmt.Sprintf("%T", c.Expression), "Operation") {
		return environment.Symbol{Value: 0}
	} else if strings.Contains(fmt.Sprintf("%T", c.Expression), "Vector") {
		return environment.Symbol{Value: 0}
	}
	return environment.Symbol{Value: c.Expression.(Variable).Id}
}
