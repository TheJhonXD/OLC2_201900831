// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElseifstmt is called when production elseifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterElseifstmt(ctx *ElseifstmtContext) {}

// ExitElseifstmt is called when production elseifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitElseifstmt(ctx *ElseifstmtContext) {}

// EnterElsestmt is called when production elsestmt is entered.
func (s *BaseSwiftGrammarListener) EnterElsestmt(ctx *ElsestmtContext) {}

// ExitElsestmt is called when production elsestmt is exited.
func (s *BaseSwiftGrammarListener) ExitElsestmt(ctx *ElsestmtContext) {}

// EnterVarstmt is called when production varstmt is entered.
func (s *BaseSwiftGrammarListener) EnterVarstmt(ctx *VarstmtContext) {}

// ExitVarstmt is called when production varstmt is exited.
func (s *BaseSwiftGrammarListener) ExitVarstmt(ctx *VarstmtContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseSwiftGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseSwiftGrammarListener) ExitTipo(ctx *TipoContext) {}

// EnterVarasgmt is called when production varasgmt is entered.
func (s *BaseSwiftGrammarListener) EnterVarasgmt(ctx *VarasgmtContext) {}

// ExitVarasgmt is called when production varasgmt is exited.
func (s *BaseSwiftGrammarListener) ExitVarasgmt(ctx *VarasgmtContext) {}

// EnterConststmt is called when production conststmt is entered.
func (s *BaseSwiftGrammarListener) EnterConststmt(ctx *ConststmtContext) {}

// ExitConststmt is called when production conststmt is exited.
func (s *BaseSwiftGrammarListener) ExitConststmt(ctx *ConststmtContext) {}

// EnterSwitchstmt is called when production switchstmt is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchstmt(ctx *SwitchstmtContext) {}

// ExitSwitchstmt is called when production switchstmt is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchstmt(ctx *SwitchstmtContext) {}

// EnterCasestmt is called when production casestmt is entered.
func (s *BaseSwiftGrammarListener) EnterCasestmt(ctx *CasestmtContext) {}

// ExitCasestmt is called when production casestmt is exited.
func (s *BaseSwiftGrammarListener) ExitCasestmt(ctx *CasestmtContext) {}

// EnterDefaultstmt is called when production defaultstmt is entered.
func (s *BaseSwiftGrammarListener) EnterDefaultstmt(ctx *DefaultstmtContext) {}

// ExitDefaultstmt is called when production defaultstmt is exited.
func (s *BaseSwiftGrammarListener) ExitDefaultstmt(ctx *DefaultstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterForstmt is called when production forstmt is entered.
func (s *BaseSwiftGrammarListener) EnterForstmt(ctx *ForstmtContext) {}

// ExitForstmt is called when production forstmt is exited.
func (s *BaseSwiftGrammarListener) ExitForstmt(ctx *ForstmtContext) {}

// EnterGuardstmt is called when production guardstmt is entered.
func (s *BaseSwiftGrammarListener) EnterGuardstmt(ctx *GuardstmtContext) {}

// ExitGuardstmt is called when production guardstmt is exited.
func (s *BaseSwiftGrammarListener) ExitGuardstmt(ctx *GuardstmtContext) {}

// EnterTransferstmt is called when production transferstmt is entered.
func (s *BaseSwiftGrammarListener) EnterTransferstmt(ctx *TransferstmtContext) {}

// ExitTransferstmt is called when production transferstmt is exited.
func (s *BaseSwiftGrammarListener) ExitTransferstmt(ctx *TransferstmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}
