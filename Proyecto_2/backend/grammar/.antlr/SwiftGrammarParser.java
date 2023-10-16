// Generated from c:/Users/TheJhonX/Documents/GitHub/OLC2_201900831/Proyecto_2/backend/grammar/SwiftGrammar.g4 by ANTLR 4.13.1

    import "Server/interfaces"
    import "Server/environment"
    import "Server/expressions"
    import "Server/instructions"
    import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, STR=4, CHARACTER=5, TRU=6, FAL=7, NIL=8, VAR=9, 
		LET=10, PRINT=11, IF=12, ELSE=13, SWITCH=14, CASE=15, DEFAULT=16, WHILE=17, 
		FOR=18, IN=19, RANGEPTS=20, GUARD=21, BREAK=22, CONTINUE=23, RETURN=24, 
		STRUCT=25, SELF=26, MUTATING=27, FUNC=28, APPEND=29, REMOVELAST=30, REMOVE=31, 
		EMPTY=32, COUNT=33, INOUT=34, CAME=35, NUMBER=36, CHAR=37, STRING=38, 
		ID=39, DIF=40, IG_IG=41, NOT=42, OR=43, AND=44, IG=45, MAY_IG=46, MEN_IG=47, 
		MAYOR=48, MENOR=49, MOD=50, MUL=51, DIV=52, ADD=53, SUB=54, PARIZQ=55, 
		PARDER=56, LLAVEIZQ=57, LLAVEDER=58, CORCHIZQ=59, CORCHDER=60, Q_MARK=61, 
		ARROW=62, COMA=63, PUNTO=64, COLON=65, AMP=66, WHITESPACE=67, COMMENT=68, 
		LINE_COMMENT=69;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_printexprlist = 4, RULE_ifstmt = 5, RULE_elseifstmt = 6, RULE_elsestmt = 7, 
		RULE_varstmt = 8, RULE_tipo = 9, RULE_varasgmt = 10, RULE_conststmt = 11, 
		RULE_switchstmt = 12, RULE_casestmt = 13, RULE_defaultstmt = 14, RULE_whilestmt = 15, 
		RULE_forstmt = 16, RULE_guardstmt = 17, RULE_transferstmt = 18, RULE_vectorstmt = 19, 
		RULE_definestmt = 20, RULE_listexpr = 21, RULE_methodvec = 22, RULE_methodvecrtrn = 23, 
		RULE_vecaccess = 24, RULE_structstmt = 25, RULE_structattrlist = 26, RULE_funcstmt = 27, 
		RULE_listparam = 28, RULE_callfunc = 29, RULE_callfuncinstruction = 30, 
		RULE_listparamcall = 31, RULE_funcembed = 32, RULE_expr = 33;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "printexprlist", "ifstmt", 
			"elseifstmt", "elsestmt", "varstmt", "tipo", "varasgmt", "conststmt", 
			"switchstmt", "casestmt", "defaultstmt", "whilestmt", "forstmt", "guardstmt", 
			"transferstmt", "vectorstmt", "definestmt", "listexpr", "methodvec", 
			"methodvecrtrn", "vecaccess", "structstmt", "structattrlist", "funcstmt", 
			"listparam", "callfunc", "callfuncinstruction", "listparamcall", "funcembed", 
			"expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", 
			"'break'", "'continue'", "'return'", "'struct'", "'self'", "'mutating'", 
			"'func'", "'append'", "'removeLast'", "'remove'", "'isEmpty'", "'count'", 
			"'inout'", "'_'", null, null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", 
			"'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'?'", "'->'", "','", 
			"'.'", "':'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL", 
			"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "IN", "RANGEPTS", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", 
			"SELF", "MUTATING", "FUNC", "APPEND", "REMOVELAST", "REMOVE", "EMPTY", 
			"COUNT", "INOUT", "CAME", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", 
			"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"CORCHIZQ", "CORCHDER", "Q_MARK", "ARROW", "COMA", "PUNTO", "COLON", 
			"AMP", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterS(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitS(this);
		}
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			((SContext)_localctx).block = block();
			setState(69);
			match(EOF);
			   
			        _localctx.code = ((SContext)_localctx).block.blk
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(73); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(72);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(75); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public PrintstmtContext printstmt;
		public CallfuncinstructionContext callfuncinstruction;
		public IfstmtContext ifstmt;
		public VarstmtContext varstmt;
		public VarasgmtContext varasgmt;
		public ConststmtContext conststmt;
		public SwitchstmtContext switchstmt;
		public WhilestmtContext whilestmt;
		public ForstmtContext forstmt;
		public GuardstmtContext guardstmt;
		public TransferstmtContext transferstmt;
		public VectorstmtContext vectorstmt;
		public MethodvecContext methodvec;
		public VecaccessContext vecaccess;
		public FuncstmtContext funcstmt;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public CallfuncinstructionContext callfuncinstruction() {
			return getRuleContext(CallfuncinstructionContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public VarstmtContext varstmt() {
			return getRuleContext(VarstmtContext.class,0);
		}
		public VarasgmtContext varasgmt() {
			return getRuleContext(VarasgmtContext.class,0);
		}
		public ConststmtContext conststmt() {
			return getRuleContext(ConststmtContext.class,0);
		}
		public SwitchstmtContext switchstmt() {
			return getRuleContext(SwitchstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public TransferstmtContext transferstmt() {
			return getRuleContext(TransferstmtContext.class,0);
		}
		public VectorstmtContext vectorstmt() {
			return getRuleContext(VectorstmtContext.class,0);
		}
		public MethodvecContext methodvec() {
			return getRuleContext(MethodvecContext.class,0);
		}
		public VecaccessContext vecaccess() {
			return getRuleContext(VecaccessContext.class,0);
		}
		public FuncstmtContext funcstmt() {
			return getRuleContext(FuncstmtContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterInstruction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitInstruction(this);
		}
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(124);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(79);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				((InstructionContext)_localctx).callfuncinstruction = callfuncinstruction();
				 _localctx.inst = ((InstructionContext)_localctx).callfuncinstruction.callfuncinstr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(85);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinstr 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(88);
				((InstructionContext)_localctx).varstmt = varstmt();
				 _localctx.inst = ((InstructionContext)_localctx).varstmt.var
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(91);
				((InstructionContext)_localctx).varasgmt = varasgmt();
				 _localctx.inst = ((InstructionContext)_localctx).varasgmt.asgmt
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(94);
				((InstructionContext)_localctx).conststmt = conststmt();
				 _localctx.inst = ((InstructionContext)_localctx).conststmt.const_
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(97);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchstmt.switchinstr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(100);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whileinstr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(103);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.forinstr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(106);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.guardinstr
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(109);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				 _localctx.inst = ((InstructionContext)_localctx).transferstmt.trns
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(112);
				((InstructionContext)_localctx).vectorstmt = vectorstmt();
				 _localctx.inst = ((InstructionContext)_localctx).vectorstmt.vectorinstr
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(115);
				((InstructionContext)_localctx).methodvec = methodvec();
				 _localctx.inst = ((InstructionContext)_localctx).methodvec.methodinstr
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(118);
				((InstructionContext)_localctx).vecaccess = vecaccess();
				 _localctx.inst = ((InstructionContext)_localctx).vecaccess.vecacc
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(121);
				((InstructionContext)_localctx).funcstmt = funcstmt();
				 _localctx.inst = ((InstructionContext)_localctx).funcstmt.funcinstr
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public PrintexprlistContext printexprlist;
		public List<PrintexprlistContext> lista = new ArrayList<PrintexprlistContext>();
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public List<PrintexprlistContext> printexprlist() {
			return getRuleContexts(PrintexprlistContext.class);
		}
		public PrintexprlistContext printexprlist(int i) {
			return getRuleContext(PrintexprlistContext.class,i);
		}
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterPrintstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitPrintstmt(this);
		}
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_printstmt);

		    var listExpr []interface{}

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(127);
			match(PARIZQ);
			setState(129); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(128);
				((PrintstmtContext)_localctx).printexprlist = printexprlist();
				((PrintstmtContext)_localctx).lista.add(((PrintstmtContext)_localctx).printexprlist);
				}
				}
				setState(131); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 54048624367108566L) != 0) );
			setState(133);
			match(PARDER);
			 
			    for _, e := range localctx.(*PrintstmtContext).GetLista() {
			        listExpr = append(listExpr, e.GetPrntexpr())
			    }
			    _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0), (((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0), listExpr)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintexprlistContext extends ParserRuleContext {
		public interfaces.Expression prntexpr;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public PrintexprlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printexprlist; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterPrintexprlist(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitPrintexprlist(this);
		}
	}

	public final PrintexprlistContext printexprlist() throws RecognitionException {
		PrintexprlistContext _localctx = new PrintexprlistContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_printexprlist);
		try {
			setState(143);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				((PrintexprlistContext)_localctx).expr = expr(0);
				setState(137);
				match(COMA);
				 _localctx.prntexpr = ((PrintexprlistContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				((PrintexprlistContext)_localctx).expr = expr(0);
				 _localctx.prntexpr = ((PrintexprlistContext)_localctx).expr.e 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ifinstr;
		public Token IF;
		public ExprContext expr;
		public BlockContext block;
		public BlockContext firstBlk;
		public ElseifstmtContext elseifstmt;
		public List<ElseifstmtContext> elif = new ArrayList<ElseifstmtContext>();
		public ElsestmtContext elsestmt;
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElsestmtContext elsestmt() {
			return getRuleContext(ElsestmtContext.class,0);
		}
		public List<ElseifstmtContext> elseifstmt() {
			return getRuleContexts(ElseifstmtContext.class);
		}
		public ElseifstmtContext elseifstmt(int i) {
			return getRuleContext(ElseifstmtContext.class,i);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterIfstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitIfstmt(this);
		}
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			setState(185);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(145);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(146);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(147);
				match(LLAVEIZQ);
				setState(148);
				((IfstmtContext)_localctx).block = block();
				setState(149);
				match(LLAVEDER);
				 _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(152);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(153);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(154);
				match(LLAVEIZQ);
				setState(155);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(156);
				match(LLAVEDER);
				setState(158); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(157);
						((IfstmtContext)_localctx).elseifstmt = elseifstmt();
						((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(160); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(162);
				((IfstmtContext)_localctx).elsestmt = elsestmt();
				 
				    var ifInterfaces []interface{}
				    // fmt.Println(((IfstmtContext)_localctx).elif)
				    for _, e := range localctx.(*IfstmtContext).GetElif() {
				        ifInterfaces = append(ifInterfaces, e.GetElifinstr())
				        fmt.Println(e.GetElifinstr())
				    }
				    _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).firstBlk.blk, ifInterfaces, ((IfstmtContext)_localctx).elsestmt.elseinstr)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(165);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(166);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(167);
				match(LLAVEIZQ);
				setState(168);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(169);
				match(LLAVEDER);
				setState(171); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(170);
					((IfstmtContext)_localctx).elseifstmt = elseifstmt();
					((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
					}
					}
					setState(173); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==ELSE );

				    var ifInterfaces []interface{}
				    // fmt.Println(((IfstmtContext)_localctx).elif)
				    for _, e := range localctx.(*IfstmtContext).GetElif() {
				        ifInterfaces = append(ifInterfaces, e.GetElifinstr())
				        fmt.Println(e.GetElifinstr())
				    }
				    _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).firstBlk.blk, ifInterfaces, nil)

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(177);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(178);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(179);
				match(LLAVEIZQ);
				setState(180);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(181);
				match(LLAVEDER);
				setState(182);
				((IfstmtContext)_localctx).elsestmt = elsestmt();
				 _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).firstBlk.blk, nil, ((IfstmtContext)_localctx).elsestmt.elseinstr) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ElseifstmtContext extends ParserRuleContext {
		public interfaces.Instruction elifinstr;
		public Token ELSE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElseifstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseifstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterElseifstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitElseifstmt(this);
		}
	}

	public final ElseifstmtContext elseifstmt() throws RecognitionException {
		ElseifstmtContext _localctx = new ElseifstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			((ElseifstmtContext)_localctx).ELSE = match(ELSE);
			setState(188);
			match(IF);
			setState(189);
			((ElseifstmtContext)_localctx).expr = expr(0);
			setState(190);
			match(LLAVEIZQ);
			setState(191);
			((ElseifstmtContext)_localctx).block = block();
			setState(192);
			match(LLAVEDER);
			 
			    _localctx.elifinstr = instructions.NewIf((((ElseifstmtContext)_localctx).ELSE!=null?((ElseifstmtContext)_localctx).ELSE.getLine():0), (((ElseifstmtContext)_localctx).ELSE!=null?((ElseifstmtContext)_localctx).ELSE.getCharPositionInLine():0), ((ElseifstmtContext)_localctx).expr.e, nil, ((ElseifstmtContext)_localctx).block.blk, nil)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ElsestmtContext extends ParserRuleContext {
		public []interface{} elseinstr;
		public BlockContext block;
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElsestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elsestmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterElsestmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitElsestmt(this);
		}
	}

	public final ElsestmtContext elsestmt() throws RecognitionException {
		ElsestmtContext _localctx = new ElsestmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_elsestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			match(ELSE);
			setState(196);
			match(LLAVEIZQ);
			setState(197);
			((ElsestmtContext)_localctx).block = block();
			setState(198);
			match(LLAVEDER);
			 _localctx.elseinstr = ((ElsestmtContext)_localctx).block.blk 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarstmtContext extends ParserRuleContext {
		public interfaces.Instruction var;
		public Token VAR;
		public Token ID;
		public TipoContext tipo;
		public ExprContext expr;
		public Token Q_MARK;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode Q_MARK() { return getToken(SwiftGrammarParser.Q_MARK, 0); }
		public VarstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterVarstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitVarstmt(this);
		}
	}

	public final VarstmtContext varstmt() throws RecognitionException {
		VarstmtContext _localctx = new VarstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varstmt);
		try {
			setState(222);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(201);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(202);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(203);
				match(COLON);
				setState(204);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(205);
				match(IG);
				setState(206);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), ((VarstmtContext)_localctx).tipo.rtipo, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(209);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(210);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(211);
				match(IG);
				setState(212);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), environment.NULL, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(215);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(216);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(217);
				match(COLON);
				setState(218);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(219);
				((VarstmtContext)_localctx).Q_MARK = match(Q_MARK);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), ((VarstmtContext)_localctx).tipo.rtipo, expressions.NewPrimitive((((VarstmtContext)_localctx).Q_MARK!=null?((VarstmtContext)_localctx).Q_MARK.getLine():0), (((VarstmtContext)_localctx).Q_MARK!=null?((VarstmtContext)_localctx).Q_MARK.getCharPositionInLine():0), environment.WILDCARD, environment.WILDCARD), false)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TipoContext extends ParserRuleContext {
		public environment.TipoExpresion rtipo;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterTipo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitTipo(this);
		}
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_tipo);
		try {
			setState(238);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(224);
				match(INT);
				_localctx.rtipo = 0
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(226);
				match(FLOAT);
				_localctx.rtipo = 1
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(228);
				match(STR);
				_localctx.rtipo = 2
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(230);
				match(BOOL);
				_localctx.rtipo = 3
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(232);
				match(CHARACTER);
				_localctx.rtipo = 5
				}
				break;
			case CORCHIZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(234);
				match(CORCHIZQ);
				setState(235);
				match(INT);
				setState(236);
				match(CORCHDER);
				_localctx.rtipo = 7
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarasgmtContext extends ParserRuleContext {
		public interfaces.Instruction asgmt;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public VarasgmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varasgmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterVarasgmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitVarasgmt(this);
		}
	}

	public final VarasgmtContext varasgmt() throws RecognitionException {
		VarasgmtContext _localctx = new VarasgmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_varasgmt);
		try {
			setState(257);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(240);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(241);
				match(IG);
				setState(242);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(246);
				match(ADD);
				setState(247);
				match(IG);
				setState(248);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewOpAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e, "+")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(251);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(252);
				match(SUB);
				setState(253);
				match(IG);
				setState(254);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewOpAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e, "-")
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConststmtContext extends ParserRuleContext {
		public interfaces.Instruction const_;
		public Token LET;
		public Token ID;
		public ExprContext expr;
		public TipoContext tipo;
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public ConststmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conststmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterConststmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitConststmt(this);
		}
	}

	public final ConststmtContext conststmt() throws RecognitionException {
		ConststmtContext _localctx = new ConststmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_conststmt);
		try {
			setState(273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(259);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(260);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(261);
				match(IG);
				setState(262);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const_ = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), environment.NULL, ((ConststmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(265);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(266);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(267);
				match(COLON);
				setState(268);
				((ConststmtContext)_localctx).tipo = tipo();
				setState(269);
				match(IG);
				setState(270);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const_ = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), ((ConststmtContext)_localctx).tipo.rtipo, ((ConststmtContext)_localctx).expr.e, true)
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchstmtContext extends ParserRuleContext {
		public interfaces.Instruction switchinstr;
		public Token SWITCH;
		public ExprContext expr;
		public CasestmtContext casestmt;
		public List<CasestmtContext> casesvar = new ArrayList<CasestmtContext>();
		public TerminalNode SWITCH() { return getToken(SwiftGrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public DefaultstmtContext defaultstmt() {
			return getRuleContext(DefaultstmtContext.class,0);
		}
		public List<CasestmtContext> casestmt() {
			return getRuleContexts(CasestmtContext.class);
		}
		public CasestmtContext casestmt(int i) {
			return getRuleContext(CasestmtContext.class,i);
		}
		public SwitchstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterSwitchstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitSwitchstmt(this);
		}
	}

	public final SwitchstmtContext switchstmt() throws RecognitionException {
		SwitchstmtContext _localctx = new SwitchstmtContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_switchstmt);

		    var switchInterfaces = []interface{}{}
		    var interfacelist []ICasestmtContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(275);
			((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
			setState(276);
			((SwitchstmtContext)_localctx).expr = expr(0);
			setState(277);
			match(LLAVEIZQ);
			setState(279); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(278);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				((SwitchstmtContext)_localctx).casesvar.add(((SwitchstmtContext)_localctx).casestmt);
				}
				}
				setState(281); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(283);
				defaultstmt();
				}
			}

			setState(286);
			match(LLAVEDER);
			 
			    interfacelist = localctx.(*SwitchstmtContext).GetCasesvar()
			    for _, e := range interfacelist {
			        switchInterfaces = append(switchInterfaces, e.GetCaseinstr())
			    }
			    _localctx.switchinstr = instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, switchInterfaces, nil)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CasestmtContext extends ParserRuleContext {
		public interfaces.Instruction caseinstr;
		public Token CASE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode CASE() { return getToken(SwiftGrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public CasestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_casestmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterCasestmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitCasestmt(this);
		}
	}

	public final CasestmtContext casestmt() throws RecognitionException {
		CasestmtContext _localctx = new CasestmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_casestmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			((CasestmtContext)_localctx).CASE = match(CASE);
			setState(290);
			((CasestmtContext)_localctx).expr = expr(0);
			setState(291);
			match(COLON);
			setState(292);
			((CasestmtContext)_localctx).block = block();
			setState(294);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(293);
				match(BREAK);
				}
			}

			 
			    _localctx.caseinstr = instructions.NewSwitch((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0), (((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e, ((CasestmtContext)_localctx).block.blk, nil) 

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefaultstmtContext extends ParserRuleContext {
		public []interface{} definstr;
		public BlockContext block;
		public TerminalNode DEFAULT() { return getToken(SwiftGrammarParser.DEFAULT, 0); }
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public DefaultstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterDefaultstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitDefaultstmt(this);
		}
	}

	public final DefaultstmtContext defaultstmt() throws RecognitionException {
		DefaultstmtContext _localctx = new DefaultstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_defaultstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			match(DEFAULT);
			setState(299);
			match(COLON);
			setState(300);
			((DefaultstmtContext)_localctx).block = block();
			setState(302);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(301);
				match(BREAK);
				}
			}

			 _localctx.definstr = ((DefaultstmtContext)_localctx).block.blk 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhilestmtContext extends ParserRuleContext {
		public interfaces.Instruction whileinstr;
		public Token WHILE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterWhilestmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitWhilestmt(this);
		}
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(307);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(308);
			match(LLAVEIZQ);
			setState(309);
			((WhilestmtContext)_localctx).block = block();
			setState(310);
			match(LLAVEDER);
			 _localctx.whileinstr = instructions.NewWhile((((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getLine():0), (((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getCharPositionInLine():0), ((WhilestmtContext)_localctx).expr.e, ((WhilestmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForstmtContext extends ParserRuleContext {
		public interfaces.Instruction forinstr;
		public Token FOR;
		public Token ID;
		public Token STRING;
		public BlockContext block;
		public ExprContext left;
		public ExprContext right;
		public Token first;
		public Token second;
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public TerminalNode RANGEPTS() { return getToken(SwiftGrammarParser.RANGEPTS, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterForstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitForstmt(this);
		}
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_forstmt);

		    var cadena interfaces.Expression
		    var str string

		try {
			setState(342);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(313);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(314);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(315);
				match(IN);
				setState(316);
				((ForstmtContext)_localctx).STRING = match(STRING);
				setState(317);
				match(LLAVEIZQ);
				setState(318);
				((ForstmtContext)_localctx).block = block();
				setState(319);
				match(LLAVEDER);
				 
				    str = (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getText():null)
				    cadena = expressions.NewPrimitive((((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getLine():0), (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1], environment.STRING)
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), cadena, nil, nil, "", ((ForstmtContext)_localctx).block.blk) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(322);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(323);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(324);
				match(IN);
				setState(325);
				((ForstmtContext)_localctx).left = expr(0);
				setState(326);
				match(RANGEPTS);
				setState(327);
				((ForstmtContext)_localctx).right = expr(0);
				setState(328);
				match(LLAVEIZQ);
				setState(329);
				((ForstmtContext)_localctx).block = block();
				setState(330);
				match(LLAVEDER);
				 
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), nil, ((ForstmtContext)_localctx).left.e, ((ForstmtContext)_localctx).right.e, "", ((ForstmtContext)_localctx).block.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(333);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(334);
				((ForstmtContext)_localctx).first = match(ID);
				setState(335);
				match(IN);
				setState(336);
				((ForstmtContext)_localctx).second = match(ID);
				setState(337);
				match(LLAVEIZQ);
				setState(338);
				((ForstmtContext)_localctx).block = block();
				setState(339);
				match(LLAVEDER);

				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).first!=null?((ForstmtContext)_localctx).first.getText():null), nil, nil, nil, (((ForstmtContext)_localctx).second!=null?((ForstmtContext)_localctx).second.getText():null), ((ForstmtContext)_localctx).block.blk)

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GuardstmtContext extends ParserRuleContext {
		public interfaces.Instruction guardinstr;
		public Token GUARD;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterGuardstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitGuardstmt(this);
		}
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(345);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(346);
			match(ELSE);
			setState(347);
			match(LLAVEIZQ);
			setState(348);
			((GuardstmtContext)_localctx).block = block();
			setState(349);
			match(LLAVEDER);
			 _localctx.guardinstr = instructions.NewGuard((((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getLine():0), (((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getCharPositionInLine():0), ((GuardstmtContext)_localctx).expr.e, ((GuardstmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TransferstmtContext extends ParserRuleContext {
		public interfaces.Instruction trns;
		public Token BREAK;
		public Token CONTINUE;
		public Token RETURN;
		public ExprContext expr;
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public TerminalNode RETURN() { return getToken(SwiftGrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TransferstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transferstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterTransferstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitTransferstmt(this);
		}
	}

	public final TransferstmtContext transferstmt() throws RecognitionException {
		TransferstmtContext _localctx = new TransferstmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_transferstmt);
		try {
			setState(362);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(352);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				 _localctx.trns = instructions.NewBreak((((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getLine():0), (((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getCharPositionInLine():0)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(354);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				 _localctx.trns = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(356);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				 _localctx.trns = instructions.NewReturn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(358);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(359);
				((TransferstmtContext)_localctx).expr = expr(0);
				 _localctx.trns = instructions.NewReturn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), ((TransferstmtContext)_localctx).expr.e) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VectorstmtContext extends ParserRuleContext {
		public interfaces.Instruction vectorinstr;
		public Token VAR;
		public Token ID;
		public TipoContext tipo;
		public DefinestmtContext definestmt;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public DefinestmtContext definestmt() {
			return getRuleContext(DefinestmtContext.class,0);
		}
		public VectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterVectorstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitVectorstmt(this);
		}
	}

	public final VectorstmtContext vectorstmt() throws RecognitionException {
		VectorstmtContext _localctx = new VectorstmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_vectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(364);
			((VectorstmtContext)_localctx).VAR = match(VAR);
			setState(365);
			((VectorstmtContext)_localctx).ID = match(ID);
			setState(366);
			match(COLON);
			setState(367);
			match(CORCHIZQ);
			setState(368);
			((VectorstmtContext)_localctx).tipo = tipo();
			setState(369);
			match(CORCHDER);
			setState(370);
			((VectorstmtContext)_localctx).definestmt = definestmt();
			 
			    _localctx.vectorinstr = instructions.NewVectorStmt((((VectorstmtContext)_localctx).VAR!=null?((VectorstmtContext)_localctx).VAR.getLine():0), (((VectorstmtContext)_localctx).VAR!=null?((VectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VectorstmtContext)_localctx).ID!=null?((VectorstmtContext)_localctx).ID.getText():null), ((VectorstmtContext)_localctx).tipo.rtipo, ((VectorstmtContext)_localctx).definestmt.defineinstr)

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefinestmtContext extends ParserRuleContext {
		public []interface{} defineinstr;
		public ListexprContext listexpr;
		public List<ListexprContext> lista = new ArrayList<ListexprContext>();
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public List<ListexprContext> listexpr() {
			return getRuleContexts(ListexprContext.class);
		}
		public ListexprContext listexpr(int i) {
			return getRuleContext(ListexprContext.class,i);
		}
		public DefinestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definestmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterDefinestmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitDefinestmt(this);
		}
	}

	public final DefinestmtContext definestmt() throws RecognitionException {
		DefinestmtContext _localctx = new DefinestmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_definestmt);

		    var defVecInterfaces []interface{}

		int _la;
		try {
			setState(387);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(373);
				match(IG);
				setState(374);
				match(CORCHIZQ);
				setState(376); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(375);
					((DefinestmtContext)_localctx).listexpr = listexpr();
					((DefinestmtContext)_localctx).lista.add(((DefinestmtContext)_localctx).listexpr);
					}
					}
					setState(378); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 54048624367108566L) != 0) );
				setState(380);
				match(CORCHDER);
				 
				    for _, e := range localctx.(*DefinestmtContext).GetLista() {
				        // fmt.Println(fmt.Sprintf("%T", e.GetListe()))
				        defVecInterfaces = append(defVecInterfaces, e.GetListe())
				    }
				    _localctx.defineinstr = defVecInterfaces

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(383);
				match(IG);
				setState(384);
				match(CORCHIZQ);
				setState(385);
				match(CORCHDER);
				 _localctx.defineinstr = nil 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListexprContext extends ParserRuleContext {
		public interfaces.Expression liste;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListexprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listexpr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterListexpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitListexpr(this);
		}
	}

	public final ListexprContext listexpr() throws RecognitionException {
		ListexprContext _localctx = new ListexprContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_listexpr);
		try {
			setState(396);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(389);
				((ListexprContext)_localctx).expr = expr(0);
				setState(390);
				match(COMA);
				 _localctx.liste = ((ListexprContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(393);
				((ListexprContext)_localctx).expr = expr(0);
				 _localctx.liste = ((ListexprContext)_localctx).expr.e 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodvecContext extends ParserRuleContext {
		public interfaces.Instruction methodinstr;
		public Token ID;
		public Token APPEND;
		public ExprContext expr;
		public Token REMOVELAST;
		public Token REMOVE;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(SwiftGrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode REMOVELAST() { return getToken(SwiftGrammarParser.REMOVELAST, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftGrammarParser.REMOVE, 0); }
		public MethodvecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodvec; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterMethodvec(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitMethodvec(this);
		}
	}

	public final MethodvecContext methodvec() throws RecognitionException {
		MethodvecContext _localctx = new MethodvecContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_methodvec);
		try {
			setState(420);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(398);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(399);
				match(PUNTO);
				setState(400);
				((MethodvecContext)_localctx).APPEND = match(APPEND);
				setState(401);
				match(PARIZQ);
				setState(402);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(403);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), ((MethodvecContext)_localctx).expr.e, (((MethodvecContext)_localctx).APPEND!=null?((MethodvecContext)_localctx).APPEND.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(406);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(407);
				match(PUNTO);
				setState(408);
				((MethodvecContext)_localctx).REMOVELAST = match(REMOVELAST);
				setState(409);
				match(PARIZQ);
				setState(410);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), nil, (((MethodvecContext)_localctx).REMOVELAST!=null?((MethodvecContext)_localctx).REMOVELAST.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(412);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(413);
				match(PUNTO);
				setState(414);
				((MethodvecContext)_localctx).REMOVE = match(REMOVE);
				setState(415);
				match(PARIZQ);
				setState(416);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(417);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), ((MethodvecContext)_localctx).expr.e, (((MethodvecContext)_localctx).REMOVE!=null?((MethodvecContext)_localctx).REMOVE.getText():null)) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodvecrtrnContext extends ParserRuleContext {
		public interfaces.Expression methodinstrtrn;
		public Token ID;
		public Token EMPTY;
		public Token COUNT;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode EMPTY() { return getToken(SwiftGrammarParser.EMPTY, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public TerminalNode CORCHIZQ() { return getToken(SwiftGrammarParser.CORCHIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORCHDER() { return getToken(SwiftGrammarParser.CORCHDER, 0); }
		public MethodvecrtrnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodvecrtrn; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterMethodvecrtrn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitMethodvecrtrn(this);
		}
	}

	public final MethodvecrtrnContext methodvecrtrn() throws RecognitionException {
		MethodvecrtrnContext _localctx = new MethodvecrtrnContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_methodvecrtrn);
		try {
			setState(436);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(422);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(423);
				match(PUNTO);
				setState(424);
				((MethodvecrtrnContext)_localctx).EMPTY = match(EMPTY);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).EMPTY!=null?((MethodvecrtrnContext)_localctx).EMPTY.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(426);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(427);
				match(PUNTO);
				setState(428);
				((MethodvecrtrnContext)_localctx).COUNT = match(COUNT);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).COUNT!=null?((MethodvecrtrnContext)_localctx).COUNT.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(430);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(431);
				match(CORCHIZQ);
				setState(432);
				((MethodvecrtrnContext)_localctx).expr = expr(0);
				setState(433);
				match(CORCHDER);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), ((MethodvecrtrnContext)_localctx).expr.e, "access") 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VecaccessContext extends ParserRuleContext {
		public interfaces.Instruction vecacc;
		public Token firstId;
		public ExprContext first;
		public Token secondId;
		public ExprContext second;
		public Token ID;
		public List<TerminalNode> CORCHIZQ() { return getTokens(SwiftGrammarParser.CORCHIZQ); }
		public TerminalNode CORCHIZQ(int i) {
			return getToken(SwiftGrammarParser.CORCHIZQ, i);
		}
		public List<TerminalNode> CORCHDER() { return getTokens(SwiftGrammarParser.CORCHDER); }
		public TerminalNode CORCHDER(int i) {
			return getToken(SwiftGrammarParser.CORCHDER, i);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public VecaccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vecaccess; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterVecaccess(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitVecaccess(this);
		}
	}

	public final VecaccessContext vecaccess() throws RecognitionException {
		VecaccessContext _localctx = new VecaccessContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_vecaccess);
		try {
			setState(457);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(438);
				((VecaccessContext)_localctx).firstId = match(ID);
				setState(439);
				match(CORCHIZQ);
				setState(440);
				((VecaccessContext)_localctx).first = expr(0);
				setState(441);
				match(CORCHDER);
				setState(442);
				match(IG);
				setState(443);
				((VecaccessContext)_localctx).secondId = match(ID);
				setState(444);
				match(CORCHIZQ);
				setState(445);
				((VecaccessContext)_localctx).second = expr(0);
				setState(446);
				match(CORCHDER);
				 
				    _localctx.vecacc = instructions.NewVectorAsgmt((((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getCharPositionInLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getText():null), ((VecaccessContext)_localctx).first.e, ((VecaccessContext)_localctx).second.e, (((VecaccessContext)_localctx).secondId!=null?((VecaccessContext)_localctx).secondId.getText():null)) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(449);
				((VecaccessContext)_localctx).ID = match(ID);
				setState(450);
				match(CORCHIZQ);
				setState(451);
				((VecaccessContext)_localctx).first = expr(0);
				setState(452);
				match(CORCHDER);
				setState(453);
				match(IG);
				setState(454);
				((VecaccessContext)_localctx).second = expr(0);
				 _localctx.vecacc = instructions.NewVectorAsgmt((((VecaccessContext)_localctx).ID!=null?((VecaccessContext)_localctx).ID.getLine():0), (((VecaccessContext)_localctx).ID!=null?((VecaccessContext)_localctx).ID.getCharPositionInLine():0), (((VecaccessContext)_localctx).ID!=null?((VecaccessContext)_localctx).ID.getText():null), ((VecaccessContext)_localctx).first.e, ((VecaccessContext)_localctx).second.e, "") 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructstmtContext extends ParserRuleContext {
		public StructattrlistContext structattrlist;
		public List<StructattrlistContext> list = new ArrayList<StructattrlistContext>();
		public TerminalNode STRUCT() { return getToken(SwiftGrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<StructattrlistContext> structattrlist() {
			return getRuleContexts(StructattrlistContext.class);
		}
		public StructattrlistContext structattrlist(int i) {
			return getRuleContext(StructattrlistContext.class,i);
		}
		public StructstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterStructstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitStructstmt(this);
		}
	}

	public final StructstmtContext structstmt() throws RecognitionException {
		StructstmtContext _localctx = new StructstmtContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_structstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(459);
			match(STRUCT);
			setState(460);
			match(ID);
			setState(461);
			match(LLAVEIZQ);
			setState(463); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(462);
				((StructstmtContext)_localctx).structattrlist = structattrlist();
				((StructstmtContext)_localctx).list.add(((StructstmtContext)_localctx).structattrlist);
				}
				}
				setState(465); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==VAR || _la==LET );
			setState(467);
			match(LLAVEDER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructattrlistContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StructattrlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structattrlist; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterStructattrlist(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitStructattrlist(this);
		}
	}

	public final StructattrlistContext structattrlist() throws RecognitionException {
		StructattrlistContext _localctx = new StructattrlistContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_structattrlist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(469);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(470);
			match(ID);
			setState(472);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 576460752303423550L) != 0)) {
				{
				setState(471);
				tipo();
				}
			}

			setState(476);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IG) {
				{
				setState(474);
				match(IG);
				setState(475);
				expr(0);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncstmtContext extends ParserRuleContext {
		public interfaces.Instruction funcinstr;
		public Token FUNC;
		public Token ID;
		public ListparamContext listparam;
		public List<ListparamContext> lista = new ArrayList<ListparamContext>();
		public TipoContext tipo;
		public BlockContext block;
		public TerminalNode FUNC() { return getToken(SwiftGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode ARROW() { return getToken(SwiftGrammarParser.ARROW, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<ListparamContext> listparam() {
			return getRuleContexts(ListparamContext.class);
		}
		public ListparamContext listparam(int i) {
			return getRuleContext(ListparamContext.class,i);
		}
		public FuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcstmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterFuncstmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitFuncstmt(this);
		}
	}

	public final FuncstmtContext funcstmt() throws RecognitionException {
		FuncstmtContext _localctx = new FuncstmtContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_funcstmt);

		    var args []interface{}

		int _la;
		try {
			setState(528);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(478);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(479);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(480);
				match(PARIZQ);
				setState(482); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(481);
					((FuncstmtContext)_localctx).listparam = listparam();
					((FuncstmtContext)_localctx).lista.add(((FuncstmtContext)_localctx).listparam);
					}
					}
					setState(484); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==CAME || _la==ID );
				setState(486);
				match(PARDER);
				setState(487);
				match(ARROW);
				setState(488);
				((FuncstmtContext)_localctx).tipo = tipo();
				setState(489);
				match(LLAVEIZQ);
				setState(490);
				((FuncstmtContext)_localctx).block = block();
				setState(491);
				match(LLAVEDER);

				    for _, e := range localctx.(*FuncstmtContext).GetLista() {
				        // fmt.Println(fmt.Sprintf("%T", e.GetListparaminstr()))
				        args = append(args, e.GetListparaminstr())
				    }
				    _localctx.funcinstr = instructions.NewFunction((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), args, ((FuncstmtContext)_localctx).tipo.rtipo, ((FuncstmtContext)_localctx).block.blk)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(494);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(495);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(496);
				match(PARIZQ);
				setState(498); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(497);
					((FuncstmtContext)_localctx).listparam = listparam();
					((FuncstmtContext)_localctx).lista.add(((FuncstmtContext)_localctx).listparam);
					}
					}
					setState(500); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==CAME || _la==ID );
				setState(502);
				match(PARDER);
				setState(503);
				match(LLAVEIZQ);
				setState(504);
				((FuncstmtContext)_localctx).block = block();
				setState(505);
				match(LLAVEDER);

				    for _, e := range localctx.(*FuncstmtContext).GetLista() {
				        // fmt.Println(fmt.Sprintf("%T", e.GetListparaminstr()))
				        args = append(args, e.GetListparaminstr())
				    }
				    _localctx.funcinstr = instructions.NewFunction((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), args, 4, ((FuncstmtContext)_localctx).block.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(508);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(509);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(510);
				match(PARIZQ);
				setState(511);
				match(PARDER);
				setState(512);
				match(ARROW);
				setState(513);
				((FuncstmtContext)_localctx).tipo = tipo();
				setState(514);
				match(LLAVEIZQ);
				setState(515);
				((FuncstmtContext)_localctx).block = block();
				setState(516);
				match(LLAVEDER);
				 
				    _localctx.funcinstr = instructions.NewFunction((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), nil, ((FuncstmtContext)_localctx).tipo.rtipo, ((FuncstmtContext)_localctx).block.blk) 

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(519);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(520);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(521);
				match(PARIZQ);
				setState(522);
				match(PARDER);
				setState(523);
				match(LLAVEIZQ);
				setState(524);
				((FuncstmtContext)_localctx).block = block();
				setState(525);
				match(LLAVEDER);
				 
				    _localctx.funcinstr = instructions.NewFunction((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), nil, 4, ((FuncstmtContext)_localctx).block.blk) 

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListparamContext extends ParserRuleContext {
		public interfaces.Instruction listparaminstr;
		public Token extr;
		public Token inter;
		public Token INOUT;
		public TipoContext tipo;
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode INOUT() { return getToken(SwiftGrammarParser.INOUT, 0); }
		public TerminalNode CAME() { return getToken(SwiftGrammarParser.CAME, 0); }
		public ListparamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listparam; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterListparam(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitListparam(this);
		}
	}

	public final ListparamContext listparam() throws RecognitionException {
		ListparamContext _localctx = new ListparamContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_listparam);
		int _la;
		try {
			setState(553);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(531);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
				case 1:
					{
					setState(530);
					((ListparamContext)_localctx).extr = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==CAME || _la==ID) ) {
						((ListparamContext)_localctx).extr = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				setState(533);
				((ListparamContext)_localctx).inter = match(ID);
				setState(534);
				match(COLON);
				setState(536);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(535);
					((ListparamContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(538);
				((ListparamContext)_localctx).tipo = tipo();
				setState(539);
				match(COMA);
				 
				    var flag bool
				    var aux string
				    if ((ListparamContext)_localctx).INOUT != nil {
				        flag = true
				    }else{
				        flag = false
				    }
				    if ((ListparamContext)_localctx).extr != nil {
				        aux = (((ListparamContext)_localctx).extr!=null?((ListparamContext)_localctx).extr.getText():null)
				    }
				    _localctx.listparaminstr = instructions.NewParam((((ListparamContext)_localctx).extr!=null?((ListparamContext)_localctx).extr.getLine():0), (((ListparamContext)_localctx).extr!=null?((ListparamContext)_localctx).extr.getCharPositionInLine():0), aux, (((ListparamContext)_localctx).inter!=null?((ListparamContext)_localctx).inter.getText():null), flag, ((ListparamContext)_localctx).tipo.rtipo) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(543);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
				case 1:
					{
					setState(542);
					((ListparamContext)_localctx).extr = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==CAME || _la==ID) ) {
						((ListparamContext)_localctx).extr = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					break;
				}
				setState(545);
				((ListparamContext)_localctx).inter = match(ID);
				setState(546);
				match(COLON);
				setState(548);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(547);
					((ListparamContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(550);
				((ListparamContext)_localctx).tipo = tipo();
				 
				    var flag bool
				    var aux string
				    if ((ListparamContext)_localctx).INOUT != nil {
				        flag = true
				    }else{
				        flag = false
				    }
				    if ((ListparamContext)_localctx).extr != nil {
				        aux = (((ListparamContext)_localctx).extr!=null?((ListparamContext)_localctx).extr.getText():null)
				    }
				    _localctx.listparaminstr = instructions.NewParam((((ListparamContext)_localctx).extr!=null?((ListparamContext)_localctx).extr.getLine():0), (((ListparamContext)_localctx).extr!=null?((ListparamContext)_localctx).extr.getCharPositionInLine():0), aux, (((ListparamContext)_localctx).inter!=null?((ListparamContext)_localctx).inter.getText():null), flag, ((ListparamContext)_localctx).tipo.rtipo) 

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CallfuncContext extends ParserRuleContext {
		public interfaces.Expression callfuncexpr;
		public Token ID;
		public ListparamcallContext listparamcall;
		public List<ListparamcallContext> lista = new ArrayList<ListparamcallContext>();
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public List<ListparamcallContext> listparamcall() {
			return getRuleContexts(ListparamcallContext.class);
		}
		public ListparamcallContext listparamcall(int i) {
			return getRuleContext(ListparamcallContext.class,i);
		}
		public CallfuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callfunc; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterCallfunc(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitCallfunc(this);
		}
	}

	public final CallfuncContext callfunc() throws RecognitionException {
		CallfuncContext _localctx = new CallfuncContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_callfunc);

		    var args []interface{}

		int _la;
		try {
			setState(569);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(555);
				((CallfuncContext)_localctx).ID = match(ID);
				setState(556);
				match(PARIZQ);
				setState(558); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(557);
					((CallfuncContext)_localctx).listparamcall = listparamcall();
					((CallfuncContext)_localctx).lista.add(((CallfuncContext)_localctx).listparamcall);
					}
					}
					setState(560); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 54048624367108566L) != 0) || _la==AMP );
				setState(562);
				match(PARDER);

				    for _, e := range localctx.(*CallfuncContext).GetLista() {
				        // fmt.Println(fmt.Sprintf("%T", e.GetListparamcallinstr()))
				        args = append(args, e.GetListparamcallinstr())
				    }
				    _localctx.callfuncexpr = expressions.NewCallFunc((((CallfuncContext)_localctx).ID!=null?((CallfuncContext)_localctx).ID.getLine():0), (((CallfuncContext)_localctx).ID!=null?((CallfuncContext)_localctx).ID.getCharPositionInLine():0), (((CallfuncContext)_localctx).ID!=null?((CallfuncContext)_localctx).ID.getText():null), args)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(565);
				((CallfuncContext)_localctx).ID = match(ID);
				setState(566);
				match(PARIZQ);
				setState(567);
				match(PARDER);
				 _localctx.callfuncexpr = expressions.NewCallFunc((((CallfuncContext)_localctx).ID!=null?((CallfuncContext)_localctx).ID.getLine():0), (((CallfuncContext)_localctx).ID!=null?((CallfuncContext)_localctx).ID.getCharPositionInLine():0), (((CallfuncContext)_localctx).ID!=null?((CallfuncContext)_localctx).ID.getText():null), nil) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CallfuncinstructionContext extends ParserRuleContext {
		public interfaces.Instruction callfuncinstr;
		public Token ID;
		public ListparamcallContext listparamcall;
		public List<ListparamcallContext> lista = new ArrayList<ListparamcallContext>();
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public List<ListparamcallContext> listparamcall() {
			return getRuleContexts(ListparamcallContext.class);
		}
		public ListparamcallContext listparamcall(int i) {
			return getRuleContext(ListparamcallContext.class,i);
		}
		public CallfuncinstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callfuncinstruction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterCallfuncinstruction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitCallfuncinstruction(this);
		}
	}

	public final CallfuncinstructionContext callfuncinstruction() throws RecognitionException {
		CallfuncinstructionContext _localctx = new CallfuncinstructionContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_callfuncinstruction);

		    var args []interface{}

		int _la;
		try {
			setState(585);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(571);
				((CallfuncinstructionContext)_localctx).ID = match(ID);
				setState(572);
				match(PARIZQ);
				setState(574); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(573);
					((CallfuncinstructionContext)_localctx).listparamcall = listparamcall();
					((CallfuncinstructionContext)_localctx).lista.add(((CallfuncinstructionContext)_localctx).listparamcall);
					}
					}
					setState(576); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 54048624367108566L) != 0) || _la==AMP );
				setState(578);
				match(PARDER);

				    for _, e := range localctx.(*CallfuncinstructionContext).GetLista() {
				        // fmt.Println(fmt.Sprintf("%T", e.GetListparamcallinstr()))
				        args = append(args, e.GetListparamcallinstr())
				    }
				    _localctx.callfuncinstr = instructions.NewCallFunc((((CallfuncinstructionContext)_localctx).ID!=null?((CallfuncinstructionContext)_localctx).ID.getLine():0), (((CallfuncinstructionContext)_localctx).ID!=null?((CallfuncinstructionContext)_localctx).ID.getCharPositionInLine():0), (((CallfuncinstructionContext)_localctx).ID!=null?((CallfuncinstructionContext)_localctx).ID.getText():null), args)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(581);
				((CallfuncinstructionContext)_localctx).ID = match(ID);
				setState(582);
				match(PARIZQ);
				setState(583);
				match(PARDER);
				 _localctx.callfuncinstr = instructions.NewCallFunc((((CallfuncinstructionContext)_localctx).ID!=null?((CallfuncinstructionContext)_localctx).ID.getLine():0), (((CallfuncinstructionContext)_localctx).ID!=null?((CallfuncinstructionContext)_localctx).ID.getCharPositionInLine():0), (((CallfuncinstructionContext)_localctx).ID!=null?((CallfuncinstructionContext)_localctx).ID.getText():null), nil) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListparamcallContext extends ParserRuleContext {
		public interfaces.Expression listparamcallinstr;
		public Token ID;
		public Token AMP;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(SwiftGrammarParser.COLON, 0); }
		public TerminalNode AMP() { return getToken(SwiftGrammarParser.AMP, 0); }
		public ListparamcallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listparamcall; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterListparamcall(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitListparamcall(this);
		}
	}

	public final ListparamcallContext listparamcall() throws RecognitionException {
		ListparamcallContext _localctx = new ListparamcallContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_listparamcall);
		int _la;
		try {
			setState(608);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(589);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
				case 1:
					{
					setState(587);
					((ListparamcallContext)_localctx).ID = match(ID);
					setState(588);
					match(COLON);
					}
					break;
				}
				setState(592);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMP) {
					{
					setState(591);
					((ListparamcallContext)_localctx).AMP = match(AMP);
					}
				}

				setState(594);
				((ListparamcallContext)_localctx).expr = expr(0);
				setState(595);
				match(COMA);
				 
				    var flag bool
				    if ((ListparamcallContext)_localctx).AMP != nil {
				        flag = true
				    }
				    _localctx.listparamcallinstr = expressions.NewCallParams((((ListparamcallContext)_localctx).expr!=null?(((ListparamcallContext)_localctx).expr.start):null).GetLine(), (((ListparamcallContext)_localctx).expr!=null?(((ListparamcallContext)_localctx).expr.start):null).GetColumn(), (((ListparamcallContext)_localctx).ID!=null?((ListparamcallContext)_localctx).ID.getText():null), flag, ((ListparamcallContext)_localctx).expr.e) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(600);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
				case 1:
					{
					setState(598);
					((ListparamcallContext)_localctx).ID = match(ID);
					setState(599);
					match(COLON);
					}
					break;
				}
				setState(603);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMP) {
					{
					setState(602);
					((ListparamcallContext)_localctx).AMP = match(AMP);
					}
				}

				setState(605);
				((ListparamcallContext)_localctx).expr = expr(0);
				 
				    var flag bool
				    if ((ListparamcallContext)_localctx).AMP != nil {
				        flag = true
				    }
				    _localctx.listparamcallinstr = expressions.NewCallParams((((ListparamcallContext)_localctx).expr!=null?(((ListparamcallContext)_localctx).expr.start):null).GetLine(), (((ListparamcallContext)_localctx).expr!=null?(((ListparamcallContext)_localctx).expr.start):null).GetColumn(), (((ListparamcallContext)_localctx).ID!=null?((ListparamcallContext)_localctx).ID.getText():null), flag, ((ListparamcallContext)_localctx).expr.e)
				 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncembedContext extends ParserRuleContext {
		public interfaces.Expression funcembedexpr;
		public Token STR;
		public ExprContext expr;
		public Token INT;
		public Token FLOAT;
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public FuncembedContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcembed; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterFuncembed(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitFuncembed(this);
		}
	}

	public final FuncembedContext funcembed() throws RecognitionException {
		FuncembedContext _localctx = new FuncembedContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_funcembed);
		try {
			setState(628);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STR:
				enterOuterAlt(_localctx, 1);
				{
				setState(610);
				((FuncembedContext)_localctx).STR = match(STR);
				setState(611);
				match(PARIZQ);
				setState(612);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(613);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewString((((FuncembedContext)_localctx).STR!=null?((FuncembedContext)_localctx).STR.getLine():0), (((FuncembedContext)_localctx).STR!=null?((FuncembedContext)_localctx).STR.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			case INT:
				enterOuterAlt(_localctx, 2);
				{
				setState(616);
				((FuncembedContext)_localctx).INT = match(INT);
				setState(617);
				match(PARIZQ);
				setState(618);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(619);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewInteger((((FuncembedContext)_localctx).INT!=null?((FuncembedContext)_localctx).INT.getLine():0), (((FuncembedContext)_localctx).INT!=null?((FuncembedContext)_localctx).INT.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(622);
				((FuncembedContext)_localctx).FLOAT = match(FLOAT);
				setState(623);
				match(PARIZQ);
				setState(624);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(625);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewFloat((((FuncembedContext)_localctx).FLOAT!=null?((FuncembedContext)_localctx).FLOAT.getLine():0), (((FuncembedContext)_localctx).FLOAT!=null?((FuncembedContext)_localctx).FLOAT.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public ExprContext expr;
		public Token NUMBER;
		public Token CHAR;
		public Token STRING;
		public Token TRU;
		public CallfuncContext callfunc;
		public FuncembedContext funcembed;
		public Token FAL;
		public Token ID;
		public MethodvecrtrnContext methodvecrtrn;
		public Token NIL;
		public Token AND;
		public Token OR;
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode NOT() { return getToken(SwiftGrammarParser.NOT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode CHAR() { return getToken(SwiftGrammarParser.CHAR, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public CallfuncContext callfunc() {
			return getRuleContext(CallfuncContext.class,0);
		}
		public FuncembedContext funcembed() {
			return getRuleContext(FuncembedContext.class,0);
		}
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public MethodvecrtrnContext methodvecrtrn() {
			return getRuleContext(MethodvecrtrnContext.class,0);
		}
		public TerminalNode NIL() { return getToken(SwiftGrammarParser.NIL, 0); }
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(SwiftGrammarParser.MOD, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).enterExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SwiftGrammarListener ) ((SwiftGrammarListener)listener).exitExpr(this);
		}
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 66;
		enterRecursionRule(_localctx, 66, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(667);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
			case 1:
				{
				setState(631);
				((ExprContext)_localctx).op = match(SUB);
				setState(632);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, "unario", expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), -1, environment.INTEGER)) 
				}
				break;
			case 2:
				{
				setState(635);
				((ExprContext)_localctx).op = match(NOT);
				setState(636);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), false, environment.BOOLEAN)) 
				}
				break;
			case 3:
				{
				setState(639);
				match(PARIZQ);
				setState(640);
				((ExprContext)_localctx).expr = expr(0);
				setState(641);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(644);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case 5:
				{
				setState(646);
				((ExprContext)_localctx).CHAR = match(CHAR);
				 
				        str := (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getLine():0), (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getCharPositionInLine():0), str[1:len(str)-1], environment.CHAR) 
				    
				}
				break;
			case 6:
				{
				setState(648);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 7:
				{
				setState(650);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(652);
				((ExprContext)_localctx).callfunc = callfunc();
				 _localctx.e = ((ExprContext)_localctx).callfunc.callfuncexpr 
				}
				break;
			case 9:
				{
				setState(655);
				((ExprContext)_localctx).funcembed = funcembed();
				 _localctx.e = ((ExprContext)_localctx).funcembed.funcembedexpr 
				}
				break;
			case 10:
				{
				setState(658);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 11:
				{
				setState(660);
				((ExprContext)_localctx).ID = match(ID);
				 _localctx.e = expressions.NewVar((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 12:
				{
				setState(662);
				((ExprContext)_localctx).methodvecrtrn = methodvecrtrn();
				 _localctx.e = ((ExprContext)_localctx).methodvecrtrn.methodinstrtrn 
				}
				break;
			case 13:
				{
				setState(665);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), "nil", environment.NULL) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(706);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,46,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(704);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(669);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(670);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7881299347898368L) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(671);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(674);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(675);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(676);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(679);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(680);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(681);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(684);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(685);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(686);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(689);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(690);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(691);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(694);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(695);
						((ExprContext)_localctx).AND = match(AND);
						setState(696);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).AND!=null?((ExprContext)_localctx).AND.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(699);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(700);
						((ExprContext)_localctx).OR = match(OR);
						setState(701);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).OR!=null?((ExprContext)_localctx).OR.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(708);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,46,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 33:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 18);
		case 1:
			return precpred(_ctx, 17);
		case 2:
			return precpred(_ctx, 16);
		case 3:
			return precpred(_ctx, 15);
		case 4:
			return precpred(_ctx, 14);
		case 5:
			return precpred(_ctx, 13);
		case 6:
			return precpred(_ctx, 12);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001E\u02c6\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0001\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0004\u0001J\b\u0001\u000b"+
		"\u0001\f\u0001K\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002}\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0004"+
		"\u0003\u0082\b\u0003\u000b\u0003\f\u0003\u0083\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004\u0090\b\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0004\u0005"+
		"\u009f\b\u0005\u000b\u0005\f\u0005\u00a0\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0004\u0005\u00ac\b\u0005\u000b\u0005\f\u0005\u00ad\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00ba\b\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00df\b\b\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0003\t\u00ef\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0003\n\u0102\b\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003"+
		"\u000b\u0112\b\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0004\f\u0118\b\f"+
		"\u000b\f\f\f\u0119\u0001\f\u0003\f\u011d\b\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u0127\b\r\u0001\r\u0001\r\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u012f\b\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u0157\b\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u016b\b\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0004"+
		"\u0014\u0179\b\u0014\u000b\u0014\f\u0014\u017a\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u0184\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0003\u0015\u018d\b\u0015\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0003\u0016\u01a5\b\u0016\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0003\u0017\u01b5\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u01ca\b\u0018\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u0019\u0004\u0019\u01d0\b\u0019\u000b\u0019"+
		"\f\u0019\u01d1\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0003\u001a\u01d9\b\u001a\u0001\u001a\u0001\u001a\u0003\u001a\u01dd\b"+
		"\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0004\u001b\u01e3"+
		"\b\u001b\u000b\u001b\f\u001b\u01e4\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0004\u001b\u01f3\b\u001b\u000b\u001b"+
		"\f\u001b\u01f4\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u0211\b\u001b\u0001\u001c"+
		"\u0003\u001c\u0214\b\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c"+
		"\u0219\b\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0003\u001c\u0220\b\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c"+
		"\u0225\b\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u022a\b"+
		"\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0004\u001d\u022f\b\u001d\u000b"+
		"\u001d\f\u001d\u0230\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0003\u001d\u023a\b\u001d\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0004\u001e\u023f\b\u001e\u000b\u001e\f\u001e\u0240"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0003\u001e\u024a\b\u001e\u0001\u001f\u0001\u001f\u0003\u001f"+
		"\u024e\b\u001f\u0001\u001f\u0003\u001f\u0251\b\u001f\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u0259"+
		"\b\u001f\u0001\u001f\u0003\u001f\u025c\b\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0003\u001f\u0261\b\u001f\u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0003 \u0275\b \u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0003!\u029c\b!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0005"+
		"!\u02c1\b!\n!\f!\u02c4\t!\u0001!\u0000\u0001B\"\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@B\u0000\u0007\u0001\u0000\t\n\u0002\u0000##\'\'\u0001\u00002"+
		"4\u0001\u000056\u0002\u0000..00\u0002\u0000//11\u0001\u0000()\u02ff\u0000"+
		"D\u0001\u0000\u0000\u0000\u0002I\u0001\u0000\u0000\u0000\u0004|\u0001"+
		"\u0000\u0000\u0000\u0006~\u0001\u0000\u0000\u0000\b\u008f\u0001\u0000"+
		"\u0000\u0000\n\u00b9\u0001\u0000\u0000\u0000\f\u00bb\u0001\u0000\u0000"+
		"\u0000\u000e\u00c3\u0001\u0000\u0000\u0000\u0010\u00de\u0001\u0000\u0000"+
		"\u0000\u0012\u00ee\u0001\u0000\u0000\u0000\u0014\u0101\u0001\u0000\u0000"+
		"\u0000\u0016\u0111\u0001\u0000\u0000\u0000\u0018\u0113\u0001\u0000\u0000"+
		"\u0000\u001a\u0121\u0001\u0000\u0000\u0000\u001c\u012a\u0001\u0000\u0000"+
		"\u0000\u001e\u0132\u0001\u0000\u0000\u0000 \u0156\u0001\u0000\u0000\u0000"+
		"\"\u0158\u0001\u0000\u0000\u0000$\u016a\u0001\u0000\u0000\u0000&\u016c"+
		"\u0001\u0000\u0000\u0000(\u0183\u0001\u0000\u0000\u0000*\u018c\u0001\u0000"+
		"\u0000\u0000,\u01a4\u0001\u0000\u0000\u0000.\u01b4\u0001\u0000\u0000\u0000"+
		"0\u01c9\u0001\u0000\u0000\u00002\u01cb\u0001\u0000\u0000\u00004\u01d5"+
		"\u0001\u0000\u0000\u00006\u0210\u0001\u0000\u0000\u00008\u0229\u0001\u0000"+
		"\u0000\u0000:\u0239\u0001\u0000\u0000\u0000<\u0249\u0001\u0000\u0000\u0000"+
		">\u0260\u0001\u0000\u0000\u0000@\u0274\u0001\u0000\u0000\u0000B\u029b"+
		"\u0001\u0000\u0000\u0000DE\u0003\u0002\u0001\u0000EF\u0005\u0000\u0000"+
		"\u0001FG\u0006\u0000\uffff\uffff\u0000G\u0001\u0001\u0000\u0000\u0000"+
		"HJ\u0003\u0004\u0002\u0000IH\u0001\u0000\u0000\u0000JK\u0001\u0000\u0000"+
		"\u0000KI\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LM\u0001\u0000"+
		"\u0000\u0000MN\u0006\u0001\uffff\uffff\u0000N\u0003\u0001\u0000\u0000"+
		"\u0000OP\u0003\u0006\u0003\u0000PQ\u0006\u0002\uffff\uffff\u0000Q}\u0001"+
		"\u0000\u0000\u0000RS\u0003<\u001e\u0000ST\u0006\u0002\uffff\uffff\u0000"+
		"T}\u0001\u0000\u0000\u0000UV\u0003\n\u0005\u0000VW\u0006\u0002\uffff\uffff"+
		"\u0000W}\u0001\u0000\u0000\u0000XY\u0003\u0010\b\u0000YZ\u0006\u0002\uffff"+
		"\uffff\u0000Z}\u0001\u0000\u0000\u0000[\\\u0003\u0014\n\u0000\\]\u0006"+
		"\u0002\uffff\uffff\u0000]}\u0001\u0000\u0000\u0000^_\u0003\u0016\u000b"+
		"\u0000_`\u0006\u0002\uffff\uffff\u0000`}\u0001\u0000\u0000\u0000ab\u0003"+
		"\u0018\f\u0000bc\u0006\u0002\uffff\uffff\u0000c}\u0001\u0000\u0000\u0000"+
		"de\u0003\u001e\u000f\u0000ef\u0006\u0002\uffff\uffff\u0000f}\u0001\u0000"+
		"\u0000\u0000gh\u0003 \u0010\u0000hi\u0006\u0002\uffff\uffff\u0000i}\u0001"+
		"\u0000\u0000\u0000jk\u0003\"\u0011\u0000kl\u0006\u0002\uffff\uffff\u0000"+
		"l}\u0001\u0000\u0000\u0000mn\u0003$\u0012\u0000no\u0006\u0002\uffff\uffff"+
		"\u0000o}\u0001\u0000\u0000\u0000pq\u0003&\u0013\u0000qr\u0006\u0002\uffff"+
		"\uffff\u0000r}\u0001\u0000\u0000\u0000st\u0003,\u0016\u0000tu\u0006\u0002"+
		"\uffff\uffff\u0000u}\u0001\u0000\u0000\u0000vw\u00030\u0018\u0000wx\u0006"+
		"\u0002\uffff\uffff\u0000x}\u0001\u0000\u0000\u0000yz\u00036\u001b\u0000"+
		"z{\u0006\u0002\uffff\uffff\u0000{}\u0001\u0000\u0000\u0000|O\u0001\u0000"+
		"\u0000\u0000|R\u0001\u0000\u0000\u0000|U\u0001\u0000\u0000\u0000|X\u0001"+
		"\u0000\u0000\u0000|[\u0001\u0000\u0000\u0000|^\u0001\u0000\u0000\u0000"+
		"|a\u0001\u0000\u0000\u0000|d\u0001\u0000\u0000\u0000|g\u0001\u0000\u0000"+
		"\u0000|j\u0001\u0000\u0000\u0000|m\u0001\u0000\u0000\u0000|p\u0001\u0000"+
		"\u0000\u0000|s\u0001\u0000\u0000\u0000|v\u0001\u0000\u0000\u0000|y\u0001"+
		"\u0000\u0000\u0000}\u0005\u0001\u0000\u0000\u0000~\u007f\u0005\u000b\u0000"+
		"\u0000\u007f\u0081\u00057\u0000\u0000\u0080\u0082\u0003\b\u0004\u0000"+
		"\u0081\u0080\u0001\u0000\u0000\u0000\u0082\u0083\u0001\u0000\u0000\u0000"+
		"\u0083\u0081\u0001\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000"+
		"\u0084\u0085\u0001\u0000\u0000\u0000\u0085\u0086\u00058\u0000\u0000\u0086"+
		"\u0087\u0006\u0003\uffff\uffff\u0000\u0087\u0007\u0001\u0000\u0000\u0000"+
		"\u0088\u0089\u0003B!\u0000\u0089\u008a\u0005?\u0000\u0000\u008a\u008b"+
		"\u0006\u0004\uffff\uffff\u0000\u008b\u0090\u0001\u0000\u0000\u0000\u008c"+
		"\u008d\u0003B!\u0000\u008d\u008e\u0006\u0004\uffff\uffff\u0000\u008e\u0090"+
		"\u0001\u0000\u0000\u0000\u008f\u0088\u0001\u0000\u0000\u0000\u008f\u008c"+
		"\u0001\u0000\u0000\u0000\u0090\t\u0001\u0000\u0000\u0000\u0091\u0092\u0005"+
		"\f\u0000\u0000\u0092\u0093\u0003B!\u0000\u0093\u0094\u00059\u0000\u0000"+
		"\u0094\u0095\u0003\u0002\u0001\u0000\u0095\u0096\u0005:\u0000\u0000\u0096"+
		"\u0097\u0006\u0005\uffff\uffff\u0000\u0097\u00ba\u0001\u0000\u0000\u0000"+
		"\u0098\u0099\u0005\f\u0000\u0000\u0099\u009a\u0003B!\u0000\u009a\u009b"+
		"\u00059\u0000\u0000\u009b\u009c\u0003\u0002\u0001\u0000\u009c\u009e\u0005"+
		":\u0000\u0000\u009d\u009f\u0003\f\u0006\u0000\u009e\u009d\u0001\u0000"+
		"\u0000\u0000\u009f\u00a0\u0001\u0000\u0000\u0000\u00a0\u009e\u0001\u0000"+
		"\u0000\u0000\u00a0\u00a1\u0001\u0000\u0000\u0000\u00a1\u00a2\u0001\u0000"+
		"\u0000\u0000\u00a2\u00a3\u0003\u000e\u0007\u0000\u00a3\u00a4\u0006\u0005"+
		"\uffff\uffff\u0000\u00a4\u00ba\u0001\u0000\u0000\u0000\u00a5\u00a6\u0005"+
		"\f\u0000\u0000\u00a6\u00a7\u0003B!\u0000\u00a7\u00a8\u00059\u0000\u0000"+
		"\u00a8\u00a9\u0003\u0002\u0001\u0000\u00a9\u00ab\u0005:\u0000\u0000\u00aa"+
		"\u00ac\u0003\f\u0006\u0000\u00ab\u00aa\u0001\u0000\u0000\u0000\u00ac\u00ad"+
		"\u0001\u0000\u0000\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae"+
		"\u0001\u0000\u0000\u0000\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b0"+
		"\u0006\u0005\uffff\uffff\u0000\u00b0\u00ba\u0001\u0000\u0000\u0000\u00b1"+
		"\u00b2\u0005\f\u0000\u0000\u00b2\u00b3\u0003B!\u0000\u00b3\u00b4\u0005"+
		"9\u0000\u0000\u00b4\u00b5\u0003\u0002\u0001\u0000\u00b5\u00b6\u0005:\u0000"+
		"\u0000\u00b6\u00b7\u0003\u000e\u0007\u0000\u00b7\u00b8\u0006\u0005\uffff"+
		"\uffff\u0000\u00b8\u00ba\u0001\u0000\u0000\u0000\u00b9\u0091\u0001\u0000"+
		"\u0000\u0000\u00b9\u0098\u0001\u0000\u0000\u0000\u00b9\u00a5\u0001\u0000"+
		"\u0000\u0000\u00b9\u00b1\u0001\u0000\u0000\u0000\u00ba\u000b\u0001\u0000"+
		"\u0000\u0000\u00bb\u00bc\u0005\r\u0000\u0000\u00bc\u00bd\u0005\f\u0000"+
		"\u0000\u00bd\u00be\u0003B!\u0000\u00be\u00bf\u00059\u0000\u0000\u00bf"+
		"\u00c0\u0003\u0002\u0001\u0000\u00c0\u00c1\u0005:\u0000\u0000\u00c1\u00c2"+
		"\u0006\u0006\uffff\uffff\u0000\u00c2\r\u0001\u0000\u0000\u0000\u00c3\u00c4"+
		"\u0005\r\u0000\u0000\u00c4\u00c5\u00059\u0000\u0000\u00c5\u00c6\u0003"+
		"\u0002\u0001\u0000\u00c6\u00c7\u0005:\u0000\u0000\u00c7\u00c8\u0006\u0007"+
		"\uffff\uffff\u0000\u00c8\u000f\u0001\u0000\u0000\u0000\u00c9\u00ca\u0005"+
		"\t\u0000\u0000\u00ca\u00cb\u0005\'\u0000\u0000\u00cb\u00cc\u0005A\u0000"+
		"\u0000\u00cc\u00cd\u0003\u0012\t\u0000\u00cd\u00ce\u0005-\u0000\u0000"+
		"\u00ce\u00cf\u0003B!\u0000\u00cf\u00d0\u0006\b\uffff\uffff\u0000\u00d0"+
		"\u00df\u0001\u0000\u0000\u0000\u00d1\u00d2\u0005\t\u0000\u0000\u00d2\u00d3"+
		"\u0005\'\u0000\u0000\u00d3\u00d4\u0005-\u0000\u0000\u00d4\u00d5\u0003"+
		"B!\u0000\u00d5\u00d6\u0006\b\uffff\uffff\u0000\u00d6\u00df\u0001\u0000"+
		"\u0000\u0000\u00d7\u00d8\u0005\t\u0000\u0000\u00d8\u00d9\u0005\'\u0000"+
		"\u0000\u00d9\u00da\u0005A\u0000\u0000\u00da\u00db\u0003\u0012\t\u0000"+
		"\u00db\u00dc\u0005=\u0000\u0000\u00dc\u00dd\u0006\b\uffff\uffff\u0000"+
		"\u00dd\u00df\u0001\u0000\u0000\u0000\u00de\u00c9\u0001\u0000\u0000\u0000"+
		"\u00de\u00d1\u0001\u0000\u0000\u0000\u00de\u00d7\u0001\u0000\u0000\u0000"+
		"\u00df\u0011\u0001\u0000\u0000\u0000\u00e0\u00e1\u0005\u0001\u0000\u0000"+
		"\u00e1\u00ef\u0006\t\uffff\uffff\u0000\u00e2\u00e3\u0005\u0002\u0000\u0000"+
		"\u00e3\u00ef\u0006\t\uffff\uffff\u0000\u00e4\u00e5\u0005\u0004\u0000\u0000"+
		"\u00e5\u00ef\u0006\t\uffff\uffff\u0000\u00e6\u00e7\u0005\u0003\u0000\u0000"+
		"\u00e7\u00ef\u0006\t\uffff\uffff\u0000\u00e8\u00e9\u0005\u0005\u0000\u0000"+
		"\u00e9\u00ef\u0006\t\uffff\uffff\u0000\u00ea\u00eb\u0005;\u0000\u0000"+
		"\u00eb\u00ec\u0005\u0001\u0000\u0000\u00ec\u00ed\u0005<\u0000\u0000\u00ed"+
		"\u00ef\u0006\t\uffff\uffff\u0000\u00ee\u00e0\u0001\u0000\u0000\u0000\u00ee"+
		"\u00e2\u0001\u0000\u0000\u0000\u00ee\u00e4\u0001\u0000\u0000\u0000\u00ee"+
		"\u00e6\u0001\u0000\u0000\u0000\u00ee\u00e8\u0001\u0000\u0000\u0000\u00ee"+
		"\u00ea\u0001\u0000\u0000\u0000\u00ef\u0013\u0001\u0000\u0000\u0000\u00f0"+
		"\u00f1\u0005\'\u0000\u0000\u00f1\u00f2\u0005-\u0000\u0000\u00f2\u00f3"+
		"\u0003B!\u0000\u00f3\u00f4\u0006\n\uffff\uffff\u0000\u00f4\u0102\u0001"+
		"\u0000\u0000\u0000\u00f5\u00f6\u0005\'\u0000\u0000\u00f6\u00f7\u00055"+
		"\u0000\u0000\u00f7\u00f8\u0005-\u0000\u0000\u00f8\u00f9\u0003B!\u0000"+
		"\u00f9\u00fa\u0006\n\uffff\uffff\u0000\u00fa\u0102\u0001\u0000\u0000\u0000"+
		"\u00fb\u00fc\u0005\'\u0000\u0000\u00fc\u00fd\u00056\u0000\u0000\u00fd"+
		"\u00fe\u0005-\u0000\u0000\u00fe\u00ff\u0003B!\u0000\u00ff\u0100\u0006"+
		"\n\uffff\uffff\u0000\u0100\u0102\u0001\u0000\u0000\u0000\u0101\u00f0\u0001"+
		"\u0000\u0000\u0000\u0101\u00f5\u0001\u0000\u0000\u0000\u0101\u00fb\u0001"+
		"\u0000\u0000\u0000\u0102\u0015\u0001\u0000\u0000\u0000\u0103\u0104\u0005"+
		"\n\u0000\u0000\u0104\u0105\u0005\'\u0000\u0000\u0105\u0106\u0005-\u0000"+
		"\u0000\u0106\u0107\u0003B!\u0000\u0107\u0108\u0006\u000b\uffff\uffff\u0000"+
		"\u0108\u0112\u0001\u0000\u0000\u0000\u0109\u010a\u0005\n\u0000\u0000\u010a"+
		"\u010b\u0005\'\u0000\u0000\u010b\u010c\u0005A\u0000\u0000\u010c\u010d"+
		"\u0003\u0012\t\u0000\u010d\u010e\u0005-\u0000\u0000\u010e\u010f\u0003"+
		"B!\u0000\u010f\u0110\u0006\u000b\uffff\uffff\u0000\u0110\u0112\u0001\u0000"+
		"\u0000\u0000\u0111\u0103\u0001\u0000\u0000\u0000\u0111\u0109\u0001\u0000"+
		"\u0000\u0000\u0112\u0017\u0001\u0000\u0000\u0000\u0113\u0114\u0005\u000e"+
		"\u0000\u0000\u0114\u0115\u0003B!\u0000\u0115\u0117\u00059\u0000\u0000"+
		"\u0116\u0118\u0003\u001a\r\u0000\u0117\u0116\u0001\u0000\u0000\u0000\u0118"+
		"\u0119\u0001\u0000\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u0119"+
		"\u011a\u0001\u0000\u0000\u0000\u011a\u011c\u0001\u0000\u0000\u0000\u011b"+
		"\u011d\u0003\u001c\u000e\u0000\u011c\u011b\u0001\u0000\u0000\u0000\u011c"+
		"\u011d\u0001\u0000\u0000\u0000\u011d\u011e\u0001\u0000\u0000\u0000\u011e"+
		"\u011f\u0005:\u0000\u0000\u011f\u0120\u0006\f\uffff\uffff\u0000\u0120"+
		"\u0019\u0001\u0000\u0000\u0000\u0121\u0122\u0005\u000f\u0000\u0000\u0122"+
		"\u0123\u0003B!\u0000\u0123\u0124\u0005A\u0000\u0000\u0124\u0126\u0003"+
		"\u0002\u0001\u0000\u0125\u0127\u0005\u0016\u0000\u0000\u0126\u0125\u0001"+
		"\u0000\u0000\u0000\u0126\u0127\u0001\u0000\u0000\u0000\u0127\u0128\u0001"+
		"\u0000\u0000\u0000\u0128\u0129\u0006\r\uffff\uffff\u0000\u0129\u001b\u0001"+
		"\u0000\u0000\u0000\u012a\u012b\u0005\u0010\u0000\u0000\u012b\u012c\u0005"+
		"A\u0000\u0000\u012c\u012e\u0003\u0002\u0001\u0000\u012d\u012f\u0005\u0016"+
		"\u0000\u0000\u012e\u012d\u0001\u0000\u0000\u0000\u012e\u012f\u0001\u0000"+
		"\u0000\u0000\u012f\u0130\u0001\u0000\u0000\u0000\u0130\u0131\u0006\u000e"+
		"\uffff\uffff\u0000\u0131\u001d\u0001\u0000\u0000\u0000\u0132\u0133\u0005"+
		"\u0011\u0000\u0000\u0133\u0134\u0003B!\u0000\u0134\u0135\u00059\u0000"+
		"\u0000\u0135\u0136\u0003\u0002\u0001\u0000\u0136\u0137\u0005:\u0000\u0000"+
		"\u0137\u0138\u0006\u000f\uffff\uffff\u0000\u0138\u001f\u0001\u0000\u0000"+
		"\u0000\u0139\u013a\u0005\u0012\u0000\u0000\u013a\u013b\u0005\'\u0000\u0000"+
		"\u013b\u013c\u0005\u0013\u0000\u0000\u013c\u013d\u0005&\u0000\u0000\u013d"+
		"\u013e\u00059\u0000\u0000\u013e\u013f\u0003\u0002\u0001\u0000\u013f\u0140"+
		"\u0005:\u0000\u0000\u0140\u0141\u0006\u0010\uffff\uffff\u0000\u0141\u0157"+
		"\u0001\u0000\u0000\u0000\u0142\u0143\u0005\u0012\u0000\u0000\u0143\u0144"+
		"\u0005\'\u0000\u0000\u0144\u0145\u0005\u0013\u0000\u0000\u0145\u0146\u0003"+
		"B!\u0000\u0146\u0147\u0005\u0014\u0000\u0000\u0147\u0148\u0003B!\u0000"+
		"\u0148\u0149\u00059\u0000\u0000\u0149\u014a\u0003\u0002\u0001\u0000\u014a"+
		"\u014b\u0005:\u0000\u0000\u014b\u014c\u0006\u0010\uffff\uffff\u0000\u014c"+
		"\u0157\u0001\u0000\u0000\u0000\u014d\u014e\u0005\u0012\u0000\u0000\u014e"+
		"\u014f\u0005\'\u0000\u0000\u014f\u0150\u0005\u0013\u0000\u0000\u0150\u0151"+
		"\u0005\'\u0000\u0000\u0151\u0152\u00059\u0000\u0000\u0152\u0153\u0003"+
		"\u0002\u0001\u0000\u0153\u0154\u0005:\u0000\u0000\u0154\u0155\u0006\u0010"+
		"\uffff\uffff\u0000\u0155\u0157\u0001\u0000\u0000\u0000\u0156\u0139\u0001"+
		"\u0000\u0000\u0000\u0156\u0142\u0001\u0000\u0000\u0000\u0156\u014d\u0001"+
		"\u0000\u0000\u0000\u0157!\u0001\u0000\u0000\u0000\u0158\u0159\u0005\u0015"+
		"\u0000\u0000\u0159\u015a\u0003B!\u0000\u015a\u015b\u0005\r\u0000\u0000"+
		"\u015b\u015c\u00059\u0000\u0000\u015c\u015d\u0003\u0002\u0001\u0000\u015d"+
		"\u015e\u0005:\u0000\u0000\u015e\u015f\u0006\u0011\uffff\uffff\u0000\u015f"+
		"#\u0001\u0000\u0000\u0000\u0160\u0161\u0005\u0016\u0000\u0000\u0161\u016b"+
		"\u0006\u0012\uffff\uffff\u0000\u0162\u0163\u0005\u0017\u0000\u0000\u0163"+
		"\u016b\u0006\u0012\uffff\uffff\u0000\u0164\u0165\u0005\u0018\u0000\u0000"+
		"\u0165\u016b\u0006\u0012\uffff\uffff\u0000\u0166\u0167\u0005\u0018\u0000"+
		"\u0000\u0167\u0168\u0003B!\u0000\u0168\u0169\u0006\u0012\uffff\uffff\u0000"+
		"\u0169\u016b\u0001\u0000\u0000\u0000\u016a\u0160\u0001\u0000\u0000\u0000"+
		"\u016a\u0162\u0001\u0000\u0000\u0000\u016a\u0164\u0001\u0000\u0000\u0000"+
		"\u016a\u0166\u0001\u0000\u0000\u0000\u016b%\u0001\u0000\u0000\u0000\u016c"+
		"\u016d\u0005\t\u0000\u0000\u016d\u016e\u0005\'\u0000\u0000\u016e\u016f"+
		"\u0005A\u0000\u0000\u016f\u0170\u0005;\u0000\u0000\u0170\u0171\u0003\u0012"+
		"\t\u0000\u0171\u0172\u0005<\u0000\u0000\u0172\u0173\u0003(\u0014\u0000"+
		"\u0173\u0174\u0006\u0013\uffff\uffff\u0000\u0174\'\u0001\u0000\u0000\u0000"+
		"\u0175\u0176\u0005-\u0000\u0000\u0176\u0178\u0005;\u0000\u0000\u0177\u0179"+
		"\u0003*\u0015\u0000\u0178\u0177\u0001\u0000\u0000\u0000\u0179\u017a\u0001"+
		"\u0000\u0000\u0000\u017a\u0178\u0001\u0000\u0000\u0000\u017a\u017b\u0001"+
		"\u0000\u0000\u0000\u017b\u017c\u0001\u0000\u0000\u0000\u017c\u017d\u0005"+
		"<\u0000\u0000\u017d\u017e\u0006\u0014\uffff\uffff\u0000\u017e\u0184\u0001"+
		"\u0000\u0000\u0000\u017f\u0180\u0005-\u0000\u0000\u0180\u0181\u0005;\u0000"+
		"\u0000\u0181\u0182\u0005<\u0000\u0000\u0182\u0184\u0006\u0014\uffff\uffff"+
		"\u0000\u0183\u0175\u0001\u0000\u0000\u0000\u0183\u017f\u0001\u0000\u0000"+
		"\u0000\u0184)\u0001\u0000\u0000\u0000\u0185\u0186\u0003B!\u0000\u0186"+
		"\u0187\u0005?\u0000\u0000\u0187\u0188\u0006\u0015\uffff\uffff\u0000\u0188"+
		"\u018d\u0001\u0000\u0000\u0000\u0189\u018a\u0003B!\u0000\u018a\u018b\u0006"+
		"\u0015\uffff\uffff\u0000\u018b\u018d\u0001\u0000\u0000\u0000\u018c\u0185"+
		"\u0001\u0000\u0000\u0000\u018c\u0189\u0001\u0000\u0000\u0000\u018d+\u0001"+
		"\u0000\u0000\u0000\u018e\u018f\u0005\'\u0000\u0000\u018f\u0190\u0005@"+
		"\u0000\u0000\u0190\u0191\u0005\u001d\u0000\u0000\u0191\u0192\u00057\u0000"+
		"\u0000\u0192\u0193\u0003B!\u0000\u0193\u0194\u00058\u0000\u0000\u0194"+
		"\u0195\u0006\u0016\uffff\uffff\u0000\u0195\u01a5\u0001\u0000\u0000\u0000"+
		"\u0196\u0197\u0005\'\u0000\u0000\u0197\u0198\u0005@\u0000\u0000\u0198"+
		"\u0199\u0005\u001e\u0000\u0000\u0199\u019a\u00057\u0000\u0000\u019a\u019b"+
		"\u00058\u0000\u0000\u019b\u01a5\u0006\u0016\uffff\uffff\u0000\u019c\u019d"+
		"\u0005\'\u0000\u0000\u019d\u019e\u0005@\u0000\u0000\u019e\u019f\u0005"+
		"\u001f\u0000\u0000\u019f\u01a0\u00057\u0000\u0000\u01a0\u01a1\u0003B!"+
		"\u0000\u01a1\u01a2\u00058\u0000\u0000\u01a2\u01a3\u0006\u0016\uffff\uffff"+
		"\u0000\u01a3\u01a5\u0001\u0000\u0000\u0000\u01a4\u018e\u0001\u0000\u0000"+
		"\u0000\u01a4\u0196\u0001\u0000\u0000\u0000\u01a4\u019c\u0001\u0000\u0000"+
		"\u0000\u01a5-\u0001\u0000\u0000\u0000\u01a6\u01a7\u0005\'\u0000\u0000"+
		"\u01a7\u01a8\u0005@\u0000\u0000\u01a8\u01a9\u0005 \u0000\u0000\u01a9\u01b5"+
		"\u0006\u0017\uffff\uffff\u0000\u01aa\u01ab\u0005\'\u0000\u0000\u01ab\u01ac"+
		"\u0005@\u0000\u0000\u01ac\u01ad\u0005!\u0000\u0000\u01ad\u01b5\u0006\u0017"+
		"\uffff\uffff\u0000\u01ae\u01af\u0005\'\u0000\u0000\u01af\u01b0\u0005;"+
		"\u0000\u0000\u01b0\u01b1\u0003B!\u0000\u01b1\u01b2\u0005<\u0000\u0000"+
		"\u01b2\u01b3\u0006\u0017\uffff\uffff\u0000\u01b3\u01b5\u0001\u0000\u0000"+
		"\u0000\u01b4\u01a6\u0001\u0000\u0000\u0000\u01b4\u01aa\u0001\u0000\u0000"+
		"\u0000\u01b4\u01ae\u0001\u0000\u0000\u0000\u01b5/\u0001\u0000\u0000\u0000"+
		"\u01b6\u01b7\u0005\'\u0000\u0000\u01b7\u01b8\u0005;\u0000\u0000\u01b8"+
		"\u01b9\u0003B!\u0000\u01b9\u01ba\u0005<\u0000\u0000\u01ba\u01bb\u0005"+
		"-\u0000\u0000\u01bb\u01bc\u0005\'\u0000\u0000\u01bc\u01bd\u0005;\u0000"+
		"\u0000\u01bd\u01be\u0003B!\u0000\u01be\u01bf\u0005<\u0000\u0000\u01bf"+
		"\u01c0\u0006\u0018\uffff\uffff\u0000\u01c0\u01ca\u0001\u0000\u0000\u0000"+
		"\u01c1\u01c2\u0005\'\u0000\u0000\u01c2\u01c3\u0005;\u0000\u0000\u01c3"+
		"\u01c4\u0003B!\u0000\u01c4\u01c5\u0005<\u0000\u0000\u01c5\u01c6\u0005"+
		"-\u0000\u0000\u01c6\u01c7\u0003B!\u0000\u01c7\u01c8\u0006\u0018\uffff"+
		"\uffff\u0000\u01c8\u01ca\u0001\u0000\u0000\u0000\u01c9\u01b6\u0001\u0000"+
		"\u0000\u0000\u01c9\u01c1\u0001\u0000\u0000\u0000\u01ca1\u0001\u0000\u0000"+
		"\u0000\u01cb\u01cc\u0005\u0019\u0000\u0000\u01cc\u01cd\u0005\'\u0000\u0000"+
		"\u01cd\u01cf\u00059\u0000\u0000\u01ce\u01d0\u00034\u001a\u0000\u01cf\u01ce"+
		"\u0001\u0000\u0000\u0000\u01d0\u01d1\u0001\u0000\u0000\u0000\u01d1\u01cf"+
		"\u0001\u0000\u0000\u0000\u01d1\u01d2\u0001\u0000\u0000\u0000\u01d2\u01d3"+
		"\u0001\u0000\u0000\u0000\u01d3\u01d4\u0005:\u0000\u0000\u01d43\u0001\u0000"+
		"\u0000\u0000\u01d5\u01d6\u0007\u0000\u0000\u0000\u01d6\u01d8\u0005\'\u0000"+
		"\u0000\u01d7\u01d9\u0003\u0012\t\u0000\u01d8\u01d7\u0001\u0000\u0000\u0000"+
		"\u01d8\u01d9\u0001\u0000\u0000\u0000\u01d9\u01dc\u0001\u0000\u0000\u0000"+
		"\u01da\u01db\u0005-\u0000\u0000\u01db\u01dd\u0003B!\u0000\u01dc\u01da"+
		"\u0001\u0000\u0000\u0000\u01dc\u01dd\u0001\u0000\u0000\u0000\u01dd5\u0001"+
		"\u0000\u0000\u0000\u01de\u01df\u0005\u001c\u0000\u0000\u01df\u01e0\u0005"+
		"\'\u0000\u0000\u01e0\u01e2\u00057\u0000\u0000\u01e1\u01e3\u00038\u001c"+
		"\u0000\u01e2\u01e1\u0001\u0000\u0000\u0000\u01e3\u01e4\u0001\u0000\u0000"+
		"\u0000\u01e4\u01e2\u0001\u0000\u0000\u0000\u01e4\u01e5\u0001\u0000\u0000"+
		"\u0000\u01e5\u01e6\u0001\u0000\u0000\u0000\u01e6\u01e7\u00058\u0000\u0000"+
		"\u01e7\u01e8\u0005>\u0000\u0000\u01e8\u01e9\u0003\u0012\t\u0000\u01e9"+
		"\u01ea\u00059\u0000\u0000\u01ea\u01eb\u0003\u0002\u0001\u0000\u01eb\u01ec"+
		"\u0005:\u0000\u0000\u01ec\u01ed\u0006\u001b\uffff\uffff\u0000\u01ed\u0211"+
		"\u0001\u0000\u0000\u0000\u01ee\u01ef\u0005\u001c\u0000\u0000\u01ef\u01f0"+
		"\u0005\'\u0000\u0000\u01f0\u01f2\u00057\u0000\u0000\u01f1\u01f3\u0003"+
		"8\u001c\u0000\u01f2\u01f1\u0001\u0000\u0000\u0000\u01f3\u01f4\u0001\u0000"+
		"\u0000\u0000\u01f4\u01f2\u0001\u0000\u0000\u0000\u01f4\u01f5\u0001\u0000"+
		"\u0000\u0000\u01f5\u01f6\u0001\u0000\u0000\u0000\u01f6\u01f7\u00058\u0000"+
		"\u0000\u01f7\u01f8\u00059\u0000\u0000\u01f8\u01f9\u0003\u0002\u0001\u0000"+
		"\u01f9\u01fa\u0005:\u0000\u0000\u01fa\u01fb\u0006\u001b\uffff\uffff\u0000"+
		"\u01fb\u0211\u0001\u0000\u0000\u0000\u01fc\u01fd\u0005\u001c\u0000\u0000"+
		"\u01fd\u01fe\u0005\'\u0000\u0000\u01fe\u01ff\u00057\u0000\u0000\u01ff"+
		"\u0200\u00058\u0000\u0000\u0200\u0201\u0005>\u0000\u0000\u0201\u0202\u0003"+
		"\u0012\t\u0000\u0202\u0203\u00059\u0000\u0000\u0203\u0204\u0003\u0002"+
		"\u0001\u0000\u0204\u0205\u0005:\u0000\u0000\u0205\u0206\u0006\u001b\uffff"+
		"\uffff\u0000\u0206\u0211\u0001\u0000\u0000\u0000\u0207\u0208\u0005\u001c"+
		"\u0000\u0000\u0208\u0209\u0005\'\u0000\u0000\u0209\u020a\u00057\u0000"+
		"\u0000\u020a\u020b\u00058\u0000\u0000\u020b\u020c\u00059\u0000\u0000\u020c"+
		"\u020d\u0003\u0002\u0001\u0000\u020d\u020e\u0005:\u0000\u0000\u020e\u020f"+
		"\u0006\u001b\uffff\uffff\u0000\u020f\u0211\u0001\u0000\u0000\u0000\u0210"+
		"\u01de\u0001\u0000\u0000\u0000\u0210\u01ee\u0001\u0000\u0000\u0000\u0210"+
		"\u01fc\u0001\u0000\u0000\u0000\u0210\u0207\u0001\u0000\u0000\u0000\u0211"+
		"7\u0001\u0000\u0000\u0000\u0212\u0214\u0007\u0001\u0000\u0000\u0213\u0212"+
		"\u0001\u0000\u0000\u0000\u0213\u0214\u0001\u0000\u0000\u0000\u0214\u0215"+
		"\u0001\u0000\u0000\u0000\u0215\u0216\u0005\'\u0000\u0000\u0216\u0218\u0005"+
		"A\u0000\u0000\u0217\u0219\u0005\"\u0000\u0000\u0218\u0217\u0001\u0000"+
		"\u0000\u0000\u0218\u0219\u0001\u0000\u0000\u0000\u0219\u021a\u0001\u0000"+
		"\u0000\u0000\u021a\u021b\u0003\u0012\t\u0000\u021b\u021c\u0005?\u0000"+
		"\u0000\u021c\u021d\u0006\u001c\uffff\uffff\u0000\u021d\u022a\u0001\u0000"+
		"\u0000\u0000\u021e\u0220\u0007\u0001\u0000\u0000\u021f\u021e\u0001\u0000"+
		"\u0000\u0000\u021f\u0220\u0001\u0000\u0000\u0000\u0220\u0221\u0001\u0000"+
		"\u0000\u0000\u0221\u0222\u0005\'\u0000\u0000\u0222\u0224\u0005A\u0000"+
		"\u0000\u0223\u0225\u0005\"\u0000\u0000\u0224\u0223\u0001\u0000\u0000\u0000"+
		"\u0224\u0225\u0001\u0000\u0000\u0000\u0225\u0226\u0001\u0000\u0000\u0000"+
		"\u0226\u0227\u0003\u0012\t\u0000\u0227\u0228\u0006\u001c\uffff\uffff\u0000"+
		"\u0228\u022a\u0001\u0000\u0000\u0000\u0229\u0213\u0001\u0000\u0000\u0000"+
		"\u0229\u021f\u0001\u0000\u0000\u0000\u022a9\u0001\u0000\u0000\u0000\u022b"+
		"\u022c\u0005\'\u0000\u0000\u022c\u022e\u00057\u0000\u0000\u022d\u022f"+
		"\u0003>\u001f\u0000\u022e\u022d\u0001\u0000\u0000\u0000\u022f\u0230\u0001"+
		"\u0000\u0000\u0000\u0230\u022e\u0001\u0000\u0000\u0000\u0230\u0231\u0001"+
		"\u0000\u0000\u0000\u0231\u0232\u0001\u0000\u0000\u0000\u0232\u0233\u0005"+
		"8\u0000\u0000\u0233\u0234\u0006\u001d\uffff\uffff\u0000\u0234\u023a\u0001"+
		"\u0000\u0000\u0000\u0235\u0236\u0005\'\u0000\u0000\u0236\u0237\u00057"+
		"\u0000\u0000\u0237\u0238\u00058\u0000\u0000\u0238\u023a\u0006\u001d\uffff"+
		"\uffff\u0000\u0239\u022b\u0001\u0000\u0000\u0000\u0239\u0235\u0001\u0000"+
		"\u0000\u0000\u023a;\u0001\u0000\u0000\u0000\u023b\u023c\u0005\'\u0000"+
		"\u0000\u023c\u023e\u00057\u0000\u0000\u023d\u023f\u0003>\u001f\u0000\u023e"+
		"\u023d\u0001\u0000\u0000\u0000\u023f\u0240\u0001\u0000\u0000\u0000\u0240"+
		"\u023e\u0001\u0000\u0000\u0000\u0240\u0241\u0001\u0000\u0000\u0000\u0241"+
		"\u0242\u0001\u0000\u0000\u0000\u0242\u0243\u00058\u0000\u0000\u0243\u0244"+
		"\u0006\u001e\uffff\uffff\u0000\u0244\u024a\u0001\u0000\u0000\u0000\u0245"+
		"\u0246\u0005\'\u0000\u0000\u0246\u0247\u00057\u0000\u0000\u0247\u0248"+
		"\u00058\u0000\u0000\u0248\u024a\u0006\u001e\uffff\uffff\u0000\u0249\u023b"+
		"\u0001\u0000\u0000\u0000\u0249\u0245\u0001\u0000\u0000\u0000\u024a=\u0001"+
		"\u0000\u0000\u0000\u024b\u024c\u0005\'\u0000\u0000\u024c\u024e\u0005A"+
		"\u0000\u0000\u024d\u024b\u0001\u0000\u0000\u0000\u024d\u024e\u0001\u0000"+
		"\u0000\u0000\u024e\u0250\u0001\u0000\u0000\u0000\u024f\u0251\u0005B\u0000"+
		"\u0000\u0250\u024f\u0001\u0000\u0000\u0000\u0250\u0251\u0001\u0000\u0000"+
		"\u0000\u0251\u0252\u0001\u0000\u0000\u0000\u0252\u0253\u0003B!\u0000\u0253"+
		"\u0254\u0005?\u0000\u0000\u0254\u0255\u0006\u001f\uffff\uffff\u0000\u0255"+
		"\u0261\u0001\u0000\u0000\u0000\u0256\u0257\u0005\'\u0000\u0000\u0257\u0259"+
		"\u0005A\u0000\u0000\u0258\u0256\u0001\u0000\u0000\u0000\u0258\u0259\u0001"+
		"\u0000\u0000\u0000\u0259\u025b\u0001\u0000\u0000\u0000\u025a\u025c\u0005"+
		"B\u0000\u0000\u025b\u025a\u0001\u0000\u0000\u0000\u025b\u025c\u0001\u0000"+
		"\u0000\u0000\u025c\u025d\u0001\u0000\u0000\u0000\u025d\u025e\u0003B!\u0000"+
		"\u025e\u025f\u0006\u001f\uffff\uffff\u0000\u025f\u0261\u0001\u0000\u0000"+
		"\u0000\u0260\u024d\u0001\u0000\u0000\u0000\u0260\u0258\u0001\u0000\u0000"+
		"\u0000\u0261?\u0001\u0000\u0000\u0000\u0262\u0263\u0005\u0004\u0000\u0000"+
		"\u0263\u0264\u00057\u0000\u0000\u0264\u0265\u0003B!\u0000\u0265\u0266"+
		"\u00058\u0000\u0000\u0266\u0267\u0006 \uffff\uffff\u0000\u0267\u0275\u0001"+
		"\u0000\u0000\u0000\u0268\u0269\u0005\u0001\u0000\u0000\u0269\u026a\u0005"+
		"7\u0000\u0000\u026a\u026b\u0003B!\u0000\u026b\u026c\u00058\u0000\u0000"+
		"\u026c\u026d\u0006 \uffff\uffff\u0000\u026d\u0275\u0001\u0000\u0000\u0000"+
		"\u026e\u026f\u0005\u0002\u0000\u0000\u026f\u0270\u00057\u0000\u0000\u0270"+
		"\u0271\u0003B!\u0000\u0271\u0272\u00058\u0000\u0000\u0272\u0273\u0006"+
		" \uffff\uffff\u0000\u0273\u0275\u0001\u0000\u0000\u0000\u0274\u0262\u0001"+
		"\u0000\u0000\u0000\u0274\u0268\u0001\u0000\u0000\u0000\u0274\u026e\u0001"+
		"\u0000\u0000\u0000\u0275A\u0001\u0000\u0000\u0000\u0276\u0277\u0006!\uffff"+
		"\uffff\u0000\u0277\u0278\u00056\u0000\u0000\u0278\u0279\u0003B!\u0014"+
		"\u0279\u027a\u0006!\uffff\uffff\u0000\u027a\u029c\u0001\u0000\u0000\u0000"+
		"\u027b\u027c\u0005*\u0000\u0000\u027c\u027d\u0003B!\u0013\u027d\u027e"+
		"\u0006!\uffff\uffff\u0000\u027e\u029c\u0001\u0000\u0000\u0000\u027f\u0280"+
		"\u00057\u0000\u0000\u0280\u0281\u0003B!\u0000\u0281\u0282\u00058\u0000"+
		"\u0000\u0282\u0283\u0006!\uffff\uffff\u0000\u0283\u029c\u0001\u0000\u0000"+
		"\u0000\u0284\u0285\u0005$\u0000\u0000\u0285\u029c\u0006!\uffff\uffff\u0000"+
		"\u0286\u0287\u0005%\u0000\u0000\u0287\u029c\u0006!\uffff\uffff\u0000\u0288"+
		"\u0289\u0005&\u0000\u0000\u0289\u029c\u0006!\uffff\uffff\u0000\u028a\u028b"+
		"\u0005\u0006\u0000\u0000\u028b\u029c\u0006!\uffff\uffff\u0000\u028c\u028d"+
		"\u0003:\u001d\u0000\u028d\u028e\u0006!\uffff\uffff\u0000\u028e\u029c\u0001"+
		"\u0000\u0000\u0000\u028f\u0290\u0003@ \u0000\u0290\u0291\u0006!\uffff"+
		"\uffff\u0000\u0291\u029c\u0001\u0000\u0000\u0000\u0292\u0293\u0005\u0007"+
		"\u0000\u0000\u0293\u029c\u0006!\uffff\uffff\u0000\u0294\u0295\u0005\'"+
		"\u0000\u0000\u0295\u029c\u0006!\uffff\uffff\u0000\u0296\u0297\u0003.\u0017"+
		"\u0000\u0297\u0298\u0006!\uffff\uffff\u0000\u0298\u029c\u0001\u0000\u0000"+
		"\u0000\u0299\u029a\u0005\b\u0000\u0000\u029a\u029c\u0006!\uffff\uffff"+
		"\u0000\u029b\u0276\u0001\u0000\u0000\u0000\u029b\u027b\u0001\u0000\u0000"+
		"\u0000\u029b\u027f\u0001\u0000\u0000\u0000\u029b\u0284\u0001\u0000\u0000"+
		"\u0000\u029b\u0286\u0001\u0000\u0000\u0000\u029b\u0288\u0001\u0000\u0000"+
		"\u0000\u029b\u028a\u0001\u0000\u0000\u0000\u029b\u028c\u0001\u0000\u0000"+
		"\u0000\u029b\u028f\u0001\u0000\u0000\u0000\u029b\u0292\u0001\u0000\u0000"+
		"\u0000\u029b\u0294\u0001\u0000\u0000\u0000\u029b\u0296\u0001\u0000\u0000"+
		"\u0000\u029b\u0299\u0001\u0000\u0000\u0000\u029c\u02c2\u0001\u0000\u0000"+
		"\u0000\u029d\u029e\n\u0012\u0000\u0000\u029e\u029f\u0007\u0002\u0000\u0000"+
		"\u029f\u02a0\u0003B!\u0013\u02a0\u02a1\u0006!\uffff\uffff\u0000\u02a1"+
		"\u02c1\u0001\u0000\u0000\u0000\u02a2\u02a3\n\u0011\u0000\u0000\u02a3\u02a4"+
		"\u0007\u0003\u0000\u0000\u02a4\u02a5\u0003B!\u0012\u02a5\u02a6\u0006!"+
		"\uffff\uffff\u0000\u02a6\u02c1\u0001\u0000\u0000\u0000\u02a7\u02a8\n\u0010"+
		"\u0000\u0000\u02a8\u02a9\u0007\u0004\u0000\u0000\u02a9\u02aa\u0003B!\u0011"+
		"\u02aa\u02ab\u0006!\uffff\uffff\u0000\u02ab\u02c1\u0001\u0000\u0000\u0000"+
		"\u02ac\u02ad\n\u000f\u0000\u0000\u02ad\u02ae\u0007\u0005\u0000\u0000\u02ae"+
		"\u02af\u0003B!\u0010\u02af\u02b0\u0006!\uffff\uffff\u0000\u02b0\u02c1"+
		"\u0001\u0000\u0000\u0000\u02b1\u02b2\n\u000e\u0000\u0000\u02b2\u02b3\u0007"+
		"\u0006\u0000\u0000\u02b3\u02b4\u0003B!\u000f\u02b4\u02b5\u0006!\uffff"+
		"\uffff\u0000\u02b5\u02c1\u0001\u0000\u0000\u0000\u02b6\u02b7\n\r\u0000"+
		"\u0000\u02b7\u02b8\u0005,\u0000\u0000\u02b8\u02b9\u0003B!\u000e\u02b9"+
		"\u02ba\u0006!\uffff\uffff\u0000\u02ba\u02c1\u0001\u0000\u0000\u0000\u02bb"+
		"\u02bc\n\f\u0000\u0000\u02bc\u02bd\u0005+\u0000\u0000\u02bd\u02be\u0003"+
		"B!\r\u02be\u02bf\u0006!\uffff\uffff\u0000\u02bf\u02c1\u0001\u0000\u0000"+
		"\u0000\u02c0\u029d\u0001\u0000\u0000\u0000\u02c0\u02a2\u0001\u0000\u0000"+
		"\u0000\u02c0\u02a7\u0001\u0000\u0000\u0000\u02c0\u02ac\u0001\u0000\u0000"+
		"\u0000\u02c0\u02b1\u0001\u0000\u0000\u0000\u02c0\u02b6\u0001\u0000\u0000"+
		"\u0000\u02c0\u02bb\u0001\u0000\u0000\u0000\u02c1\u02c4\u0001\u0000\u0000"+
		"\u0000\u02c2\u02c0\u0001\u0000\u0000\u0000\u02c2\u02c3\u0001\u0000\u0000"+
		"\u0000\u02c3C\u0001\u0000\u0000\u0000\u02c4\u02c2\u0001\u0000\u0000\u0000"+
		"/K|\u0083\u008f\u00a0\u00ad\u00b9\u00de\u00ee\u0101\u0111\u0119\u011c"+
		"\u0126\u012e\u0156\u016a\u017a\u0183\u018c\u01a4\u01b4\u01c9\u01d1\u01d8"+
		"\u01dc\u01e4\u01f4\u0210\u0213\u0218\u021f\u0224\u0229\u0230\u0239\u0240"+
		"\u0249\u024d\u0250\u0258\u025b\u0260\u0274\u029b\u02c0\u02c2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}