package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type Float struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewFloat(line int, col int, value interfaces.Expression) Float {
	return Float{Line: line, Col: col, Value: value}
}

func (f Float) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := f.Value.Ejecutar(ast, env)
	if result.Type == environment.STRING {
		floatNumber, err := strconv.ParseFloat(result.Value.(string), 64)
		if err != nil {
			fmt.Println("Error: No se puede convertir el valor a Float")
			result.Value = 0
			result.Type = environment.NULL
			return result
		}
		result.Value = floatNumber
		result.Type = environment.FLOAT
	} else if result.Type == environment.FLOAT {
		result.Type = environment.FLOAT
	} else {
		result.Value = 0
		result.Type = environment.NULL
		fmt.Println("Error: No se puede convertir el valor a Float")
	}
	return result
}
