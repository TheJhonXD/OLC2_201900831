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

	table := [5][5]environment.TipoExpresion{
		/* 			INTEGER				FLOAT				STRING 				BOOLEAN				 */
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL}, // FLOAT
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},      // STRING
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},  // BOOLEAN
		{environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL},  // NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	var op1, op2 environment.Symbol
	op1 = o.Op_izq.Ejecutar(ast, env)
	op2 = o.Op_der.Ejecutar(ast, env)

	switch o.Operator {
	case "+":
		{
			dominante = table[op1.Type][op2.Type]

			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) + op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 + val2}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: r1 + r2}
			} else {
				ast.SetError("ERROR: No es posible sumar")
			}
		}
	case "-":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) - op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 - val2}
			} else {
				ast.SetError("ERROR: No es posible restar")
			}
		}
	case "*":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) * op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 * val2}
			} else {
				ast.SetError("ERROR: No es posible multiplicar")
			}
		}
	case "/":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: op1.Value.(int) / op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: dominante, Value: val1 / val2}
			} else {
				ast.SetError("ERROR: No es posible dividir")
			}
		}
	case "<":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) < op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 < val2}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case ">":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) > op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 > val2}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case "<=":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) <= op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 <= val2}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case ">=":
		{
			dominante = table[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) >= op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: val1 >= val2}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case "==":
		{
			if op1.Type == op2.Type {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value == op2.Value}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case "!=":
		{
			if op1.Type == op2.Type {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value != op2.Value}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case "&&":
		{
			if op1.Type == environment.BOOLEAN && op2.Type == environment.BOOLEAN {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) && op2.Value.(bool)}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
	case "||":
		{
			if op1.Type == environment.BOOLEAN && op2.Type == environment.BOOLEAN {
				return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) || op2.Value.(bool)}
			} else {
				ast.SetError("ERROR: No es posible comparar")
			}
		}
		//!Podría llevar un default
	}
	var result interface{}
	return environment.Symbol{Line: o.Line, Col: o.Col, Type: environment.NULL, Value: result}
}
