package environment

type Value struct {
	Value      string
	IsTmp      bool
	Type       string
	TrueLabel  []interface{}
	FalseLabel []interface{}
	OutLabel   []interface{}
	IntValue   int
}

func NewValue(val string, isTmp bool, tipo string) Value {
	return Value{Value: val, IsTmp: isTmp, Type: tipo, IntValue: 0}
}
