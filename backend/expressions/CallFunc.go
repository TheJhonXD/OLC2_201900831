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
				funcEnv.SaveVar(param.InternalId, environment.Symbol{Line: param.Line, Col: param.Col, Type: param.Type, Value: 0, Const: false})
			}

		}
		// Asignar valores a las variaables recientemente guardadas
		for i, a := range c.Args {
			if a.(CallParams).Id == "" {
				var_name := a.(CallParams).Ejecutar(ast, env)
				param = result.Args[i].(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
				var_param := funcEnv.GetVar(param.InternalId)
				var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
				funcEnv.SetVar(param.InternalId, var_param)
				if a.(CallParams).Amp == true && param.Reference == true {
					var_reference[param.InternalId] = var_name.Value.(string)
				} else {
					flagrun = false
					fmt.Println("Error: No se puede asignar una variable por valor a una variable por referencia")
				}
			} else {
				for _, p := range result.Args {
					param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
					if param.ExternalId == a.(CallParams).Id {
						var_name := a.(CallParams).Ejecutar(ast, env)
						var_param := funcEnv.GetVar(param.InternalId)
						var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
						funcEnv.SetVar(param.InternalId, var_param)
						if a.(CallParams).Amp == true && param.Reference == true {
							var_reference[param.InternalId] = var_name.Value.(string)
						} else {
							flagrun = false
							fmt.Println("Error: No se puede asignar una variable por valor a una variable por referencia")
						}
						break
					}
				}
			}
		}
		//Ejecutar las instrucciones de la funcion
		if flagrun {
			for _, inst := range result.Block {
				inst.(interfaces.Instruction).Ejecutar(ast, funcEnv)
				if result2.ReturnFlag == true {
					fmt.Println("Error: La funcion no tiene retorno")
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
			if param.Reference == true {
				var_reference[param.InternalId] = ""
			}
			funcEnv.SaveVar(param.InternalId, environment.Symbol{Line: param.Line, Col: param.Col, Type: param.Type, Value: 0, Const: false})
		}
		// Asignar valores a las variaables recientemente guardadas
		for i, a := range c.Args {
			if a.(CallParams).Id == "" {
				var_name := a.(CallParams).Ejecutar(ast, env)
				param = result.Args[i].(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
				var_param := funcEnv.GetVar(param.InternalId)
				var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
				funcEnv.SetVar(param.InternalId, var_param)
				if a.(CallParams).Amp == true && param.Reference == true {
					var_reference[param.InternalId] = var_name.Value.(string)
				} else {
					flagrun = false
					fmt.Println("Error: No se puede asignar una variable por valor a una variable por referencia")
				}
			} else {
				for _, p := range result.Args {
					fmt.Println("RESULT: ")
					param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
					if param.ExternalId == a.(CallParams).Id {
						fmt.Println("RESULT2: ")
						var_name := a.(CallParams).Id
						fmt.Println("NAME: ", var_name)
						var_param := funcEnv.GetVar(param.InternalId)
						var_param.Value = a.(CallParams).Expression.Ejecutar(ast, env).Value
						funcEnv.SetVar(param.InternalId, var_param)
						if a.(CallParams).Amp == true && param.Reference == true {
							var_reference[param.InternalId] = var_name
						} else if a.(CallParams).Amp != param.Reference {
							flagrun = false
							fmt.Println("Error: No se puede asignar una variable por valor a una variable por referencia")
						}
						break
					}
				}
			}
		}
		fmt.Println("EFFFFFE")
		//Ejecutar las instrucciones de la funcion
		if flagrun {
			for _, inst := range result.Block {
				result2 = inst.(interfaces.Instruction).Ejecutar(ast, funcEnv).(environment.Symbol)
				for key, value := range var_reference {
					varaux1 := funcEnv.GetVar(key)
					varaux2 := env.(environment.Env).GetVar(value)
					varaux2.Value = varaux1.Value
					env.(environment.Env).SetVar(value, varaux2)
				}
				if result2.ReturnFlag == true {
					if result2.Type == result.RtnType {
						return result2
					} else {
						fmt.Println("Error: El tipo de retorno no coincide")
					}
				}
			}
		}
	}
	result2 = environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
	return result2
}
