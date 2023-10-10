// Code generated from SwiftGrammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterPrintexprlist is called when entering the printexprlist production.
	EnterPrintexprlist(c *PrintexprlistContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElseifstmt is called when entering the elseifstmt production.
	EnterElseifstmt(c *ElseifstmtContext)

	// EnterElsestmt is called when entering the elsestmt production.
	EnterElsestmt(c *ElsestmtContext)

	// EnterVarstmt is called when entering the varstmt production.
	EnterVarstmt(c *VarstmtContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterVarasgmt is called when entering the varasgmt production.
	EnterVarasgmt(c *VarasgmtContext)

	// EnterConststmt is called when entering the conststmt production.
	EnterConststmt(c *ConststmtContext)

	// EnterSwitchstmt is called when entering the switchstmt production.
	EnterSwitchstmt(c *SwitchstmtContext)

	// EnterCasestmt is called when entering the casestmt production.
	EnterCasestmt(c *CasestmtContext)

	// EnterDefaultstmt is called when entering the defaultstmt production.
	EnterDefaultstmt(c *DefaultstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterGuardstmt is called when entering the guardstmt production.
	EnterGuardstmt(c *GuardstmtContext)

	// EnterTransferstmt is called when entering the transferstmt production.
	EnterTransferstmt(c *TransferstmtContext)

	// EnterVectorstmt is called when entering the vectorstmt production.
	EnterVectorstmt(c *VectorstmtContext)

	// EnterDefinestmt is called when entering the definestmt production.
	EnterDefinestmt(c *DefinestmtContext)

	// EnterListexpr is called when entering the listexpr production.
	EnterListexpr(c *ListexprContext)

	// EnterMethodvec is called when entering the methodvec production.
	EnterMethodvec(c *MethodvecContext)

	// EnterMethodvecrtrn is called when entering the methodvecrtrn production.
	EnterMethodvecrtrn(c *MethodvecrtrnContext)

	// EnterVecaccess is called when entering the vecaccess production.
	EnterVecaccess(c *VecaccessContext)

	// EnterStructstmt is called when entering the structstmt production.
	EnterStructstmt(c *StructstmtContext)

	// EnterStructattrlist is called when entering the structattrlist production.
	EnterStructattrlist(c *StructattrlistContext)

	// EnterFuncstmt is called when entering the funcstmt production.
	EnterFuncstmt(c *FuncstmtContext)

	// EnterListparam is called when entering the listparam production.
	EnterListparam(c *ListparamContext)

	// EnterCallfunc is called when entering the callfunc production.
	EnterCallfunc(c *CallfuncContext)

	// EnterCallfuncinstruction is called when entering the callfuncinstruction production.
	EnterCallfuncinstruction(c *CallfuncinstructionContext)

	// EnterListparamcall is called when entering the listparamcall production.
	EnterListparamcall(c *ListparamcallContext)

	// EnterFuncembed is called when entering the funcembed production.
	EnterFuncembed(c *FuncembedContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitPrintexprlist is called when exiting the printexprlist production.
	ExitPrintexprlist(c *PrintexprlistContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElseifstmt is called when exiting the elseifstmt production.
	ExitElseifstmt(c *ElseifstmtContext)

	// ExitElsestmt is called when exiting the elsestmt production.
	ExitElsestmt(c *ElsestmtContext)

	// ExitVarstmt is called when exiting the varstmt production.
	ExitVarstmt(c *VarstmtContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitVarasgmt is called when exiting the varasgmt production.
	ExitVarasgmt(c *VarasgmtContext)

	// ExitConststmt is called when exiting the conststmt production.
	ExitConststmt(c *ConststmtContext)

	// ExitSwitchstmt is called when exiting the switchstmt production.
	ExitSwitchstmt(c *SwitchstmtContext)

	// ExitCasestmt is called when exiting the casestmt production.
	ExitCasestmt(c *CasestmtContext)

	// ExitDefaultstmt is called when exiting the defaultstmt production.
	ExitDefaultstmt(c *DefaultstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitGuardstmt is called when exiting the guardstmt production.
	ExitGuardstmt(c *GuardstmtContext)

	// ExitTransferstmt is called when exiting the transferstmt production.
	ExitTransferstmt(c *TransferstmtContext)

	// ExitVectorstmt is called when exiting the vectorstmt production.
	ExitVectorstmt(c *VectorstmtContext)

	// ExitDefinestmt is called when exiting the definestmt production.
	ExitDefinestmt(c *DefinestmtContext)

	// ExitListexpr is called when exiting the listexpr production.
	ExitListexpr(c *ListexprContext)

	// ExitMethodvec is called when exiting the methodvec production.
	ExitMethodvec(c *MethodvecContext)

	// ExitMethodvecrtrn is called when exiting the methodvecrtrn production.
	ExitMethodvecrtrn(c *MethodvecrtrnContext)

	// ExitVecaccess is called when exiting the vecaccess production.
	ExitVecaccess(c *VecaccessContext)

	// ExitStructstmt is called when exiting the structstmt production.
	ExitStructstmt(c *StructstmtContext)

	// ExitStructattrlist is called when exiting the structattrlist production.
	ExitStructattrlist(c *StructattrlistContext)

	// ExitFuncstmt is called when exiting the funcstmt production.
	ExitFuncstmt(c *FuncstmtContext)

	// ExitListparam is called when exiting the listparam production.
	ExitListparam(c *ListparamContext)

	// ExitCallfunc is called when exiting the callfunc production.
	ExitCallfunc(c *CallfuncContext)

	// ExitCallfuncinstruction is called when exiting the callfuncinstruction production.
	ExitCallfuncinstruction(c *CallfuncinstructionContext)

	// ExitListparamcall is called when exiting the listparamcall production.
	ExitListparamcall(c *ListparamcallContext)

	// ExitFuncembed is called when exiting the funcembed production.
	ExitFuncembed(c *FuncembedContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
