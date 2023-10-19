package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type Statement struct {
	Line  int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interface{}
	Const bool
}

func NewStmt(line int, col int, name string, tipo environment.TipoExpresion, value interface{}, constant bool) Statement {
	return Statement{line, col, name, tipo, value, constant}
}

func (v Statement) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	var newVar environment.Symbol
	result = v.Value.(interfaces.Expression).Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	fmt.Println("Agregando una declaracion")
	//Agregando el tipo de variable
	if v.Type == environment.NULL {
		v.Type = result.Type
	} else if v.Type != result.Type {
		ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "Los tipos no coinciden")
		return result
	}
	fmt.Println("VAR: ", result.Value)
	newVar = env.(environment.Env).SaveVar(v.Name, v.Type, v.Const)

	if result.Type == environment.BOOLEAN {
		//si no es temp (boolean)
		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Pos), "1")
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Pos), "0")
		gen.AddGoto(newLabel)
		gen.AddLabel(newLabel)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(newVar.Pos), result.Value)
		gen.AddBr()
	}

	return result
}
