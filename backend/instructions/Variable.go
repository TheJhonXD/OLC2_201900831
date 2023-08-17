package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type Variable struct {
	Line  int
	Col   int
	Value interface{}
}

func NewVar(line int, col int, value interface{}) Variable {
	return Variable{line, col, value}
}

func (v Variable) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	val := v.Value.(interfaces.Expression).Ejecutar(ast, env)
	consoleOut := fmt.Sprintf("%v", val.Value)
	ast.SetPrint(consoleOut)
	return nil
}
