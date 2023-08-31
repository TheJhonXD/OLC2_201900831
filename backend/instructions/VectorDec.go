package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
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
		fmt.Println("VectorDec")
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
			} else {
				ast.SetError("Error Semantico: El vector \"" + v.Id + "\" ya existe")
			}
		} else {
			result.Type = environment.NULL
			ast.SetError("Error Semantico: Los valores del vector no coinciden con el tipo del vector")
		}

	}
	return result
}
