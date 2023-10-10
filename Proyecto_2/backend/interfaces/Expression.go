package interfaces

import (
	"Server/environment"
	"Server/generator"
)

type Expression interface {
	Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value
}
