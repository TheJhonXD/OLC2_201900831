package instructions

import (
	"Server/environment"
	"Server/generator"
)

type Continue struct {
	Line int
	Col  int
}

func NewContinue(line int, col int) Continue {
	return Continue{line, col}
}

func (c Continue) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	result.ContinueFlag = true
	return result
}
