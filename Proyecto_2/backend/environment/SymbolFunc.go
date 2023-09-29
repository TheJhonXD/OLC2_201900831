package environment

type SymbolFunc struct {
	Line    int
	Col     int
	Id      string
	Args    []interface{}
	RtnType TipoExpresion
	Block   []interface{}
}
