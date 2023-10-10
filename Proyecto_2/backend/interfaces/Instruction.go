package interfaces

import (
	"Server/environment"
	"Server/generator"
)

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{}
}
