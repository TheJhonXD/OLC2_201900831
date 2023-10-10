package instructions

import (
	"Server/environment"
	"Server/generator"
)

type Function struct {
	Line    int
	Col     int
	Id      string
	Args    []interface{}
	RtnType environment.TipoExpresion
	Block   []interface{}
}

func NewFunction(line int, col int, ide string, args []interface{}, rtnType environment.TipoExpresion, block []interface{}) Function {
	return Function{Line: line, Col: col, Id: ide, Args: args, RtnType: rtnType, Block: block}
}

func (f Function) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
