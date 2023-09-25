package instructions

import (
	"Server/environment"
	"Server/interfaces"
)

type VectorDec struct {
	Line    int
	Col     int
	Id      string
	Type    environment.TipoExpresion
	ExpList []interface{}
}

func NewVectorStmt(line int, col int, id string, tipo environment.TipoExpresion, expList []interface{}) VectorDec {
	return VectorDec{line, col, id, tipo, expList}
}

func (v VectorDec) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol
	var values []interface{}
	mismatchFlag := false
	if v.ExpList != nil {
		for _, exp := range v.ExpList {
			result = exp.(interfaces.Expression).Ejecutar(ast, env)
			if result.Type != v.Type {
				mismatchFlag = true
				break
			}
			values = append(values, result.Value)
		}
		result = environment.Symbol{Line: v.Line, Col: v.Col, Type: v.Type, Value: values, Const: false}

		if !mismatchFlag {
			resultEnv := env.(environment.Env).GetVar(v.Id)
			if resultEnv.Type == environment.NULL {
				env.(environment.Env).SaveVar(v.Id, result)
				ast.AddSymbol(v.Line, v.Col, v.Type, "Vector", env.(environment.Env).Id, v.Id)
			} else {
				ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El vector \""+v.Id+"\" ya existe")
			}
		} else {
			result.Type = environment.NULL
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "Los tipos de datos no coinciden")
		}

	} else {
		result = environment.Symbol{Line: v.Line, Col: v.Col, Type: v.Type, Value: values, Const: false}
		resultEnv := env.(environment.Env).GetVar(v.Id)
		if resultEnv.Type == environment.NULL {
			values = append(values, nil)
			env.(environment.Env).SaveVar(v.Id, result)
			ast.AddSymbol(v.Line, v.Col, v.Type, "Vector", env.(environment.Env).Id, v.Id)
		} else {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "El vector \""+v.Id+"\" ya existe")
		}
	}
	return result
}
