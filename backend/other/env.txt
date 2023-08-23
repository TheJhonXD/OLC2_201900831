package environment

import "fmt"

type Env struct {
	SymbolTable map[string]Symbol
	Next        *Env
}

type LinkedList struct {
	First *Env
	Last  *Env
}

func (ll *LinkedList) Add(newEnv *Env) {
	if ll.First == nil {
		ll.First = newEnv
		ll.Last = newEnv
	} else {
		ll.Last.Next = newEnv
		ll.Last = newEnv
	}
}

func (ll *LinkedList) GetLast() *Env {
	return ll.Last
}

func (ll *LinkedList) GetFirst() *Env {
	return ll.First
}

// Muestra toda la lista
func (ll *LinkedList) GetList() {
	actual := ll.First
	for actual != nil {
		fmt.Println(actual.SymbolTable)
		actual = actual.Next
	}
}
