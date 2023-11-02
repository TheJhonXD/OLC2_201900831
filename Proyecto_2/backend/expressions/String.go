package expressions

import (
	"Server/environment"
	"Server/generator"
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

func (s String) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	result = s.Value.Ejecutar(ast, env, gen)
	if result.Type == environment.INTEGER {
		newVal := strconv.Itoa(result.IntValue)
		newP := NewPrimitive(s.Line, s.Col, newVal, environment.STRING)
		result = newP.Ejecutar(ast, env, gen)
		result.Type = environment.STRING
	} else if result.Type == environment.FLOAT {
		cadena := strconv.FormatFloat(result.FloatValue, 'f', -1, 64)
		fmt.Println(cadena)
		newP := NewPrimitive(s.Line, s.Col, cadena, environment.STRING)
		result = newP.Ejecutar(ast, env, gen)
		result.Type = environment.STRING
	}
	return result
}
