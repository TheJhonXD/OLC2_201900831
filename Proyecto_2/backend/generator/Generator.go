package generator

import "fmt"

type Generator struct {
	Tmp           int
	Label         int
	Code          []interface{}
	FinalCode     []interface{}
	Natives       []interface{}
	FuncCode      []interface{}
	TmpList       []interface{}
	BreakLabel    string
	ContinueLabel string
	MainCode      bool
}

func NewGenerator() Generator {
	return Generator{
		Tmp:           0,
		Label:         0,
		BreakLabel:    "",
		ContinueLabel: "",
		MainCode:      true,
	}
}

func (g *Generator) GetCode() []interface{} {
	return g.Code
}

func (g *Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

// Add break label
func (g *Generator) AddBreakLabel(label string) {
	g.BreakLabel = label
}

// Add continue label
func (g *Generator) AddContinueLabel(label string) {
	g.ContinueLabel = label
}

// Generate new tmp
func (g *Generator) NewTmp() string {
	temp := "t" + fmt.Sprintf("%v", g.Tmp)
	g.Tmp = g.Tmp + 1
	//Lo guardamos para declararlo
	g.TmpList = append(g.TmpList, temp)
	return temp
}

// Generate new label
func (g *Generator) NewLabel() string {
	temp := g.Label
	g.Label = g.Label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

// Add label to code
func (g *Generator) AddLabel(label string) {
	if g.MainCode {
		g.Code = append(g.Code, label+":\n")
	} else {
		g.FuncCode = append(g.FuncCode, label+":\n")
	}
}

func (g *Generator) AddIf(left string, right string, operator string, label string) {
	if g.MainCode {
		g.Code = append(g.Code, "if("+left+" "+operator+" "+right+") goto "+label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "if("+left+" "+operator+" "+right+") goto "+label+";\n")
	}
}

func (g *Generator) AddGoto(label string) {
	if g.MainCode {
		g.Code = append(g.Code, "goto "+label+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "goto "+label+";\n")
	}
}

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+left+" "+operator+" "+right+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+left+" "+operator+" "+right+";\n")
	}
}

func (g *Generator) AddAssign(target string, val string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = "+val+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = "+val+";\n")
	}
}

func (g *Generator) AddPrintf(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

func (g *Generator) AddBr() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

func (g *Generator) AddComment(target string) {
	if g.MainCode {
		g.Code = append(g.Code, "//"+target+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "//"+target+"\n")
	}
}
