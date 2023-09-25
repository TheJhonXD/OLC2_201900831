package environment

type TipoExpresion int

const (
	INTEGER  TipoExpresion = iota // 0
	FLOAT                         // 1
	STRING                        // 2
	BOOLEAN                       // 3
	NULL                          // 4
	CHAR                          // 5
	WILDCARD                      // 6
	VECTOR                        // 7
)
