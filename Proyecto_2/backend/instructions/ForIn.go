package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type ForIn struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Op_left    interfaces.Expression
	Op_right   interfaces.Expression
	VecId      string
	Block      []interface{}
}

func NewForIn(line int, col int, ide string, expression interfaces.Expression, left interfaces.Expression, right interfaces.Expression, vecId string, block []interface{}) ForIn {
	return ForIn{line, col, ide, expression, left, right, vecId, block}
}

func (f ForIn) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
