package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
)

type Float struct {
	Line  int
	Col   int
	Value interfaces.Expression
}

func NewFloat(line int, col int, value interfaces.Expression) Float {
	return Float{Line: line, Col: col, Value: value}
}

func (f Float) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	result = f.Value.Ejecutar(ast, env, gen)

	if result.Type == environment.INTEGER {
		nuevo := gen.NewTmp()
		gen.AddComment("Convirtiendo a entero")
		gen.AddExpression(nuevo, "(float)"+result.Value, "0", "+")
		result.Value = nuevo
		result.Type = environment.FLOAT
	}

	return result
}
