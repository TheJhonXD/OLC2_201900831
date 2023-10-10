package instructions

import (
	"Server/environment"
	"Server/generator"
)

type Statement struct {
	Line  int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interface{}
	Const bool
}

func NewStmt(line int, col int, name string, tipo environment.TipoExpresion, value interface{}, constant bool) Statement {
	return Statement{line, col, name, tipo, value, constant}
}

func (v Statement) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return environment.Value{}
}
