package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type Return struct {
	Line       int
	Col        int
	Expression interfaces.Expression
}

func NewReturn(line int, col int, exp interfaces.Expression) Return {
	return Return{line, col, exp}
}

func (r Return) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre Return")
	var result environment.Symbol
	if env.(environment.Env).IsFunc() {
		if r.Expression != nil {
			val := r.Expression.Ejecutar(ast, env)
			result = environment.Symbol{Line: r.Line, Col: r.Col, Type: val.Type, Value: val.Value, ReturnFlag: true}
		} else {
			result = environment.Symbol{Line: r.Line, Col: r.Col, Type: environment.NULL, Value: nil, ReturnFlag: true}
		}

	}
	return result
}
