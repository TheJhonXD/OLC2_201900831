package instructions

import (
	"Server/environment"
	"Server/interfaces"
)

type VectorMethod struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Method     string
}

func NewVectorMethod(line int, col int, ide string, expression interfaces.Expression, method string) VectorMethod {
	return VectorMethod{Line: line, Col: col, Id: ide, Expression: expression, Method: method}
}

func (v VectorMethod) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	result := env.(environment.Env).GetVar(v.Id)
	if result.Type == environment.NULL {
		return environment.Symbol{Line: v.Line, Col: v.Col, Type: environment.NULL, Value: result, Const: result.Const}
	}
	if v.Method == "append" {
		var valuesAux []interface{}
		for _, val := range result.Value.([]interface{}) {
			valuesAux = append(valuesAux, val)
		}
		valuesAux = append(valuesAux, v.Expression.Ejecutar(ast, env))
		result.Value = valuesAux
		//Actualizo la variable
		env.(environment.Env).SetVar(v.Id, result)
	} else if v.Method == "removeLast" {
		var valuesAux []interface{}
		if len(result.Value.([]interface{})) > 0 {
			for _, val := range result.Value.([]interface{}) {
				valuesAux = append(valuesAux, val)
			}
			valuesAux = valuesAux[:len(valuesAux)-1]
			result.Value = valuesAux
			//Actualizo la variable
			env.(environment.Env).SetVar(v.Id, result)
		} else {
			ast.SetError("Error Semantico: El vector \"" + v.Id + "\" esta vacio")
		}
	} else if v.Method == "remove" {
		if len(result.Value.([]interface{})) > 0 {
			var valuesAux []interface{}
			pos := v.Expression.Ejecutar(ast, env).Value.(int)
			if len(result.Value.([]interface{})) > pos {
				for _, val := range result.Value.([]interface{}) {
					valuesAux = append(valuesAux, val)
				}
				newSlice := append(valuesAux[:pos], valuesAux[pos+1:])
				result.Value = newSlice
				//Actualizo la variable
				env.(environment.Env).SetVar(v.Id, result)
			} else {
				ast.SetError("Error Semantico: El indice \"" + v.Id + "\" esta fuera de rango")
			}
		} else {
			ast.SetError("Error Semantico: El vector \"" + v.Id + "\" esta vacio")
		}
	}
	return result
}
