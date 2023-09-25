package instructions

import (
	"Server/environment"
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

func (i If) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	condition := i.Expression.Ejecutar(ast, env)
	var result environment.Symbol
	if condition.Type != environment.BOOLEAN {
		ast.AddError(i.Line, i.Col, env.(environment.Env).Id, "La expresion de la sentencia if debe ser booleana")
		return nil
	}

	if condition.Value == true {
		fmt.Println("Entre IIIIIIf")
		var ifEnv environment.Env
		ifEnv = environment.NewEnv(env.(environment.Env), "IF")
		for _, inst := range i.Block {
			result = inst.(interfaces.Instruction).Ejecutar(ast, ifEnv).(environment.Symbol)
			if result.BreakFlag {
				break
			}
			if result.ContinueFlag {
				break
			}
			if result.ReturnFlag {
				// result.ReturnFlag = false
				return result
			}
		}
		return result
	} else if i.ElseIf != nil {
		fmt.Println("Entre ELSE IF")
		for _, elseIfInstr := range i.ElseIf {
			var condition_elif environment.Symbol
			condition_elif = elseIfInstr.(If).Expression.Ejecutar(ast, env)
			if condition_elif.Value == true {
				var elseIfEnv environment.Env
				elseIfEnv = environment.NewEnv(env.(environment.Env), "ELSE IF")
				for _, instr := range elseIfInstr.(If).ElseIf {
					result = instr.(interfaces.Instruction).Ejecutar(ast, elseIfEnv).(environment.Symbol)
					if result.BreakFlag {
						break
					}
					if result.ContinueFlag {
						break
					}
					if result.ReturnFlag {
						// result.ReturnFlag = false
						return result
					}
				}
				return result
			}
		}
		if i.ElseInstr != nil {
			fmt.Println("ENTREE??????")
			var elseEnv environment.Env
			elseEnv = environment.NewEnv(env.(environment.Env), "ELSE")
			for _, instr := range i.ElseInstr {
				result = instr.(interfaces.Instruction).Ejecutar(ast, elseEnv).(environment.Symbol)
				if result.BreakFlag {
					break
				}
				if result.ContinueFlag {
					break
				}
				if result.ReturnFlag {
					// result.ReturnFlag = false
					return result
				}
			}
			return result
		}
	} else {
		fmt.Println("Entre ELSE")
		var elseEnv environment.Env
		elseEnv = environment.NewEnv(env.(environment.Env), "ELSE")
		for _, inst := range i.ElseInstr {
			result = inst.(interfaces.Instruction).Ejecutar(ast, elseEnv).(environment.Symbol)
			if result.BreakFlag {
				break
			}
			if result.ContinueFlag {
				break
			}
			if result.ReturnFlag {
				// result.ReturnFlag = false
				return result
			}
		}
		return result
	}
	return result
}
