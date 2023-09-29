package environment

type SymbolVar struct {
	Line  int
	Col   int
	Type  TipoExpresion
	Value interface{}
}
