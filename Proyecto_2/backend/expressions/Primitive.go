package expressions

import (
	"Server/environment"
	"Server/generator"
	"fmt"
	"strconv"
)

type Primitive struct {
	Line int
	Col  int
	Val  interface{}
	Type environment.TipoExpresion
}

func NewPrimitive(line int, col int, val interface{}, tipo environment.TipoExpresion) Primitive {
	return Primitive{Line: line, Col: col, Val: val, Type: tipo}
}

func (p Primitive) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	fmt.Println("Primitivo")
	var result environment.Value
	if p.Type == environment.INTEGER {
		result = environment.NewValue(fmt.Sprintf("%v", p.Val), false, p.Type)
		result.IntValue = p.Val.(int)
	} else if p.Type == environment.FLOAT {
		result = environment.NewValue(fmt.Sprintf("%v", p.Val), false, p.Type)
		result.FloatValue = p.Val.(float64)
	} else if p.Type == environment.STRING {
		newTmp := gen.NewTmp()
		gen.AddAssign(newTmp, "H")
		myString := p.Val.(string)

		/* Recorrer el strin en ascii */
		byteArray := []byte(myString)
		for _, asc := range byteArray {
			// Se agrega ascii al heap
			gen.AddSetHeap("(int)H", strconv.Itoa(int(asc)))
			// Suma 1 al heap pointer
			gen.AddExpression("H", "H", "1", "+")
		}
		// Caracteres de escape
		gen.AddSetHeap("(int)H", "-1")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddBr()
		result = environment.NewValue(newTmp, true, p.Type)
	} else if p.Type == environment.BOOLEAN {
		gen.AddComment("Primitivo bool")
		trueLabel := gen.NewLabel()
		falseLabel := gen.NewLabel()
		if p.Val.(bool) {
			gen.AddGoto(trueLabel)
		} else {
			gen.AddGoto(falseLabel)
		}
		result = environment.NewValue("", false, environment.BOOLEAN)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)
	} else if p.Type == environment.NULL {
		// Imprimir nil
		gen.AddPrintf("c", "(char)110")
		gen.AddPrintf("c", "(char)105")
		gen.AddPrintf("c", "(char)108")
		gen.AddPrintf("c", "(char)10")
		gen.AddBr()
		result = environment.NewValue("", false, environment.NULL)
	}

	return result
}
