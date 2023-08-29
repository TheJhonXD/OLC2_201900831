package expressions

import (
	"Server/environment"
)

type Variable struct {
	Line int
	Col  int
	Id   string
}

func NewVar(line int, col int, ide string) Variable {
	return Variable{Line: line, Col: col, Id: ide}
}

func (v Variable) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Env).GetVar(v.Id)

	if result.Type == environment.NULL {
		return environment.Symbol{Line: v.Line, Col: v.Col, Type: environment.NULL, Value: result, Const: result.Const}
	}

	return result
}
