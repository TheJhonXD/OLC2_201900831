package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
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
	fmt.Println("Generando asignacion")
	// Buscando variable en entorno
	variable := env.(environment.Env).GetVar(a.Id)

	// Ejecutando valor
	result = a.Value.(interfaces.Expression).Ejecutar(ast, env, gen)
	if result.Type == variable.Type {
		// Comprobando si es constante
		if !variable.Const {
			// Realizando asignacion
			fmt.Println("Entro aqui")
			gen.AddSetStack(strconv.Itoa(variable.Pos), result.Value)
			gen.AddBr()
		} else {
			ast.AddError(a.Line, a.Col, env.(environment.Env).Id, "La variable es constante")
		}
	} else {
		ast.AddError(a.Line, a.Col, env.(environment.Env).Id, "Los tipos no coinciden")
	}
	return result
}

func (o OPASSIGMENT) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
