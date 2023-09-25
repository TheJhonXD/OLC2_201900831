package instructions

import (
	"Server/environment"
	"Server/expressions"
	"Server/interfaces"
)

type VectorAsgmt struct {
	Line        int
	Col         int
	Id          string
	Position    interfaces.Expression
	Expression  interfaces.Expression
	Id2ndVector string
}

func NewVectorAsgmt(line int, col int, ide string, position interfaces.Expression, expression interfaces.Expression, idvector string) VectorAsgmt {
	return VectorAsgmt{Line: line, Col: col, Id: ide, Position: position, Expression: expression, Id2ndVector: idvector}
}

func (v VectorAsgmt) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	result := env.(environment.Env).GetVar(v.Id)
	if result.Type == environment.NULL {
		ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "La variable \""+v.Id+"\" no existe")
		return result
	}
	if v.Id2ndVector != "" {
		pos := v.Position.Ejecutar(ast, env)
		if pos.Type == environment.INTEGER {
			if len(result.Value.([]interface{})) > pos.Value.(int) {
				varAux := expressions.NewVector(v.Line, v.Col, v.Id2ndVector, v.Expression, "access")
				result2 := varAux.Ejecutar(ast, env)
				if result2.Type != environment.NULL {
					if result2.Type == result.Type {
						result.Value.([]interface{})[pos.Value.(int)] = result2.Value
						env.(environment.Env).SetVar(v.Id, result)
					} else {
						ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El tipo de dato \""+v.Id2ndVector+"\" no coincide con el tipo de dato del vector \""+v.Id+"\"")
					}
				}
			} else {
				ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El indice \""+v.Id+"\" esta fuera de rango")
			}
		} else {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El indice \""+v.Id+"\" debe ser de tipo INTEGER")
		}
	} else {
		pos := v.Position.Ejecutar(ast, env)
		if pos.Type == environment.INTEGER {
			if len(result.Value.([]interface{})) > pos.Value.(int) {
				result2 := v.Expression.Ejecutar(ast, env)
				if result2.Type != environment.NULL {
					if result2.Type == result.Type {
						result.Value.([]interface{})[pos.Value.(int)] = result2.Value
						env.(environment.Env).SetVar(v.Id, result)
					} else {
						ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El tipo de dato \""+v.Id2ndVector+"\" no coincide con el tipo de dato del vector \""+v.Id+"\"")
					}
				}
			} else {
				ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El indice \""+v.Id+"\" esta fuera de rango")
			}
		} else {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El indice \""+v.Id+"\" debe ser de tipo INTEGER")
		}
	}
	return result
}
