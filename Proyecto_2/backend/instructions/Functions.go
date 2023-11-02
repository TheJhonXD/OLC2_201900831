package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
	"strings"
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

func (f Function) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	tmpRetorno := gen.NewTmp()
	// env.(environment.Env).SaveVar(f.Id+"dx", f.RtnType, false)

	gen.SetMainFlag(false)
	gen.AddComment("******** Funcion " + f.Id)
	gen.AddTitle(f.Id)
	endLabel := gen.NewLabel()

	//entorno
	var envFunc environment.Env
	envFunc = environment.NewEnv(env.(environment.Env), f.Id)
	envFunc.Size["size"] = envFunc.Size["size"] + 1
	//variables
	for _, s := range f.Args {
		fmt.Println("Si guardo variables")
		res := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
		fmt.Println(res)
		envFunc.SaveVar(res.(environment.Value).Value, res.(environment.Value).Type, false)
	}
	//instrucciones func
	for _, s := range f.Block {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, envFunc, gen)
			if resInst != nil {
				//agregando etiquetas de salida
				for _, lvl := range resInst.(environment.Value).OutLabel {
					gen.AddLabel(lvl.(string))
				}
				if resInst.(environment.Value).ReturnFlag {
					gen.AddComment("Generacion return")
					gen.AddExpression(tmpRetorno, resInst.(environment.Value).Value, "0", "+")
					gen.AddKeyNValue(f.Id, tmpRetorno)

					/* nuevo := environment.NewValue(nuevoTmp, false, resInst.(environment.Value).Type)
					nuevo.ReturnFlag = true
					return resInst */
					// gen.AddOther("return " + resInst.(environment.Value).Value + ";")
				}
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Ejecutar(ast, envFunc, gen)
			//agregando etiquetas de salida
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
			if result.ReturnFlag {
				gen.AddComment("Generacion return")
				gen.AddExpression(tmpRetorno, result.Value, "0", "+")
				gen.AddKeyNValue(f.Id, tmpRetorno)
				/* nuevo := environment.NewValue(nuevoTmp, false, result.Type)
				nuevo.ReturnFlag = true
				return nuevo */
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}
	//etiqueta de salida
	gen.AddLabel(endLabel)
	gen.AddEnd()
	gen.SetMainFlag(true)
	return nil
}
