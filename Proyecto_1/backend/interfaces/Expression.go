package interfaces

import "Server/environment"

type Expression interface {
	Ejecutar(ast *environment.AST, env interface{}) environment.Symbol
}
