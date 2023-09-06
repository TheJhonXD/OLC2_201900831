package environment

type SymbolFuncParam struct {
	Line       int
	Col        int
	ExternalId string
	InternalId string
	Reference  bool
	Type       TipoExpresion
}
