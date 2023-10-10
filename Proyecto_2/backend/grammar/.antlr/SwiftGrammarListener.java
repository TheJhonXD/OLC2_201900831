// Generated from c:/Users/TheJhonX/Documents/GitHub/OLC2_201900831/Proyecto_2/backend/grammar/SwiftGrammar.g4 by ANTLR 4.13.1

    import "Server/interfaces"
    import "Server/environment"
    import "Server/expressions"
    import "Server/instructions"
    import "strings"

import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SwiftGrammarParser}.
 */
public interface SwiftGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(SwiftGrammarParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(SwiftGrammarParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(SwiftGrammarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void enterInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#instruction}.
	 * @param ctx the parse tree
	 */
	void exitInstruction(SwiftGrammarParser.InstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#printstmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintstmt(SwiftGrammarParser.PrintstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#printexprlist}.
	 * @param ctx the parse tree
	 */
	void enterPrintexprlist(SwiftGrammarParser.PrintexprlistContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#printexprlist}.
	 * @param ctx the parse tree
	 */
	void exitPrintexprlist(SwiftGrammarParser.PrintexprlistContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void enterIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#ifstmt}.
	 * @param ctx the parse tree
	 */
	void exitIfstmt(SwiftGrammarParser.IfstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#elseifstmt}.
	 * @param ctx the parse tree
	 */
	void enterElseifstmt(SwiftGrammarParser.ElseifstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#elseifstmt}.
	 * @param ctx the parse tree
	 */
	void exitElseifstmt(SwiftGrammarParser.ElseifstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#elsestmt}.
	 * @param ctx the parse tree
	 */
	void enterElsestmt(SwiftGrammarParser.ElsestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#elsestmt}.
	 * @param ctx the parse tree
	 */
	void exitElsestmt(SwiftGrammarParser.ElsestmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#varstmt}.
	 * @param ctx the parse tree
	 */
	void enterVarstmt(SwiftGrammarParser.VarstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#varstmt}.
	 * @param ctx the parse tree
	 */
	void exitVarstmt(SwiftGrammarParser.VarstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(SwiftGrammarParser.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(SwiftGrammarParser.TipoContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#varasgmt}.
	 * @param ctx the parse tree
	 */
	void enterVarasgmt(SwiftGrammarParser.VarasgmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#varasgmt}.
	 * @param ctx the parse tree
	 */
	void exitVarasgmt(SwiftGrammarParser.VarasgmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#conststmt}.
	 * @param ctx the parse tree
	 */
	void enterConststmt(SwiftGrammarParser.ConststmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#conststmt}.
	 * @param ctx the parse tree
	 */
	void exitConststmt(SwiftGrammarParser.ConststmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#switchstmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchstmt(SwiftGrammarParser.SwitchstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#switchstmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchstmt(SwiftGrammarParser.SwitchstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#casestmt}.
	 * @param ctx the parse tree
	 */
	void enterCasestmt(SwiftGrammarParser.CasestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#casestmt}.
	 * @param ctx the parse tree
	 */
	void exitCasestmt(SwiftGrammarParser.CasestmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#defaultstmt}.
	 * @param ctx the parse tree
	 */
	void enterDefaultstmt(SwiftGrammarParser.DefaultstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#defaultstmt}.
	 * @param ctx the parse tree
	 */
	void exitDefaultstmt(SwiftGrammarParser.DefaultstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#whilestmt}.
	 * @param ctx the parse tree
	 */
	void enterWhilestmt(SwiftGrammarParser.WhilestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#whilestmt}.
	 * @param ctx the parse tree
	 */
	void exitWhilestmt(SwiftGrammarParser.WhilestmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#forstmt}.
	 * @param ctx the parse tree
	 */
	void enterForstmt(SwiftGrammarParser.ForstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#forstmt}.
	 * @param ctx the parse tree
	 */
	void exitForstmt(SwiftGrammarParser.ForstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#guardstmt}.
	 * @param ctx the parse tree
	 */
	void enterGuardstmt(SwiftGrammarParser.GuardstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#guardstmt}.
	 * @param ctx the parse tree
	 */
	void exitGuardstmt(SwiftGrammarParser.GuardstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#transferstmt}.
	 * @param ctx the parse tree
	 */
	void enterTransferstmt(SwiftGrammarParser.TransferstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#transferstmt}.
	 * @param ctx the parse tree
	 */
	void exitTransferstmt(SwiftGrammarParser.TransferstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vectorstmt}.
	 * @param ctx the parse tree
	 */
	void enterVectorstmt(SwiftGrammarParser.VectorstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vectorstmt}.
	 * @param ctx the parse tree
	 */
	void exitVectorstmt(SwiftGrammarParser.VectorstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#definestmt}.
	 * @param ctx the parse tree
	 */
	void enterDefinestmt(SwiftGrammarParser.DefinestmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#definestmt}.
	 * @param ctx the parse tree
	 */
	void exitDefinestmt(SwiftGrammarParser.DefinestmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listexpr}.
	 * @param ctx the parse tree
	 */
	void enterListexpr(SwiftGrammarParser.ListexprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listexpr}.
	 * @param ctx the parse tree
	 */
	void exitListexpr(SwiftGrammarParser.ListexprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#methodvec}.
	 * @param ctx the parse tree
	 */
	void enterMethodvec(SwiftGrammarParser.MethodvecContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#methodvec}.
	 * @param ctx the parse tree
	 */
	void exitMethodvec(SwiftGrammarParser.MethodvecContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#methodvecrtrn}.
	 * @param ctx the parse tree
	 */
	void enterMethodvecrtrn(SwiftGrammarParser.MethodvecrtrnContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#methodvecrtrn}.
	 * @param ctx the parse tree
	 */
	void exitMethodvecrtrn(SwiftGrammarParser.MethodvecrtrnContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#vecaccess}.
	 * @param ctx the parse tree
	 */
	void enterVecaccess(SwiftGrammarParser.VecaccessContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#vecaccess}.
	 * @param ctx the parse tree
	 */
	void exitVecaccess(SwiftGrammarParser.VecaccessContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#structstmt}.
	 * @param ctx the parse tree
	 */
	void enterStructstmt(SwiftGrammarParser.StructstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#structstmt}.
	 * @param ctx the parse tree
	 */
	void exitStructstmt(SwiftGrammarParser.StructstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#structattrlist}.
	 * @param ctx the parse tree
	 */
	void enterStructattrlist(SwiftGrammarParser.StructattrlistContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#structattrlist}.
	 * @param ctx the parse tree
	 */
	void exitStructattrlist(SwiftGrammarParser.StructattrlistContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#funcstmt}.
	 * @param ctx the parse tree
	 */
	void enterFuncstmt(SwiftGrammarParser.FuncstmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#funcstmt}.
	 * @param ctx the parse tree
	 */
	void exitFuncstmt(SwiftGrammarParser.FuncstmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listparam}.
	 * @param ctx the parse tree
	 */
	void enterListparam(SwiftGrammarParser.ListparamContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listparam}.
	 * @param ctx the parse tree
	 */
	void exitListparam(SwiftGrammarParser.ListparamContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#callfunc}.
	 * @param ctx the parse tree
	 */
	void enterCallfunc(SwiftGrammarParser.CallfuncContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#callfunc}.
	 * @param ctx the parse tree
	 */
	void exitCallfunc(SwiftGrammarParser.CallfuncContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#callfuncinstruction}.
	 * @param ctx the parse tree
	 */
	void enterCallfuncinstruction(SwiftGrammarParser.CallfuncinstructionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#callfuncinstruction}.
	 * @param ctx the parse tree
	 */
	void exitCallfuncinstruction(SwiftGrammarParser.CallfuncinstructionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#listparamcall}.
	 * @param ctx the parse tree
	 */
	void enterListparamcall(SwiftGrammarParser.ListparamcallContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#listparamcall}.
	 * @param ctx the parse tree
	 */
	void exitListparamcall(SwiftGrammarParser.ListparamcallContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#funcembed}.
	 * @param ctx the parse tree
	 */
	void enterFuncembed(SwiftGrammarParser.FuncembedContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#funcembed}.
	 * @param ctx the parse tree
	 */
	void exitFuncembed(SwiftGrammarParser.FuncembedContext ctx);
	/**
	 * Enter a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(SwiftGrammarParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SwiftGrammarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(SwiftGrammarParser.ExprContext ctx);
}