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
	result := v.Value.(interfaces.Expression).Ejecutar(ast, env)

	/* CASTEO A PATITA */
	if v.Type == environment.FLOAT && result.Type == environment.INTEGER { //Casteo a pata
		result.Value, _ = strconv.ParseFloat(fmt.Sprintf("%v", result.Value), 64)
		result.Type = environment.FLOAT
	} else if v.Type == environment.INTEGER && result.Type == environment.FLOAT { //Casteo a oata
		entero, _ := strconv.ParseInt(fmt.Sprintf("%v", result.Value), 10, 64)
		flotante, _ := strconv.ParseFloat(fmt.Sprintf("%v", result.Value), 64)

		if entero < int64(flotante) {
			fmt.Println("Error de casteo")
			return result
		}
		result.Value, _ = strconv.ParseInt(fmt.Sprintf("%v", result.Value), 10, 64)
		// result.Value = int(result.Value.(float32))
		result.Type = environment.INTEGER
	}

	// fmt.Println("RESULTADO: ", result.Type)
	// fmt.Println("VARIABLE: ", v.Type)
	result.Const = v.Const

	// Si el tipo de la variable es igual al tipo de la expresion o si la variable es un wildcard se guarda
	if result.Type == v.Type {
		env.(environment.Env).SaveVar(v.Name, result)
	} else if result.Type == environment.WILDCARD {
		env.(environment.Env).SaveVar(v.Name, result)
	} else if v.Type == environment.NULL {
		v.Type = result.Type
		env.(environment.Env).SaveVar(v.Name, result)
	} else {
		fmt.Println("Los tipos no coinciden")
	}

	return result
}
