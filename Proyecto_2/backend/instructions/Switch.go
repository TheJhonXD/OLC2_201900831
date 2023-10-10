package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Switch struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Cases      []interface{}
	Default    []interface{}
}

func NewSwitch(line int, col int, expression interfaces.Expression, cases []interface{}, def []interface{}) Switch {
	return Switch{line, col, expression, cases, def}
}

func (s Switch) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
