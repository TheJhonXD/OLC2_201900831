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
		ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "La variable \""+v.Id+"\" no existe")
		return environment.Symbol{Line: v.Line, Col: v.Col, Type: environment.NULL, Value: result, Const: result.Const}
	}
	if v.Method == "append" {
		var valuesAux []interface{}
		if len(result.Value.([]interface{})) > 0 {
			for _, val := range result.Value.([]interface{}) {
				valuesAux = append(valuesAux, val)
			}
		}
		valuesAux = append(valuesAux, v.Expression.Ejecutar(ast, env).Value)
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
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El vector \""+v.Id+"\" esta vacio")
		}
	} else if v.Method == "remove" {
		if len(result.Value.([]interface{})) > 0 {
			var valuesAux []interface{}
			pos := v.Expression.Ejecutar(ast, env).Value.(int)
			if len(result.Value.([]interface{})) > pos {
				for i, val := range result.Value.([]interface{}) {
					if i != pos {
						valuesAux = append(valuesAux, val)
					}
				}

				result.Value = valuesAux
				//Actualizo la variable
				env.(environment.Env).SetVar(v.Id, result)
			} else {
				ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El indice \""+v.Id+"\" esta fuera de rango")
			}
		} else {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El vector \""+v.Id+"\" esta vacio")
		}
	}
	return result
}
