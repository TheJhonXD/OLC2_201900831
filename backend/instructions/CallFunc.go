package instructions

import (
	"Server/environment"
	"Server/expressions"
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

func (c CallFunc) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre CallFunc Instruction")
	result := env.(environment.Env).GetFunc(c.Id)
	var result2 environment.Symbol
	if result.RtnType == environment.NULL {
		//Crear un nuevo entorno para la funcion
		funcEnv := environment.NewEnv(env.(environment.Env), "FUNC "+c.Id)

		/* Asignar los valores de entrada a los parametros de la funcion */
		// Guardar la variables de la función
		var param environment.SymbolFuncParam
		for _, p := range result.Args {
			param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
			funcEnv.SaveVar(param.InternalId, environment.Symbol{Line: param.Line, Col: param.Col, Type: param.Type, Value: 0, Const: false})
		}
		// Asignar valores a las variaables recientemente guardadas
		for i, a := range c.Args {
			if a.(expressions.CallParams).Id == "" {
				param = result.Args[i].(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
				var_param := funcEnv.GetVar(param.InternalId)
				var_param.Value = a.(expressions.CallParams).Expression.Ejecutar(ast, env).Value
				funcEnv.SetVar(param.InternalId, var_param)
			} else {
				for _, p := range result.Args {
					param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
					if param.ExternalId == a.(expressions.CallParams).Id {
						var_param := funcEnv.GetVar(param.InternalId)
						var_param.Value = a.(expressions.CallParams).Expression.Ejecutar(ast, env).Value
						funcEnv.SetVar(param.InternalId, var_param)
						break
					}
				}
			}
		}
		//Ejecutar las instrucciones de la funcion
		for _, inst := range result.Block {
			inst.(interfaces.Instruction).Ejecutar(ast, funcEnv)
			if result2.ReturnFlag == true {
				fmt.Println("Error: La funcion no tiene retorno")
			}
		}
	} else {
		//Crear un nuevo entorno para la funcion
		funcEnv := environment.NewEnv(env.(environment.Env), "FUNC "+c.Id)

		/* Asignar los valores de entrada a los parametros de la funcion */
		// Guardar la variables de la función
		var param environment.SymbolFuncParam
		for _, p := range result.Args {
			param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
			funcEnv.SaveVar(param.InternalId, environment.Symbol{Line: param.Line, Col: param.Col, Type: param.Type, Value: 0, Const: false})
		}
		// Asignar valores a las variaables recientemente guardadas
		for i, a := range c.Args {
			if a.(expressions.CallParams).Id == "" {
				param = result.Args[i].(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
				var_param := funcEnv.GetVar(param.InternalId)
				var_param.Value = a.(expressions.CallParams).Expression.Ejecutar(ast, env).Value
				funcEnv.SetVar(param.InternalId, var_param)
			} else {
				for _, p := range result.Args {
					param = p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
					if param.ExternalId == a.(expressions.CallParams).Id {
						var_param := funcEnv.GetVar(param.InternalId)
						var_param.Value = a.(expressions.CallParams).Expression.Ejecutar(ast, env).Value
						funcEnv.SetVar(param.InternalId, var_param)
						break
					}
				}
			}
		}
		//Ejecutar las instrucciones de la funcion
		for _, inst := range result.Block {
			result2 = inst.(interfaces.Instruction).Ejecutar(ast, funcEnv).(environment.Symbol)
			if result2.ReturnFlag == true {
				if result2.Type == result.RtnType {
					return result2
				} else {
					fmt.Println("Error: El tipo de retorno no coincide")
				}
			}
		}
	}
	result2 = environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, Const: false}
	return result2
}
