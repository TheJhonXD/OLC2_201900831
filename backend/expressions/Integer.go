package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type Integer struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewInteger(line int, col int, value interfaces.Expression) Integer {
	return Integer{Line: line, Col: col, Value: value}
}

func (i Integer) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := i.Value.Ejecutar(ast, env)
	if result.Type == environment.INTEGER {
		result.Type = environment.INTEGER
	} else if result.Type == environment.STRING {
		entero, err := strconv.Atoi(result.Value.(string))
		if err != nil {
			fmt.Println("Error: No se puede convertir el valor a Integer")
			result.Value = 0
			result.Type = environment.NULL
			return result
		}
		result.Value = entero
		result.Type = environment.INTEGER
	} else if result.Type == environment.FLOAT {
		result.Value = int(result.Value.(float64))
		result.Type = environment.INTEGER
	} else {
		result.Value = 0
		result.Type = environment.NULL
		fmt.Println("Error: No se puede convertir el valor a Integer")
	}
	return result
}
