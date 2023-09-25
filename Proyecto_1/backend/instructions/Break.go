package instructions

import (
	"Server/environment"
)

type Break struct {
	Line int
	Col  int
}

func NewBreak(line int, col int) Break {
	return Break{line, col}
}

func (b Break) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var result environment.Symbol

	if env.(environment.Env).IsLoop() || env.(environment.Env).IsSwitch() {
		result = environment.Symbol{Line: b.Line, Col: b.Col, Type: environment.NULL, Value: 0, ReturnFlag: true}
	} else {
		ast.AddError(b.Line, b.Col, env.(environment.Env).Id, "La sentencia break solo puede ser usada dentro de un ciclo o un switch")
	}
	return result
}
