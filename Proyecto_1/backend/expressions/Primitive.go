package expressions

import (
	"Server/environment"
)

type Primitive struct {
	Line int
	Col  int
	Val  interface{}
	Type environment.TipoExpresion
}

func NewPrimitive(line int, col int, val interface{}, tipo environment.TipoExpresion) Primitive {
	return Primitive{Line: line, Col: col, Val: val, Type: tipo}
}

func (p Primitive) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{Line: p.Line, Col: p.Col, Type: p.Type, Value: p.Val}
}
