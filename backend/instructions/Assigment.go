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
	env.(environment.Env).SetVar(a.Id, result)
	return result
}

func (o OPASSIGMENT) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre OPASSIGMENT")
	result := o.Value.(interfaces.Expression).Ejecutar(ast, env)
	// Obtengo el valor de la variable
	valVar := env.(environment.Env).GetVar(o.Id)
	// Creo una nueva variable primitiva con el valor de la variable
	newVar := expressions.NewPrimitive(o.Line, o.Col, valVar.Value, valVar.Type)
	// Creo una operacion con el valor de la variable y el valor de la expresion
	tmpResult := expressions.NewOperation(o.Line, o.Col, newVar, o.Op, o.Value.(interfaces.Expression)).Ejecutar(ast, env)
	result.Value = tmpResult.Value
	// Guardo el resultado en la variable
	env.(environment.Env).SetVar(o.Id, result)
	return result
}
