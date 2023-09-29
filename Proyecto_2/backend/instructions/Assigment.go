package instructions

import (
	"Server/environment"
	"Server/expressions"
	"Server/interfaces"
	"fmt"
)

type ASSIGMENT struct {
	Line  int
	Col   int
	Id    string
	Value interface{}
}

type OPASSIGMENT struct {
	Line  int
	Col   int
	Id    string
	Value interface{}
	Op    string
}

func NewAsgmt(line int, col int, ide string, val interface{}) ASSIGMENT {
	return ASSIGMENT{Line: line, Col: col, Id: ide, Value: val}
}

func NewOpAsgmt(line int, col int, ide string, val interface{}, Op string) OPASSIGMENT {
	return OPASSIGMENT{Line: line, Col: col, Id: ide, Value: val, Op: Op}
}

func (a ASSIGMENT) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre ASSIGMENT")
	result := a.Value.(interfaces.Expression).Ejecutar(ast, env)
	if result.Type == environment.NULL {
		ast.AddError(a.Line, a.Col, env.(environment.Env).Id, "No se puede operar con un valor nulo")
		return result
	}
	result_env := env.(environment.Env).SetVar(a.Id, result)
	if result_env.Type == environment.NULL {
		if result_env.Const {
			ast.AddError(a.Line, a.Col, env.(environment.Env).Id, "La variable \""+a.Id+"\" es constante")

		} else {
			ast.AddError(a.Line, a.Col, env.(environment.Env).Id, "La variable \""+a.Id+"\" no existe")
		}
		return result
	}
	return result
}

func (o OPASSIGMENT) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre OPASSIGMENT")
	result := o.Value.(interfaces.Expression).Ejecutar(ast, env)
	if result.Type == environment.NULL {
		ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No se puede operar con un valor nulo")
		return result
	}
	// Obtengo el valor de la variable
	valVar := env.(environment.Env).GetVar(o.Id)
	if valVar.Type == environment.NULL {
		ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "La variable \""+o.Id+"\" no existe")
		return result
	}
	// Creo una nueva variable primitiva con el valor de la variable
	newVar := expressions.NewPrimitive(o.Line, o.Col, valVar.Value, valVar.Type)
	// Creo una operacion con el valor de la variable y el valor de la expresion
	tmpResult := expressions.NewOperation(o.Line, o.Col, newVar, o.Op, o.Value.(interfaces.Expression)).Ejecutar(ast, env)
	result.Value = tmpResult.Value
	// Guardo el resultado en la variable
	result_env := env.(environment.Env).SetVar(o.Id, result)
	if result_env.Type == environment.NULL {
		if result_env.Const {
			ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "La variable \""+o.Id+"\" es constante")
		} else {
			ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "La variable \""+o.Id+"\" no existe")
		}
		return result
	}
	return result
}
