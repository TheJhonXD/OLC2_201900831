package instructions

import (
	"Server/environment"
	"Server/interfaces"
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

func (g Guard) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol
	flagTransfer := false
	condition := g.Expression.Ejecutar(ast, env)
	if condition.Type == environment.BOOLEAN {
		if condition.Value == false {
			for _, inst := range g.Block {
				result = inst.(interfaces.Instruction).Ejecutar(ast, env).(environment.Symbol)
				if result.BreakFlag == true {
					flagTransfer = true
					return result
				} else if result.ContinueFlag == true {
					flagTransfer = true
					return result
				} else if result.ReturnFlag == true {
					flagTransfer = true
					return result
				}
			}
			if !flagTransfer {
				ast.AddError(g.Line, g.Col, env.(environment.Env).Id, "Guard debe tener una sentencia de transferencia")
			}
		}
	} else {
		ast.AddError(g.Line, g.Col, env.(environment.Env).Id, "La expresion de la sentencia guard debe ser booleana")
		return result
	}
	return result
}
