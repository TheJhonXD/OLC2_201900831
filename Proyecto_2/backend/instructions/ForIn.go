package instructions

import (
	"Server/environment"
	"Server/expressions"
	"Server/generator"
	"Server/interfaces"
	"fmt"
	"strconv"
)

type ForIn struct {
	Line       int
	Col        int
	Id         string
	Expression interfaces.Expression
	Op_left    interfaces.Expression
	Op_right   interfaces.Expression
	VecId      string
	Block      []interface{}
}

func NewForIn(line int, col int, ide string, expression interfaces.Expression, left interfaces.Expression, right interfaces.Expression, vecId string, block []interface{}) ForIn {
	return ForIn{line, col, ide, expression, left, right, vecId, block}
}

func (f ForIn) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generacion For In")
	fmt.Println("Generacion For In")
	var condition environment.Value
	var result environment.Value
	// var flag bool

	if f.Expression != nil {
		init := gen.NewLabel()
		gen.AddLabel(init)
		condition = f.Expression.Ejecutar(ast, env, gen)
		if condition.Type == environment.STRING {
			var forEnv environment.Env
			forEnv = environment.NewEnv(env.(environment.Env), "FOR IN")
			var_exists := env.(environment.Env).GetVar(f.Id)
			if var_exists.Type == environment.NULL {
				var newValue Statement
				newValue = NewStmt(f.Line, f.Col, f.Id, environment.STRING, "", false)
				newValue.Ejecutar(ast, forEnv, gen)

				// newOperation := expressions.NewOperation(f.Line, f.Col, expressions.NewPrimitive(f.Line, f.Col, f.Id, 2), ">", expressions.NewPrimitive(f.Line, f.Col, f.Id, 2))
				newTemp1 := gen.NewTmp()
				newTemp2 := gen.NewTmp()
				size := strconv.Itoa(env.(environment.Env).Size["size"])
				gen.AddExpression(newTemp1, "P", size, "+")     //nuevo temporal en pos vacia
				gen.AddExpression(newTemp1, newTemp1, "1", "+") //se deja espacio de retorno
				//! Result es diferente
				gen.AddSetStack("(int)"+newTemp1, result.Value) //se coloca string en parametro que se manda
				gen.AddExpression("P", "P", size, "+")          // cambio de entorno
				gen.AddCall("dbrust_assign")                    //Llamada
				gen.AddGetStack(newTemp2, "(int)P")             //obtencion retorno
				gen.AddExpression("P", "P", size, "-")          //regreso del entorno
				newTmp := gen.NewTmp()
				lbl1 := gen.NewLabel()
				lbl2 := gen.NewLabel()
				gen.AddIf(newTmp, strconv.Itoa(len(condition.Value)), ">", lbl1)
				gen.AddGoto(lbl2)
				/* newAsgmt := NewAsgmt(f.Line, f.Col, f.Id, string(values))
				newAsgmt.Ejecutar(ast, forEnv, gen) */
				gen.AddExpression(newTmp, newTmp, "1", "+")
				for _, instr := range f.Block {
					result = instr.(interfaces.Instruction).Ejecutar(ast, forEnv, gen).(environment.Value)

				}

				gen.AddLabel(lbl2)
			}
		}
	} else if f.Op_left != nil && f.Op_right != nil {
		gen.AddComment("For In con rango")
		fmt.Println("For In con rango")
		var left, right environment.Value
		left = f.Op_left.Ejecutar(ast, env, gen)
		right = f.Op_right.Ejecutar(ast, env, gen)

		// Comprobando que sean enteros
		if left.Type == environment.INTEGER && right.Type == environment.INTEGER {
			var forEnv environment.Env
			forEnv = environment.NewEnv(env.(environment.Env), "FOR IN")
			var_exists := env.(environment.Env).GetVar(f.Id)
			// Comprobando que la variable no exista en el entorno
			if var_exists.Type == environment.NULL {
				// Creando nueva variable
				var newStmt Statement
				newStmt = NewStmt(f.Line, f.Col, f.Id, environment.INTEGER, expressions.NewPrimitive(f.Line, f.Col, 0, environment.INTEGER), false)
				newStmt.Ejecutar(ast, forEnv, gen)

				var newOp expressions.Operation
				var newOp2 expressions.Operation
				// Nueva operacion para aumentar en uno el valor de la variable
				newOp = expressions.NewOperation(f.Line, f.Col, expressions.NewVar(f.Line, f.Col, f.Id), "+", expressions.NewPrimitive(f.Line, f.Col, 1, environment.INTEGER))
				// Nueva operacion para inicializar la variable
				newOp2 = expressions.NewOperation(f.Line, f.Col, expressions.NewVar(f.Line, f.Col, f.Id), "+", expressions.NewPrimitive(f.Line, f.Col, left.IntValue, environment.INTEGER))

				var newAsgmt ASSIGMENT
				// Asignando el valor inicial (left)
				newAsgmt = NewAsgmt(f.Line, f.Col, f.Id, newOp2)
				newAsgmt.Ejecutar(ast, forEnv, gen)

				/* ****** FOR START ****** */
				init := gen.NewLabel()
				gen.AddLabel(init)
				// Nueva operacion para comprobar si la variable es menor o igual que el valor final (right)
				var newCondition expressions.Operation
				newCondition = expressions.NewOperation(f.Line, f.Col, expressions.NewVar(f.Line, f.Col, f.Id), "<=", expressions.NewPrimitive(f.Line, f.Col, right.IntValue, environment.INTEGER))
				condition = newCondition.Ejecutar(ast, forEnv, gen)
				// Etiqueta para ejecutar si se cumple la condicion
				gen.AddLabel(condition.TrueLabel[0].(string))
				// Ejecutando bloque
				for _, instr := range f.Block {
					result = instr.(interfaces.Instruction).Ejecutar(ast, forEnv, gen).(environment.Value)
				}
				// Aumentando en uno el valor de la variable
				newAsgmt = NewAsgmt(f.Line, f.Col, f.Id, newOp)
				newAsgmt.Ejecutar(ast, forEnv, gen)
				// Volviendo a la etiqueta de inicio del for
				gen.AddGoto(init)
				// Etiqueta para salir del for
				gen.AddLabel(condition.FalseLabel[0].(string))
			} else {
				ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "La variable \""+f.Id+"\" ya existe en el ambito padre")
			}
		} else {
			ast.AddError(f.Line, f.Col, env.(environment.Env).Id, "Los tipos de las expresiones no son enteros")
		}
	}

	return result
}
