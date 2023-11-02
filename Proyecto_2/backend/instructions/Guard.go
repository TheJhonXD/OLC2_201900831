package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
)

type Guard struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Block      []interface{}
}

func NewGuard(line int, col int, expression interfaces.Expression, block []interface{}) Guard {
	return Guard{line, col, expression, block}
}

func (g Guard) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generacion Guard")
	fmt.Println("Generacion Guard")
	var condition environment.Value
	var result environment.Value
	flagTransfer := false
	condition = g.Expression.Ejecutar(ast, env, gen)
	lblOut := gen.NewLabel()
	/* *************** CONDICIONAL IF *************** */
	if condition.Type == environment.BOOLEAN {
		for _, lbl := range condition.TrueLabel {
			gen.AddLabel(lbl.(string))
			gen.AddGoto(lblOut)
			// gen.AddLabel(condition.TrueLabel[0].(string))
		}
		// gen.AddLabel(condition.FalseLabel[len(condition.FalseLabel)-1].(string))

		for _, lbl := range condition.FalseLabel {
			gen.AddLabel(lbl.(string))
			for _, inst := range g.Block {
				result = inst.(interfaces.Instruction).Ejecutar(ast, env, gen).(environment.Value)
				if result.BreakFlag {
					break
				}
				if result.ContinueFlag {
					break
				}
				if result.ReturnFlag {
					return result
				}
			}
			if !flagTransfer {
				ast.AddError(g.Line, g.Col, env.(environment.Env).Id, "Guard debe tener una sentencia de transferencia")
			}
			gen.AddGoto(lblOut)
		}

		gen.AddLabel(lblOut)
	} else {
		ast.AddError(g.Line, g.Col, env.(environment.Env).Id, "La expresion de la sentencia guard debe ser booleana")
		return result
	}
	return result
}
