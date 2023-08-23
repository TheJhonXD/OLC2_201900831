package server

import (
	"Server/environment"
	"Server/grammar/parser"
	"Server/interfaces"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseSwiftGrammarListener
	Code []interface{}
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitS(ctx *parser.SContext) {
	this.Code = ctx.GetCode()
}

func analyzer(code string) string {
	// Leyendo entrada
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Creacion de parser
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	// return tree.ToStringTree(nil, p)
	//listener
	var listener *TreeShapeListener = NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	Code := listener.Code
	//create ast
	var Ast environment.AST
	//create environment
	env := environment.NewEnv(nil, "GLOBAL")
	//ejecución del codigo
	for _, inst := range Code {
		// fmt.Println(reflect.TypeOf(inst))
		inst.(interfaces.Instruction).Ejecutar(&Ast, env)
	}
	return Ast.GetPrint()
}
