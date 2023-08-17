package environment

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       string
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
