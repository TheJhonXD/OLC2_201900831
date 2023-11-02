package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type CallFunc struct {
	Line int
	Col  int
	Id   string
	Args []interface{}
}

func NewCallFunc(line int, col int, ide string, args []interface{}) CallFunc {
	return CallFunc{Line: line, Col: col, Id: ide, Args: args}
}

func (c CallFunc) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	size := env.(environment.Env).Size["size"]
	gen.AddComment("Inicio de llamada")
	if len(c.Args) > 0 {
		tmp1 := gen.NewTmp()
		gen.AddExpression(tmp1, "P", strconv.Itoa(size+1), "+")
		for i := 0; i < len(c.Args); i++ {
			//ejecutando parametros
			if strings.Contains(fmt.Sprintf("%T", c.Args[i]), "expressions") {
				result = c.Args[i].(interfaces.Expression).Ejecutar(ast, env, gen)
			} else {
				fmt.Println("Error en bloque")
			}
			//guardando
			gen.AddSetStack("(int)"+tmp1, result.Value)
			if (len(c.Args) - 1) != i {
				gen.AddExpression(tmp1, tmp1, "1", "+")
			}
		}

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(c.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	} else {
		fmt.Println("entro aqui")
		tmp1 := gen.NewTmp()

		gen.AddExpression("P", "P", strconv.Itoa(size), "+")
		gen.AddCall(c.Id)
		gen.AddGetStack(tmp1, "(int)P")
		gen.AddExpression("P", "P", strconv.Itoa(size), "-")

	}
	gen.AddComment("Final de llamada")
	result.Value = gen.GetKeyNValue(c.Id)
	return result
}
