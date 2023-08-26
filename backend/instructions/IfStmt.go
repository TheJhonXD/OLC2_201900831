package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type If struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Block      []interface{}
	ElseIf     []interface{}
	ElseInstr  []interface{}
}

func NewIf(line int, col int, expression interfaces.Expression, block []interface{}, elseif []interface{}, elseinstr []interface{}) If {
	return If{line, col, expression, block, elseif, elseinstr}
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
			fmt.Printf("%T\n", inst)
			inst.(interfaces.Instruction).Ejecutar(ast, ifEnv)
		}
		return nil
	} else if i.ElseIf != nil {
		fmt.Println("ELSE IF!!!!!!!!!!!!")
		/* fmt.Printf("%T\n", i.ElseIf)
		fmt.Printf("%T\n", i.Block) */
		for _, elseIfInstr := range i.ElseIf {
			// fmt.Println("ELSE IF")
			// fmt.Printf("%T\n", elseIfInstr)
			var condition_elif environment.Symbol
			condition_elif = elseIfInstr.(If).Expression.Ejecutar(ast, env)
			// fmt.Println(condition_elif.Value)
			if condition_elif.Value == true {
				// fmt.Println("SI")
				var elseIfEnv environment.Env
				elseIfEnv = environment.NewEnv(env.(environment.Env), "ELSE IF")
				// fmt.Println(elseIfInstr.(If).ElseIf)
				for _, instr := range elseIfInstr.(If).ElseIf {
					instr.(interfaces.Instruction).Ejecutar(ast, elseIfEnv)
				}
				return nil
			}
		}
		if i.ElseInstr != nil {
			fmt.Println("ELSE!!!!!!!!!!!!")
			var elseEnv environment.Env
			elseEnv = environment.NewEnv(env.(environment.Env), "ELSE")
			for _, instr := range i.ElseInstr {
				instr.(interfaces.Instruction).Ejecutar(ast, elseEnv)
			}
		}
	} else {
		var elseEnv environment.Env
		elseEnv = environment.NewEnv(env.(environment.Env), "ELSE")
		for _, inst := range i.ElseInstr {
			inst.(interfaces.Instruction).Ejecutar(ast, elseEnv)
		}
	}

	return nil
}
