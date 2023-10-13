package expressions

import (
	"Server/environment"
	"Server/generator"
	"fmt"
	"strconv"
)

type Variable struct {
	Line int
	Col  int
	Id   string
}

func NewVar(line int, col int, ide string) Variable {
	return Variable{Line: line, Col: col, Id: ide}
}

func (v Variable) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("Llamando una variable")
	fmt.Println("Llamando variable")
	var result environment.Value
	retSym := env.(environment.Env).GetVar(v.Id)

	if retSym.Type == environment.NULL {
		ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "La variable \""+v.Id+"\" no existe")
		return result
	}
	newTemp := gen.NewTmp()
	newTemp2 := gen.NewTmp()

	if gen.MainCode {
		gen.AddGetStack(newTemp2, strconv.Itoa(retSym.Pos))
	} else {
		gen.AddExpression(newTemp, "P", strconv.Itoa(retSym.Pos), "+")
		gen.AddGetStack(newTemp2, "(int)"+newTemp)
	}

	if retSym.Type == environment.BOOLEAN {
		trueLabel := gen.NewLabel()
		falseLabel := gen.NewLabel()
		gen.AddIf(newTemp2, "1", "==", trueLabel)
		gen.AddGoto(falseLabel)
		result = environment.NewValue("", false, environment.BOOLEAN)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	} else {
		result = environment.NewValue(newTemp2, true, retSym.Type)
	}
	return result
}
