package instructions

import (
	"Server/environment"
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

func (v Statement) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	// Si no es nulo ejecuta la expresion
	fmt.Println(v.Value.(interfaces.Expression))
	result := v.Value.(interfaces.Expression).Ejecutar(ast, env)

	/* CASTEO A PATITA */
	if v.Type == environment.FLOAT && result.Type == environment.INTEGER { //Casteo a pata
		result.Value, _ = strconv.ParseFloat(fmt.Sprintf("%v", result.Value), 64)
		result.Type = environment.FLOAT
	} else if v.Type == environment.INTEGER && result.Type == environment.FLOAT { //Casteo a oata
		entero, _ := strconv.ParseInt(fmt.Sprintf("%v", result.Value), 10, 64)
		flotante, _ := strconv.ParseFloat(fmt.Sprintf("%v", result.Value), 64)

		if entero < int64(flotante) {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "No se puedo convertir de float a int")
			return result
		}
		result.Value, _ = strconv.ParseInt(fmt.Sprintf("%v", result.Value), 10, 64)
		// result.Value = int(result.Value.(float32))
		result.Type = environment.INTEGER
	}
	result.Const = v.Const
	// Si el tipo de la variable es igual al tipo de la expresion o si la variable es un wildcard se guarda
	if result.Type == v.Type {

		if !env.(environment.Env).SaveVar(v.Name, result) {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "La variable ya existe")
		} else {
			ast.AddSymbol(v.Line, v.Col, result.Type, "Variable", env.(environment.Env).Id, v.Name)
		}
	} else if result.Type == environment.WILDCARD {

		if !env.(environment.Env).SaveVar(v.Name, result) {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "La variable ya existe")
		} else {
			ast.AddSymbol(v.Line, v.Col, result.Type, "Variable", env.(environment.Env).Id, v.Name)
		}
	} else if v.Type == environment.NULL {
		v.Type = result.Type
		if !env.(environment.Env).SaveVar(v.Name, result) {
			ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "La variable ya existe")
		} else {

			ast.AddSymbol(v.Line, v.Col, result.Type, "Variable", env.(environment.Env).Id, v.Name)
		}
	} else {
		fmt.Println("Entre?")
		ast.AddError(v.Line, v.Col, env.(environment.Env).Id, "Los tipos no coinciden")
	}
	fmt.Println("AAAA")
	return result
}
