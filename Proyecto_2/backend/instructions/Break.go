package instructions

import (
	"Server/environment"
	"Server/generator"
)

type Break struct {
	Line int
	Col  int
}

func NewBreak(line int, col int) Break {
	return Break{line, col}
}

func (b Break) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	result.BreakFlag = true
	return result
}
