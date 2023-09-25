package environment

type SymbolT struct {
	Line       int
	Col        int
	DataType   TipoExpresion
	SymbolType string
	Scope      string
	Id         string
}

func NewSymbolT(line int, col int, dataType TipoExpresion, symbolType string, scope string, id string) SymbolT {
	return SymbolT{Line: line, Col: col, DataType: dataType, SymbolType: symbolType, Scope: scope, Id: id}
}
