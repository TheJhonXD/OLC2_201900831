package instructions

import (
	"Server/environment"
	"Server/interfaces"
)

type ForIn struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Op_left    interfaces.Expression
	Op_right   interfaces.Expression
	Block      []interface{}
}

func NewForIn(line int, col int, ide string, expression interfaces.Expression, left interfaces.Expression, right interfaces.Expression, block []interface{}) ForIn {
	return ForIn{line, col, ide, expression, left, right, block}
}

func (f ForIn) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	// fmt.Println(fmt.Sprintf("%T", env.(environment.Env).GetVar(f.Id)))
	if f.Expression != nil {
		iter_var := f.Expression.Ejecutar(ast, env)
		if iter_var.Type == environment.STRING {
			var forEnv environment.Env
			forEnv = environment.NewEnv(env.(environment.Env), "FOR IN")
			var_exists := env.(environment.Env).GetVar(f.Id)
			if var_exists.Type == environment.NULL {
				var newSymbol environment.Symbol
				newSymbol = environment.Symbol{Line: f.Line + 3, Col: f.Col + 3, Type: environment.CHAR, Value: 0, Const: true}
				forEnv.SaveVar(f.Id, newSymbol)

				for _, var_aux_i := range iter_var.Value.(string) {
					newSymbol.Value = string(var_aux_i)
					forEnv.SetConstVar(f.Id, newSymbol)
					for _, inst := range f.Block {
						inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
					}
				}
			}
		} else {
			//Todo: la variable no es un string
		}
	} else {
		var left, right environment.Symbol
		left = f.Op_left.Ejecutar(ast, env)
		right = f.Op_right.Ejecutar(ast, env)
		if left.Type == environment.INTEGER && right.Type == environment.INTEGER {
			var forEnv environment.Env
			forEnv = environment.NewEnv(env.(environment.Env), "FOR IN")
			var_exists := env.(environment.Env).GetVar(f.Id)
			if var_exists.Type == environment.NULL {
				var newSymbol environment.Symbol
				newSymbol = environment.Symbol{Line: f.Line + 3, Col: f.Col + 3, Type: environment.CHAR, Value: 0, Const: true}
				forEnv.SaveVar(f.Id, newSymbol)
				if left.Value.(int) < right.Value.(int) {
					//! El for cambia
					for var_i := left.Value.(int); var_i <= right.Value.(int); var_i++ {
						newSymbol.Value = var_i
						forEnv.SetConstVar(f.Id, newSymbol)
						for _, inst := range f.Block {
							inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
						}
					}
				}
			}

		} else {
			//Todo: error de tipos
		}
	}
	return nil
}
