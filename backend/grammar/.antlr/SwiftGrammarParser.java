// Generated from c:\Users\TheJhonX\Documents\GitHub\OLC2_201900831\backend\grammar\SwiftGrammar.g4 by ANTLR 4.9.2

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

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

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
		RULE_vecaccess = 24, RULE_funcstmt = 25, RULE_listparam = 26, RULE_callfunc = 27, 
		RULE_callfuncinstruction = 28, RULE_listparamcall = 29, RULE_funcembed = 30, 
		RULE_expr = 31;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "printexprlist", "ifstmt", 
			"elseifstmt", "elsestmt", "varstmt", "tipo", "varasgmt", "conststmt", 
			"switchstmt", "casestmt", "defaultstmt", "whilestmt", "forstmt", "guardstmt", 
			"transferstmt", "vectorstmt", "definestmt", "listexpr", "methodvec", 
			"methodvecrtrn", "vecaccess", "funcstmt", "listparam", "callfunc", "callfuncinstruction", 
			"listparamcall", "funcembed", "expr"
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
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			((SContext)_localctx).block = block();
			setState(65);
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
			setState(69); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(68);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(71); 
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
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(120);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(78);
				((InstructionContext)_localctx).callfuncinstruction = callfuncinstruction();
				 _localctx.inst = ((InstructionContext)_localctx).callfuncinstruction.callfuncinstr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(81);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinstr 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(84);
				((InstructionContext)_localctx).varstmt = varstmt();
				 _localctx.inst = ((InstructionContext)_localctx).varstmt.var
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(87);
				((InstructionContext)_localctx).varasgmt = varasgmt();
				 _localctx.inst = ((InstructionContext)_localctx).varasgmt.asgmt
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(90);
				((InstructionContext)_localctx).conststmt = conststmt();
				 _localctx.inst = ((InstructionContext)_localctx).conststmt.const
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(93);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchstmt.switchinstr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(96);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whileinstr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(99);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.forinstr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(102);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.guardinstr
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(105);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				 _localctx.inst = ((InstructionContext)_localctx).transferstmt.trns
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(108);
				((InstructionContext)_localctx).vectorstmt = vectorstmt();
				 _localctx.inst = ((InstructionContext)_localctx).vectorstmt.vectorinstr
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(111);
				((InstructionContext)_localctx).methodvec = methodvec();
				 _localctx.inst = ((InstructionContext)_localctx).methodvec.methodinstr
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(114);
				((InstructionContext)_localctx).vecaccess = vecaccess();
				 _localctx.inst = ((InstructionContext)_localctx).vecaccess.vecacc
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(117);
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
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_printstmt);

		    var listExpr []interface{}

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(123);
			match(PARIZQ);
			setState(125); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(124);
				((PrintstmtContext)_localctx).printexprlist = printexprlist();
				((PrintstmtContext)_localctx).lista.add(((PrintstmtContext)_localctx).printexprlist);
				}
				}
				setState(127); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0) );
			setState(129);
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
	}

	public final PrintexprlistContext printexprlist() throws RecognitionException {
		PrintexprlistContext _localctx = new PrintexprlistContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_printexprlist);
		try {
			setState(139);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				((PrintexprlistContext)_localctx).expr = expr(0);
				setState(133);
				match(COMA);
				 _localctx.prntexpr = ((PrintexprlistContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(136);
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
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
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
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			setState(173);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(142);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(143);
				match(LLAVEIZQ);
				setState(144);
				((IfstmtContext)_localctx).block = block();
				setState(145);
				match(LLAVEDER);
				 _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(149);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(150);
				match(LLAVEIZQ);
				setState(151);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(152);
				match(LLAVEDER);
				setState(154); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(153);
						((IfstmtContext)_localctx).elseifstmt = elseifstmt();
						((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(156); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(158);
				match(ELSE);
				setState(160);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LLAVEIZQ) {
					{
					setState(159);
					((IfstmtContext)_localctx).elsestmt = elsestmt();
					}
				}

				 
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
				setState(164);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(165);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(166);
				match(LLAVEIZQ);
				setState(167);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(168);
				match(LLAVEDER);
				setState(169);
				match(ELSE);
				setState(170);
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
	}

	public final ElseifstmtContext elseifstmt() throws RecognitionException {
		ElseifstmtContext _localctx = new ElseifstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			((ElseifstmtContext)_localctx).ELSE = match(ELSE);
			setState(176);
			match(IF);
			setState(177);
			((ElseifstmtContext)_localctx).expr = expr(0);
			setState(178);
			match(LLAVEIZQ);
			setState(179);
			((ElseifstmtContext)_localctx).block = block();
			setState(180);
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

	public static class ElsestmtContext extends ParserRuleContext {
		public []interface{} elseinstr;
		public BlockContext block;
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElsestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elsestmt; }
	}

	public final ElsestmtContext elsestmt() throws RecognitionException {
		ElsestmtContext _localctx = new ElsestmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_elsestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(LLAVEIZQ);
			setState(184);
			((ElsestmtContext)_localctx).block = block();
			setState(185);
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
	}

	public final VarstmtContext varstmt() throws RecognitionException {
		VarstmtContext _localctx = new VarstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_varstmt);
		try {
			setState(209);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(188);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(189);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(190);
				match(COLON);
				setState(191);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(192);
				match(IG);
				setState(193);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), ((VarstmtContext)_localctx).tipo.rtipo, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(196);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(197);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(198);
				match(IG);
				setState(199);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), environment.NULL, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(202);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(203);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(204);
				match(COLON);
				setState(205);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(206);
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

	public static class TipoContext extends ParserRuleContext {
		public environment.TipoExpresion rtipo;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(SwiftGrammarParser.CHARACTER, 0); }
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_tipo);
		try {
			setState(221);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(211);
				match(INT);
				_localctx.rtipo = 0
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(213);
				match(FLOAT);
				_localctx.rtipo = 1
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(215);
				match(STR);
				_localctx.rtipo = 2
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(217);
				match(BOOL);
				_localctx.rtipo = 3
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(219);
				match(CHARACTER);
				_localctx.rtipo = 5
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
	}

	public final VarasgmtContext varasgmt() throws RecognitionException {
		VarasgmtContext _localctx = new VarasgmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_varasgmt);
		try {
			setState(240);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(223);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(224);
				match(IG);
				setState(225);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(229);
				match(ADD);
				setState(230);
				match(IG);
				setState(231);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewOpAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e, "+")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(234);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(235);
				match(SUB);
				setState(236);
				match(IG);
				setState(237);
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

	public static class ConststmtContext extends ParserRuleContext {
		public interfaces.Instruction const;
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
	}

	public final ConststmtContext conststmt() throws RecognitionException {
		ConststmtContext _localctx = new ConststmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_conststmt);
		try {
			setState(256);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(243);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(244);
				match(IG);
				setState(245);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), environment.NULL, ((ConststmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(248);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(249);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(250);
				match(COLON);
				setState(251);
				((ConststmtContext)_localctx).tipo = tipo();
				setState(252);
				match(IG);
				setState(253);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), ((ConststmtContext)_localctx).tipo.rtipo, ((ConststmtContext)_localctx).expr.e, true)
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
			setState(258);
			((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
			setState(259);
			((SwitchstmtContext)_localctx).expr = expr(0);
			setState(260);
			match(LLAVEIZQ);
			setState(262); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(261);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				((SwitchstmtContext)_localctx).casesvar.add(((SwitchstmtContext)_localctx).casestmt);
				}
				}
				setState(264); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(266);
				defaultstmt();
				}
			}

			setState(269);
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
	}

	public final CasestmtContext casestmt() throws RecognitionException {
		CasestmtContext _localctx = new CasestmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_casestmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			((CasestmtContext)_localctx).CASE = match(CASE);
			setState(273);
			((CasestmtContext)_localctx).expr = expr(0);
			setState(274);
			match(COLON);
			setState(275);
			((CasestmtContext)_localctx).block = block();
			setState(277);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(276);
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
	}

	public final DefaultstmtContext defaultstmt() throws RecognitionException {
		DefaultstmtContext _localctx = new DefaultstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_defaultstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			match(DEFAULT);
			setState(282);
			match(COLON);
			setState(283);
			((DefaultstmtContext)_localctx).block = block();
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(284);
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
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(290);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(291);
			match(LLAVEIZQ);
			setState(292);
			((WhilestmtContext)_localctx).block = block();
			setState(293);
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
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_forstmt);

		    var cadena interfaces.Expression
		    var str string

		try {
			setState(325);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(296);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(297);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(298);
				match(IN);
				setState(299);
				((ForstmtContext)_localctx).STRING = match(STRING);
				setState(300);
				match(LLAVEIZQ);
				setState(301);
				((ForstmtContext)_localctx).block = block();
				setState(302);
				match(LLAVEDER);
				 
				    str = (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getText():null)
				    cadena = expressions.NewPrimitive((((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getLine():0), (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1], environment.STRING)
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), cadena, nil, nil, "", ((ForstmtContext)_localctx).block.blk) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(306);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(307);
				match(IN);
				setState(308);
				((ForstmtContext)_localctx).left = expr(0);
				setState(309);
				match(RANGEPTS);
				setState(310);
				((ForstmtContext)_localctx).right = expr(0);
				setState(311);
				match(LLAVEIZQ);
				setState(312);
				((ForstmtContext)_localctx).block = block();
				setState(313);
				match(LLAVEDER);
				 
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), nil, ((ForstmtContext)_localctx).left.e, ((ForstmtContext)_localctx).right.e, "", ((ForstmtContext)_localctx).block.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(316);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(317);
				((ForstmtContext)_localctx).first = match(ID);
				setState(318);
				match(IN);
				setState(319);
				((ForstmtContext)_localctx).second = match(ID);
				setState(320);
				match(LLAVEIZQ);
				setState(321);
				((ForstmtContext)_localctx).block = block();
				setState(322);
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
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(327);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(328);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(329);
			match(ELSE);
			setState(330);
			match(LLAVEIZQ);
			setState(331);
			((GuardstmtContext)_localctx).block = block();
			setState(332);
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
	}

	public final TransferstmtContext transferstmt() throws RecognitionException {
		TransferstmtContext _localctx = new TransferstmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_transferstmt);
		try {
			setState(345);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(335);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				 _localctx.trns = instructions.NewBreak((((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getLine():0), (((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getCharPositionInLine():0)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(337);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				 _localctx.trns = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(339);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				 _localctx.trns = instructions.NewReturn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(341);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(342);
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
	}

	public final VectorstmtContext vectorstmt() throws RecognitionException {
		VectorstmtContext _localctx = new VectorstmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_vectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(347);
			((VectorstmtContext)_localctx).VAR = match(VAR);
			setState(348);
			((VectorstmtContext)_localctx).ID = match(ID);
			setState(349);
			match(COLON);
			setState(350);
			match(CORCHIZQ);
			setState(351);
			((VectorstmtContext)_localctx).tipo = tipo();
			setState(352);
			match(CORCHDER);
			setState(353);
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
	}

	public final DefinestmtContext definestmt() throws RecognitionException {
		DefinestmtContext _localctx = new DefinestmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_definestmt);

		    var defVecInterfaces []interface{}

		int _la;
		try {
			setState(370);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(356);
				match(IG);
				setState(357);
				match(CORCHIZQ);
				setState(359); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(358);
					((DefinestmtContext)_localctx).listexpr = listexpr();
					((DefinestmtContext)_localctx).lista.add(((DefinestmtContext)_localctx).listexpr);
					}
					}
					setState(361); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0) );
				setState(363);
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
				setState(366);
				match(IG);
				setState(367);
				match(CORCHIZQ);
				setState(368);
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
	}

	public final ListexprContext listexpr() throws RecognitionException {
		ListexprContext _localctx = new ListexprContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_listexpr);
		try {
			setState(379);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(372);
				((ListexprContext)_localctx).expr = expr(0);
				setState(373);
				match(COMA);
				 _localctx.liste = ((ListexprContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(376);
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
	}

	public final MethodvecContext methodvec() throws RecognitionException {
		MethodvecContext _localctx = new MethodvecContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_methodvec);
		try {
			setState(403);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(381);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(382);
				match(PUNTO);
				setState(383);
				((MethodvecContext)_localctx).APPEND = match(APPEND);
				setState(384);
				match(PARIZQ);
				setState(385);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(386);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), ((MethodvecContext)_localctx).expr.e, (((MethodvecContext)_localctx).APPEND!=null?((MethodvecContext)_localctx).APPEND.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(389);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(390);
				match(PUNTO);
				setState(391);
				((MethodvecContext)_localctx).REMOVELAST = match(REMOVELAST);
				setState(392);
				match(PARIZQ);
				setState(393);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), nil, (((MethodvecContext)_localctx).REMOVELAST!=null?((MethodvecContext)_localctx).REMOVELAST.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(395);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(396);
				match(PUNTO);
				setState(397);
				((MethodvecContext)_localctx).REMOVE = match(REMOVE);
				setState(398);
				match(PARIZQ);
				setState(399);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(400);
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
	}

	public final MethodvecrtrnContext methodvecrtrn() throws RecognitionException {
		MethodvecrtrnContext _localctx = new MethodvecrtrnContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_methodvecrtrn);
		try {
			setState(419);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(405);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(406);
				match(PUNTO);
				setState(407);
				((MethodvecrtrnContext)_localctx).EMPTY = match(EMPTY);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).EMPTY!=null?((MethodvecrtrnContext)_localctx).EMPTY.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(409);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(410);
				match(PUNTO);
				setState(411);
				((MethodvecrtrnContext)_localctx).COUNT = match(COUNT);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).COUNT!=null?((MethodvecrtrnContext)_localctx).COUNT.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(413);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(414);
				match(CORCHIZQ);
				setState(415);
				((MethodvecrtrnContext)_localctx).expr = expr(0);
				setState(416);
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
	}

	public final VecaccessContext vecaccess() throws RecognitionException {
		VecaccessContext _localctx = new VecaccessContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_vecaccess);
		try {
			setState(440);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(421);
				((VecaccessContext)_localctx).firstId = match(ID);
				setState(422);
				match(CORCHIZQ);
				setState(423);
				((VecaccessContext)_localctx).first = expr(0);
				setState(424);
				match(CORCHDER);
				setState(425);
				match(IG);
				setState(426);
				((VecaccessContext)_localctx).secondId = match(ID);
				setState(427);
				match(CORCHIZQ);
				setState(428);
				((VecaccessContext)_localctx).second = expr(0);
				setState(429);
				match(CORCHDER);
				 
				    _localctx.vecacc = instructions.NewVectorAsgmt((((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getCharPositionInLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getText():null), ((VecaccessContext)_localctx).first.e, ((VecaccessContext)_localctx).second.e, (((VecaccessContext)_localctx).secondId!=null?((VecaccessContext)_localctx).secondId.getText():null)) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(432);
				((VecaccessContext)_localctx).ID = match(ID);
				setState(433);
				match(CORCHIZQ);
				setState(434);
				((VecaccessContext)_localctx).first = expr(0);
				setState(435);
				match(CORCHDER);
				setState(436);
				match(IG);
				setState(437);
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
	}

	public final FuncstmtContext funcstmt() throws RecognitionException {
		FuncstmtContext _localctx = new FuncstmtContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_funcstmt);

		    var args []interface{}

		int _la;
		try {
			setState(492);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(443);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(444);
				match(PARIZQ);
				setState(446); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(445);
					((FuncstmtContext)_localctx).listparam = listparam();
					((FuncstmtContext)_localctx).lista.add(((FuncstmtContext)_localctx).listparam);
					}
					}
					setState(448); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==CAME || _la==ID );
				setState(450);
				match(PARDER);
				setState(451);
				match(ARROW);
				setState(452);
				((FuncstmtContext)_localctx).tipo = tipo();
				setState(453);
				match(LLAVEIZQ);
				setState(454);
				((FuncstmtContext)_localctx).block = block();
				setState(455);
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
				setState(458);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(459);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(460);
				match(PARIZQ);
				setState(462); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(461);
					((FuncstmtContext)_localctx).listparam = listparam();
					((FuncstmtContext)_localctx).lista.add(((FuncstmtContext)_localctx).listparam);
					}
					}
					setState(464); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==CAME || _la==ID );
				setState(466);
				match(PARDER);
				setState(467);
				match(LLAVEIZQ);
				setState(468);
				((FuncstmtContext)_localctx).block = block();
				setState(469);
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
				setState(472);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(473);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(474);
				match(PARIZQ);
				setState(475);
				match(PARDER);
				setState(476);
				match(ARROW);
				setState(477);
				((FuncstmtContext)_localctx).tipo = tipo();
				setState(478);
				match(LLAVEIZQ);
				setState(479);
				((FuncstmtContext)_localctx).block = block();
				setState(480);
				match(LLAVEDER);
				 
				    _localctx.funcinstr = instructions.NewFunction((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), nil, ((FuncstmtContext)_localctx).tipo.rtipo, ((FuncstmtContext)_localctx).block.blk) 

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(483);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(484);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(485);
				match(PARIZQ);
				setState(486);
				match(PARDER);
				setState(487);
				match(LLAVEIZQ);
				setState(488);
				((FuncstmtContext)_localctx).block = block();
				setState(489);
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
	}

	public final ListparamContext listparam() throws RecognitionException {
		ListparamContext _localctx = new ListparamContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_listparam);
		int _la;
		try {
			setState(517);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(495);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
				case 1:
					{
					setState(494);
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
				setState(497);
				((ListparamContext)_localctx).inter = match(ID);
				setState(498);
				match(COLON);
				setState(500);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(499);
					((ListparamContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(502);
				((ListparamContext)_localctx).tipo = tipo();
				setState(503);
				match(COMA);
				 
				    var flag bool
				    var aux string
				    if ((ListparamContext)_localctx).INOUT != nil {
				        flag = true
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
				setState(507);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
				case 1:
					{
					setState(506);
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
				setState(509);
				((ListparamContext)_localctx).inter = match(ID);
				setState(510);
				match(COLON);
				setState(512);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(511);
					((ListparamContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(514);
				((ListparamContext)_localctx).tipo = tipo();
				 
				    var flag bool
				    var aux string
				    if ((ListparamContext)_localctx).INOUT != nil {
				        flag = true
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
	}

	public final CallfuncContext callfunc() throws RecognitionException {
		CallfuncContext _localctx = new CallfuncContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_callfunc);

		    var args []interface{}

		int _la;
		try {
			setState(533);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(519);
				((CallfuncContext)_localctx).ID = match(ID);
				setState(520);
				match(PARIZQ);
				setState(522); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(521);
					((CallfuncContext)_localctx).listparamcall = listparamcall();
					((CallfuncContext)_localctx).lista.add(((CallfuncContext)_localctx).listparamcall);
					}
					}
					setState(524); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0) || _la==AMP );
				setState(526);
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
				setState(529);
				((CallfuncContext)_localctx).ID = match(ID);
				setState(530);
				match(PARIZQ);
				setState(531);
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
	}

	public final CallfuncinstructionContext callfuncinstruction() throws RecognitionException {
		CallfuncinstructionContext _localctx = new CallfuncinstructionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_callfuncinstruction);

		    var args []interface{}

		int _la;
		try {
			setState(549);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(535);
				((CallfuncinstructionContext)_localctx).ID = match(ID);
				setState(536);
				match(PARIZQ);
				setState(538); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(537);
					((CallfuncinstructionContext)_localctx).listparamcall = listparamcall();
					((CallfuncinstructionContext)_localctx).lista.add(((CallfuncinstructionContext)_localctx).listparamcall);
					}
					}
					setState(540); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0) || _la==AMP );
				setState(542);
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
				setState(545);
				((CallfuncinstructionContext)_localctx).ID = match(ID);
				setState(546);
				match(PARIZQ);
				setState(547);
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
	}

	public final ListparamcallContext listparamcall() throws RecognitionException {
		ListparamcallContext _localctx = new ListparamcallContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_listparamcall);
		int _la;
		try {
			setState(572);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(553);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
				case 1:
					{
					setState(551);
					((ListparamcallContext)_localctx).ID = match(ID);
					setState(552);
					match(COLON);
					}
					break;
				}
				setState(556);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMP) {
					{
					setState(555);
					((ListparamcallContext)_localctx).AMP = match(AMP);
					}
				}

				setState(558);
				((ListparamcallContext)_localctx).expr = expr(0);
				setState(559);
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
				setState(564);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
				case 1:
					{
					setState(562);
					((ListparamcallContext)_localctx).ID = match(ID);
					setState(563);
					match(COLON);
					}
					break;
				}
				setState(567);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMP) {
					{
					setState(566);
					((ListparamcallContext)_localctx).AMP = match(AMP);
					}
				}

				setState(569);
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
	}

	public final FuncembedContext funcembed() throws RecognitionException {
		FuncembedContext _localctx = new FuncembedContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_funcembed);
		try {
			setState(592);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STR:
				enterOuterAlt(_localctx, 1);
				{
				setState(574);
				((FuncembedContext)_localctx).STR = match(STR);
				setState(575);
				match(PARIZQ);
				setState(576);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(577);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewString((((FuncembedContext)_localctx).STR!=null?((FuncembedContext)_localctx).STR.getLine():0), (((FuncembedContext)_localctx).STR!=null?((FuncembedContext)_localctx).STR.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			case INT:
				enterOuterAlt(_localctx, 2);
				{
				setState(580);
				((FuncembedContext)_localctx).INT = match(INT);
				setState(581);
				match(PARIZQ);
				setState(582);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(583);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewInteger((((FuncembedContext)_localctx).INT!=null?((FuncembedContext)_localctx).INT.getLine():0), (((FuncembedContext)_localctx).INT!=null?((FuncembedContext)_localctx).INT.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(586);
				((FuncembedContext)_localctx).FLOAT = match(FLOAT);
				setState(587);
				match(PARIZQ);
				setState(588);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(589);
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
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
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
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 62;
		enterRecursionRule(_localctx, 62, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(625);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				{
				setState(595);
				((ExprContext)_localctx).op = match(SUB);
				setState(596);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, "unario", expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), -1, environment.INTEGER)) 
				}
				break;
			case 2:
				{
				setState(599);
				match(PARIZQ);
				setState(600);
				((ExprContext)_localctx).expr = expr(0);
				setState(601);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 3:
				{
				setState(604);
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
			case 4:
				{
				setState(606);
				((ExprContext)_localctx).CHAR = match(CHAR);
				 
				        str := (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getLine():0), (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getCharPositionInLine():0), str[1:len(str)-1], environment.CHAR) 
				    
				}
				break;
			case 5:
				{
				setState(608);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(610);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(612);
				((ExprContext)_localctx).callfunc = callfunc();
				 _localctx.e = ((ExprContext)_localctx).callfunc.callfuncexpr 
				}
				break;
			case 8:
				{
				setState(615);
				((ExprContext)_localctx).funcembed = funcembed();
				 _localctx.e = ((ExprContext)_localctx).funcembed.funcembedexpr 
				}
				break;
			case 9:
				{
				setState(618);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 10:
				{
				setState(620);
				((ExprContext)_localctx).ID = match(ID);
				 _localctx.e = expressions.NewVar((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 11:
				{
				setState(622);
				((ExprContext)_localctx).methodvecrtrn = methodvecrtrn();
				 _localctx.e = ((ExprContext)_localctx).methodvecrtrn.methodinstrtrn 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(664);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(662);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(627);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(628);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MOD) | (1L << MUL) | (1L << DIV))) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(629);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(632);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(633);
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
						setState(634);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(637);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(638);
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
						setState(639);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(642);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(643);
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
						setState(644);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(647);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(648);
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
						setState(649);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(652);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(653);
						match(AND);
						setState(654);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(657);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(658);
						match(OR);
						setState(659);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(666);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
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
		case 31:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 17);
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		case 5:
			return precpred(_ctx, 12);
		case 6:
			return precpred(_ctx, 11);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3G\u029e\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\3\2\3\2\3\2\3\2\3\3\6\3H\n\3\r\3\16\3I\3\3\3\3\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\5\4{\n\4\3\5\3\5\3\5\6\5\u0080\n\5\r\5\16\5\u0081"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008e\n\6\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\6\7\u009d\n\7\r\7\16\7\u009e\3\7"+
		"\3\7\5\7\u00a3\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00b0"+
		"\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\5\n\u00d4\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u00e0\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\5\f\u00f3\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\5\r\u0103\n\r\3\16\3\16\3\16\3\16\6\16\u0109\n\16\r\16\16\16"+
		"\u010a\3\16\5\16\u010e\n\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\5"+
		"\17\u0118\n\17\3\17\3\17\3\20\3\20\3\20\3\20\5\20\u0120\n\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0148\n\22\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\5\24\u015c\n\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\6\26\u016a\n\26\r\26\16\26\u016b\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\5\26\u0175\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u017e"+
		"\n\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\5\30\u0196\n\30\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u01a6"+
		"\n\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u01bb\n\32\3\33\3\33\3\33\3\33\6\33"+
		"\u01c1\n\33\r\33\16\33\u01c2\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\6\33\u01d1\n\33\r\33\16\33\u01d2\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u01ef\n\33\3\34\5\34\u01f2"+
		"\n\34\3\34\3\34\3\34\5\34\u01f7\n\34\3\34\3\34\3\34\3\34\3\34\5\34\u01fe"+
		"\n\34\3\34\3\34\3\34\5\34\u0203\n\34\3\34\3\34\3\34\5\34\u0208\n\34\3"+
		"\35\3\35\3\35\6\35\u020d\n\35\r\35\16\35\u020e\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\5\35\u0218\n\35\3\36\3\36\3\36\6\36\u021d\n\36\r\36\16\36\u021e"+
		"\3\36\3\36\3\36\3\36\3\36\3\36\3\36\5\36\u0228\n\36\3\37\3\37\5\37\u022c"+
		"\n\37\3\37\5\37\u022f\n\37\3\37\3\37\3\37\3\37\3\37\3\37\5\37\u0237\n"+
		"\37\3\37\5\37\u023a\n\37\3\37\3\37\3\37\5\37\u023f\n\37\3 \3 \3 \3 \3"+
		" \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \5 \u0253\n \3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\5!\u0274\n!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\7!\u0299\n!\f"+
		"!\16!\u029c\13!\3!\2\3@\"\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&"+
		"(*,.\60\62\64\668:<>@\2\b\4\2%%))\3\2\64\66\3\2\678\4\2\60\60\62\62\4"+
		"\2\61\61\63\63\3\2*+\2\u02d2\2B\3\2\2\2\4G\3\2\2\2\6z\3\2\2\2\b|\3\2\2"+
		"\2\n\u008d\3\2\2\2\f\u00af\3\2\2\2\16\u00b1\3\2\2\2\20\u00b9\3\2\2\2\22"+
		"\u00d3\3\2\2\2\24\u00df\3\2\2\2\26\u00f2\3\2\2\2\30\u0102\3\2\2\2\32\u0104"+
		"\3\2\2\2\34\u0112\3\2\2\2\36\u011b\3\2\2\2 \u0123\3\2\2\2\"\u0147\3\2"+
		"\2\2$\u0149\3\2\2\2&\u015b\3\2\2\2(\u015d\3\2\2\2*\u0174\3\2\2\2,\u017d"+
		"\3\2\2\2.\u0195\3\2\2\2\60\u01a5\3\2\2\2\62\u01ba\3\2\2\2\64\u01ee\3\2"+
		"\2\2\66\u0207\3\2\2\28\u0217\3\2\2\2:\u0227\3\2\2\2<\u023e\3\2\2\2>\u0252"+
		"\3\2\2\2@\u0273\3\2\2\2BC\5\4\3\2CD\7\2\2\3DE\b\2\1\2E\3\3\2\2\2FH\5\6"+
		"\4\2GF\3\2\2\2HI\3\2\2\2IG\3\2\2\2IJ\3\2\2\2JK\3\2\2\2KL\b\3\1\2L\5\3"+
		"\2\2\2MN\5\b\5\2NO\b\4\1\2O{\3\2\2\2PQ\5:\36\2QR\b\4\1\2R{\3\2\2\2ST\5"+
		"\f\7\2TU\b\4\1\2U{\3\2\2\2VW\5\22\n\2WX\b\4\1\2X{\3\2\2\2YZ\5\26\f\2Z"+
		"[\b\4\1\2[{\3\2\2\2\\]\5\30\r\2]^\b\4\1\2^{\3\2\2\2_`\5\32\16\2`a\b\4"+
		"\1\2a{\3\2\2\2bc\5 \21\2cd\b\4\1\2d{\3\2\2\2ef\5\"\22\2fg\b\4\1\2g{\3"+
		"\2\2\2hi\5$\23\2ij\b\4\1\2j{\3\2\2\2kl\5&\24\2lm\b\4\1\2m{\3\2\2\2no\5"+
		"(\25\2op\b\4\1\2p{\3\2\2\2qr\5.\30\2rs\b\4\1\2s{\3\2\2\2tu\5\62\32\2u"+
		"v\b\4\1\2v{\3\2\2\2wx\5\64\33\2xy\b\4\1\2y{\3\2\2\2zM\3\2\2\2zP\3\2\2"+
		"\2zS\3\2\2\2zV\3\2\2\2zY\3\2\2\2z\\\3\2\2\2z_\3\2\2\2zb\3\2\2\2ze\3\2"+
		"\2\2zh\3\2\2\2zk\3\2\2\2zn\3\2\2\2zq\3\2\2\2zt\3\2\2\2zw\3\2\2\2{\7\3"+
		"\2\2\2|}\7\r\2\2}\177\79\2\2~\u0080\5\n\6\2\177~\3\2\2\2\u0080\u0081\3"+
		"\2\2\2\u0081\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0083\3\2\2\2\u0083"+
		"\u0084\7:\2\2\u0084\u0085\b\5\1\2\u0085\t\3\2\2\2\u0086\u0087\5@!\2\u0087"+
		"\u0088\7A\2\2\u0088\u0089\b\6\1\2\u0089\u008e\3\2\2\2\u008a\u008b\5@!"+
		"\2\u008b\u008c\b\6\1\2\u008c\u008e\3\2\2\2\u008d\u0086\3\2\2\2\u008d\u008a"+
		"\3\2\2\2\u008e\13\3\2\2\2\u008f\u0090\7\16\2\2\u0090\u0091\5@!\2\u0091"+
		"\u0092\7;\2\2\u0092\u0093\5\4\3\2\u0093\u0094\7<\2\2\u0094\u0095\b\7\1"+
		"\2\u0095\u00b0\3\2\2\2\u0096\u0097\7\16\2\2\u0097\u0098\5@!\2\u0098\u0099"+
		"\7;\2\2\u0099\u009a\5\4\3\2\u009a\u009c\7<\2\2\u009b\u009d\5\16\b\2\u009c"+
		"\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f\3\2"+
		"\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a2\7\17\2\2\u00a1\u00a3\5\20\t\2\u00a2"+
		"\u00a1\3\2\2\2\u00a2\u00a3\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a5\b\7"+
		"\1\2\u00a5\u00b0\3\2\2\2\u00a6\u00a7\7\16\2\2\u00a7\u00a8\5@!\2\u00a8"+
		"\u00a9\7;\2\2\u00a9\u00aa\5\4\3\2\u00aa\u00ab\7<\2\2\u00ab\u00ac\7\17"+
		"\2\2\u00ac\u00ad\5\20\t\2\u00ad\u00ae\b\7\1\2\u00ae\u00b0\3\2\2\2\u00af"+
		"\u008f\3\2\2\2\u00af\u0096\3\2\2\2\u00af\u00a6\3\2\2\2\u00b0\r\3\2\2\2"+
		"\u00b1\u00b2\7\17\2\2\u00b2\u00b3\7\16\2\2\u00b3\u00b4\5@!\2\u00b4\u00b5"+
		"\7;\2\2\u00b5\u00b6\5\4\3\2\u00b6\u00b7\7<\2\2\u00b7\u00b8\b\b\1\2\u00b8"+
		"\17\3\2\2\2\u00b9\u00ba\7;\2\2\u00ba\u00bb\5\4\3\2\u00bb\u00bc\7<\2\2"+
		"\u00bc\u00bd\b\t\1\2\u00bd\21\3\2\2\2\u00be\u00bf\7\13\2\2\u00bf\u00c0"+
		"\7)\2\2\u00c0\u00c1\7C\2\2\u00c1\u00c2\5\24\13\2\u00c2\u00c3\7/\2\2\u00c3"+
		"\u00c4\5@!\2\u00c4\u00c5\b\n\1\2\u00c5\u00d4\3\2\2\2\u00c6\u00c7\7\13"+
		"\2\2\u00c7\u00c8\7)\2\2\u00c8\u00c9\7/\2\2\u00c9\u00ca\5@!\2\u00ca\u00cb"+
		"\b\n\1\2\u00cb\u00d4\3\2\2\2\u00cc\u00cd\7\13\2\2\u00cd\u00ce\7)\2\2\u00ce"+
		"\u00cf\7C\2\2\u00cf\u00d0\5\24\13\2\u00d0\u00d1\7?\2\2\u00d1\u00d2\b\n"+
		"\1\2\u00d2\u00d4\3\2\2\2\u00d3\u00be\3\2\2\2\u00d3\u00c6\3\2\2\2\u00d3"+
		"\u00cc\3\2\2\2\u00d4\23\3\2\2\2\u00d5\u00d6\7\3\2\2\u00d6\u00e0\b\13\1"+
		"\2\u00d7\u00d8\7\4\2\2\u00d8\u00e0\b\13\1\2\u00d9\u00da\7\6\2\2\u00da"+
		"\u00e0\b\13\1\2\u00db\u00dc\7\5\2\2\u00dc\u00e0\b\13\1\2\u00dd\u00de\7"+
		"\7\2\2\u00de\u00e0\b\13\1\2\u00df\u00d5\3\2\2\2\u00df\u00d7\3\2\2\2\u00df"+
		"\u00d9\3\2\2\2\u00df\u00db\3\2\2\2\u00df\u00dd\3\2\2\2\u00e0\25\3\2\2"+
		"\2\u00e1\u00e2\7)\2\2\u00e2\u00e3\7/\2\2\u00e3\u00e4\5@!\2\u00e4\u00e5"+
		"\b\f\1\2\u00e5\u00f3\3\2\2\2\u00e6\u00e7\7)\2\2\u00e7\u00e8\7\67\2\2\u00e8"+
		"\u00e9\7/\2\2\u00e9\u00ea\5@!\2\u00ea\u00eb\b\f\1\2\u00eb\u00f3\3\2\2"+
		"\2\u00ec\u00ed\7)\2\2\u00ed\u00ee\78\2\2\u00ee\u00ef\7/\2\2\u00ef\u00f0"+
		"\5@!\2\u00f0\u00f1\b\f\1\2\u00f1\u00f3\3\2\2\2\u00f2\u00e1\3\2\2\2\u00f2"+
		"\u00e6\3\2\2\2\u00f2\u00ec\3\2\2\2\u00f3\27\3\2\2\2\u00f4\u00f5\7\f\2"+
		"\2\u00f5\u00f6\7)\2\2\u00f6\u00f7\7/\2\2\u00f7\u00f8\5@!\2\u00f8\u00f9"+
		"\b\r\1\2\u00f9\u0103\3\2\2\2\u00fa\u00fb\7\f\2\2\u00fb\u00fc\7)\2\2\u00fc"+
		"\u00fd\7C\2\2\u00fd\u00fe\5\24\13\2\u00fe\u00ff\7/\2\2\u00ff\u0100\5@"+
		"!\2\u0100\u0101\b\r\1\2\u0101\u0103\3\2\2\2\u0102\u00f4\3\2\2\2\u0102"+
		"\u00fa\3\2\2\2\u0103\31\3\2\2\2\u0104\u0105\7\20\2\2\u0105\u0106\5@!\2"+
		"\u0106\u0108\7;\2\2\u0107\u0109\5\34\17\2\u0108\u0107\3\2\2\2\u0109\u010a"+
		"\3\2\2\2\u010a\u0108\3\2\2\2\u010a\u010b\3\2\2\2\u010b\u010d\3\2\2\2\u010c"+
		"\u010e\5\36\20\2\u010d\u010c\3\2\2\2\u010d\u010e\3\2\2\2\u010e\u010f\3"+
		"\2\2\2\u010f\u0110\7<\2\2\u0110\u0111\b\16\1\2\u0111\33\3\2\2\2\u0112"+
		"\u0113\7\21\2\2\u0113\u0114\5@!\2\u0114\u0115\7C\2\2\u0115\u0117\5\4\3"+
		"\2\u0116\u0118\7\30\2\2\u0117\u0116\3\2\2\2\u0117\u0118\3\2\2\2\u0118"+
		"\u0119\3\2\2\2\u0119\u011a\b\17\1\2\u011a\35\3\2\2\2\u011b\u011c\7\22"+
		"\2\2\u011c\u011d\7C\2\2\u011d\u011f\5\4\3\2\u011e\u0120\7\30\2\2\u011f"+
		"\u011e\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u0121\3\2\2\2\u0121\u0122\b\20"+
		"\1\2\u0122\37\3\2\2\2\u0123\u0124\7\23\2\2\u0124\u0125\5@!\2\u0125\u0126"+
		"\7;\2\2\u0126\u0127\5\4\3\2\u0127\u0128\7<\2\2\u0128\u0129\b\21\1\2\u0129"+
		"!\3\2\2\2\u012a\u012b\7\24\2\2\u012b\u012c\7)\2\2\u012c\u012d\7\25\2\2"+
		"\u012d\u012e\7(\2\2\u012e\u012f\7;\2\2\u012f\u0130\5\4\3\2\u0130\u0131"+
		"\7<\2\2\u0131\u0132\b\22\1\2\u0132\u0148\3\2\2\2\u0133\u0134\7\24\2\2"+
		"\u0134\u0135\7)\2\2\u0135\u0136\7\25\2\2\u0136\u0137\5@!\2\u0137\u0138"+
		"\7\26\2\2\u0138\u0139\5@!\2\u0139\u013a\7;\2\2\u013a\u013b\5\4\3\2\u013b"+
		"\u013c\7<\2\2\u013c\u013d\b\22\1\2\u013d\u0148\3\2\2\2\u013e\u013f\7\24"+
		"\2\2\u013f\u0140\7)\2\2\u0140\u0141\7\25\2\2\u0141\u0142\7)\2\2\u0142"+
		"\u0143\7;\2\2\u0143\u0144\5\4\3\2\u0144\u0145\7<\2\2\u0145\u0146\b\22"+
		"\1\2\u0146\u0148\3\2\2\2\u0147\u012a\3\2\2\2\u0147\u0133\3\2\2\2\u0147"+
		"\u013e\3\2\2\2\u0148#\3\2\2\2\u0149\u014a\7\27\2\2\u014a\u014b\5@!\2\u014b"+
		"\u014c\7\17\2\2\u014c\u014d\7;\2\2\u014d\u014e\5\4\3\2\u014e\u014f\7<"+
		"\2\2\u014f\u0150\b\23\1\2\u0150%\3\2\2\2\u0151\u0152\7\30\2\2\u0152\u015c"+
		"\b\24\1\2\u0153\u0154\7\31\2\2\u0154\u015c\b\24\1\2\u0155\u0156\7\32\2"+
		"\2\u0156\u015c\b\24\1\2\u0157\u0158\7\32\2\2\u0158\u0159\5@!\2\u0159\u015a"+
		"\b\24\1\2\u015a\u015c\3\2\2\2\u015b\u0151\3\2\2\2\u015b\u0153\3\2\2\2"+
		"\u015b\u0155\3\2\2\2\u015b\u0157\3\2\2\2\u015c\'\3\2\2\2\u015d\u015e\7"+
		"\13\2\2\u015e\u015f\7)\2\2\u015f\u0160\7C\2\2\u0160\u0161\7=\2\2\u0161"+
		"\u0162\5\24\13\2\u0162\u0163\7>\2\2\u0163\u0164\5*\26\2\u0164\u0165\b"+
		"\25\1\2\u0165)\3\2\2\2\u0166\u0167\7/\2\2\u0167\u0169\7=\2\2\u0168\u016a"+
		"\5,\27\2\u0169\u0168\3\2\2\2\u016a\u016b\3\2\2\2\u016b\u0169\3\2\2\2\u016b"+
		"\u016c\3\2\2\2\u016c\u016d\3\2\2\2\u016d\u016e\7>\2\2\u016e\u016f\b\26"+
		"\1\2\u016f\u0175\3\2\2\2\u0170\u0171\7/\2\2\u0171\u0172\7=\2\2\u0172\u0173"+
		"\7>\2\2\u0173\u0175\b\26\1\2\u0174\u0166\3\2\2\2\u0174\u0170\3\2\2\2\u0175"+
		"+\3\2\2\2\u0176\u0177\5@!\2\u0177\u0178\7A\2\2\u0178\u0179\b\27\1\2\u0179"+
		"\u017e\3\2\2\2\u017a\u017b\5@!\2\u017b\u017c\b\27\1\2\u017c\u017e\3\2"+
		"\2\2\u017d\u0176\3\2\2\2\u017d\u017a\3\2\2\2\u017e-\3\2\2\2\u017f\u0180"+
		"\7)\2\2\u0180\u0181\7B\2\2\u0181\u0182\7\37\2\2\u0182\u0183\79\2\2\u0183"+
		"\u0184\5@!\2\u0184\u0185\7:\2\2\u0185\u0186\b\30\1\2\u0186\u0196\3\2\2"+
		"\2\u0187\u0188\7)\2\2\u0188\u0189\7B\2\2\u0189\u018a\7 \2\2\u018a\u018b"+
		"\79\2\2\u018b\u018c\7:\2\2\u018c\u0196\b\30\1\2\u018d\u018e\7)\2\2\u018e"+
		"\u018f\7B\2\2\u018f\u0190\7!\2\2\u0190\u0191\79\2\2\u0191\u0192\5@!\2"+
		"\u0192\u0193\7:\2\2\u0193\u0194\b\30\1\2\u0194\u0196\3\2\2\2\u0195\u017f"+
		"\3\2\2\2\u0195\u0187\3\2\2\2\u0195\u018d\3\2\2\2\u0196/\3\2\2\2\u0197"+
		"\u0198\7)\2\2\u0198\u0199\7B\2\2\u0199\u019a\7\"\2\2\u019a\u01a6\b\31"+
		"\1\2\u019b\u019c\7)\2\2\u019c\u019d\7B\2\2\u019d\u019e\7#\2\2\u019e\u01a6"+
		"\b\31\1\2\u019f\u01a0\7)\2\2\u01a0\u01a1\7=\2\2\u01a1\u01a2\5@!\2\u01a2"+
		"\u01a3\7>\2\2\u01a3\u01a4\b\31\1\2\u01a4\u01a6\3\2\2\2\u01a5\u0197\3\2"+
		"\2\2\u01a5\u019b\3\2\2\2\u01a5\u019f\3\2\2\2\u01a6\61\3\2\2\2\u01a7\u01a8"+
		"\7)\2\2\u01a8\u01a9\7=\2\2\u01a9\u01aa\5@!\2\u01aa\u01ab\7>\2\2\u01ab"+
		"\u01ac\7/\2\2\u01ac\u01ad\7)\2\2\u01ad\u01ae\7=\2\2\u01ae\u01af\5@!\2"+
		"\u01af\u01b0\7>\2\2\u01b0\u01b1\b\32\1\2\u01b1\u01bb\3\2\2\2\u01b2\u01b3"+
		"\7)\2\2\u01b3\u01b4\7=\2\2\u01b4\u01b5\5@!\2\u01b5\u01b6\7>\2\2\u01b6"+
		"\u01b7\7/\2\2\u01b7\u01b8\5@!\2\u01b8\u01b9\b\32\1\2\u01b9\u01bb\3\2\2"+
		"\2\u01ba\u01a7\3\2\2\2\u01ba\u01b2\3\2\2\2\u01bb\63\3\2\2\2\u01bc\u01bd"+
		"\7\36\2\2\u01bd\u01be\7)\2\2\u01be\u01c0\79\2\2\u01bf\u01c1\5\66\34\2"+
		"\u01c0\u01bf\3\2\2\2\u01c1\u01c2\3\2\2\2\u01c2\u01c0\3\2\2\2\u01c2\u01c3"+
		"\3\2\2\2\u01c3\u01c4\3\2\2\2\u01c4\u01c5\7:\2\2\u01c5\u01c6\7@\2\2\u01c6"+
		"\u01c7\5\24\13\2\u01c7\u01c8\7;\2\2\u01c8\u01c9\5\4\3\2\u01c9\u01ca\7"+
		"<\2\2\u01ca\u01cb\b\33\1\2\u01cb\u01ef\3\2\2\2\u01cc\u01cd\7\36\2\2\u01cd"+
		"\u01ce\7)\2\2\u01ce\u01d0\79\2\2\u01cf\u01d1\5\66\34\2\u01d0\u01cf\3\2"+
		"\2\2\u01d1\u01d2\3\2\2\2\u01d2\u01d0\3\2\2\2\u01d2\u01d3\3\2\2\2\u01d3"+
		"\u01d4\3\2\2\2\u01d4\u01d5\7:\2\2\u01d5\u01d6\7;\2\2\u01d6\u01d7\5\4\3"+
		"\2\u01d7\u01d8\7<\2\2\u01d8\u01d9\b\33\1\2\u01d9\u01ef\3\2\2\2\u01da\u01db"+
		"\7\36\2\2\u01db\u01dc\7)\2\2\u01dc\u01dd\79\2\2\u01dd\u01de\7:\2\2\u01de"+
		"\u01df\7@\2\2\u01df\u01e0\5\24\13\2\u01e0\u01e1\7;\2\2\u01e1\u01e2\5\4"+
		"\3\2\u01e2\u01e3\7<\2\2\u01e3\u01e4\b\33\1\2\u01e4\u01ef\3\2\2\2\u01e5"+
		"\u01e6\7\36\2\2\u01e6\u01e7\7)\2\2\u01e7\u01e8\79\2\2\u01e8\u01e9\7:\2"+
		"\2\u01e9\u01ea\7;\2\2\u01ea\u01eb\5\4\3\2\u01eb\u01ec\7<\2\2\u01ec\u01ed"+
		"\b\33\1\2\u01ed\u01ef\3\2\2\2\u01ee\u01bc\3\2\2\2\u01ee\u01cc\3\2\2\2"+
		"\u01ee\u01da\3\2\2\2\u01ee\u01e5\3\2\2\2\u01ef\65\3\2\2\2\u01f0\u01f2"+
		"\t\2\2\2\u01f1\u01f0\3\2\2\2\u01f1\u01f2\3\2\2\2\u01f2\u01f3\3\2\2\2\u01f3"+
		"\u01f4\7)\2\2\u01f4\u01f6\7C\2\2\u01f5\u01f7\7$\2\2\u01f6\u01f5\3\2\2"+
		"\2\u01f6\u01f7\3\2\2\2\u01f7\u01f8\3\2\2\2\u01f8\u01f9\5\24\13\2\u01f9"+
		"\u01fa\7A\2\2\u01fa\u01fb\b\34\1\2\u01fb\u0208\3\2\2\2\u01fc\u01fe\t\2"+
		"\2\2\u01fd\u01fc\3\2\2\2\u01fd\u01fe\3\2\2\2\u01fe\u01ff\3\2\2\2\u01ff"+
		"\u0200\7)\2\2\u0200\u0202\7C\2\2\u0201\u0203\7$\2\2\u0202\u0201\3\2\2"+
		"\2\u0202\u0203\3\2\2\2\u0203\u0204\3\2\2\2\u0204\u0205\5\24\13\2\u0205"+
		"\u0206\b\34\1\2\u0206\u0208\3\2\2\2\u0207\u01f1\3\2\2\2\u0207\u01fd\3"+
		"\2\2\2\u0208\67\3\2\2\2\u0209\u020a\7)\2\2\u020a\u020c\79\2\2\u020b\u020d"+
		"\5<\37\2\u020c\u020b\3\2\2\2\u020d\u020e\3\2\2\2\u020e\u020c\3\2\2\2\u020e"+
		"\u020f\3\2\2\2\u020f\u0210\3\2\2\2\u0210\u0211\7:\2\2\u0211\u0212\b\35"+
		"\1\2\u0212\u0218\3\2\2\2\u0213\u0214\7)\2\2\u0214\u0215\79\2\2\u0215\u0216"+
		"\7:\2\2\u0216\u0218\b\35\1\2\u0217\u0209\3\2\2\2\u0217\u0213\3\2\2\2\u0218"+
		"9\3\2\2\2\u0219\u021a\7)\2\2\u021a\u021c\79\2\2\u021b\u021d\5<\37\2\u021c"+
		"\u021b\3\2\2\2\u021d\u021e\3\2\2\2\u021e\u021c\3\2\2\2\u021e\u021f\3\2"+
		"\2\2\u021f\u0220\3\2\2\2\u0220\u0221\7:\2\2\u0221\u0222\b\36\1\2\u0222"+
		"\u0228\3\2\2\2\u0223\u0224\7)\2\2\u0224\u0225\79\2\2\u0225\u0226\7:\2"+
		"\2\u0226\u0228\b\36\1\2\u0227\u0219\3\2\2\2\u0227\u0223\3\2\2\2\u0228"+
		";\3\2\2\2\u0229\u022a\7)\2\2\u022a\u022c\7C\2\2\u022b\u0229\3\2\2\2\u022b"+
		"\u022c\3\2\2\2\u022c\u022e\3\2\2\2\u022d\u022f\7D\2\2\u022e\u022d\3\2"+
		"\2\2\u022e\u022f\3\2\2\2\u022f\u0230\3\2\2\2\u0230\u0231\5@!\2\u0231\u0232"+
		"\7A\2\2\u0232\u0233\b\37\1\2\u0233\u023f\3\2\2\2\u0234\u0235\7)\2\2\u0235"+
		"\u0237\7C\2\2\u0236\u0234\3\2\2\2\u0236\u0237\3\2\2\2\u0237\u0239\3\2"+
		"\2\2\u0238\u023a\7D\2\2\u0239\u0238\3\2\2\2\u0239\u023a\3\2\2\2\u023a"+
		"\u023b\3\2\2\2\u023b\u023c\5@!\2\u023c\u023d\b\37\1\2\u023d\u023f\3\2"+
		"\2\2\u023e\u022b\3\2\2\2\u023e\u0236\3\2\2\2\u023f=\3\2\2\2\u0240\u0241"+
		"\7\6\2\2\u0241\u0242\79\2\2\u0242\u0243\5@!\2\u0243\u0244\7:\2\2\u0244"+
		"\u0245\b \1\2\u0245\u0253\3\2\2\2\u0246\u0247\7\3\2\2\u0247\u0248\79\2"+
		"\2\u0248\u0249\5@!\2\u0249\u024a\7:\2\2\u024a\u024b\b \1\2\u024b\u0253"+
		"\3\2\2\2\u024c\u024d\7\4\2\2\u024d\u024e\79\2\2\u024e\u024f\5@!\2\u024f"+
		"\u0250\7:\2\2\u0250\u0251\b \1\2\u0251\u0253\3\2\2\2\u0252\u0240\3\2\2"+
		"\2\u0252\u0246\3\2\2\2\u0252\u024c\3\2\2\2\u0253?\3\2\2\2\u0254\u0255"+
		"\b!\1\2\u0255\u0256\78\2\2\u0256\u0257\5@!\24\u0257\u0258\b!\1\2\u0258"+
		"\u0274\3\2\2\2\u0259\u025a\79\2\2\u025a\u025b\5@!\2\u025b\u025c\7:\2\2"+
		"\u025c\u025d\b!\1\2\u025d\u0274\3\2\2\2\u025e\u025f\7&\2\2\u025f\u0274"+
		"\b!\1\2\u0260\u0261\7\'\2\2\u0261\u0274\b!\1\2\u0262\u0263\7(\2\2\u0263"+
		"\u0274\b!\1\2\u0264\u0265\7\b\2\2\u0265\u0274\b!\1\2\u0266\u0267\58\35"+
		"\2\u0267\u0268\b!\1\2\u0268\u0274\3\2\2\2\u0269\u026a\5> \2\u026a\u026b"+
		"\b!\1\2\u026b\u0274\3\2\2\2\u026c\u026d\7\t\2\2\u026d\u0274\b!\1\2\u026e"+
		"\u026f\7)\2\2\u026f\u0274\b!\1\2\u0270\u0271\5\60\31\2\u0271\u0272\b!"+
		"\1\2\u0272\u0274\3\2\2\2\u0273\u0254\3\2\2\2\u0273\u0259\3\2\2\2\u0273"+
		"\u025e\3\2\2\2\u0273\u0260\3\2\2\2\u0273\u0262\3\2\2\2\u0273\u0264\3\2"+
		"\2\2\u0273\u0266\3\2\2\2\u0273\u0269\3\2\2\2\u0273\u026c\3\2\2\2\u0273"+
		"\u026e\3\2\2\2\u0273\u0270\3\2\2\2\u0274\u029a\3\2\2\2\u0275\u0276\f\23"+
		"\2\2\u0276\u0277\t\3\2\2\u0277\u0278\5@!\24\u0278\u0279\b!\1\2\u0279\u0299"+
		"\3\2\2\2\u027a\u027b\f\22\2\2\u027b\u027c\t\4\2\2\u027c\u027d\5@!\23\u027d"+
		"\u027e\b!\1\2\u027e\u0299\3\2\2\2\u027f\u0280\f\21\2\2\u0280\u0281\t\5"+
		"\2\2\u0281\u0282\5@!\22\u0282\u0283\b!\1\2\u0283\u0299\3\2\2\2\u0284\u0285"+
		"\f\20\2\2\u0285\u0286\t\6\2\2\u0286\u0287\5@!\21\u0287\u0288\b!\1\2\u0288"+
		"\u0299\3\2\2\2\u0289\u028a\f\17\2\2\u028a\u028b\t\7\2\2\u028b\u028c\5"+
		"@!\20\u028c\u028d\b!\1\2\u028d\u0299\3\2\2\2\u028e\u028f\f\16\2\2\u028f"+
		"\u0290\7.\2\2\u0290\u0291\5@!\17\u0291\u0292\b!\1\2\u0292\u0299\3\2\2"+
		"\2\u0293\u0294\f\r\2\2\u0294\u0295\7-\2\2\u0295\u0296\5@!\16\u0296\u0297"+
		"\b!\1\2\u0297\u0299\3\2\2\2\u0298\u0275\3\2\2\2\u0298\u027a\3\2\2\2\u0298"+
		"\u027f\3\2\2\2\u0298\u0284\3\2\2\2\u0298\u0289\3\2\2\2\u0298\u028e\3\2"+
		"\2\2\u0298\u0293\3\2\2\2\u0299\u029c\3\2\2\2\u029a\u0298\3\2\2\2\u029a"+
		"\u029b\3\2\2\2\u029bA\3\2\2\2\u029c\u029a\3\2\2\2.Iz\u0081\u008d\u009e"+
		"\u00a2\u00af\u00d3\u00df\u00f2\u0102\u010a\u010d\u0117\u011f\u0147\u015b"+
		"\u016b\u0174\u017d\u0195\u01a5\u01ba\u01c2\u01d2\u01ee\u01f1\u01f6\u01fd"+
		"\u0202\u0207\u020e\u0217\u021e\u0227\u022b\u022e\u0236\u0239\u023e\u0252"+
		"\u0273\u0298\u029a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}