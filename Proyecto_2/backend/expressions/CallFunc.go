package expressions

import (
	"Server/environment"
	"Server/generator"
)

type CallFunc struct {
	Line int
	Col  int
	Id   string
	Args []interface{}
}

func NewCallFunc(line int, col int, ide string, args []interface{}) CallFunc {
	return CallFunc{Line: line, Col: col, Id: ide, Args: args}
}

func (c CallFunc) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}
