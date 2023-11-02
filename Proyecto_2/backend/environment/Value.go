package environment

type Value struct {
	Value        string
	IsTmp        bool
	Type         TipoExpresion
	TrueLabel    []interface{}
	FalseLabel   []interface{}
	OutLabel     []interface{}
	IntValue     int
	FloatValue   float64
	BreakFlag    bool
	ContinueFlag bool
	ReturnFlag   bool
}

func NewValue(val string, isTmp bool, tipo TipoExpresion) Value {
	return Value{Value: val, IsTmp: isTmp, Type: tipo, IntValue: 0}
}
