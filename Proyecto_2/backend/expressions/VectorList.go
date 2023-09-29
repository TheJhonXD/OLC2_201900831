package expressions

import "Server/environment"

type VectorList struct {
	Line   int
	Col    int
	Values []interface{}
}

func NewVectorList(line int, col int, values []interface{}) VectorList {
	return VectorList{line, col, values}
}

func (v VectorList) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	return v.Values
}
