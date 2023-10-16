package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
)

type If struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Block      []interface{}
	ElseIf     []interface{}
	ElseInstr  []interface{}
}

func NewIf(line int, col int, expression interfaces.Expression, block []interface{}, elseif []interface{}, elseinstr []interface{}) If {
	return If{line, col, expression, block, elseif, elseinstr}
}

func (i If) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generacion if")
	fmt.Println("Generacion if")
	var condition environment.Value
	var result environment.Value
	condition = i.Expression.Ejecutar(ast, env, gen)
	lblOut := gen.NewLabel()
	/* *************** CONDICIONAL IF *************** */
	var ifEnv environment.Env
	ifEnv = environment.NewEnv(env.(environment.Env), "IF")
	for _, lbl := range condition.TrueLabel {
		gen.AddLabel(lbl.(string))
		// gen.AddLabel(condition.TrueLabel[0].(string))
		for _, inst := range i.Block {
			result = inst.(interfaces.Instruction).Ejecutar(ast, ifEnv, gen).(environment.Value)
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
		gen.AddGoto(lblOut)
	}

	gen.AddLabel(condition.FalseLabel[len(condition.FalseLabel)-1].(string))

	/* *************** CONDICIONAL ELSE IF *************** */
	if i.ElseIf != nil {
		var elseifEnv environment.Env
		elseifEnv = environment.NewEnv(env.(environment.Env), "ELSE IF")
		var condition_elif environment.Value
		for _, elseifInstr := range i.ElseIf {
			condition_elif = elseifInstr.(If).Expression.Ejecutar(ast, env, gen)
			for _, lbl := range condition_elif.TrueLabel {
				gen.AddLabel(lbl.(string))
				for _, inst := range elseifInstr.(If).ElseIf {
					result = inst.(interfaces.Instruction).Ejecutar(ast, elseifEnv, gen).(environment.Value)
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
				gen.AddGoto(lblOut)
			}

			for _, lbl := range condition_elif.FalseLabel[:len(condition_elif.FalseLabel)-1] {
				gen.AddLabel(lbl.(string))
			}

			gen.AddLabel(condition_elif.FalseLabel[len(condition_elif.FalseLabel)-1].(string))
		}
	}

	/* *************** CONDICIONAL ELSE *************** */
	// gen.AddLabel(condition.FalseLabel[0].(string))
	var elseEnv environment.Env
	elseEnv = environment.NewEnv(env.(environment.Env), "ELSE")
	for _, inst := range i.ElseInstr {
		result = inst.(interfaces.Instruction).Ejecutar(ast, elseEnv, gen).(environment.Value)
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
	for _, lbl := range condition.FalseLabel[:len(condition.FalseLabel)-1] {
		gen.AddLabel(lbl.(string))
	}

	gen.AddLabel(lblOut)
	return result
}
