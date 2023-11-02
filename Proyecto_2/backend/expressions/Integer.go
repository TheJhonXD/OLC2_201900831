package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
)

type Integer struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewInteger(line int, col int, value interfaces.Expression) Integer {
	return Integer{Line: line, Col: col, Value: value}
}

func (i Integer) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	result = i.Value.Ejecutar(ast, env, gen)

	if result.Type == environment.FLOAT {
		nuevo := gen.NewTmp()
		gen.AddComment("Convirtiendo a entero")
		gen.AddExpression(nuevo, "(int)"+result.Value, "0", "+")
		result.Value = nuevo
		result.Type = environment.INTEGER
	} else if result.Type == environment.STRING {
		fmt.Println("Entro aqui:", result.Value)
		result.Type = environment.INTEGER
	}

	return result
}
