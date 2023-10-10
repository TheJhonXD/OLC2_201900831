package environment

import (
	"fmt"
	"strings"
)

type Env struct {
	SymbolTable map[string]Symbol
	FuncTable   map[string]SymbolFunc
	Prev        interface{}
	Id          string
	Size        map[string]int
}

func NewEnv(prev interface{}, ide string) Env {
	return Env{SymbolTable: make(map[string]Symbol), FuncTable: make(map[string]SymbolFunc), Prev: prev, Id: ide, Size: make(map[string]int)}
}

func (env Env) SaveVar(id string, value Symbol) bool {
	if variable, ok := env.SymbolTable[id]; ok {
		fmt.Println("Error: La variable \""+id+"\" ya existe, con valor: ", variable.Value)
		return false
	}
	env.SymbolTable[id] = value
	return true
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

func (env Env) SetConstVar(id string, value Symbol) Symbol {
	var tmpEnv Env
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.SymbolTable[id]; ok {
			if tmpEnv.SymbolTable[id].Type == value.Type {
				tmpEnv.SymbolTable[id] = value
				return variable
			} else {
				fmt.Println("Error: Los tipos no coinciden")
				return Symbol{Line: 0, Col: 0, Type: NULL, Value: 0}
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

func (env Env) SaveFunc(id string, value SymbolFunc) {
	if v, ok := env.FuncTable[id]; ok {
		fmt.Println("Error: La función \""+id+"\" ya existe, con valor: ", v)
	}
	env.FuncTable[id] = value
}

func (env Env) GetFunc(id string) (SymbolFunc, bool) {
	var tmp Env
	tmp = env
	for {
		if v, ok := tmp.FuncTable[id]; ok {
			return v, true
		}
		if tmp.Prev == nil {
			break
		} else {
			tmp = tmp.Prev.(Env)
		}
	}
	fmt.Println("La función \"", id, "\" no existe")
	return SymbolFunc{RtnType: NULL}, false
}

func (env Env) IsLoop() bool {
	var tmp Env
	tmp = env
	for {
		if tmp.Id == "FOR IN" || tmp.Id == "WHILE" {
			return true
		}
		if tmp.Prev != nil {
			tmp = tmp.Prev.(Env)
		} else {
			break
		}
	}
	fmt.Println("Error: No se puede usar break o continue fuera de un ciclo")
	return false
}

func (env Env) IsSwitch() bool {
	var tmp Env
	tmp = env
	for {
		if tmp.Id == "SWITCH" {
			return true
		}
		if tmp.Prev != nil {
			tmp = tmp.Prev.(Env)
		} else {
			break
		}
	}
	fmt.Println("Error: No se puede usar break o continue fuera de un switch")
	return false
}

func (env Env) IsFunc() bool {
	var tmp Env
	tmp = env
	for {
		if strings.Contains(tmp.Id, "FUNC") {
			return true
		}
		if tmp.Prev != nil {
			tmp = tmp.Prev.(Env)
		} else {
			break
		}
	}
	return false
}
