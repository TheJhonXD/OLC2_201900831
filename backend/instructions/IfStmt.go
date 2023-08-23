package instructions

import (
	"Server/environment"
	"Server/interfaces"
)

type If struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Block      []interface{}
}

func NewIf(line int, col int, expression interfaces.Expression, block []interface{}) If {
	return If{line, col, expression, block}
}

func (i If) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	condition := i.Expression.Ejecutar(ast, env)
	if condition.Type != environment.BOOLEAN {
		ast.SetError("Error de tipos en la condicion del if")
		return nil
	}

	if condition.Value == true {
		var ifEnv environment.Env
		ifEnv = environment.NewEnv(env.(environment.Env), "IF")
		for _, inst := range i.Block {
			inst.(interfaces.Instruction).Ejecutar(ast, ifEnv)
		}
		return nil
	}

	return nil
}
