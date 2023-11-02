package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
)

type While struct {
	Line         int
	Col          int
	Expression   interfaces.Expression
	Instructions []interface{}
}

func NewWhile(line int, col int, expression interfaces.Expression, instructions []interface{}) While {
	return While{line, col, expression, instructions}
}

func (w While) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generacion While")
	fmt.Println("Generacion While")
	var condition environment.Value
	var result environment.Value

	init := gen.NewLabel()
	gen.AddLabel(init)
	condition = w.Expression.Ejecutar(ast, env, gen)
	// Condicion verdadera
	gen.AddContinueLabel(init)
	gen.AddBreakLabel(condition.FalseLabel[0].(string))
	for _, lbl := range condition.TrueLabel {
		gen.AddLabel(lbl.(string))
		for _, inst := range w.Instructions {
			result = inst.(interfaces.Instruction).Ejecutar(ast, env, gen).(environment.Value)
			if result.BreakFlag {
				// break
			}
			if result.ReturnFlag {
				return result
			}
		}
		gen.AddGoto(init)
	}
	// Condicion falsa
	// gen.AddLabel(condition.FalseLabel[len(condition.FalseLabel)-1].(string))
	for _, lbl := range condition.FalseLabel {
		gen.AddLabel(lbl.(string))
	}
	return result
}
