package interfaces

import "Server/environment"

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}) interface{}
}
