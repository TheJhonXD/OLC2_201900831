package environment

import "fmt"

type Env struct {
	SymbolTable map[string]Symbol
	Prev        interface{}
	Id          string
}

func NewEnv(prev interface{}, ide string) Env {
	return Env{SymbolTable: make(map[string]Symbol), Prev: prev, Id: ide}
}

func (env Env) SaveVar(id string, value Symbol) {
	if variable, ok := env.SymbolTable[id]; ok {
		fmt.Println("Error: La variable \""+id+"\" ya existe, con valor: ", variable.Value.(int))
		return
	}
	env.SymbolTable[id] = value
}

func (env Env) GetVar(id string) Symbol {
	var tmpEnv Env
	tmpEnv = env

	for {
		if variable, ok := tmpEnv.SymbolTable[id]; ok {
			return variable
		}
		if tmpEnv.Prev == nil {
			break
		} else {
			tmpEnv = tmpEnv.Prev.(Env)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Line: 0, Col: 0, Type: NULL, Value: 0, Const: false}
}

func (env Env) SetVar(id string, value Symbol) Symbol {
	var tmpEnv Env
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.SymbolTable[id]; ok {
			if !tmpEnv.SymbolTable[id].Const {
				if tmpEnv.SymbolTable[id].Type == value.Type {
					tmpEnv.SymbolTable[id] = value
					return variable
				} else {
					fmt.Println("Error: Los tipos no coinciden")
					return Symbol{Line: 0, Col: 0, Type: NULL, Value: 0}
				}
			} else {
				fmt.Println("Error: La variable \"" + id + "\" es constante")
				return Symbol{Line: 0, Col: 0, Type: NULL, Value: 0, Const: true}
			}
		}
		if tmpEnv.Prev == nil {
			break
		} else {
			tmpEnv = tmpEnv.Prev.(Env)
		}
	}
	fmt.Println("La variable \"", id, "\" no existe")
	return Symbol{Line: 0, Col: 0, Type: NULL, Value: 0}
}

/* func (env *Env) Search(nombre string) (SymbolVar, bool) {
	sym, ok := env.SymbolTable[nombre]
	if !ok && env.Prev != nil {
		return env.Prev.Search(nombre)
	}
	return sym, ok
}
*/
