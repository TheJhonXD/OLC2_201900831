package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"strconv"
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

func (a ASSIGMENT) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	gen.AddComment("Generando asignacion")
	//buscando variable en entorno
	variable := env.(environment.Env).GetVar(a.Id)
	//ejecutando valor
	result = a.Value.(interfaces.Expression).Ejecutar(ast, env, gen)
	//realizando asignacion
	gen.AddSetStack(strconv.Itoa(variable.Pos), result.Value)
	gen.AddBr()
	return result
}

func (o OPASSIGMENT) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
