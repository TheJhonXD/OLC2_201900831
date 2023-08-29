package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type Print struct {
	Line  int
	Col   int
	Value interface{}
}

func NewPrint(line int, col int, val interface{}) Print {
	return Print{line, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	valToPrint := p.Value.(interfaces.Expression).Ejecutar(ast, env)
	consoleOut := fmt.Sprintf("%v", valToPrint.Value)
	ast.SetPrint(consoleOut)
	return valToPrint
}
