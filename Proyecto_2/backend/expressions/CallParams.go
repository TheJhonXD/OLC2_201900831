package expressions

import (
	"Server/environment"
	"Server/generator"
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

func (c CallParams) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	fmt.Println("entro a call params")
	if strings.Contains(fmt.Sprintf("%T", c.Expression), "Primitive") {
		// fmt.Println("Entre CallParams EXP PRIMITIVE")
		return environment.Value{Value: c.Expression.Ejecutar(ast, env, gen).Value}
	} else if strings.Contains(fmt.Sprintf("%T", c.Expression), "Operation") {
		return environment.Value{Value: "0"}
	} else if strings.Contains(fmt.Sprintf("%T", c.Expression), "Vector") {
		return environment.Value{Value: "0"}
	}
	return environment.Value{Value: c.Expression.(Variable).Ejecutar(ast, env, gen).Value}
}
