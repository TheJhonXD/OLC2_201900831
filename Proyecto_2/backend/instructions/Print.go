package instructions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type Print struct {
	Line  int
	Col   int
	Value []interface{}
}

func NewPrint(line int, col int, val []interface{}) Print {
	return Print{line, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	for _, val := range p.Value {
		result = val.(interfaces.Expression).Ejecutar(ast, env, gen)
		if result.Type == environment.INTEGER {
			gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value))
			// gen.AddPrintf("c", "10")
			gen.AddBr()

		} else if result.Type == environment.FLOAT {
			gen.AddPrintf(".2f", "(float)"+fmt.Sprintf("%v", result.Value)+"")
			// gen.AddPrintf("c", "10")
			gen.AddBr()
		} else if result.Type == environment.BOOLEAN {
			if result.IsTmp {
				// cuando es variable
			}
			newLabel := gen.NewLabel()
			for _, lbl := range result.TrueLabel {
				gen.AddLabel(lbl.(string))
			}
			gen.AddPrintf("c", "(char)116")
			gen.AddPrintf("c", "(char)114")
			gen.AddPrintf("c", "(char)117")
			gen.AddPrintf("c", "(char)101")
			gen.AddGoto(newLabel)
			// Add labels
			for _, lbl := range result.FalseLabel {
				gen.AddLabel(lbl.(string))
			}
			gen.AddPrintf("c", "(char)102")
			gen.AddPrintf("c", "(char)97")
			gen.AddPrintf("c", "(char)108")
			gen.AddPrintf("c", "(char)115")
			gen.AddPrintf("c", "(char)101")
			gen.AddLabel(newLabel)
			// gen.AddPrintf("c", "10")
			gen.AddBr()
		} else if result.Type == environment.STRING {
			gen.AddComment("Print string")
			fmt.Println("Imprimir string")
			//agregar codigo en el main
			newTemp1 := gen.NewTmp()
			newTemp2 := gen.NewTmp()
			size := strconv.Itoa(env.(environment.Env).Size["size"])
			gen.AddExpression(newTemp1, "P", size, "+")     //nuevo temporal en pos vacia
			gen.AddExpression(newTemp1, newTemp1, "1", "+") //se deja espacio de retorno
			gen.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
			// fmt.Println("Result: ", result.Value)
			gen.AddExpression("P", "P", size, "+") // cambio de entorno
			gen.AddCall("dbrust_printString")      //Llamada
			gen.AddGetStack(newTemp2, "(int)P")    //obtencion retorno
			gen.AddExpression("P", "P", size, "-") //regreso del entorno
			// gen.AddPrintf("c", "10")                        //salto de linea
			gen.AddBr()
		}
		// Añado un espacio entre cada parametro a imprimir
		gen.AddPrintf("c", "32")
	}
	// Añado un salto de linea al final del print
	gen.AddPrintf("c", "10")
	return result
}
