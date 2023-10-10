package expressions

import (
	"Server/environment"
	"Server/generator"
)

type Variable struct {
	Line int
	Col  int
	Id   string
}

func NewVar(line int, col int, ide string) Variable {
	return Variable{Line: line, Col: col, Id: ide}
}

func (v Variable) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
