package instructions

import (
	"Server/environment"
	"Server/interfaces"
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

func (s Switch) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	switch_var := s.Expression.Ejecutar(ast, env)
	var result environment.Symbol
	if switch_var.Type != environment.NULL {
		for _, case_ := range s.Cases {
			case_var := case_.(Switch).Expression.Ejecutar(ast, env)
			if case_var.Value == switch_var.Value {
				var caseEnv environment.Env
				caseEnv = environment.NewEnv(env.(environment.Env), "SWITCH")
				for _, instr := range case_.(Switch).Cases {
					result = instr.(interfaces.Instruction).Ejecutar(ast, caseEnv).(environment.Symbol)
					if result.ContinueFlag {
						ast.AddError(s.Line, s.Col, caseEnv.Id, "No se puede usar continue en un switch")
					}
					if result.ReturnFlag {
						// result.ReturnFlag = false
						return result
					}
				}
				return result
			}
		}
		if s.Default != nil {
			for _, instr := range s.Default {
				result = instr.(interfaces.Instruction).Ejecutar(ast, env).(environment.Symbol)
				if result.ContinueFlag {
					ast.AddError(s.Line, s.Col, env.(environment.Env).Id, "No se puede usar continue en un switch")
				}
				if result.ReturnFlag {
					// result.ReturnFlag = false
					return result
				}
			}
		}
	} else {
		ast.AddError(s.Line, s.Col, env.(environment.Env).Id, "No se puede evaluar la expresion del switch")
	}
	return nil
}
