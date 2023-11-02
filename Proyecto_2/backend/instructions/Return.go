package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type Return struct {
	Line       int
	Col        int
	Expression interfaces.Expression
}

func NewReturn(line int, col int, exp interfaces.Expression) Return {
	return Return{line, col, exp}
}

func (r Return) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	fmt.Println("entro a return")
	var result environment.Value
	if r.Expression != nil {
		result = r.Expression.Ejecutar(ast, env, gen)
		fmt.Println("Valor de retorno Int: " + strconv.Itoa(result.IntValue))
		// result = environment.NewValue(result.Value, false, result.Type)
		fmt.Println("Valor de retorno: " + result.Value)
		result.ReturnFlag = true
		// return environment.Value{Value: result.Value, Type: result.Type}
	} else {
		result.ReturnFlag = true
	}
	return result
}
