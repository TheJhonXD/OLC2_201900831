package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type ForIn struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Op_left    interfaces.Expression
	Op_right   interfaces.Expression
	VecId      string
	Block      []interface{}
}

func NewForIn(line int, col int, ide string, expression interfaces.Expression, left interfaces.Expression, right interfaces.Expression, vecId string, block []interface{}) ForIn {
	return ForIn{line, col, ide, expression, left, right, vecId, block}
}

func (f ForIn) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	// fmt.Println(fmt.Sprintf("%T", env.(environment.Env).GetVar(f.Id)))
	var result environment.Symbol
	var flag bool
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
						result = inst.(interfaces.Instruction).Ejecutar(ast, forEnv).(environment.Symbol)
						if result.BreakFlag {
							flag = true
							break
						}
						if result.ContinueFlag {
							break
						}
						if result.ReturnFlag {
							result.ReturnFlag = false
							return result
						}
					}
					if flag {
						break
					}
				}
			} else {
				ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "La variable ya existe en el ambito padre")
			}
		} else {
			ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "El tipo de la expresion no es string")
		}
	} else if f.Op_left != nil && f.Op_right != nil {
		var left, right environment.Symbol
		left = f.Op_left.Ejecutar(ast, env)
		right = f.Op_right.Ejecutar(ast, env)
		if left.Type == environment.INTEGER && right.Type == environment.INTEGER {
			var forEnv environment.Env
			forEnv = environment.NewEnv(env.(environment.Env), "FOR IN")
			var_exists := env.(environment.Env).GetVar(f.Id)
			if var_exists.Type == environment.NULL {
				var newSymbol environment.Symbol
				newSymbol = environment.Symbol{Line: f.Line + 3, Col: f.Col + 3, Type: environment.INTEGER, Value: 0, Const: true}
				forEnv.SaveVar(f.Id, newSymbol)
				if left.Value.(int) < right.Value.(int) {
					fmt.Println("Si entre")
					for var_i := left.Value.(int); var_i <= right.Value.(int); var_i++ {
						newSymbol.Value = var_i
						forEnv.SetConstVar(f.Id, newSymbol)
						for _, inst := range f.Block {
							result = inst.(interfaces.Instruction).Ejecutar(ast, forEnv).(environment.Symbol)
							if result.BreakFlag {
								flag = true
								break
							}
							if result.ContinueFlag {
								break
							}
							if result.ReturnFlag {
								result.ReturnFlag = false
								return result
							}
						}
						if flag {
							break
						}
					}
				} else {
					ast.AddError(f.Line, f.Col, forEnv.Id, "El limite inferior es mayor al limite superior")
				}
			} else {
				ast.AddError(f.Line, f.Col, forEnv.Id, "La variable ya existe en el ambito padre")
			}

		} else {
			ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "Los tipos de las expresiones no son enteros")
		}
	} else if f.VecId != "" {
		var forEnv environment.Env
		forEnv = environment.NewEnv(env.(environment.Env), "FOR IN")
		var_exist := env.(environment.Env).GetVar(f.Id)
		if var_exist.Type == environment.NULL {
			var newSymbol environment.Symbol
			newSymbol = environment.Symbol{Line: f.Line + 3, Col: f.Col + 3, Type: environment.CHAR, Value: 0, Const: true}
			forEnv.SaveVar(f.Id, newSymbol)
			listVec := env.(environment.Env).GetVar(f.VecId)
			if listVec.Type != environment.NULL {
				for _, val := range listVec.Value.([]interface{}) {
					newSymbol.Value = val
					forEnv.SetConstVar(f.Id, newSymbol)
					for _, inst := range f.Block {
						result = inst.(interfaces.Instruction).Ejecutar(ast, forEnv).(environment.Symbol)
						if result.BreakFlag {
							flag = true
							break
						}
						if result.ContinueFlag {
							break
						}
						if result.ReturnFlag {
							result.ReturnFlag = false
							return result
						}
					}
					if flag {
						break
					}
				}
			} else {
				ast.AddError(f.Line, f.Col, forEnv.Id, "El vector no existe")
			}
		} else {
			ast.AddError(f.Line, f.Col, forEnv.Id, "La variable ya existe en el ambito padre")
		}
	}
	return result
}
