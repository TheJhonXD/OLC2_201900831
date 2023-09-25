package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type Print struct {
	Line  int
	Col   int
	Value []interface{}
}

func NewPrint(line int, col int, val []interface{}) Print {
	return Print{line, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre Print")
	var result interface{}
	var consoleOut string
	for _, val := range p.Value {
		valToPrint := val.(interfaces.Expression).Ejecutar(ast, env)
		result = valToPrint
		consoleOut += fmt.Sprintf("%v", valToPrint.Value) + " "
	}
	ast.SetPrint(consoleOut)
	return result
}
