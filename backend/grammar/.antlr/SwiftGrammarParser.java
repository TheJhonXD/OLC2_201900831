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
		EMPTY=32, COUNT=33, NUMBER=34, CHAR=35, STRING=36, ID=37, DIF=38, IG_IG=39, 
		NOT=40, OR=41, AND=42, IG=43, MAY_IG=44, MEN_IG=45, MAYOR=46, MENOR=47, 
		MOD=48, MUL=49, DIV=50, ADD=51, SUB=52, PARIZQ=53, PARDER=54, LLAVEIZQ=55, 
		LLAVEDER=56, CORCHIZQ=57, CORCHDER=58, Q_MARK=59, ARROW=60, COMA=61, PUNTO=62, 
		COLON=63, WHITESPACE=64, COMMENT=65, LINE_COMMENT=66;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_ifstmt = 4, RULE_elseifstmt = 5, RULE_elsestmt = 6, RULE_varstmt = 7, 
		RULE_tipo = 8, RULE_varasgmt = 9, RULE_conststmt = 10, RULE_switchstmt = 11, 
		RULE_casestmt = 12, RULE_defaultstmt = 13, RULE_whilestmt = 14, RULE_forstmt = 15, 
		RULE_guardstmt = 16, RULE_transferstmt = 17, RULE_vectorstmt = 18, RULE_definestmt = 19, 
		RULE_listexpr = 20, RULE_methodvec = 21, RULE_methodvecrtrn = 22, RULE_vecaccess = 23, 
		RULE_expr = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "ifstmt", "elseifstmt", "elsestmt", 
			"varstmt", "tipo", "varasgmt", "conststmt", "switchstmt", "casestmt", 
			"defaultstmt", "whilestmt", "forstmt", "guardstmt", "transferstmt", "vectorstmt", 
			"definestmt", "listexpr", "methodvec", "methodvecrtrn", "vecaccess", 
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
			null, null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", 
			"'>='", "'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'", "'?'", "'->'", "','", "'.'", "':'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL", 
			"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "IN", "RANGEPTS", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", 
			"SELF", "MUTATING", "FUNC", "APPEND", "REMOVELAST", "REMOVE", "EMPTY", 
			"COUNT", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", "MUL", "DIV", 
			"ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "Q_MARK", "ARROW", "COMA", "PUNTO", "COLON", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
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
			setState(50);
			((SContext)_localctx).block = block();
			setState(51);
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
			setState(55); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(54);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(57); 
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
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
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
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(100);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(61);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinstr 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(67);
				((InstructionContext)_localctx).varstmt = varstmt();
				 _localctx.inst = ((InstructionContext)_localctx).varstmt.var
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(70);
				((InstructionContext)_localctx).varasgmt = varasgmt();
				 _localctx.inst = ((InstructionContext)_localctx).varasgmt.asgmt
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(73);
				((InstructionContext)_localctx).conststmt = conststmt();
				 _localctx.inst = ((InstructionContext)_localctx).conststmt.const
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(76);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchstmt.switchinstr
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(79);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whileinstr
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(82);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.forinstr
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(85);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.guardinstr
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(88);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				 _localctx.inst = ((InstructionContext)_localctx).transferstmt.trns
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(91);
				((InstructionContext)_localctx).vectorstmt = vectorstmt();
				 _localctx.inst = ((InstructionContext)_localctx).vectorstmt.vectorinstr
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(94);
				((InstructionContext)_localctx).methodvec = methodvec();
				 _localctx.inst = ((InstructionContext)_localctx).methodvec.methodinstr
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(97);
				((InstructionContext)_localctx).vecaccess = vecaccess();
				 _localctx.inst = ((InstructionContext)_localctx).vecaccess.vecacc
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
		public ExprContext expr;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(103);
			match(PARIZQ);
			setState(104);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(105);
			match(PARDER);
			 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
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
		enterRule(_localctx, 8, RULE_ifstmt);
		int _la;
		try {
			int _alt;
			setState(140);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(108);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(109);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(110);
				match(LLAVEIZQ);
				setState(111);
				((IfstmtContext)_localctx).block = block();
				setState(112);
				match(LLAVEDER);
				 _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(115);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(116);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(117);
				match(LLAVEIZQ);
				setState(118);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(119);
				match(LLAVEDER);
				setState(121); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(120);
						((IfstmtContext)_localctx).elseifstmt = elseifstmt();
						((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(123); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(125);
				match(ELSE);
				setState(127);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LLAVEIZQ) {
					{
					setState(126);
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
				setState(131);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(132);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(133);
				match(LLAVEIZQ);
				setState(134);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(135);
				match(LLAVEDER);
				setState(136);
				match(ELSE);
				setState(137);
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
		enterRule(_localctx, 10, RULE_elseifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			((ElseifstmtContext)_localctx).ELSE = match(ELSE);
			setState(143);
			match(IF);
			setState(144);
			((ElseifstmtContext)_localctx).expr = expr(0);
			setState(145);
			match(LLAVEIZQ);
			setState(146);
			((ElseifstmtContext)_localctx).block = block();
			setState(147);
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
		enterRule(_localctx, 12, RULE_elsestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			match(LLAVEIZQ);
			setState(151);
			((ElsestmtContext)_localctx).block = block();
			setState(152);
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
		enterRule(_localctx, 14, RULE_varstmt);
		try {
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(155);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(156);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(157);
				match(COLON);
				setState(158);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(159);
				match(IG);
				setState(160);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), ((VarstmtContext)_localctx).tipo.rtipo, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(163);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(164);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(165);
				match(IG);
				setState(166);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), environment.NULL, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(169);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(170);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(171);
				match(COLON);
				setState(172);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(173);
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
		enterRule(_localctx, 16, RULE_tipo);
		try {
			setState(188);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(178);
				match(INT);
				_localctx.rtipo = 0
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(180);
				match(FLOAT);
				_localctx.rtipo = 1
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(182);
				match(STR);
				_localctx.rtipo = 2
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(184);
				match(BOOL);
				_localctx.rtipo = 3
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(186);
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
		enterRule(_localctx, 18, RULE_varasgmt);
		try {
			setState(207);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(191);
				match(IG);
				setState(192);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(195);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(196);
				match(ADD);
				setState(197);
				match(IG);
				setState(198);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewOpAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e, "+")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(201);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(202);
				match(SUB);
				setState(203);
				match(IG);
				setState(204);
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
		enterRule(_localctx, 20, RULE_conststmt);
		try {
			setState(223);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(209);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(210);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(211);
				match(IG);
				setState(212);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), environment.NULL, ((ConststmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(215);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(216);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(217);
				match(COLON);
				setState(218);
				((ConststmtContext)_localctx).tipo = tipo();
				setState(219);
				match(IG);
				setState(220);
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
		enterRule(_localctx, 22, RULE_switchstmt);

		    var switchInterfaces = []interface{}{}
		    var interfacelist []ICasestmtContext

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
			setState(226);
			((SwitchstmtContext)_localctx).expr = expr(0);
			setState(227);
			match(LLAVEIZQ);
			setState(229); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(228);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				((SwitchstmtContext)_localctx).casesvar.add(((SwitchstmtContext)_localctx).casestmt);
				}
				}
				setState(231); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(234);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(233);
				defaultstmt();
				}
			}

			setState(236);
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
		enterRule(_localctx, 24, RULE_casestmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			((CasestmtContext)_localctx).CASE = match(CASE);
			setState(240);
			((CasestmtContext)_localctx).expr = expr(0);
			setState(241);
			match(COLON);
			setState(242);
			((CasestmtContext)_localctx).block = block();
			setState(244);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(243);
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
		enterRule(_localctx, 26, RULE_defaultstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(248);
			match(DEFAULT);
			setState(249);
			match(COLON);
			setState(250);
			((DefaultstmtContext)_localctx).block = block();
			setState(252);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(251);
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
		enterRule(_localctx, 28, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(257);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(258);
			match(LLAVEIZQ);
			setState(259);
			((WhilestmtContext)_localctx).block = block();
			setState(260);
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
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
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
		enterRule(_localctx, 30, RULE_forstmt);

		    var cadena interfaces.Expression
		    var str string

		try {
			setState(283);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(263);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(264);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(265);
				match(IN);
				setState(266);
				((ForstmtContext)_localctx).STRING = match(STRING);
				setState(267);
				match(LLAVEIZQ);
				setState(268);
				((ForstmtContext)_localctx).block = block();
				setState(269);
				match(LLAVEDER);
				 
				    str = (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getText():null)
				    cadena = expressions.NewPrimitive((((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getLine():0), (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1], environment.STRING)
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), cadena, nil, nil, ((ForstmtContext)_localctx).block.blk) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(272);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(273);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(274);
				match(IN);
				setState(275);
				((ForstmtContext)_localctx).left = expr(0);
				setState(276);
				match(RANGEPTS);
				setState(277);
				((ForstmtContext)_localctx).right = expr(0);
				setState(278);
				match(LLAVEIZQ);
				setState(279);
				((ForstmtContext)_localctx).block = block();
				setState(280);
				match(LLAVEDER);
				 
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), nil, ((ForstmtContext)_localctx).left.e, ((ForstmtContext)_localctx).right.e, ((ForstmtContext)_localctx).block.blk)

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
		enterRule(_localctx, 32, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(286);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(287);
			match(ELSE);
			setState(288);
			match(LLAVEIZQ);
			setState(289);
			((GuardstmtContext)_localctx).block = block();
			setState(290);
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
		enterRule(_localctx, 34, RULE_transferstmt);
		try {
			setState(303);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(293);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				 _localctx.trns = instructions.NewBreak((((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getLine():0), (((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getCharPositionInLine():0)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(295);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				 _localctx.trns = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(297);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				 _localctx.trns = instructions.NewReturn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(299);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(300);
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
		enterRule(_localctx, 36, RULE_vectorstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(305);
			((VectorstmtContext)_localctx).VAR = match(VAR);
			setState(306);
			((VectorstmtContext)_localctx).ID = match(ID);
			setState(307);
			match(COLON);
			setState(308);
			match(CORCHIZQ);
			setState(309);
			((VectorstmtContext)_localctx).tipo = tipo();
			setState(310);
			match(CORCHDER);
			setState(311);
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
		enterRule(_localctx, 38, RULE_definestmt);

		    var defVecInterfaces []interface{}

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
			match(IG);
			setState(315);
			match(CORCHIZQ);
			setState(317); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(316);
				((DefinestmtContext)_localctx).listexpr = listexpr();
				((DefinestmtContext)_localctx).lista.add(((DefinestmtContext)_localctx).listexpr);
				}
				}
				setState(319); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TRU) | (1L << FAL) | (1L << NUMBER) | (1L << CHAR) | (1L << STRING) | (1L << ID) | (1L << SUB) | (1L << PARIZQ))) != 0) );
			setState(321);
			match(CORCHDER);
			 
			    for _, e := range localctx.(*DefinestmtContext).GetLista() {
			        // fmt.Println(fmt.Sprintf("%T", e.GetListe()))
			        defVecInterfaces = append(defVecInterfaces, e.GetListe())
			    }
			    _localctx.defineinstr = defVecInterfaces

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
		enterRule(_localctx, 40, RULE_listexpr);
		try {
			setState(331);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(324);
				((ListexprContext)_localctx).expr = expr(0);
				setState(325);
				match(COMA);
				 _localctx.liste = ((ListexprContext)_localctx).expr.e 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(328);
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
		enterRule(_localctx, 42, RULE_methodvec);
		try {
			setState(355);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(333);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(334);
				match(PUNTO);
				setState(335);
				((MethodvecContext)_localctx).APPEND = match(APPEND);
				setState(336);
				match(PARIZQ);
				setState(337);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(338);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), ((MethodvecContext)_localctx).expr.e, (((MethodvecContext)_localctx).APPEND!=null?((MethodvecContext)_localctx).APPEND.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(341);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(342);
				match(PUNTO);
				setState(343);
				((MethodvecContext)_localctx).REMOVELAST = match(REMOVELAST);
				setState(344);
				match(PARIZQ);
				setState(345);
				match(PARDER);
				 _localctx.methodinstr = instructions.NewVectorMethod((((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecContext)_localctx).ID!=null?((MethodvecContext)_localctx).ID.getText():null), nil, (((MethodvecContext)_localctx).REMOVELAST!=null?((MethodvecContext)_localctx).REMOVELAST.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(347);
				((MethodvecContext)_localctx).ID = match(ID);
				setState(348);
				match(PUNTO);
				setState(349);
				((MethodvecContext)_localctx).REMOVE = match(REMOVE);
				setState(350);
				match(PARIZQ);
				setState(351);
				((MethodvecContext)_localctx).expr = expr(0);
				setState(352);
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
		enterRule(_localctx, 44, RULE_methodvecrtrn);
		try {
			setState(371);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(357);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(358);
				match(PUNTO);
				setState(359);
				((MethodvecrtrnContext)_localctx).EMPTY = match(EMPTY);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).EMPTY!=null?((MethodvecrtrnContext)_localctx).EMPTY.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(361);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(362);
				match(PUNTO);
				setState(363);
				((MethodvecrtrnContext)_localctx).COUNT = match(COUNT);
				 _localctx.methodinstrtrn = expressions.NewVector((((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getCharPositionInLine():0), (((MethodvecrtrnContext)_localctx).ID!=null?((MethodvecrtrnContext)_localctx).ID.getText():null), nil, (((MethodvecrtrnContext)_localctx).COUNT!=null?((MethodvecrtrnContext)_localctx).COUNT.getText():null)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(365);
				((MethodvecrtrnContext)_localctx).ID = match(ID);
				setState(366);
				match(CORCHIZQ);
				setState(367);
				((MethodvecrtrnContext)_localctx).expr = expr(0);
				setState(368);
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
		enterRule(_localctx, 46, RULE_vecaccess);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(373);
			((VecaccessContext)_localctx).firstId = match(ID);
			setState(374);
			match(CORCHIZQ);
			setState(375);
			((VecaccessContext)_localctx).first = expr(0);
			setState(376);
			match(CORCHDER);
			setState(377);
			match(IG);
			setState(378);
			((VecaccessContext)_localctx).secondId = match(ID);
			setState(379);
			match(CORCHIZQ);
			setState(380);
			((VecaccessContext)_localctx).second = expr(0);
			setState(381);
			match(CORCHDER);
			 
			    _localctx.vecacc = instructions.NewVectorAsgmt((((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getCharPositionInLine():0), (((VecaccessContext)_localctx).firstId!=null?((VecaccessContext)_localctx).firstId.getText():null), ((VecaccessContext)_localctx).first.e, ((VecaccessContext)_localctx).second.e, (((VecaccessContext)_localctx).secondId!=null?((VecaccessContext)_localctx).secondId.getText():null)) 

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
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(409);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				{
				setState(385);
				((ExprContext)_localctx).op = match(SUB);
				setState(386);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, "unario", expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), -1, environment.INTEGER)) 
				}
				break;
			case 2:
				{
				setState(389);
				match(PARIZQ);
				setState(390);
				((ExprContext)_localctx).expr = expr(0);
				setState(391);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 3:
				{
				setState(394);
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
				setState(396);
				((ExprContext)_localctx).CHAR = match(CHAR);
				 
				        str := (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getLine():0), (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getCharPositionInLine():0), str[1:len(str)-1], environment.CHAR) 
				    
				}
				break;
			case 5:
				{
				setState(398);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(400);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(402);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(404);
				((ExprContext)_localctx).ID = match(ID);
				 _localctx.e = expressions.NewVar((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 9:
				{
				setState(406);
				((ExprContext)_localctx).methodvecrtrn = methodvecrtrn();
				 _localctx.e = ((ExprContext)_localctx).methodvecrtrn.methodinstrtrn 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(448);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(446);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(411);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(412);
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
						setState(413);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(416);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(417);
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
						setState(418);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(421);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(422);
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
						setState(423);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(426);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(427);
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
						setState(428);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(431);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(432);
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
						setState(433);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(436);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(437);
						match(AND);
						setState(438);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(441);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(442);
						match(OR);
						setState(443);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(10);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(450);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
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
		case 24:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3D\u01c6\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\3\2\3\2\3\2\3\2\3\3\6\3:\n\3\r\3\16\3;\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\5\4g\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\6\6|\n\6\r\6\16\6}\3\6\3\6\5\6\u0082\n\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008f\n\6\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00b3\n\t\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00bf\n\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u00d2"+
		"\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00e2"+
		"\n\f\3\r\3\r\3\r\3\r\6\r\u00e8\n\r\r\r\16\r\u00e9\3\r\5\r\u00ed\n\r\3"+
		"\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\5\16\u00f7\n\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\17\5\17\u00ff\n\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u011e\n\21\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23"+
		"\u0132\n\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25"+
		"\6\25\u0140\n\25\r\25\16\25\u0141\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3"+
		"\26\3\26\3\26\5\26\u014e\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\5\27\u0166\n\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\5\30\u0176\n\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\5\32\u019c\n\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\7\32\u01c1\n\32\f\32"+
		"\16\32\u01c4\13\32\3\32\2\3\62\33\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*,.\60\62\2\7\3\2\62\64\3\2\65\66\4\2..\60\60\4\2//\61\61\3\2"+
		"()\2\u01e3\2\64\3\2\2\2\49\3\2\2\2\6f\3\2\2\2\bh\3\2\2\2\n\u008e\3\2\2"+
		"\2\f\u0090\3\2\2\2\16\u0098\3\2\2\2\20\u00b2\3\2\2\2\22\u00be\3\2\2\2"+
		"\24\u00d1\3\2\2\2\26\u00e1\3\2\2\2\30\u00e3\3\2\2\2\32\u00f1\3\2\2\2\34"+
		"\u00fa\3\2\2\2\36\u0102\3\2\2\2 \u011d\3\2\2\2\"\u011f\3\2\2\2$\u0131"+
		"\3\2\2\2&\u0133\3\2\2\2(\u013c\3\2\2\2*\u014d\3\2\2\2,\u0165\3\2\2\2."+
		"\u0175\3\2\2\2\60\u0177\3\2\2\2\62\u019b\3\2\2\2\64\65\5\4\3\2\65\66\7"+
		"\2\2\3\66\67\b\2\1\2\67\3\3\2\2\28:\5\6\4\298\3\2\2\2:;\3\2\2\2;9\3\2"+
		"\2\2;<\3\2\2\2<=\3\2\2\2=>\b\3\1\2>\5\3\2\2\2?@\5\b\5\2@A\b\4\1\2Ag\3"+
		"\2\2\2BC\5\n\6\2CD\b\4\1\2Dg\3\2\2\2EF\5\20\t\2FG\b\4\1\2Gg\3\2\2\2HI"+
		"\5\24\13\2IJ\b\4\1\2Jg\3\2\2\2KL\5\26\f\2LM\b\4\1\2Mg\3\2\2\2NO\5\30\r"+
		"\2OP\b\4\1\2Pg\3\2\2\2QR\5\36\20\2RS\b\4\1\2Sg\3\2\2\2TU\5 \21\2UV\b\4"+
		"\1\2Vg\3\2\2\2WX\5\"\22\2XY\b\4\1\2Yg\3\2\2\2Z[\5$\23\2[\\\b\4\1\2\\g"+
		"\3\2\2\2]^\5&\24\2^_\b\4\1\2_g\3\2\2\2`a\5,\27\2ab\b\4\1\2bg\3\2\2\2c"+
		"d\5\60\31\2de\b\4\1\2eg\3\2\2\2f?\3\2\2\2fB\3\2\2\2fE\3\2\2\2fH\3\2\2"+
		"\2fK\3\2\2\2fN\3\2\2\2fQ\3\2\2\2fT\3\2\2\2fW\3\2\2\2fZ\3\2\2\2f]\3\2\2"+
		"\2f`\3\2\2\2fc\3\2\2\2g\7\3\2\2\2hi\7\r\2\2ij\7\67\2\2jk\5\62\32\2kl\7"+
		"8\2\2lm\b\5\1\2m\t\3\2\2\2no\7\16\2\2op\5\62\32\2pq\79\2\2qr\5\4\3\2r"+
		"s\7:\2\2st\b\6\1\2t\u008f\3\2\2\2uv\7\16\2\2vw\5\62\32\2wx\79\2\2xy\5"+
		"\4\3\2y{\7:\2\2z|\5\f\7\2{z\3\2\2\2|}\3\2\2\2}{\3\2\2\2}~\3\2\2\2~\177"+
		"\3\2\2\2\177\u0081\7\17\2\2\u0080\u0082\5\16\b\2\u0081\u0080\3\2\2\2\u0081"+
		"\u0082\3\2\2\2\u0082\u0083\3\2\2\2\u0083\u0084\b\6\1\2\u0084\u008f\3\2"+
		"\2\2\u0085\u0086\7\16\2\2\u0086\u0087\5\62\32\2\u0087\u0088\79\2\2\u0088"+
		"\u0089\5\4\3\2\u0089\u008a\7:\2\2\u008a\u008b\7\17\2\2\u008b\u008c\5\16"+
		"\b\2\u008c\u008d\b\6\1\2\u008d\u008f\3\2\2\2\u008en\3\2\2\2\u008eu\3\2"+
		"\2\2\u008e\u0085\3\2\2\2\u008f\13\3\2\2\2\u0090\u0091\7\17\2\2\u0091\u0092"+
		"\7\16\2\2\u0092\u0093\5\62\32\2\u0093\u0094\79\2\2\u0094\u0095\5\4\3\2"+
		"\u0095\u0096\7:\2\2\u0096\u0097\b\7\1\2\u0097\r\3\2\2\2\u0098\u0099\7"+
		"9\2\2\u0099\u009a\5\4\3\2\u009a\u009b\7:\2\2\u009b\u009c\b\b\1\2\u009c"+
		"\17\3\2\2\2\u009d\u009e\7\13\2\2\u009e\u009f\7\'\2\2\u009f\u00a0\7A\2"+
		"\2\u00a0\u00a1\5\22\n\2\u00a1\u00a2\7-\2\2\u00a2\u00a3\5\62\32\2\u00a3"+
		"\u00a4\b\t\1\2\u00a4\u00b3\3\2\2\2\u00a5\u00a6\7\13\2\2\u00a6\u00a7\7"+
		"\'\2\2\u00a7\u00a8\7-\2\2\u00a8\u00a9\5\62\32\2\u00a9\u00aa\b\t\1\2\u00aa"+
		"\u00b3\3\2\2\2\u00ab\u00ac\7\13\2\2\u00ac\u00ad\7\'\2\2\u00ad\u00ae\7"+
		"A\2\2\u00ae\u00af\5\22\n\2\u00af\u00b0\7=\2\2\u00b0\u00b1\b\t\1\2\u00b1"+
		"\u00b3\3\2\2\2\u00b2\u009d\3\2\2\2\u00b2\u00a5\3\2\2\2\u00b2\u00ab\3\2"+
		"\2\2\u00b3\21\3\2\2\2\u00b4\u00b5\7\3\2\2\u00b5\u00bf\b\n\1\2\u00b6\u00b7"+
		"\7\4\2\2\u00b7\u00bf\b\n\1\2\u00b8\u00b9\7\6\2\2\u00b9\u00bf\b\n\1\2\u00ba"+
		"\u00bb\7\5\2\2\u00bb\u00bf\b\n\1\2\u00bc\u00bd\7\7\2\2\u00bd\u00bf\b\n"+
		"\1\2\u00be\u00b4\3\2\2\2\u00be\u00b6\3\2\2\2\u00be\u00b8\3\2\2\2\u00be"+
		"\u00ba\3\2\2\2\u00be\u00bc\3\2\2\2\u00bf\23\3\2\2\2\u00c0\u00c1\7\'\2"+
		"\2\u00c1\u00c2\7-\2\2\u00c2\u00c3\5\62\32\2\u00c3\u00c4\b\13\1\2\u00c4"+
		"\u00d2\3\2\2\2\u00c5\u00c6\7\'\2\2\u00c6\u00c7\7\65\2\2\u00c7\u00c8\7"+
		"-\2\2\u00c8\u00c9\5\62\32\2\u00c9\u00ca\b\13\1\2\u00ca\u00d2\3\2\2\2\u00cb"+
		"\u00cc\7\'\2\2\u00cc\u00cd\7\66\2\2\u00cd\u00ce\7-\2\2\u00ce\u00cf\5\62"+
		"\32\2\u00cf\u00d0\b\13\1\2\u00d0\u00d2\3\2\2\2\u00d1\u00c0\3\2\2\2\u00d1"+
		"\u00c5\3\2\2\2\u00d1\u00cb\3\2\2\2\u00d2\25\3\2\2\2\u00d3\u00d4\7\f\2"+
		"\2\u00d4\u00d5\7\'\2\2\u00d5\u00d6\7-\2\2\u00d6\u00d7\5\62\32\2\u00d7"+
		"\u00d8\b\f\1\2\u00d8\u00e2\3\2\2\2\u00d9\u00da\7\f\2\2\u00da\u00db\7\'"+
		"\2\2\u00db\u00dc\7A\2\2\u00dc\u00dd\5\22\n\2\u00dd\u00de\7-\2\2\u00de"+
		"\u00df\5\62\32\2\u00df\u00e0\b\f\1\2\u00e0\u00e2\3\2\2\2\u00e1\u00d3\3"+
		"\2\2\2\u00e1\u00d9\3\2\2\2\u00e2\27\3\2\2\2\u00e3\u00e4\7\20\2\2\u00e4"+
		"\u00e5\5\62\32\2\u00e5\u00e7\79\2\2\u00e6\u00e8\5\32\16\2\u00e7\u00e6"+
		"\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea"+
		"\u00ec\3\2\2\2\u00eb\u00ed\5\34\17\2\u00ec\u00eb\3\2\2\2\u00ec\u00ed\3"+
		"\2\2\2\u00ed\u00ee\3\2\2\2\u00ee\u00ef\7:\2\2\u00ef\u00f0\b\r\1\2\u00f0"+
		"\31\3\2\2\2\u00f1\u00f2\7\21\2\2\u00f2\u00f3\5\62\32\2\u00f3\u00f4\7A"+
		"\2\2\u00f4\u00f6\5\4\3\2\u00f5\u00f7\7\30\2\2\u00f6\u00f5\3\2\2\2\u00f6"+
		"\u00f7\3\2\2\2\u00f7\u00f8\3\2\2\2\u00f8\u00f9\b\16\1\2\u00f9\33\3\2\2"+
		"\2\u00fa\u00fb\7\22\2\2\u00fb\u00fc\7A\2\2\u00fc\u00fe\5\4\3\2\u00fd\u00ff"+
		"\7\30\2\2\u00fe\u00fd\3\2\2\2\u00fe\u00ff\3\2\2\2\u00ff\u0100\3\2\2\2"+
		"\u0100\u0101\b\17\1\2\u0101\35\3\2\2\2\u0102\u0103\7\23\2\2\u0103\u0104"+
		"\5\62\32\2\u0104\u0105\79\2\2\u0105\u0106\5\4\3\2\u0106\u0107\7:\2\2\u0107"+
		"\u0108\b\20\1\2\u0108\37\3\2\2\2\u0109\u010a\7\24\2\2\u010a\u010b\7\'"+
		"\2\2\u010b\u010c\7\25\2\2\u010c\u010d\7&\2\2\u010d\u010e\79\2\2\u010e"+
		"\u010f\5\4\3\2\u010f\u0110\7:\2\2\u0110\u0111\b\21\1\2\u0111\u011e\3\2"+
		"\2\2\u0112\u0113\7\24\2\2\u0113\u0114\7\'\2\2\u0114\u0115\7\25\2\2\u0115"+
		"\u0116\5\62\32\2\u0116\u0117\7\26\2\2\u0117\u0118\5\62\32\2\u0118\u0119"+
		"\79\2\2\u0119\u011a\5\4\3\2\u011a\u011b\7:\2\2\u011b\u011c\b\21\1\2\u011c"+
		"\u011e\3\2\2\2\u011d\u0109\3\2\2\2\u011d\u0112\3\2\2\2\u011e!\3\2\2\2"+
		"\u011f\u0120\7\27\2\2\u0120\u0121\5\62\32\2\u0121\u0122\7\17\2\2\u0122"+
		"\u0123\79\2\2\u0123\u0124\5\4\3\2\u0124\u0125\7:\2\2\u0125\u0126\b\22"+
		"\1\2\u0126#\3\2\2\2\u0127\u0128\7\30\2\2\u0128\u0132\b\23\1\2\u0129\u012a"+
		"\7\31\2\2\u012a\u0132\b\23\1\2\u012b\u012c\7\32\2\2\u012c\u0132\b\23\1"+
		"\2\u012d\u012e\7\32\2\2\u012e\u012f\5\62\32\2\u012f\u0130\b\23\1\2\u0130"+
		"\u0132\3\2\2\2\u0131\u0127\3\2\2\2\u0131\u0129\3\2\2\2\u0131\u012b\3\2"+
		"\2\2\u0131\u012d\3\2\2\2\u0132%\3\2\2\2\u0133\u0134\7\13\2\2\u0134\u0135"+
		"\7\'\2\2\u0135\u0136\7A\2\2\u0136\u0137\7;\2\2\u0137\u0138\5\22\n\2\u0138"+
		"\u0139\7<\2\2\u0139\u013a\5(\25\2\u013a\u013b\b\24\1\2\u013b\'\3\2\2\2"+
		"\u013c\u013d\7-\2\2\u013d\u013f\7;\2\2\u013e\u0140\5*\26\2\u013f\u013e"+
		"\3\2\2\2\u0140\u0141\3\2\2\2\u0141\u013f\3\2\2\2\u0141\u0142\3\2\2\2\u0142"+
		"\u0143\3\2\2\2\u0143\u0144\7<\2\2\u0144\u0145\b\25\1\2\u0145)\3\2\2\2"+
		"\u0146\u0147\5\62\32\2\u0147\u0148\7?\2\2\u0148\u0149\b\26\1\2\u0149\u014e"+
		"\3\2\2\2\u014a\u014b\5\62\32\2\u014b\u014c\b\26\1\2\u014c\u014e\3\2\2"+
		"\2\u014d\u0146\3\2\2\2\u014d\u014a\3\2\2\2\u014e+\3\2\2\2\u014f\u0150"+
		"\7\'\2\2\u0150\u0151\7@\2\2\u0151\u0152\7\37\2\2\u0152\u0153\7\67\2\2"+
		"\u0153\u0154\5\62\32\2\u0154\u0155\78\2\2\u0155\u0156\b\27\1\2\u0156\u0166"+
		"\3\2\2\2\u0157\u0158\7\'\2\2\u0158\u0159\7@\2\2\u0159\u015a\7 \2\2\u015a"+
		"\u015b\7\67\2\2\u015b\u015c\78\2\2\u015c\u0166\b\27\1\2\u015d\u015e\7"+
		"\'\2\2\u015e\u015f\7@\2\2\u015f\u0160\7!\2\2\u0160\u0161\7\67\2\2\u0161"+
		"\u0162\5\62\32\2\u0162\u0163\78\2\2\u0163\u0164\b\27\1\2\u0164\u0166\3"+
		"\2\2\2\u0165\u014f\3\2\2\2\u0165\u0157\3\2\2\2\u0165\u015d\3\2\2\2\u0166"+
		"-\3\2\2\2\u0167\u0168\7\'\2\2\u0168\u0169\7@\2\2\u0169\u016a\7\"\2\2\u016a"+
		"\u0176\b\30\1\2\u016b\u016c\7\'\2\2\u016c\u016d\7@\2\2\u016d\u016e\7#"+
		"\2\2\u016e\u0176\b\30\1\2\u016f\u0170\7\'\2\2\u0170\u0171\7;\2\2\u0171"+
		"\u0172\5\62\32\2\u0172\u0173\7<\2\2\u0173\u0174\b\30\1\2\u0174\u0176\3"+
		"\2\2\2\u0175\u0167\3\2\2\2\u0175\u016b\3\2\2\2\u0175\u016f\3\2\2\2\u0176"+
		"/\3\2\2\2\u0177\u0178\7\'\2\2\u0178\u0179\7;\2\2\u0179\u017a\5\62\32\2"+
		"\u017a\u017b\7<\2\2\u017b\u017c\7-\2\2\u017c\u017d\7\'\2\2\u017d\u017e"+
		"\7;\2\2\u017e\u017f\5\62\32\2\u017f\u0180\7<\2\2\u0180\u0181\b\31\1\2"+
		"\u0181\61\3\2\2\2\u0182\u0183\b\32\1\2\u0183\u0184\7\66\2\2\u0184\u0185"+
		"\5\62\32\22\u0185\u0186\b\32\1\2\u0186\u019c\3\2\2\2\u0187\u0188\7\67"+
		"\2\2\u0188\u0189\5\62\32\2\u0189\u018a\78\2\2\u018a\u018b\b\32\1\2\u018b"+
		"\u019c\3\2\2\2\u018c\u018d\7$\2\2\u018d\u019c\b\32\1\2\u018e\u018f\7%"+
		"\2\2\u018f\u019c\b\32\1\2\u0190\u0191\7&\2\2\u0191\u019c\b\32\1\2\u0192"+
		"\u0193\7\b\2\2\u0193\u019c\b\32\1\2\u0194\u0195\7\t\2\2\u0195\u019c\b"+
		"\32\1\2\u0196\u0197\7\'\2\2\u0197\u019c\b\32\1\2\u0198\u0199\5.\30\2\u0199"+
		"\u019a\b\32\1\2\u019a\u019c\3\2\2\2\u019b\u0182\3\2\2\2\u019b\u0187\3"+
		"\2\2\2\u019b\u018c\3\2\2\2\u019b\u018e\3\2\2\2\u019b\u0190\3\2\2\2\u019b"+
		"\u0192\3\2\2\2\u019b\u0194\3\2\2\2\u019b\u0196\3\2\2\2\u019b\u0198\3\2"+
		"\2\2\u019c\u01c2\3\2\2\2\u019d\u019e\f\21\2\2\u019e\u019f\t\2\2\2\u019f"+
		"\u01a0\5\62\32\22\u01a0\u01a1\b\32\1\2\u01a1\u01c1\3\2\2\2\u01a2\u01a3"+
		"\f\20\2\2\u01a3\u01a4\t\3\2\2\u01a4\u01a5\5\62\32\21\u01a5\u01a6\b\32"+
		"\1\2\u01a6\u01c1\3\2\2\2\u01a7\u01a8\f\17\2\2\u01a8\u01a9\t\4\2\2\u01a9"+
		"\u01aa\5\62\32\20\u01aa\u01ab\b\32\1\2\u01ab\u01c1\3\2\2\2\u01ac\u01ad"+
		"\f\16\2\2\u01ad\u01ae\t\5\2\2\u01ae\u01af\5\62\32\17\u01af\u01b0\b\32"+
		"\1\2\u01b0\u01c1\3\2\2\2\u01b1\u01b2\f\r\2\2\u01b2\u01b3\t\6\2\2\u01b3"+
		"\u01b4\5\62\32\16\u01b4\u01b5\b\32\1\2\u01b5\u01c1\3\2\2\2\u01b6\u01b7"+
		"\f\f\2\2\u01b7\u01b8\7,\2\2\u01b8\u01b9\5\62\32\r\u01b9\u01ba\b\32\1\2"+
		"\u01ba\u01c1\3\2\2\2\u01bb\u01bc\f\13\2\2\u01bc\u01bd\7+\2\2\u01bd\u01be"+
		"\5\62\32\f\u01be\u01bf\b\32\1\2\u01bf\u01c1\3\2\2\2\u01c0\u019d\3\2\2"+
		"\2\u01c0\u01a2\3\2\2\2\u01c0\u01a7\3\2\2\2\u01c0\u01ac\3\2\2\2\u01c0\u01b1"+
		"\3\2\2\2\u01c0\u01b6\3\2\2\2\u01c0\u01bb\3\2\2\2\u01c1\u01c4\3\2\2\2\u01c2"+
		"\u01c0\3\2\2\2\u01c2\u01c3\3\2\2\2\u01c3\63\3\2\2\2\u01c4\u01c2\3\2\2"+
		"\2\30;f}\u0081\u008e\u00b2\u00be\u00d1\u00e1\u00e9\u00ec\u00f6\u00fe\u011d"+
		"\u0131\u0141\u014d\u0165\u0175\u019b\u01c0\u01c2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}