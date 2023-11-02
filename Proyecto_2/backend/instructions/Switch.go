package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
)

type Switch struct {
	Line       int
	Col        int
	Expression interfaces.Expression
	Cases      []interface{}
	Default    []interface{}
}

func NewSwitch(line int, col int, expression interfaces.Expression, cases []interface{}, def []interface{}) Switch {
	return Switch{line, col, expression, cases, def}
}

func (s Switch) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generacion Switch")
	fmt.Println("Generacion Switch")
	var condition environment.Value
	var result environment.Value

	condition = s.Expression.Ejecutar(ast, env, gen)
	var caseEnv environment.Env
	caseEnv = environment.NewEnv(env.(environment.Env), "SWITCH")

	for _, c := range s.Cases {
		case_var := c.(Switch).Expression.Ejecutar(ast, caseEnv, gen)
		falseLabel := gen.NewLabel()

		gen.AddIf(condition.Value, case_var.Value, "!=", falseLabel)
		for _, ins := range c.(Switch).Cases {
			result = ins.(interfaces.Instruction).Ejecutar(ast, caseEnv, gen).(environment.Value)
		}
		gen.AddLabel(falseLabel)

	}

	if s.Default != nil {
		for _, ins := range s.Default {
			result = ins.(interfaces.Instruction).Ejecutar(ast, caseEnv, gen).(environment.Value)
		}
	}

	return result
}
