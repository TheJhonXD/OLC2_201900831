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
	visited := make(map[string]bool)
	var function environment.SymbolFunc
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
	function = environment.SymbolFunc{Line: f.Line, Col: f.Col, Id: f.Id, Args: f.Args, RtnType: f.RtnType, Block: f.Block}
	aux, existsFunc := env.(environment.Env).GetFunc(f.Id)
	if existsFunc {
		if aux.RtnType == environment.NULL {
			env.(environment.Env).SaveFunc(f.Id, function)
		}
	} else {
		ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "La funcion \""+f.Id+"\" ya existe")
	}
	return function
}
