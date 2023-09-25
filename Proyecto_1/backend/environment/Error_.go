package environment

type Error_ struct {
	Line        int
	Column      int
	Scope       string
	Description string
}

func NewError(line int, column int, scope string, description string) Error_ {
	err := Error_{Line: line, Column: column, Scope: scope, Description: description}
	return err
}
