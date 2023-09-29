package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type CallFunc struct {
	Line int
	Col  int
	Id   string
	Args []interface{}
}

func NewCallFunc(line int, col int, ide string, args []interface{}) CallFunc {
	return CallFunc{Line: line, Col: col, Id: ide, Args: args}
}

func (c CallFunc) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	fmt.Println("Entre CallFunc Expressions")
	flagReference := false
	result, funcExists := env.(environment.Env).GetFunc(c.Id)
	if !funcExists {
		ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "La funcion \""+c.Id+"\" no existe")
		return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
	}
	var flagrun bool = true
	var result2 environment.Symbol
	if result.RtnType == environment.NULL {
		//Crear un nuevo entorno para la funcion
		funcEnv := environment.NewEnv(env.(environment.Env), "FUNC "+c.Id)
		var_reference := make(map[string]string)
		/* Asignar los valores de entrada a los parametros de la funcion */
		// Guardar la variables de la función
		var param environment.SymbolFuncParam
		for _, p := range result.Args {
			param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
			if param.Type != environment.NULL {
				if param.Reference == true {
					var_reference[param.InternalId] = ""
				}
				if !funcEnv.SaveVar(param.InternalId, environment.Symbol{Line: param.Line, Col: param.Col, Type: param.Type, Value: 0, Const: false}) {
					ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" ya existe")
					return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
				}
			} else {
				ast.AddError(c.Line, c.Col, funcEnv.Id, "Error en la asignacion de parametros")
				return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
			}

		}
		// Asignar valores a las variaables recientemente guardadas
		for i, a := range c.Args {
			if a.(CallParams).Id == "" {
				var_name := a.(CallParams).Ejecutar(ast, env)
				param = result.Args[i].(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
				var_param := funcEnv.GetVar(param.InternalId)
				if var_param.Type != environment.NULL {
					if var_param.Type == a.(CallParams).Expression.Ejecutar(ast, env).Type {
						var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
						if funcEnv.SetVar(param.InternalId, var_param).Type == environment.NULL {
							ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" no existe, no se puede asignar")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						if a.(CallParams).Amp == true && param.Reference == true {
							flagReference = true
							var_reference[param.InternalId] = var_name.Value.(string)
						} else {
							flagrun = false
							ast.AddError(c.Line, c.Col, funcEnv.Id, "No se puede asignar una variable por valor a una variable por referencia")
						}
					} else {
						ast.AddError(c.Line, c.Col, funcEnv.Id, "Los tipos no coinciden")
						return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
					}
				} else {
					ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" no existe")
					return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
				}
			} else {
				for _, p := range result.Args {
					param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
					if param.ExternalId == a.(CallParams).Id {
						var_name := a.(CallParams).Ejecutar(ast, env)
						var_param := funcEnv.GetVar(param.InternalId)
						if var_param.Type != environment.NULL {
							if var_param.Type == a.(CallParams).Expression.Ejecutar(ast, env).Type {
								var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
								if funcEnv.SetVar(param.InternalId, var_param).Type == environment.NULL {
									ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" no existe, no se puede asignar")
									return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
								}
								if a.(CallParams).Amp == true && param.Reference == true {
									flagReference = true
									var_reference[param.InternalId] = var_name.Value.(string)
								} else {
									flagrun = false
									ast.AddError(c.Line, c.Col, funcEnv.Id, "No se puede asignar una variable por valor a una variable por referencia")
								}
								break
							} else {
								ast.AddError(c.Line, c.Col, funcEnv.Id, "Los tipos no coinciden")
								return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
							}
						} else {
							ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
					}
				}
			}
		}
		//Ejecutar las instrucciones de la funcion
		if flagrun {
			for _, inst := range result.Block {
				inst.(interfaces.Instruction).Ejecutar(ast, funcEnv)
				if flagReference {
					for key, value := range var_reference {
						varaux1 := funcEnv.GetVar(key)
						if varaux1.Type == environment.NULL {
							ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+key+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						varaux2 := env.(environment.Env).GetVar(value)
						if varaux2.Type == environment.NULL {
							ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "La variable \""+value+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						varaux2.Value = varaux1.Value
						if varaux2.Type != varaux1.Type {
							ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "Los tipos no coinciden")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						if env.(environment.Env).SetVar(value, varaux2).Type == environment.NULL {
							ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "La variable \""+value+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
					}
				}
				if result2.ReturnFlag == true {
					ast.AddError(c.Line, c.Col, funcEnv.Id, "La funcion no debe retornar ningun valor")
				}
			}
		}
	} else {
		//Crear un nuevo entorno para la funcion
		funcEnv := environment.NewEnv(env.(environment.Env), "FUNC "+c.Id)
		var_reference := make(map[string]string)
		/* Asignar los valores de entrada a los parametros de la funcion */
		// Guardar la variables de la función
		var param environment.SymbolFuncParam
		for _, p := range result.Args {
			param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
			if param.Type != environment.NULL {
				if param.Reference == true {
					var_reference[param.InternalId] = ""
				}
				if !funcEnv.SaveVar(param.InternalId, environment.Symbol{Line: param.Line, Col: param.Col, Type: param.Type, Value: 0, Const: false}) {
					ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" ya existe")
					return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
				}
			} else {
				ast.AddError(c.Line, c.Col, funcEnv.Id, "Error en la asignacion de parametros")
				return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
			}
		}

		// Asignar valores a las variaables recientemente guardadas
		for i, a := range c.Args {
			if a.(CallParams).Id == "" {
				var_name := a.(CallParams).Ejecutar(ast, env)
				fmt.Println("EFE")
				param = result.Args[i].(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
				var_param := funcEnv.GetVar(param.InternalId)
				if var_param.Type != environment.NULL {
					var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
					if var_param.Type == a.(CallParams).Expression.Ejecutar(ast, env).Type {
						fmt.Println("EFE2")
						fmt.Println(flagrun)
						if funcEnv.SetVar(param.InternalId, var_param).Type == environment.NULL {
							ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" no existe, no se puede asignar")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						if a.(CallParams).Amp == true && param.Reference == true {
							flagReference = true
							var_reference[param.InternalId] = var_name.Value.(string)
						} else if a.(CallParams).Amp != param.Reference {
							flagrun = false
							ast.AddError(c.Line, c.Col, funcEnv.Id, "No se puede asignar una variable por valor a una variable por referencia")
						}
					} else {
						ast.AddError(c.Line, c.Col, funcEnv.Id, "Los tipos no coinciden")
					}
				} else {
					ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+param.InternalId+"\" no existe")
				}
			} else {
				for _, p := range result.Args {
					param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
					if param.ExternalId == a.(CallParams).Id {
						var_name := a.(CallParams).Id
						var_param := funcEnv.GetVar(param.InternalId)
						var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
						funcEnv.SetVar(param.InternalId, var_param)
						if a.(CallParams).Amp == true && param.Reference == true {
							flagReference = true
							var_reference[param.InternalId] = var_name
						} else if a.(CallParams).Amp != param.Reference {
							flagrun = false
							ast.AddError(c.Line, c.Col, funcEnv.Id, "No se puede asignar una variable por valor a una variable por referencia")
						}
						break
					}
				}
			}
		}

		//Ejecutar las instrucciones de la funcion
		if flagrun {
			for _, inst := range result.Block {
				result2 = inst.(interfaces.Instruction).Ejecutar(ast, funcEnv).(environment.Symbol)
				if flagReference {
					fmt.Println("?????????")
					for key, value := range var_reference {
						varaux1 := funcEnv.GetVar(key)
						if varaux1.Type == environment.NULL {
							ast.AddError(c.Line, c.Col, funcEnv.Id, "La variable \""+key+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						varaux2 := env.(environment.Env).GetVar(value)
						if varaux2.Type == environment.NULL {
							ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "La variable \""+value+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						varaux2.Value = varaux1.Value
						if varaux2.Type != varaux1.Type {
							ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "Los tipos no coinciden")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
						if env.(environment.Env).SetVar(value, varaux2).Type == environment.NULL {
							ast.AddError(c.Line, c.Col, env.(environment.Env).Id, "La variable \""+value+"\" no existe")
							return environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
						}
					}
				}
				if result2.ReturnFlag == true {
					if result2.Type == result.RtnType {
						return result2
					} else {
						ast.AddError(c.Line, c.Col, funcEnv.Id, "El tipo de retorno no coincide")
					}
				}
			}
		}
	}
	result2 = environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
	return result2
}
