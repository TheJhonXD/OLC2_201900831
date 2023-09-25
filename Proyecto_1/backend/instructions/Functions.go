package instructions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
)

type Function struct {
	Line    int
	Col     int
	Id      string
	Args    []interface{}
	RtnType environment.TipoExpresion
	Block   []interface{}
}

func NewFunction(line int, col int, ide string, args []interface{}, rtnType environment.TipoExpresion, block []interface{}) Function {
	return Function{Line: line, Col: col, Id: ide, Args: args, RtnType: rtnType, Block: block}
}

func (f Function) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("Entre Function")
	if env.(environment.Env).Id != "GLOBAL" {
		ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "Las funciones solo pueden ser declaradas en el ambito global")
		return environment.SymbolFunc{}
	}
	visited := make(map[string]bool)
	var function environment.SymbolFunc
	// Comprueba que no existan parametros repetidos
	for _, p := range f.Args {
		param := p.(interfaces.Instruction).Ejecutar(ast, env).(environment.SymbolFuncParam)
		if param.ExternalId != "_" {
			if visited[param.ExternalId] {
				ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "Parametro externo repetido")
				return nil
			}
			visited[param.ExternalId] = true
		}

		if param.InternalId != "_" {
			if visited[param.InternalId] {
				ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "Parametro interno repetido")
				return nil
			}
			visited[param.InternalId] = true
		}
	}
	// Crea un nuevo simbolo de funcion
	function = environment.SymbolFunc{Line: f.Line, Col: f.Col, Id: f.Id, Args: f.Args, RtnType: f.RtnType, Block: f.Block}
	// Comprueba que la funcion no exista y la guarda
	_, existsFunc := env.(environment.Env).GetFunc(f.Id)
	if !existsFunc {
		env.(environment.Env).SaveFunc(f.Id, function)
		ast.AddSymbol(f.Line, f.Col, f.RtnType, "Funcion", env.(environment.Env).Id, f.Id)
	} else {
		ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "La funcion \""+f.Id+"\" ya existe")
	}
	return function
}
