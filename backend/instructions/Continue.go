package instructions

import "Server/environment"

type Continue struct {
	Line int
	Col  int
}

func NewContinue(line int, col int) Continue {
	return Continue{line, col}
}

func (c Continue) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol
	if env.(environment.Env).IsLoop() {
		result = environment.Symbol{Line: c.Line, Col: c.Col, Type: environment.NULL, Value: 0, ContinueFlag: true}
	}
	return result
}
