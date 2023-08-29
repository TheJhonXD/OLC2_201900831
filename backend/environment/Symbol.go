package environment

type Symbol struct {
	Line         int
	Col          int
	Type         TipoExpresion
	Value        interface{}
	Const        bool
	ReturnFlag   bool
	BreakFlag    bool
	ContinueFlag bool
}
