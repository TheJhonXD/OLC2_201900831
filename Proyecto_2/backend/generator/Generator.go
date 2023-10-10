package generator

import "fmt"

type Generator struct {
	Tmp             int
	Label           int
	Code            []interface{}
	FinalCode       []interface{}
	Natives         []interface{}
	FuncCode        []interface{}
	TmpList         []interface{}
	PrintStringFlag bool
	BreakLabel      string
	ContinueLabel   string
	MainCode        bool
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

func (g *Generator) GetTmps() []interface{} {
	return g.TmpList
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

func (g *Generator) AddSetHeap(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "heap["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "heap["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetHeap(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = heap["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = heap["+index+"];\n")
	}
}

func (g *Generator) AddSetStack(index string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "stack["+index+"] = "+value+";\n")
	} else {
		g.FuncCode = append(g.FuncCode, "stack["+index+"] = "+value+";\n")
	}
}

func (g *Generator) AddGetStack(target string, index string) {
	if g.MainCode {
		g.Code = append(g.Code, target+" = stack["+index+"];\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+" = stack["+index+"];\n")
	}
}

func (g *Generator) AddCall(target string) {
	if g.MainCode {
		g.Code = append(g.Code, target+"();\n")
	} else {
		g.FuncCode = append(g.FuncCode, target+"();\n")
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

func (g *Generator) GenerateFinalCode() {
	/* ********** HEADER ********** */
	g.FinalCode = append(g.FinalCode, "#include <stdio.h>\n")
	g.FinalCode = append(g.FinalCode, "#include <math.h>\n")
	g.FinalCode = append(g.FinalCode, "double heap[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double stack[30101999];\n")
	g.FinalCode = append(g.FinalCode, "double P;\n")
	g.FinalCode = append(g.FinalCode, "double H;\n")

	/* ********** TEMPORAL DECLARATION ********** */
	tmpList := g.GetTmps()
	if len(tmpList) > 0 {
		g.FinalCode = append(g.FinalCode, "double ")
		tmpDec := fmt.Sprintf("%v", tmpList[0])
		tmpList = tmpList[1:]
		for _, tmp := range tmpList {
			tmpDec += ", " + fmt.Sprintf("%v", tmp)
		}
		tmpDec += ";\n\n"
		g.FinalCode = append(g.FinalCode, tmpDec)
	}

	/* ********** NATIVE FUNCTIONS ********** */
	if len(g.Natives) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------NATIVES------*/\n")
		for _, s := range g.Natives {
			g.FinalCode = append(g.FinalCode, s)
		}
	}

	/* ********** FUNCTIONS ********** */
	if len(g.FuncCode) > 0 {
		g.FinalCode = append(g.FinalCode, "/*------FUNCTIONS------*/\n")
		for _, s := range g.FuncCode {
			g.FinalCode = append(g.FinalCode, s)
		}
	}

	/* ********** MAIN ********** */
	g.FinalCode = append(g.FinalCode, "/*------MAIN------*/\n")
	g.FinalCode = append(g.FinalCode, "void main() {\n")
	g.FinalCode = append(g.FinalCode, "\tP = 0; H = 0;\n\n")
	for _, s := range g.Code {
		g.FinalCode = append(g.FinalCode, "\t"+s.(string))
	}
	g.FinalCode = append(g.FinalCode, "\n\treturn;\n}\n")
}

func (g *Generator) GeneratePrintString() {
	if g.PrintStringFlag {
		// Generando temporales y etiquetas
		newTmp1 := g.NewTmp()
		newTmp2 := g.NewTmp()
		newTmp3 := g.NewTmp()
		newLbl1 := g.NewLabel()
		newLbl2 := g.NewLabel()
		// Se genera la funcion printstring
		g.Natives = append(g.Natives, "void dbrust_printString() {\n")
		g.Natives = append(g.Natives, "\t"+newTmp1+" = P + 1;\n")
		g.Natives = append(g.Natives, "\t"+newTmp2+" = stack[(int)"+newTmp1+"];\n")
		g.Natives = append(g.Natives, "\t"+newLbl2+":\n")
		g.Natives = append(g.Natives, "\t"+newTmp3+" = heap[(int)"+newTmp2+"];\n")
		g.Natives = append(g.Natives, "\tif("+newTmp3+" == -1) goto "+newLbl1+";\n")
		g.Natives = append(g.Natives, "\tprintf(\"%c\", (char)"+newTmp3+");\n")
		g.Natives = append(g.Natives, "\t"+newTmp2+" = "+newTmp2+" + 1;\n")
		g.Natives = append(g.Natives, "\tgoto "+newLbl2+";\n")
		g.Natives = append(g.Natives, "\t"+newLbl1+":\n")
		g.Natives = append(g.Natives, "\treturn;\n")
		g.Natives = append(g.Natives, "}\n\n")
		g.PrintStringFlag = false
	}
}
