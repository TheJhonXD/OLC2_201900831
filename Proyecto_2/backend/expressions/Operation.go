package expressions

import (
	"Server/environment"
	"Server/generator"
	"Server/interfaces"
	"fmt"
)

type Operation struct {
	Line     int
	Col      int
	Operator string
	Op_izq   interfaces.Expression
	Op_der   interfaces.Expression
}

func NewOperation(line int, col int, op1 interfaces.Expression, operator string, op2 interfaces.Expression) Operation {
	exp := Operation{Line: line, Col: col, Operator: operator, Op_izq: op1, Op_der: op2}
	return exp
}

func (o Operation) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var dominante environment.TipoExpresion

	/* table := [5][5]environment.TipoExpresion{
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL},
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		{environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	} */

	sum_table := [6][6]environment.TipoExpresion{
		/* 			INTEGER				FLOAT				STRING 				BOOLEAN			NULL 				CHAR		 */
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL}, // INTEGER
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},   // FLOAT
		{environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL, environment.STRING}, // STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},     // BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},     // NULL
		{environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL, environment.STRING}, // CHAR
	}

	sub_mult_div_table := [6][6]environment.TipoExpresion{
		/* 			INTEGER				FLOAT				STRING 				BOOLEAN			NULL 				CHAR		 */
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL}, // INTEGER
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},   // FLOAT
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},     // STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},     // BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},     // NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},     // CHAR
	}

	mod_table := [6][6]environment.TipoExpresion{
		/* 			INTEGER				FLOAT				STRING 				BOOLEAN			NULL 				CHAR		 */
		{environment.INTEGER, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL}, // INTEGER
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // FLOAT
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // CHAR
	}

	rel_table := [6][6]environment.TipoExpresion{
		/* 			INTEGER				FLOAT				STRING 				BOOLEAN			NULL 				CHAR		 */
		{environment.INTEGER, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL}, // INTEGER
		{environment.NULL, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},   // FLOAT
		{environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL, environment.NULL},  // STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},    // NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.CHAR},    // CHAR
	}

	var op1, op2, result environment.Value

	newTmp := gen.NewTmp()
	// fmt.Println("OP1: ", op1.Type, " OP2: ", op2.Type)
	switch o.Operator {
	case "+":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = sum_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTmp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTmp, true, dominante)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.STRING {
				fmt.Println("Concatenando")
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible sumar")
			}
		}
	case "-":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = sub_mult_div_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTmp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTmp, true, dominante)
				result.IntValue = op1.IntValue - op2.IntValue
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible restar")
			}
		}
	case "*":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = sub_mult_div_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTmp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTmp, true, dominante)
				result.IntValue = op1.IntValue * op2.IntValue
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible multiplicar")
			}
		}
	case "/":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = sub_mult_div_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTmp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTmp, op1.Value, op2.Value, "/")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTmp, true, dominante)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible dividir")
			}
		}
	case "%":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = mod_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTmp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTmp, op1.Value, op2.Value, "/")
				newTmp2 := gen.NewTmp()
				gen.AddExpression(newTmp2, "(int)"+newTmp, op2.Value, "*")
				newTmp3 := gen.NewTmp()
				gen.AddExpression(newTmp3, op1.Value, newTmp2, "-")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTmp3, true, dominante)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible aplicar modulo")
			}
		}

	case "unario":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			if op1.Type == environment.INTEGER || op1.Type == environment.FLOAT {
				gen.AddExpression(newTmp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTmp, true, dominante)
				result.IntValue = 0 - op1.IntValue
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible aplicar menos unario")
			}
		}
	case "<":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = rel_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es menor")
			}
		}
	case ">":
		{
			fmt.Println("jsjsjs")
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = rel_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es mayor")
			}
		}
	case "<=":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = rel_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es menor o igual")
			}
		}
	case ">=":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = rel_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es mayor o igual")
			}
		}
	case "==":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = rel_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "==", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es igual")
			}
		}
	case "!=":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			dominante = rel_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "!=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es desigual")
			}
		}
	case "&&":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			//add op1 labels
			for _, lvl := range op1.TrueLabel {
				gen.AddLabel(lvl.(string))
			}
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			result = environment.NewValue("", false, environment.BOOLEAN)
			fmt.Println("op1: ", op1.FalseLabel, " op2: ", op2.FalseLabel)
			result.TrueLabel = append(result.TrueLabel, op2.TrueLabel...)
			result.FalseLabel = append(result.FalseLabel, op1.FalseLabel...)
			result.FalseLabel = append(result.FalseLabel, op2.FalseLabel...)
			return result
		}
	case "||":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			//add op1 labels
			for _, lvl := range op1.FalseLabel {
				gen.AddLabel(lvl.(string))
			}
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			result = environment.NewValue("", false, environment.BOOLEAN)
			result.TrueLabel = append(result.TrueLabel, op1.TrueLabel...)
			result.TrueLabel = append(result.TrueLabel, op2.TrueLabel...)
			result.FalseLabel = append(result.FalseLabel, op2.FalseLabel...)
			return result
		}
	case "!":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			if op1.Type == environment.BOOLEAN {
				result = environment.NewValue("", false, environment.BOOLEAN)
				result.TrueLabel = append(result.TrueLabel, op1.FalseLabel...)
				result.FalseLabel = append(result.FalseLabel, op1.TrueLabel...)
				return result
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible aplicar negacion logica")
			}
		}
	}

	gen.AddBr()
	return environment.Value{}
}
