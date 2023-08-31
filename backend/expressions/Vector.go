package expressions

import (
	"Server/environment"
	"Server/interfaces"
)

type Vector struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Method     string
}

func NewVector(line int, col int, ide string, expression interfaces.Expression, method string) Vector {
	return Vector{Line: line, Col: col, Id: ide, Expression: expression, Method: method}
}

func (v Vector) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Env).GetVar(v.Id)
	if result.Type == environment.NULL {
		return environment.Symbol{Line: v.Line, Col: v.Col, Type: environment.NULL, Value: result, Const: result.Const}
	}

	if v.Method == "isEmpty" {
		result.Type = environment.BOOLEAN
		if len(result.Value.([]interface{})) == 0 {
			result.Value = true
		} else {
			result.Value = false
		}
	} else if v.Method == "count" {
		result.Type = environment.INTEGER
		result.Value = len(result.Value.([]interface{}))
	} else if v.Method == "access" {
		if len(result.Value.([]interface{})) > 0 {
			pos := v.Expression.Ejecutar(ast, env).Value.(int)
			if len(result.Value.([]interface{})) > pos {
				for i, val := range result.Value.([]interface{}) {
					if pos == i {
						result.Value = val
						break
					}
				}
			} else {
				ast.SetError("Error Semantico: El indice \"" + v.Id + "\" esta fuera de rango")
			}
		} else {
			ast.SetError("Error Semantico: El vector \"" + v.Id + "\" esta vacio")
		}
	}
	return result
}
