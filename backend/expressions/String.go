package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type String struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewString(line int, col int, value interfaces.Expression) String {
	return String{Line: line, Col: col, Value: value}
}

func (s String) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := s.Value.Ejecutar(ast, env)
	if result.Type == environment.INTEGER {
		result.Value = result.Value.(int)
		result.Type = environment.STRING
	} else if result.Type == environment.FLOAT {
		result.Value = strconv.FormatFloat(result.Value.(float64), 'f', 4, 64)
		result.Type = environment.STRING
	} else if result.Type == environment.BOOLEAN {
		result.Value = strconv.FormatBool(result.Value.(bool))
		result.Type = environment.STRING
	} else {
		result.Value = ""
		result.Type = environment.NULL
		fmt.Println("Error: No se puede convertir el valor a String")
	}
	return result
}
