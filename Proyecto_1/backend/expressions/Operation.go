package expressions

import (
	"Server/environment"
	"Server/interfaces"
	"fmt"
	"strconv"
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

func (o Operation) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
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

	var op1, op2 environment.Symbol
	op1 = o.Op_izq.Ejecutar(ast, env)
	op2 = o.Op_der.Ejecutar(ast, env)
	// fmt.Println("OP1: ", op1.Type, " OP2: ", op2.Type)
	switch o.Operator {
	case "+":
		{
			dominante = sum_table[op1.Type][op2.Type]

			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) + op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 + val2, Const: false}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: r1 + r2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible sumar")
			}
		}
	case "-":
		{
			dominante = sub_mult_div_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) - op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 - val2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible restar")
			}
		}
	case "*":
		{
			dominante = sub_mult_div_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) * op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 * val2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible multiplicar")
			}
		}
	case "/":
		{
			dominante = sub_mult_div_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				if op2.Value.(int) != 0 {
					return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) / op2.Value.(int), Const: false}
				} else {
					ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible dividir entre 0")
				}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				if val2 != 0 {
					return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 / val2, Const: false}
				} else {
					ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible dividir entre 0")
				}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible dividir")
			}
		}
	case "%":
		{
			dominante = mod_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				if op2.Value.(int) != 0 {
					return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) % op2.Value.(int), Const: false}
				} else {
					ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible aplicar modulo")
				}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible aplicar modulo")
			}
		}

	case "unario":
		{
			if op1.Type == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.INTEGER, Value: 0 - op1.Value.(int), Const: false}
			} else if op1.Type == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.FLOAT, Value: 0 - val1, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible aplicar menos unario")
			}
		}
	case "<":
		{
			dominante = rel_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) < op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 < val2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es menor")
			}
		}
	case ">":
		{
			dominante = rel_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) > op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 > val2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es mayor")
			}
		}
	case "<=":
		{
			dominante = rel_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) <= op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 <= val2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es menor o igual")
			}
		}
	case ">=":
		{
			dominante = rel_table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) >= op2.Value.(int), Const: false}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 >= val2, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar si es mayor o igual")
			}
		}
	case "==":
		{
			if op1.Type == op2.Type {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value == op2.Value, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar la igualdad")
			}
		}
	case "!=":
		{
			if op1.Type == op2.Type {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value != op2.Value, Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es operar la desigualdad")
			}
		}
	case "&&":
		{
			if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) && op2.Value.(bool), Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar")
			}
		}
	case "||":
		{
			if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) || op2.Value.(bool), Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible comparar")
			}
		}
	case "!":
		{
			if op1.Type == environment.BOOLEAN {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: !op1.Value.(bool), Const: false}
			} else {
				ast.AddError(o.Line, o.Col, env.(environment.Env).Id, "No es posible negar")
			}
		}
	}

	var result interface{}
	return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.NULL, Value: result}
}
