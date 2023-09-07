package environment

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       string
	ErroresAux   []Error_
}

func NewAST(inst []interface{}, print string, err string) AST {
	ast := AST{Instructions: inst, Print: print, Errors: err}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(text string) {
	a.Print = a.Print + text + "\n"
}

func (a *AST) GetErrors() string {
	return a.Errors
}

func (a *AST) SetError(err string) {
	a.Errors = a.Errors + err
}

func (a *AST) AddError(line int, column int, scope string, description string) {
	err := NewError(line, column, scope, description)
	a.ErroresAux = append(a.ErroresAux, err)
}

func (a *AST) GetErrorsAux() []Error_ {
	return a.ErroresAux
}
