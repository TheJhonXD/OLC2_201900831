package instructions

import (
	"Server/environment"
	"Server/interfaces"
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

func (w While) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	breakAux := false
	var result environment.Symbol
	condition := w.Expression.Ejecutar(ast, env)
	if condition.Type != environment.BOOLEAN {
		ast.SetError("Error de tipos en la condicion del while")
		return result
	}

	if condition.Value == true {
		var whileEnv environment.Env
		whileEnv = environment.NewEnv(env.(environment.Env), "WHILE")
		for {
			condition = w.Expression.Ejecutar(ast, whileEnv)
			if condition.Value != true {
				break
			}
			for _, inst := range w.Instructions {
				result = inst.(interfaces.Instruction).Ejecutar(ast, whileEnv).(environment.Symbol)

				if result.BreakFlag == true {
					result.BreakFlag = false
					breakAux = true
					break
				} else if result.ContinueFlag == true {
					result.ContinueFlag = false
					break
				} else if result.ReturnFlag == true {
					result.ReturnFlag = false
					return result
				}

			}
			if breakAux {
				break
			}
		}
	}
	return result
}
