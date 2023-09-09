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
		ARROW=62, COMA=63, PUNTO=64, COLON=65, AMP=66, PUNTOCOMA=67, WHITESPACE=68, 
		COMMENT=69, LINE_COMMENT=70;
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
			"'.'", "':'", "'&'", "';'"
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
			"AMP", "PUNTOCOMA", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		public TerminalNode PUNTOCOMA() { return getToken(SwiftGrammarParser.PUNTOCOMA, 0); }
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
		int _la;
		try {
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				((InstructionContext)_localctx).printstmt = printstmt();
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(76);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(81);
				((InstructionContext)_localctx).callfuncinstruction = callfuncinstruction();
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(82);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).callfuncinstruction.callfuncinstr
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(88);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinstr 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(93);
				((InstructionContext)_localctx).varstmt = varstmt();
				setState(95);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(94);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).varstmt.var
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(99);
				((InstructionContext)_localctx).varasgmt = varasgmt();
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(100);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).varasgmt.asgmt
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(105);
				((InstructionContext)_localctx).conststmt = conststmt();
				setState(107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(106);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).conststmt.const
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(111);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(112);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).switchstmt.switchinstr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(117);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(118);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whileinstr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(123);
				((InstructionContext)_localctx).forstmt = forstmt();
				setState(125);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(124);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).forstmt.forinstr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(129);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				setState(131);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(130);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.guardinstr
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(135);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				setState(137);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(136);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).transferstmt.trns
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(141);
				((InstructionContext)_localctx).vectorstmt = vectorstmt();
				setState(143);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(142);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).vectorstmt.vectorinstr
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(147);
				((InstructionContext)_localctx).methodvec = methodvec();
				setState(149);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(148);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).methodvec.methodinstr
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(153);
				((InstructionContext)_localctx).vecaccess = vecaccess();
				setState(155);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(154);
					match(PUNTOCOMA);
					}
				}

				 _localctx.inst = ((InstructionContext)_localctx).vecaccess.vecacc
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(159);
				((InstructionContext)_localctx).funcstmt = funcstmt();
				setState(161);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTOCOMA) {
					{
					setState(160);
					match(PUNTOCOMA);
					}
				}

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
			setState(167);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(168);
			match(PARIZQ);
			setState(170); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(169);
				((PrintstmtContext)_localctx).printexprlist = printexprlist();
				((PrintstmtContext)_localctx).lista.add(((PrintstmtContext)_localctx).printexprlist);
				}
				}
				setState(172); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NIL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << NOT) | (1L << SUB) | (1L << PARIZQ))) != 0) );
			setState(174);
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
			setState(184);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(177);
				((PrintexprlistContext)_localctx).expr = expr(0);
				setState(178);
				match(COMA);
				 _localctx.prntexpr = ((PrintexprlistContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(181);
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
			setState(226);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(186);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(187);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(188);
				match(LLAVEIZQ);
				setState(189);
				((IfstmtContext)_localctx).block = block();
				setState(190);
				match(LLAVEDER);
				 _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(193);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(194);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(195);
				match(LLAVEIZQ);
				setState(196);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(197);
				match(LLAVEDER);
				setState(199); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(198);
						((IfstmtContext)_localctx).elseifstmt = elseifstmt();
						((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(201); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(203);
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
				setState(206);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(207);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(208);
				match(LLAVEIZQ);
				setState(209);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(210);
				match(LLAVEDER);
				setState(212); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(211);
					((IfstmtContext)_localctx).elseifstmt = elseifstmt();
					((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
					}
					}
					setState(214); 
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
				setState(218);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(219);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(220);
				match(LLAVEIZQ);
				setState(221);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(222);
				match(LLAVEDER);
				setState(223);
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
			setState(228);
			((ElseifstmtContext)_localctx).ELSE = match(ELSE);
			setState(229);
			match(IF);
			setState(230);
			((ElseifstmtContext)_localctx).expr = expr(0);
			setState(231);
			match(LLAVEIZQ);
			setState(232);
			((ElseifstmtContext)_localctx).block = block();
			setState(233);
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
	}

	public final ElsestmtContext elsestmt() throws RecognitionException {
		ElsestmtContext _localctx = new ElsestmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_elsestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			match(ELSE);
			setState(237);
			match(LLAVEIZQ);
			setState(238);
			((ElsestmtContext)_localctx).block = block();
			setState(239);
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
			setState(263);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(243);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(244);
				match(COLON);
				setState(245);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(246);
				match(IG);
				setState(247);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), ((VarstmtContext)_localctx).tipo.rtipo, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(250);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(251);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(252);
				match(IG);
				setState(253);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), environment.NULL, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(256);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(257);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(258);
				match(COLON);
				setState(259);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(260);
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
			setState(275);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(265);
				match(INT);
				_localctx.rtipo = 0
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(267);
				match(FLOAT);
				_localctx.rtipo = 1
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(269);
				match(STR);
				_localctx.rtipo = 2
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(271);
				match(BOOL);
				_localctx.rtipo = 3
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(273);
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
			setState(294);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(277);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(278);
				match(IG);
				setState(279);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(282);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(283);
				match(ADD);
				setState(284);
				match(IG);
				setState(285);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewOpAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e, "+")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(288);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(289);
				match(SUB);
				setState(290);
				match(IG);
				setState(291);
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
			setState(310);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(296);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(297);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(298);
				match(IG);
				setState(299);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), environment.NULL, ((ConststmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(302);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(303);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(304);
				match(COLON);
				setState(305);
				((ConststmtContext)_localctx).tipo = tipo();
				setState(306);
				match(IG);
				setState(307);
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
			setState(312);
			((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
			setState(313);
			((SwitchstmtContext)_localctx).expr = expr(0);
			setState(314);
			match(LLAVEIZQ);
			setState(316); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(315);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				((SwitchstmtContext)_localctx).casesvar.add(((SwitchstmtContext)_localctx).casestmt);
				}
				}
				setState(318); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(321);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(320);
				defaultstmt();
				}
			}

			setState(323);
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
			setState(326);
			((CasestmtContext)_localctx).CASE = match(CASE);
			setState(327);
			((CasestmtContext)_localctx).expr = expr(0);
			setState(328);
			match(COLON);
			setState(329);
			((CasestmtContext)_localctx).block = block();
			setState(331);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(330);
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
			setState(335);
			match(DEFAULT);
			setState(336);
			match(COLON);
			setState(337);
			((DefaultstmtContext)_localctx).block = block();
			setState(339);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(338);
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
			setState(343);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(344);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(345);
			match(LLAVEIZQ);
			setState(346);
			((WhilestmtContext)_localctx).block = block();
			setState(347);
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
			setState(379);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(350);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(351);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(352);
				match(IN);
				setState(353);
				((ForstmtContext)_localctx).STRING = match(STRING);
				setState(354);
				match(LLAVEIZQ);
				setState(355);
				((ForstmtContext)_localctx).block = block();
				setState(356);
				match(LLAVEDER);
				 
				    str = (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getText():null)
				    cadena = expressions.NewPrimitive((((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getLine():0), (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1], environment.STRING)
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), cadena, nil, nil, "", ((ForstmtContext)_localctx).block.blk) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(359);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(360);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(361);
				match(IN);
				setState(362);
				((ForstmtContext)_localctx).left = expr(0);
				setState(363);
				match(RANGEPTS);
				setState(364);
				((ForstmtContext)_localctx).right = expr(0);
				setState(365);
				match(LLAVEIZQ);
				setState(366);
				((ForstmtContext)_localctx).block = block();
				setState(367);
				match(LLAVEDER);
				 
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), nil, ((ForstmtContext)_localctx).left.e, ((ForstmtContext)_localctx).right.e, "", ((ForstmtContext)_localctx).block.blk)

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(370);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(371);
				((ForstmtContext)_localctx).first = match(ID);
				setState(372);
				match(IN);
				setState(373);
				((ForstmtContext)_localctx).second = match(ID);
				setState(374);
				match(LLAVEIZQ);
				setState(375);
				((ForstmtContext)_localctx).block = block();
				setState(376);
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
			setState(381);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(382);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(383);
			match(ELSE);
			setState(384);
			match(LLAVEIZQ);
			setState(385);
			((GuardstmtContext)_localctx).block = block();
			setState(386);
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
			setState(399);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(389);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				 _localctx.trns = instructions.NewBreak((((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getLine():0), (((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getCharPositionInLine():0)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(391);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				 _localctx.trns = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(393);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				 _localctx.trns = instructions.NewReturn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(395);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(396);
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
			setState(401);
			((VectorstmtContext)_localctx).VAR = match(VAR);
			setState(402);
			((VectorstmtContext)_localctx).ID = match(ID);
			setState(403);
			match(COLON);
			setState(404);
			match(CORCHIZQ);
			setState(405);
			((VectorstmtContext)_localctx).tipo = tipo();
			setState(406);
			match(CORCHDER);
			setState(407);
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
			setState(424);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(410);
				match(IG);
				setState(411);
				match(CORCHIZQ);
				setState(413); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(412);
					((DefinestmtContext)_localctx).listexpr = listexpr();
					((DefinestmtContext)_localctx).lista.add(((DefinestmtContext)_localctx).listexpr);
					}
					}
					setState(415); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NIL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << NOT) | (1L << SUB) | (1L << PARIZQ))) != 0) );
				setState(417);
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
				setState(420);
				match(IG);
				setState(421);
				match(CORCHIZQ);
				setState(422);
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
			setState(433);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(426);
				((ListexprContext)_localctx).expr = expr(0);
				setState(427);
				match(COMA);
				 _localctx.liste = ((ListexprContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(430);
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
			setState(457);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(435);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(436);
				match(PUNTO);
				setState(437);
				((MethodvecContext)_localctx).APPEND = match(APPEND);
				setState(438);
				match(PARIZQ);
				setState(439);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(440);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), ((MethodvecContext)_localctx).expr.e, (((MethodvecContext)_localctx).APPEND!=null?((MethodvecContext)_localctx).APPEND.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(443);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(444);
				match(PUNTO);
				setState(445);
				((MethodvecContext)_localctx).REMOVELAST = match(REMOVELAST);
				setState(446);
				match(PARIZQ);
				setState(447);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), nil, (((MethodvecContext)_localctx).REMOVELAST!=null?((MethodvecContext)_localctx).REMOVELAST.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(449);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(450);
				match(PUNTO);
				setState(451);
				((MethodvecContext)_localctx).REMOVE = match(REMOVE);
				setState(452);
				match(PARIZQ);
				setState(453);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(454);
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
			setState(473);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(459);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(460);
				match(PUNTO);
				setState(461);
				((MethodvecrtrnContext)_localctx).EMPTY = match(EMPTY);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).EMPTY!=null?((MethodvecrtrnContext)_localctx).EMPTY.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(463);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(464);
				match(PUNTO);
				setState(465);
				((MethodvecrtrnContext)_localctx).COUNT = match(COUNT);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).COUNT!=null?((MethodvecrtrnContext)_localctx).COUNT.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(467);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(468);
				match(CORCHIZQ);
				setState(469);
				((MethodvecrtrnContext)_localctx).expr = expr(0);
				setState(470);
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
			setState(494);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(475);
				((VecaccessContext)_localctx).firstId = match(ID);
				setState(476);
				match(CORCHIZQ);
				setState(477);
				((VecaccessContext)_localctx).first = expr(0);
				setState(478);
				match(CORCHDER);
				setState(479);
				match(IG);
				setState(480);
				((VecaccessContext)_localctx).secondId = match(ID);
				setState(481);
				match(CORCHIZQ);
				setState(482);
				((VecaccessContext)_localctx).second = expr(0);
				setState(483);
				match(CORCHDER);
				 
				    _localctx.vecacc = instructions.NewVectorAsgmt((((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getCharPositionInLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getText():null), ((VecaccessContext)_localctx).first.e, ((VecaccessContext)_localctx).second.e, (((VecaccessContext)_localctx).secondId!=null?((VecaccessContext)_localctx).secondId.getText():null)) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(486);
				((VecaccessContext)_localctx).ID = match(ID);
				setState(487);
				match(CORCHIZQ);
				setState(488);
				((VecaccessContext)_localctx).first = expr(0);
				setState(489);
				match(CORCHDER);
				setState(490);
				match(IG);
				setState(491);
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
			setState(546);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(496);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(497);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(498);
				match(PARIZQ);
				setState(500); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(499);
					((FuncstmtContext)_localctx).listparam = listparam();
					((FuncstmtContext)_localctx).lista.add(((FuncstmtContext)_localctx).listparam);
					}
					}
					setState(502); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==CAME || _la==ID );
				setState(504);
				match(PARDER);
				setState(505);
				match(ARROW);
				setState(506);
				((FuncstmtContext)_localctx).tipo = tipo();
				setState(507);
				match(LLAVEIZQ);
				setState(508);
				((FuncstmtContext)_localctx).block = block();
				setState(509);
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
				setState(512);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(513);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(514);
				match(PARIZQ);
				setState(516); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(515);
					((FuncstmtContext)_localctx).listparam = listparam();
					((FuncstmtContext)_localctx).lista.add(((FuncstmtContext)_localctx).listparam);
					}
					}
					setState(518); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==CAME || _la==ID );
				setState(520);
				match(PARDER);
				setState(521);
				match(LLAVEIZQ);
				setState(522);
				((FuncstmtContext)_localctx).block = block();
				setState(523);
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
				setState(526);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(527);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(528);
				match(PARIZQ);
				setState(529);
				match(PARDER);
				setState(530);
				match(ARROW);
				setState(531);
				((FuncstmtContext)_localctx).tipo = tipo();
				setState(532);
				match(LLAVEIZQ);
				setState(533);
				((FuncstmtContext)_localctx).block = block();
				setState(534);
				match(LLAVEDER);
				 
				    _localctx.funcinstr = instructions.NewFunction((((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getLine():0), (((FuncstmtContext)_localctx).FUNC!=null?((FuncstmtContext)_localctx).FUNC.getCharPositionInLine():0), (((FuncstmtContext)_localctx).ID!=null?((FuncstmtContext)_localctx).ID.getText():null), nil, ((FuncstmtContext)_localctx).tipo.rtipo, ((FuncstmtContext)_localctx).block.blk) 

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(537);
				((FuncstmtContext)_localctx).FUNC = match(FUNC);
				setState(538);
				((FuncstmtContext)_localctx).ID = match(ID);
				setState(539);
				match(PARIZQ);
				setState(540);
				match(PARDER);
				setState(541);
				match(LLAVEIZQ);
				setState(542);
				((FuncstmtContext)_localctx).block = block();
				setState(543);
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
			setState(571);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,45,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(549);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
				case 1:
					{
					setState(548);
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
				setState(551);
				((ListparamContext)_localctx).inter = match(ID);
				setState(552);
				match(COLON);
				setState(554);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(553);
					((ListparamContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(556);
				((ListparamContext)_localctx).tipo = tipo();
				setState(557);
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
				setState(561);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
				case 1:
					{
					setState(560);
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
				setState(563);
				((ListparamContext)_localctx).inter = match(ID);
				setState(564);
				match(COLON);
				setState(566);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(565);
					((ListparamContext)_localctx).INOUT = match(INOUT);
					}
				}

				setState(568);
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
			setState(587);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,47,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(573);
				((CallfuncContext)_localctx).ID = match(ID);
				setState(574);
				match(PARIZQ);
				setState(576); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(575);
					((CallfuncContext)_localctx).listparamcall = listparamcall();
					((CallfuncContext)_localctx).lista.add(((CallfuncContext)_localctx).listparamcall);
					}
					}
					setState(578); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NIL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << NOT) | (1L << SUB) | (1L << PARIZQ))) != 0) || _la==AMP );
				setState(580);
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
				setState(583);
				((CallfuncContext)_localctx).ID = match(ID);
				setState(584);
				match(PARIZQ);
				setState(585);
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
			setState(603);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(589);
				((CallfuncinstructionContext)_localctx).ID = match(ID);
				setState(590);
				match(PARIZQ);
				setState(592); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(591);
					((CallfuncinstructionContext)_localctx).listparamcall = listparamcall();
					((CallfuncinstructionContext)_localctx).lista.add(((CallfuncinstructionContext)_localctx).listparamcall);
					}
					}
					setState(594); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << FLOAT) | (1L << STR) | (1L << TRU) | (1L << FAL) | (1L << NIL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << NOT) | (1L << SUB) | (1L << PARIZQ))) != 0) || _la==AMP );
				setState(596);
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
				setState(599);
				((CallfuncinstructionContext)_localctx).ID = match(ID);
				setState(600);
				match(PARIZQ);
				setState(601);
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
			setState(626);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(607);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,50,_ctx) ) {
				case 1:
					{
					setState(605);
					((ListparamcallContext)_localctx).ID = match(ID);
					setState(606);
					match(COLON);
					}
					break;
				}
				setState(610);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMP) {
					{
					setState(609);
					((ListparamcallContext)_localctx).AMP = match(AMP);
					}
				}

				setState(612);
				((ListparamcallContext)_localctx).expr = expr(0);
				setState(613);
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
				setState(618);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
				case 1:
					{
					setState(616);
					((ListparamcallContext)_localctx).ID = match(ID);
					setState(617);
					match(COLON);
					}
					break;
				}
				setState(621);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==AMP) {
					{
					setState(620);
					((ListparamcallContext)_localctx).AMP = match(AMP);
					}
				}

				setState(623);
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
			setState(646);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STR:
				enterOuterAlt(_localctx, 1);
				{
				setState(628);
				((FuncembedContext)_localctx).STR = match(STR);
				setState(629);
				match(PARIZQ);
				setState(630);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(631);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewString((((FuncembedContext)_localctx).STR!=null?((FuncembedContext)_localctx).STR.getLine():0), (((FuncembedContext)_localctx).STR!=null?((FuncembedContext)_localctx).STR.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			case INT:
				enterOuterAlt(_localctx, 2);
				{
				setState(634);
				((FuncembedContext)_localctx).INT = match(INT);
				setState(635);
				match(PARIZQ);
				setState(636);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(637);
				match(PARDER);
				 _localctx.funcembedexpr = expressions.NewInteger((((FuncembedContext)_localctx).INT!=null?((FuncembedContext)_localctx).INT.getLine():0), (((FuncembedContext)_localctx).INT!=null?((FuncembedContext)_localctx).INT.getCharPositionInLine():0), ((FuncembedContext)_localctx).expr.e) 
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(640);
				((FuncembedContext)_localctx).FLOAT = match(FLOAT);
				setState(641);
				match(PARIZQ);
				setState(642);
				((FuncembedContext)_localctx).expr = expr(0);
				setState(643);
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
			setState(685);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,56,_ctx) ) {
			case 1:
				{
				setState(649);
				((ExprContext)_localctx).op = match(SUB);
				setState(650);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, "unario", expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), -1, environment.INTEGER)) 
				}
				break;
			case 2:
				{
				setState(653);
				((ExprContext)_localctx).op = match(NOT);
				setState(654);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), false, environment.BOOLEAN)) 
				}
				break;
			case 3:
				{
				setState(657);
				match(PARIZQ);
				setState(658);
				((ExprContext)_localctx).expr = expr(0);
				setState(659);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(662);
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
				setState(664);
				((ExprContext)_localctx).CHAR = match(CHAR);
				 
				        str := (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getLine():0), (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getCharPositionInLine():0), str[1:len(str)-1], environment.CHAR) 
				    
				}
				break;
			case 6:
				{
				setState(666);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 7:
				{
				setState(668);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(670);
				((ExprContext)_localctx).callfunc = callfunc();
				 _localctx.e = ((ExprContext)_localctx).callfunc.callfuncexpr 
				}
				break;
			case 9:
				{
				setState(673);
				((ExprContext)_localctx).funcembed = funcembed();
				 _localctx.e = ((ExprContext)_localctx).funcembed.funcembedexpr 
				}
				break;
			case 10:
				{
				setState(676);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 11:
				{
				setState(678);
				((ExprContext)_localctx).ID = match(ID);
				 _localctx.e = expressions.NewVar((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 12:
				{
				setState(680);
				((ExprContext)_localctx).methodvecrtrn = methodvecrtrn();
				 _localctx.e = ((ExprContext)_localctx).methodvecrtrn.methodinstrtrn 
				}
				break;
			case 13:
				{
				setState(683);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), "nil", environment.NULL) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(724);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,58,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(722);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,57,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(687);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(688);
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
						setState(689);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(692);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(693);
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
						setState(694);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(697);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(698);
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
						setState(699);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(702);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(703);
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
						setState(704);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(707);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(708);
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
						setState(709);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(712);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(713);
						((ExprContext)_localctx).AND = match(AND);
						setState(714);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).AND!=null?((ExprContext)_localctx).AND.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(717);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(718);
						((ExprContext)_localctx).OR = match(OR);
						setState(719);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).OR!=null?((ExprContext)_localctx).OR.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(726);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,58,_ctx);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3H\u02da\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\3\2\3\2\3\2\3\2\3\3\6\3H\n\3\r\3\16\3I\3\3\3\3\3\4\3\4\5\4P\n\4\3"+
		"\4\3\4\3\4\3\4\5\4V\n\4\3\4\3\4\3\4\3\4\5\4\\\n\4\3\4\3\4\3\4\3\4\5\4"+
		"b\n\4\3\4\3\4\3\4\3\4\5\4h\n\4\3\4\3\4\3\4\3\4\5\4n\n\4\3\4\3\4\3\4\3"+
		"\4\5\4t\n\4\3\4\3\4\3\4\3\4\5\4z\n\4\3\4\3\4\3\4\3\4\5\4\u0080\n\4\3\4"+
		"\3\4\3\4\3\4\5\4\u0086\n\4\3\4\3\4\3\4\3\4\5\4\u008c\n\4\3\4\3\4\3\4\3"+
		"\4\5\4\u0092\n\4\3\4\3\4\3\4\3\4\5\4\u0098\n\4\3\4\3\4\3\4\3\4\5\4\u009e"+
		"\n\4\3\4\3\4\3\4\3\4\5\4\u00a4\n\4\3\4\3\4\5\4\u00a8\n\4\3\5\3\5\3\5\6"+
		"\5\u00ad\n\5\r\5\16\5\u00ae\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5"+
		"\6\u00bb\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\6\7\u00ca"+
		"\n\7\r\7\16\7\u00cb\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\6\7\u00d7\n\7"+
		"\r\7\16\7\u00d8\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00e5\n\7"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\5\n\u010a\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u0116\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\5\f\u0129\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\5\r\u0139\n\r\3\16\3\16\3\16\3\16\6\16\u013f\n\16\r\16\16\16"+
		"\u0140\3\16\5\16\u0144\n\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\5"+
		"\17\u014e\n\17\3\17\3\17\3\20\3\20\3\20\3\20\5\20\u0156\n\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u017e\n\22\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\5\24\u0192\n\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\6\26\u01a0\n\26\r\26\16\26\u01a1\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\5\26\u01ab\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\5\27\u01b4"+
		"\n\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\5\30\u01cc\n\30\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u01dc"+
		"\n\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u01f1\n\32\3\33\3\33\3\33\3\33\6\33"+
		"\u01f7\n\33\r\33\16\33\u01f8\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\6\33\u0207\n\33\r\33\16\33\u0208\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u0225\n\33\3\34\5\34\u0228"+
		"\n\34\3\34\3\34\3\34\5\34\u022d\n\34\3\34\3\34\3\34\3\34\3\34\5\34\u0234"+
		"\n\34\3\34\3\34\3\34\5\34\u0239\n\34\3\34\3\34\3\34\5\34\u023e\n\34\3"+
		"\35\3\35\3\35\6\35\u0243\n\35\r\35\16\35\u0244\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\5\35\u024e\n\35\3\36\3\36\3\36\6\36\u0253\n\36\r\36\16\36\u0254"+
		"\3\36\3\36\3\36\3\36\3\36\3\36\3\36\5\36\u025e\n\36\3\37\3\37\5\37\u0262"+
		"\n\37\3\37\5\37\u0265\n\37\3\37\3\37\3\37\3\37\3\37\3\37\5\37\u026d\n"+
		"\37\3\37\5\37\u0270\n\37\3\37\3\37\3\37\5\37\u0275\n\37\3 \3 \3 \3 \3"+
		" \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \5 \u0289\n \3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\5!\u02b0\n!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\7!\u02d5\n!\f!\16!\u02d8\13!\3!\2\3@\"\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\62\64\668:<>@\2\b\4\2%%))\3\2\64\66\3\2\678\4"+
		"\2\60\60\62\62\4\2\61\61\63\63\3\2*+\2\u0320\2B\3\2\2\2\4G\3\2\2\2\6\u00a7"+
		"\3\2\2\2\b\u00a9\3\2\2\2\n\u00ba\3\2\2\2\f\u00e4\3\2\2\2\16\u00e6\3\2"+
		"\2\2\20\u00ee\3\2\2\2\22\u0109\3\2\2\2\24\u0115\3\2\2\2\26\u0128\3\2\2"+
		"\2\30\u0138\3\2\2\2\32\u013a\3\2\2\2\34\u0148\3\2\2\2\36\u0151\3\2\2\2"+
		" \u0159\3\2\2\2\"\u017d\3\2\2\2$\u017f\3\2\2\2&\u0191\3\2\2\2(\u0193\3"+
		"\2\2\2*\u01aa\3\2\2\2,\u01b3\3\2\2\2.\u01cb\3\2\2\2\60\u01db\3\2\2\2\62"+
		"\u01f0\3\2\2\2\64\u0224\3\2\2\2\66\u023d\3\2\2\28\u024d\3\2\2\2:\u025d"+
		"\3\2\2\2<\u0274\3\2\2\2>\u0288\3\2\2\2@\u02af\3\2\2\2BC\5\4\3\2CD\7\2"+
		"\2\3DE\b\2\1\2E\3\3\2\2\2FH\5\6\4\2GF\3\2\2\2HI\3\2\2\2IG\3\2\2\2IJ\3"+
		"\2\2\2JK\3\2\2\2KL\b\3\1\2L\5\3\2\2\2MO\5\b\5\2NP\7E\2\2ON\3\2\2\2OP\3"+
		"\2\2\2PQ\3\2\2\2QR\b\4\1\2R\u00a8\3\2\2\2SU\5:\36\2TV\7E\2\2UT\3\2\2\2"+
		"UV\3\2\2\2VW\3\2\2\2WX\b\4\1\2X\u00a8\3\2\2\2Y[\5\f\7\2Z\\\7E\2\2[Z\3"+
		"\2\2\2[\\\3\2\2\2\\]\3\2\2\2]^\b\4\1\2^\u00a8\3\2\2\2_a\5\22\n\2`b\7E"+
		"\2\2a`\3\2\2\2ab\3\2\2\2bc\3\2\2\2cd\b\4\1\2d\u00a8\3\2\2\2eg\5\26\f\2"+
		"fh\7E\2\2gf\3\2\2\2gh\3\2\2\2hi\3\2\2\2ij\b\4\1\2j\u00a8\3\2\2\2km\5\30"+
		"\r\2ln\7E\2\2ml\3\2\2\2mn\3\2\2\2no\3\2\2\2op\b\4\1\2p\u00a8\3\2\2\2q"+
		"s\5\32\16\2rt\7E\2\2sr\3\2\2\2st\3\2\2\2tu\3\2\2\2uv\b\4\1\2v\u00a8\3"+
		"\2\2\2wy\5 \21\2xz\7E\2\2yx\3\2\2\2yz\3\2\2\2z{\3\2\2\2{|\b\4\1\2|\u00a8"+
		"\3\2\2\2}\177\5\"\22\2~\u0080\7E\2\2\177~\3\2\2\2\177\u0080\3\2\2\2\u0080"+
		"\u0081\3\2\2\2\u0081\u0082\b\4\1\2\u0082\u00a8\3\2\2\2\u0083\u0085\5$"+
		"\23\2\u0084\u0086\7E\2\2\u0085\u0084\3\2\2\2\u0085\u0086\3\2\2\2\u0086"+
		"\u0087\3\2\2\2\u0087\u0088\b\4\1\2\u0088\u00a8\3\2\2\2\u0089\u008b\5&"+
		"\24\2\u008a\u008c\7E\2\2\u008b\u008a\3\2\2\2\u008b\u008c\3\2\2\2\u008c"+
		"\u008d\3\2\2\2\u008d\u008e\b\4\1\2\u008e\u00a8\3\2\2\2\u008f\u0091\5("+
		"\25\2\u0090\u0092\7E\2\2\u0091\u0090\3\2\2\2\u0091\u0092\3\2\2\2\u0092"+
		"\u0093\3\2\2\2\u0093\u0094\b\4\1\2\u0094\u00a8\3\2\2\2\u0095\u0097\5."+
		"\30\2\u0096\u0098\7E\2\2\u0097\u0096\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\u0099\3\2\2\2\u0099\u009a\b\4\1\2\u009a\u00a8\3\2\2\2\u009b\u009d\5\62"+
		"\32\2\u009c\u009e\7E\2\2\u009d\u009c\3\2\2\2\u009d\u009e\3\2\2\2\u009e"+
		"\u009f\3\2\2\2\u009f\u00a0\b\4\1\2\u00a0\u00a8\3\2\2\2\u00a1\u00a3\5\64"+
		"\33\2\u00a2\u00a4\7E\2\2\u00a3\u00a2\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4"+
		"\u00a5\3\2\2\2\u00a5\u00a6\b\4\1\2\u00a6\u00a8\3\2\2\2\u00a7M\3\2\2\2"+
		"\u00a7S\3\2\2\2\u00a7Y\3\2\2\2\u00a7_\3\2\2\2\u00a7e\3\2\2\2\u00a7k\3"+
		"\2\2\2\u00a7q\3\2\2\2\u00a7w\3\2\2\2\u00a7}\3\2\2\2\u00a7\u0083\3\2\2"+
		"\2\u00a7\u0089\3\2\2\2\u00a7\u008f\3\2\2\2\u00a7\u0095\3\2\2\2\u00a7\u009b"+
		"\3\2\2\2\u00a7\u00a1\3\2\2\2\u00a8\7\3\2\2\2\u00a9\u00aa\7\r\2\2\u00aa"+
		"\u00ac\79\2\2\u00ab\u00ad\5\n\6\2\u00ac\u00ab\3\2\2\2\u00ad\u00ae\3\2"+
		"\2\2\u00ae\u00ac\3\2\2\2\u00ae\u00af\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0"+
		"\u00b1\7:\2\2\u00b1\u00b2\b\5\1\2\u00b2\t\3\2\2\2\u00b3\u00b4\5@!\2\u00b4"+
		"\u00b5\7A\2\2\u00b5\u00b6\b\6\1\2\u00b6\u00bb\3\2\2\2\u00b7\u00b8\5@!"+
		"\2\u00b8\u00b9\b\6\1\2\u00b9\u00bb\3\2\2\2\u00ba\u00b3\3\2\2\2\u00ba\u00b7"+
		"\3\2\2\2\u00bb\13\3\2\2\2\u00bc\u00bd\7\16\2\2\u00bd\u00be\5@!\2\u00be"+
		"\u00bf\7;\2\2\u00bf\u00c0\5\4\3\2\u00c0\u00c1\7<\2\2\u00c1\u00c2\b\7\1"+
		"\2\u00c2\u00e5\3\2\2\2\u00c3\u00c4\7\16\2\2\u00c4\u00c5\5@!\2\u00c5\u00c6"+
		"\7;\2\2\u00c6\u00c7\5\4\3\2\u00c7\u00c9\7<\2\2\u00c8\u00ca\5\16\b\2\u00c9"+
		"\u00c8\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\u00c9\3\2\2\2\u00cb\u00cc\3\2"+
		"\2\2\u00cc\u00cd\3\2\2\2\u00cd\u00ce\5\20\t\2\u00ce\u00cf\b\7\1\2\u00cf"+
		"\u00e5\3\2\2\2\u00d0\u00d1\7\16\2\2\u00d1\u00d2\5@!\2\u00d2\u00d3\7;\2"+
		"\2\u00d3\u00d4\5\4\3\2\u00d4\u00d6\7<\2\2\u00d5\u00d7\5\16\b\2\u00d6\u00d5"+
		"\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9"+
		"\u00da\3\2\2\2\u00da\u00db\b\7\1\2\u00db\u00e5\3\2\2\2\u00dc\u00dd\7\16"+
		"\2\2\u00dd\u00de\5@!\2\u00de\u00df\7;\2\2\u00df\u00e0\5\4\3\2\u00e0\u00e1"+
		"\7<\2\2\u00e1\u00e2\5\20\t\2\u00e2\u00e3\b\7\1\2\u00e3\u00e5\3\2\2\2\u00e4"+
		"\u00bc\3\2\2\2\u00e4\u00c3\3\2\2\2\u00e4\u00d0\3\2\2\2\u00e4\u00dc\3\2"+
		"\2\2\u00e5\r\3\2\2\2\u00e6\u00e7\7\17\2\2\u00e7\u00e8\7\16\2\2\u00e8\u00e9"+
		"\5@!\2\u00e9\u00ea\7;\2\2\u00ea\u00eb\5\4\3\2\u00eb\u00ec\7<\2\2\u00ec"+
		"\u00ed\b\b\1\2\u00ed\17\3\2\2\2\u00ee\u00ef\7\17\2\2\u00ef\u00f0\7;\2"+
		"\2\u00f0\u00f1\5\4\3\2\u00f1\u00f2\7<\2\2\u00f2\u00f3\b\t\1\2\u00f3\21"+
		"\3\2\2\2\u00f4\u00f5\7\13\2\2\u00f5\u00f6\7)\2\2\u00f6\u00f7\7C\2\2\u00f7"+
		"\u00f8\5\24\13\2\u00f8\u00f9\7/\2\2\u00f9\u00fa\5@!\2\u00fa\u00fb\b\n"+
		"\1\2\u00fb\u010a\3\2\2\2\u00fc\u00fd\7\13\2\2\u00fd\u00fe\7)\2\2\u00fe"+
		"\u00ff\7/\2\2\u00ff\u0100\5@!\2\u0100\u0101\b\n\1\2\u0101\u010a\3\2\2"+
		"\2\u0102\u0103\7\13\2\2\u0103\u0104\7)\2\2\u0104\u0105\7C\2\2\u0105\u0106"+
		"\5\24\13\2\u0106\u0107\7?\2\2\u0107\u0108\b\n\1\2\u0108\u010a\3\2\2\2"+
		"\u0109\u00f4\3\2\2\2\u0109\u00fc\3\2\2\2\u0109\u0102\3\2\2\2\u010a\23"+
		"\3\2\2\2\u010b\u010c\7\3\2\2\u010c\u0116\b\13\1\2\u010d\u010e\7\4\2\2"+
		"\u010e\u0116\b\13\1\2\u010f\u0110\7\6\2\2\u0110\u0116\b\13\1\2\u0111\u0112"+
		"\7\5\2\2\u0112\u0116\b\13\1\2\u0113\u0114\7\7\2\2\u0114\u0116\b\13\1\2"+
		"\u0115\u010b\3\2\2\2\u0115\u010d\3\2\2\2\u0115\u010f\3\2\2\2\u0115\u0111"+
		"\3\2\2\2\u0115\u0113\3\2\2\2\u0116\25\3\2\2\2\u0117\u0118\7)\2\2\u0118"+
		"\u0119\7/\2\2\u0119\u011a\5@!\2\u011a\u011b\b\f\1\2\u011b\u0129\3\2\2"+
		"\2\u011c\u011d\7)\2\2\u011d\u011e\7\67\2\2\u011e\u011f\7/\2\2\u011f\u0120"+
		"\5@!\2\u0120\u0121\b\f\1\2\u0121\u0129\3\2\2\2\u0122\u0123\7)\2\2\u0123"+
		"\u0124\78\2\2\u0124\u0125\7/\2\2\u0125\u0126\5@!\2\u0126\u0127\b\f\1\2"+
		"\u0127\u0129\3\2\2\2\u0128\u0117\3\2\2\2\u0128\u011c\3\2\2\2\u0128\u0122"+
		"\3\2\2\2\u0129\27\3\2\2\2\u012a\u012b\7\f\2\2\u012b\u012c\7)\2\2\u012c"+
		"\u012d\7/\2\2\u012d\u012e\5@!\2\u012e\u012f\b\r\1\2\u012f\u0139\3\2\2"+
		"\2\u0130\u0131\7\f\2\2\u0131\u0132\7)\2\2\u0132\u0133\7C\2\2\u0133\u0134"+
		"\5\24\13\2\u0134\u0135\7/\2\2\u0135\u0136\5@!\2\u0136\u0137\b\r\1\2\u0137"+
		"\u0139\3\2\2\2\u0138\u012a\3\2\2\2\u0138\u0130\3\2\2\2\u0139\31\3\2\2"+
		"\2\u013a\u013b\7\20\2\2\u013b\u013c\5@!\2\u013c\u013e\7;\2\2\u013d\u013f"+
		"\5\34\17\2\u013e\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u013e\3\2\2\2"+
		"\u0140\u0141\3\2\2\2\u0141\u0143\3\2\2\2\u0142\u0144\5\36\20\2\u0143\u0142"+
		"\3\2\2\2\u0143\u0144\3\2\2\2\u0144\u0145\3\2\2\2\u0145\u0146\7<\2\2\u0146"+
		"\u0147\b\16\1\2\u0147\33\3\2\2\2\u0148\u0149\7\21\2\2\u0149\u014a\5@!"+
		"\2\u014a\u014b\7C\2\2\u014b\u014d\5\4\3\2\u014c\u014e\7\30\2\2\u014d\u014c"+
		"\3\2\2\2\u014d\u014e\3\2\2\2\u014e\u014f\3\2\2\2\u014f\u0150\b\17\1\2"+
		"\u0150\35\3\2\2\2\u0151\u0152\7\22\2\2\u0152\u0153\7C\2\2\u0153\u0155"+
		"\5\4\3\2\u0154\u0156\7\30\2\2\u0155\u0154\3\2\2\2\u0155\u0156\3\2\2\2"+
		"\u0156\u0157\3\2\2\2\u0157\u0158\b\20\1\2\u0158\37\3\2\2\2\u0159\u015a"+
		"\7\23\2\2\u015a\u015b\5@!\2\u015b\u015c\7;\2\2\u015c\u015d\5\4\3\2\u015d"+
		"\u015e\7<\2\2\u015e\u015f\b\21\1\2\u015f!\3\2\2\2\u0160\u0161\7\24\2\2"+
		"\u0161\u0162\7)\2\2\u0162\u0163\7\25\2\2\u0163\u0164\7(\2\2\u0164\u0165"+
		"\7;\2\2\u0165\u0166\5\4\3\2\u0166\u0167\7<\2\2\u0167\u0168\b\22\1\2\u0168"+
		"\u017e\3\2\2\2\u0169\u016a\7\24\2\2\u016a\u016b\7)\2\2\u016b\u016c\7\25"+
		"\2\2\u016c\u016d\5@!\2\u016d\u016e\7\26\2\2\u016e\u016f\5@!\2\u016f\u0170"+
		"\7;\2\2\u0170\u0171\5\4\3\2\u0171\u0172\7<\2\2\u0172\u0173\b\22\1\2\u0173"+
		"\u017e\3\2\2\2\u0174\u0175\7\24\2\2\u0175\u0176\7)\2\2\u0176\u0177\7\25"+
		"\2\2\u0177\u0178\7)\2\2\u0178\u0179\7;\2\2\u0179\u017a\5\4\3\2\u017a\u017b"+
		"\7<\2\2\u017b\u017c\b\22\1\2\u017c\u017e\3\2\2\2\u017d\u0160\3\2\2\2\u017d"+
		"\u0169\3\2\2\2\u017d\u0174\3\2\2\2\u017e#\3\2\2\2\u017f\u0180\7\27\2\2"+
		"\u0180\u0181\5@!\2\u0181\u0182\7\17\2\2\u0182\u0183\7;\2\2\u0183\u0184"+
		"\5\4\3\2\u0184\u0185\7<\2\2\u0185\u0186\b\23\1\2\u0186%\3\2\2\2\u0187"+
		"\u0188\7\30\2\2\u0188\u0192\b\24\1\2\u0189\u018a\7\31\2\2\u018a\u0192"+
		"\b\24\1\2\u018b\u018c\7\32\2\2\u018c\u0192\b\24\1\2\u018d\u018e\7\32\2"+
		"\2\u018e\u018f\5@!\2\u018f\u0190\b\24\1\2\u0190\u0192\3\2\2\2\u0191\u0187"+
		"\3\2\2\2\u0191\u0189\3\2\2\2\u0191\u018b\3\2\2\2\u0191\u018d\3\2\2\2\u0192"+
		"\'\3\2\2\2\u0193\u0194\7\13\2\2\u0194\u0195\7)\2\2\u0195\u0196\7C\2\2"+
		"\u0196\u0197\7=\2\2\u0197\u0198\5\24\13\2\u0198\u0199\7>\2\2\u0199\u019a"+
		"\5*\26\2\u019a\u019b\b\25\1\2\u019b)\3\2\2\2\u019c\u019d\7/\2\2\u019d"+
		"\u019f\7=\2\2\u019e\u01a0\5,\27\2\u019f\u019e\3\2\2\2\u01a0\u01a1\3\2"+
		"\2\2\u01a1\u019f\3\2\2\2\u01a1\u01a2\3\2\2\2\u01a2\u01a3\3\2\2\2\u01a3"+
		"\u01a4\7>\2\2\u01a4\u01a5\b\26\1\2\u01a5\u01ab\3\2\2\2\u01a6\u01a7\7/"+
		"\2\2\u01a7\u01a8\7=\2\2\u01a8\u01a9\7>\2\2\u01a9\u01ab\b\26\1\2\u01aa"+
		"\u019c\3\2\2\2\u01aa\u01a6\3\2\2\2\u01ab+\3\2\2\2\u01ac\u01ad\5@!\2\u01ad"+
		"\u01ae\7A\2\2\u01ae\u01af\b\27\1\2\u01af\u01b4\3\2\2\2\u01b0\u01b1\5@"+
		"!\2\u01b1\u01b2\b\27\1\2\u01b2\u01b4\3\2\2\2\u01b3\u01ac\3\2\2\2\u01b3"+
		"\u01b0\3\2\2\2\u01b4-\3\2\2\2\u01b5\u01b6\7)\2\2\u01b6\u01b7\7B\2\2\u01b7"+
		"\u01b8\7\37\2\2\u01b8\u01b9\79\2\2\u01b9\u01ba\5@!\2\u01ba\u01bb\7:\2"+
		"\2\u01bb\u01bc\b\30\1\2\u01bc\u01cc\3\2\2\2\u01bd\u01be\7)\2\2\u01be\u01bf"+
		"\7B\2\2\u01bf\u01c0\7 \2\2\u01c0\u01c1\79\2\2\u01c1\u01c2\7:\2\2\u01c2"+
		"\u01cc\b\30\1\2\u01c3\u01c4\7)\2\2\u01c4\u01c5\7B\2\2\u01c5\u01c6\7!\2"+
		"\2\u01c6\u01c7\79\2\2\u01c7\u01c8\5@!\2\u01c8\u01c9\7:\2\2\u01c9\u01ca"+
		"\b\30\1\2\u01ca\u01cc\3\2\2\2\u01cb\u01b5\3\2\2\2\u01cb\u01bd\3\2\2\2"+
		"\u01cb\u01c3\3\2\2\2\u01cc/\3\2\2\2\u01cd\u01ce\7)\2\2\u01ce\u01cf\7B"+
		"\2\2\u01cf\u01d0\7\"\2\2\u01d0\u01dc\b\31\1\2\u01d1\u01d2\7)\2\2\u01d2"+
		"\u01d3\7B\2\2\u01d3\u01d4\7#\2\2\u01d4\u01dc\b\31\1\2\u01d5\u01d6\7)\2"+
		"\2\u01d6\u01d7\7=\2\2\u01d7\u01d8\5@!\2\u01d8\u01d9\7>\2\2\u01d9\u01da"+
		"\b\31\1\2\u01da\u01dc\3\2\2\2\u01db\u01cd\3\2\2\2\u01db\u01d1\3\2\2\2"+
		"\u01db\u01d5\3\2\2\2\u01dc\61\3\2\2\2\u01dd\u01de\7)\2\2\u01de\u01df\7"+
		"=\2\2\u01df\u01e0\5@!\2\u01e0\u01e1\7>\2\2\u01e1\u01e2\7/\2\2\u01e2\u01e3"+
		"\7)\2\2\u01e3\u01e4\7=\2\2\u01e4\u01e5\5@!\2\u01e5\u01e6\7>\2\2\u01e6"+
		"\u01e7\b\32\1\2\u01e7\u01f1\3\2\2\2\u01e8\u01e9\7)\2\2\u01e9\u01ea\7="+
		"\2\2\u01ea\u01eb\5@!\2\u01eb\u01ec\7>\2\2\u01ec\u01ed\7/\2\2\u01ed\u01ee"+
		"\5@!\2\u01ee\u01ef\b\32\1\2\u01ef\u01f1\3\2\2\2\u01f0\u01dd\3\2\2\2\u01f0"+
		"\u01e8\3\2\2\2\u01f1\63\3\2\2\2\u01f2\u01f3\7\36\2\2\u01f3\u01f4\7)\2"+
		"\2\u01f4\u01f6\79\2\2\u01f5\u01f7\5\66\34\2\u01f6\u01f5\3\2\2\2\u01f7"+
		"\u01f8\3\2\2\2\u01f8\u01f6\3\2\2\2\u01f8\u01f9\3\2\2\2\u01f9\u01fa\3\2"+
		"\2\2\u01fa\u01fb\7:\2\2\u01fb\u01fc\7@\2\2\u01fc\u01fd\5\24\13\2\u01fd"+
		"\u01fe\7;\2\2\u01fe\u01ff\5\4\3\2\u01ff\u0200\7<\2\2\u0200\u0201\b\33"+
		"\1\2\u0201\u0225\3\2\2\2\u0202\u0203\7\36\2\2\u0203\u0204\7)\2\2\u0204"+
		"\u0206\79\2\2\u0205\u0207\5\66\34\2\u0206\u0205\3\2\2\2\u0207\u0208\3"+
		"\2\2\2\u0208\u0206\3\2\2\2\u0208\u0209\3\2\2\2\u0209\u020a\3\2\2\2\u020a"+
		"\u020b\7:\2\2\u020b\u020c\7;\2\2\u020c\u020d\5\4\3\2\u020d\u020e\7<\2"+
		"\2\u020e\u020f\b\33\1\2\u020f\u0225\3\2\2\2\u0210\u0211\7\36\2\2\u0211"+
		"\u0212\7)\2\2\u0212\u0213\79\2\2\u0213\u0214\7:\2\2\u0214\u0215\7@\2\2"+
		"\u0215\u0216\5\24\13\2\u0216\u0217\7;\2\2\u0217\u0218\5\4\3\2\u0218\u0219"+
		"\7<\2\2\u0219\u021a\b\33\1\2\u021a\u0225\3\2\2\2\u021b\u021c\7\36\2\2"+
		"\u021c\u021d\7)\2\2\u021d\u021e\79\2\2\u021e\u021f\7:\2\2\u021f\u0220"+
		"\7;\2\2\u0220\u0221\5\4\3\2\u0221\u0222\7<\2\2\u0222\u0223\b\33\1\2\u0223"+
		"\u0225\3\2\2\2\u0224\u01f2\3\2\2\2\u0224\u0202\3\2\2\2\u0224\u0210\3\2"+
		"\2\2\u0224\u021b\3\2\2\2\u0225\65\3\2\2\2\u0226\u0228\t\2\2\2\u0227\u0226"+
		"\3\2\2\2\u0227\u0228\3\2\2\2\u0228\u0229\3\2\2\2\u0229\u022a\7)\2\2\u022a"+
		"\u022c\7C\2\2\u022b\u022d\7$\2\2\u022c\u022b\3\2\2\2\u022c\u022d\3\2\2"+
		"\2\u022d\u022e\3\2\2\2\u022e\u022f\5\24\13\2\u022f\u0230\7A\2\2\u0230"+
		"\u0231\b\34\1\2\u0231\u023e\3\2\2\2\u0232\u0234\t\2\2\2\u0233\u0232\3"+
		"\2\2\2\u0233\u0234\3\2\2\2\u0234\u0235\3\2\2\2\u0235\u0236\7)\2\2\u0236"+
		"\u0238\7C\2\2\u0237\u0239\7$\2\2\u0238\u0237\3\2\2\2\u0238\u0239\3\2\2"+
		"\2\u0239\u023a\3\2\2\2\u023a\u023b\5\24\13\2\u023b\u023c\b\34\1\2\u023c"+
		"\u023e\3\2\2\2\u023d\u0227\3\2\2\2\u023d\u0233\3\2\2\2\u023e\67\3\2\2"+
		"\2\u023f\u0240\7)\2\2\u0240\u0242\79\2\2\u0241\u0243\5<\37\2\u0242\u0241"+
		"\3\2\2\2\u0243\u0244\3\2\2\2\u0244\u0242\3\2\2\2\u0244\u0245\3\2\2\2\u0245"+
		"\u0246\3\2\2\2\u0246\u0247\7:\2\2\u0247\u0248\b\35\1\2\u0248\u024e\3\2"+
		"\2\2\u0249\u024a\7)\2\2\u024a\u024b\79\2\2\u024b\u024c\7:\2\2\u024c\u024e"+
		"\b\35\1\2\u024d\u023f\3\2\2\2\u024d\u0249\3\2\2\2\u024e9\3\2\2\2\u024f"+
		"\u0250\7)\2\2\u0250\u0252\79\2\2\u0251\u0253\5<\37\2\u0252\u0251\3\2\2"+
		"\2\u0253\u0254\3\2\2\2\u0254\u0252\3\2\2\2\u0254\u0255\3\2\2\2\u0255\u0256"+
		"\3\2\2\2\u0256\u0257\7:\2\2\u0257\u0258\b\36\1\2\u0258\u025e\3\2\2\2\u0259"+
		"\u025a\7)\2\2\u025a\u025b\79\2\2\u025b\u025c\7:\2\2\u025c\u025e\b\36\1"+
		"\2\u025d\u024f\3\2\2\2\u025d\u0259\3\2\2\2\u025e;\3\2\2\2\u025f\u0260"+
		"\7)\2\2\u0260\u0262\7C\2\2\u0261\u025f\3\2\2\2\u0261\u0262\3\2\2\2\u0262"+
		"\u0264\3\2\2\2\u0263\u0265\7D\2\2\u0264\u0263\3\2\2\2\u0264\u0265\3\2"+
		"\2\2\u0265\u0266\3\2\2\2\u0266\u0267\5@!\2\u0267\u0268\7A\2\2\u0268\u0269"+
		"\b\37\1\2\u0269\u0275\3\2\2\2\u026a\u026b\7)\2\2\u026b\u026d\7C\2\2\u026c"+
		"\u026a\3\2\2\2\u026c\u026d\3\2\2\2\u026d\u026f\3\2\2\2\u026e\u0270\7D"+
		"\2\2\u026f\u026e\3\2\2\2\u026f\u0270\3\2\2\2\u0270\u0271\3\2\2\2\u0271"+
		"\u0272\5@!\2\u0272\u0273\b\37\1\2\u0273\u0275\3\2\2\2\u0274\u0261\3\2"+
		"\2\2\u0274\u026c\3\2\2\2\u0275=\3\2\2\2\u0276\u0277\7\6\2\2\u0277\u0278"+
		"\79\2\2\u0278\u0279\5@!\2\u0279\u027a\7:\2\2\u027a\u027b\b \1\2\u027b"+
		"\u0289\3\2\2\2\u027c\u027d\7\3\2\2\u027d\u027e\79\2\2\u027e\u027f\5@!"+
		"\2\u027f\u0280\7:\2\2\u0280\u0281\b \1\2\u0281\u0289\3\2\2\2\u0282\u0283"+
		"\7\4\2\2\u0283\u0284\79\2\2\u0284\u0285\5@!\2\u0285\u0286\7:\2\2\u0286"+
		"\u0287\b \1\2\u0287\u0289\3\2\2\2\u0288\u0276\3\2\2\2\u0288\u027c\3\2"+
		"\2\2\u0288\u0282\3\2\2\2\u0289?\3\2\2\2\u028a\u028b\b!\1\2\u028b\u028c"+
		"\78\2\2\u028c\u028d\5@!\26\u028d\u028e\b!\1\2\u028e\u02b0\3\2\2\2\u028f"+
		"\u0290\7,\2\2\u0290\u0291\5@!\25\u0291\u0292\b!\1\2\u0292\u02b0\3\2\2"+
		"\2\u0293\u0294\79\2\2\u0294\u0295\5@!\2\u0295\u0296\7:\2\2\u0296\u0297"+
		"\b!\1\2\u0297\u02b0\3\2\2\2\u0298\u0299\7&\2\2\u0299\u02b0\b!\1\2\u029a"+
		"\u029b\7\'\2\2\u029b\u02b0\b!\1\2\u029c\u029d\7(\2\2\u029d\u02b0\b!\1"+
		"\2\u029e\u029f\7\b\2\2\u029f\u02b0\b!\1\2\u02a0\u02a1\58\35\2\u02a1\u02a2"+
		"\b!\1\2\u02a2\u02b0\3\2\2\2\u02a3\u02a4\5> \2\u02a4\u02a5\b!\1\2\u02a5"+
		"\u02b0\3\2\2\2\u02a6\u02a7\7\t\2\2\u02a7\u02b0\b!\1\2\u02a8\u02a9\7)\2"+
		"\2\u02a9\u02b0\b!\1\2\u02aa\u02ab\5\60\31\2\u02ab\u02ac\b!\1\2\u02ac\u02b0"+
		"\3\2\2\2\u02ad\u02ae\7\n\2\2\u02ae\u02b0\b!\1\2\u02af\u028a\3\2\2\2\u02af"+
		"\u028f\3\2\2\2\u02af\u0293\3\2\2\2\u02af\u0298\3\2\2\2\u02af\u029a\3\2"+
		"\2\2\u02af\u029c\3\2\2\2\u02af\u029e\3\2\2\2\u02af\u02a0\3\2\2\2\u02af"+
		"\u02a3\3\2\2\2\u02af\u02a6\3\2\2\2\u02af\u02a8\3\2\2\2\u02af\u02aa\3\2"+
		"\2\2\u02af\u02ad\3\2\2\2\u02b0\u02d6\3\2\2\2\u02b1\u02b2\f\24\2\2\u02b2"+
		"\u02b3\t\3\2\2\u02b3\u02b4\5@!\25\u02b4\u02b5\b!\1\2\u02b5\u02d5\3\2\2"+
		"\2\u02b6\u02b7\f\23\2\2\u02b7\u02b8\t\4\2\2\u02b8\u02b9\5@!\24\u02b9\u02ba"+
		"\b!\1\2\u02ba\u02d5\3\2\2\2\u02bb\u02bc\f\22\2\2\u02bc\u02bd\t\5\2\2\u02bd"+
		"\u02be\5@!\23\u02be\u02bf\b!\1\2\u02bf\u02d5\3\2\2\2\u02c0\u02c1\f\21"+
		"\2\2\u02c1\u02c2\t\6\2\2\u02c2\u02c3\5@!\22\u02c3\u02c4\b!\1\2\u02c4\u02d5"+
		"\3\2\2\2\u02c5\u02c6\f\20\2\2\u02c6\u02c7\t\7\2\2\u02c7\u02c8\5@!\21\u02c8"+
		"\u02c9\b!\1\2\u02c9\u02d5\3\2\2\2\u02ca\u02cb\f\17\2\2\u02cb\u02cc\7."+
		"\2\2\u02cc\u02cd\5@!\20\u02cd\u02ce\b!\1\2\u02ce\u02d5\3\2\2\2\u02cf\u02d0"+
		"\f\16\2\2\u02d0\u02d1\7-\2\2\u02d1\u02d2\5@!\17\u02d2\u02d3\b!\1\2\u02d3"+
		"\u02d5\3\2\2\2\u02d4\u02b1\3\2\2\2\u02d4\u02b6\3\2\2\2\u02d4\u02bb\3\2"+
		"\2\2\u02d4\u02c0\3\2\2\2\u02d4\u02c5\3\2\2\2\u02d4\u02ca\3\2\2\2\u02d4"+
		"\u02cf\3\2\2\2\u02d5\u02d8\3\2\2\2\u02d6\u02d4\3\2\2\2\u02d6\u02d7\3\2"+
		"\2\2\u02d7A\3\2\2\2\u02d8\u02d6\3\2\2\2=IOU[agmsy\177\u0085\u008b\u0091"+
		"\u0097\u009d\u00a3\u00a7\u00ae\u00ba\u00cb\u00d8\u00e4\u0109\u0115\u0128"+
		"\u0138\u0140\u0143\u014d\u0155\u017d\u0191\u01a1\u01aa\u01b3\u01cb\u01db"+
		"\u01f0\u01f8\u0208\u0224\u0227\u022c\u0233\u0238\u023d\u0244\u024d\u0254"+
		"\u025d\u0261\u0264\u026c\u026f\u0274\u0288\u02af\u02d4\u02d6";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}