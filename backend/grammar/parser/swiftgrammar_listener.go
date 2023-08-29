// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
