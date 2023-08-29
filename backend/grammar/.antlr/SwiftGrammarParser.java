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
		STRUCT=25, SELF=26, MUTATING=27, FUNC=28, NUMBER=29, CHAR=30, STRING=31, 
		ID=32, DIF=33, IG_IG=34, NOT=35, OR=36, AND=37, IG=38, MAY_IG=39, MEN_IG=40, 
		MAYOR=41, MENOR=42, MOD=43, MUL=44, DIV=45, ADD=46, SUB=47, PARIZQ=48, 
		PARDER=49, LLAVEIZQ=50, LLAVEDER=51, Q_MARK=52, ARROW=53, COMA=54, PUNTO=55, 
		COLON=56, WHITESPACE=57, COMMENT=58, LINE_COMMENT=59;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_ifstmt = 4, RULE_elseifstmt = 5, RULE_elsestmt = 6, RULE_varstmt = 7, 
		RULE_tipo = 8, RULE_varasgmt = 9, RULE_conststmt = 10, RULE_switchstmt = 11, 
		RULE_casestmt = 12, RULE_defaultstmt = 13, RULE_whilestmt = 14, RULE_forstmt = 15, 
		RULE_guardstmt = 16, RULE_transferstmt = 17, RULE_expr = 18;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "ifstmt", "elseifstmt", "elsestmt", 
			"varstmt", "tipo", "varasgmt", "conststmt", "switchstmt", "casestmt", 
			"defaultstmt", "whilestmt", "forstmt", "guardstmt", "transferstmt", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'", 
			"'break'", "'continue'", "'return'", "'struct'", "'self'", "'mutating'", 
			"'func'", null, null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", 
			"'='", "'>='", "'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", 
			"'('", "')'", "'{'", "'}'", "'?'", "'->'", "','", "'.'", "':'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL", 
			"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "IN", "RANGEPTS", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", 
			"SELF", "MUTATING", "FUNC", "NUMBER", "CHAR", "STRING", "ID", "DIF", 
			"IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", 
			"MOD", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"Q_MARK", "ARROW", "COMA", "PUNTO", "COLON", "WHITESPACE", "COMMENT", 
			"LINE_COMMENT"
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
			setState(38);
			((SContext)_localctx).block = block();
			setState(39);
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
			setState(43); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(42);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(45); 
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
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(79);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(49);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(52);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinstr 
				}
				break;
			case VAR:
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				((InstructionContext)_localctx).varstmt = varstmt();
				 _localctx.inst = ((InstructionContext)_localctx).varstmt.var
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 4);
				{
				setState(58);
				((InstructionContext)_localctx).varasgmt = varasgmt();
				 _localctx.inst = ((InstructionContext)_localctx).varasgmt.asgmt
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 5);
				{
				setState(61);
				((InstructionContext)_localctx).conststmt = conststmt();
				 _localctx.inst = ((InstructionContext)_localctx).conststmt.const
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 6);
				{
				setState(64);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchstmt.switchinstr
				}
				break;
			case WHILE:
				enterOuterAlt(_localctx, 7);
				{
				setState(67);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whileinstr
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 8);
				{
				setState(70);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.forinstr
				}
				break;
			case GUARD:
				enterOuterAlt(_localctx, 9);
				{
				setState(73);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardstmt.guardinstr
				}
				break;
			case BREAK:
			case CONTINUE:
			case RETURN:
				enterOuterAlt(_localctx, 10);
				{
				setState(76);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				 _localctx.inst = ((InstructionContext)_localctx).transferstmt.trns
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
			setState(81);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(82);
			match(PARIZQ);
			setState(83);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(84);
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
			setState(119);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(88);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(89);
				match(LLAVEIZQ);
				setState(90);
				((IfstmtContext)_localctx).block = block();
				setState(91);
				match(LLAVEDER);
				 _localctx.ifinstr = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(94);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(95);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(96);
				match(LLAVEIZQ);
				setState(97);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(98);
				match(LLAVEDER);
				setState(100); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(99);
						((IfstmtContext)_localctx).elseifstmt = elseifstmt();
						((IfstmtContext)_localctx).elif.add(((IfstmtContext)_localctx).elseifstmt);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(102); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(104);
				match(ELSE);
				setState(106);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==LLAVEIZQ) {
					{
					setState(105);
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
				setState(110);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(111);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(112);
				match(LLAVEIZQ);
				setState(113);
				((IfstmtContext)_localctx).firstBlk = block();
				setState(114);
				match(LLAVEDER);
				setState(115);
				match(ELSE);
				setState(116);
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
			setState(121);
			((ElseifstmtContext)_localctx).ELSE = match(ELSE);
			setState(122);
			match(IF);
			setState(123);
			((ElseifstmtContext)_localctx).expr = expr(0);
			setState(124);
			match(LLAVEIZQ);
			setState(125);
			((ElseifstmtContext)_localctx).block = block();
			setState(126);
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
			setState(129);
			match(LLAVEIZQ);
			setState(130);
			((ElsestmtContext)_localctx).block = block();
			setState(131);
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
			setState(155);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(134);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(135);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(136);
				match(COLON);
				setState(137);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(138);
				match(IG);
				setState(139);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), ((VarstmtContext)_localctx).tipo.rtipo, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(142);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(143);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(144);
				match(IG);
				setState(145);
				((VarstmtContext)_localctx).expr = expr(0);
				_localctx.var = instructions.NewStmt((((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getLine():0), (((VarstmtContext)_localctx).VAR!=null?((VarstmtContext)_localctx).VAR.getCharPositionInLine():0), (((VarstmtContext)_localctx).ID!=null?((VarstmtContext)_localctx).ID.getText():null), environment.NULL, ((VarstmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(148);
				((VarstmtContext)_localctx).VAR = match(VAR);
				setState(149);
				((VarstmtContext)_localctx).ID = match(ID);
				setState(150);
				match(COLON);
				setState(151);
				((VarstmtContext)_localctx).tipo = tipo();
				setState(152);
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
			setState(167);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(157);
				match(INT);
				_localctx.rtipo = 0
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(159);
				match(FLOAT);
				_localctx.rtipo = 1
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(161);
				match(STR);
				_localctx.rtipo = 2
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 4);
				{
				setState(163);
				match(BOOL);
				_localctx.rtipo = 3
				}
				break;
			case CHARACTER:
				enterOuterAlt(_localctx, 5);
				{
				setState(165);
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
			setState(186);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(170);
				match(IG);
				setState(171);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(175);
				match(ADD);
				setState(176);
				match(IG);
				setState(177);
				((VarasgmtContext)_localctx).expr = expr(0);
				_localctx.asgmt = instructions.NewOpAsgmt((((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getCharPositionInLine():0), (((VarasgmtContext)_localctx).ID!=null?((VarasgmtContext)_localctx).ID.getText():null), ((VarasgmtContext)_localctx).expr.e, "+")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(180);
				((VarasgmtContext)_localctx).ID = match(ID);
				setState(181);
				match(SUB);
				setState(182);
				match(IG);
				setState(183);
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
			setState(202);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(188);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(189);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(190);
				match(IG);
				setState(191);
				((ConststmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewStmt((((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getLine():0), (((ConststmtContext)_localctx).LET!=null?((ConststmtContext)_localctx).LET.getCharPositionInLine():0), (((ConststmtContext)_localctx).ID!=null?((ConststmtContext)_localctx).ID.getText():null), environment.NULL, ((ConststmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				((ConststmtContext)_localctx).LET = match(LET);
				setState(195);
				((ConststmtContext)_localctx).ID = match(ID);
				setState(196);
				match(COLON);
				setState(197);
				((ConststmtContext)_localctx).tipo = tipo();
				setState(198);
				match(IG);
				setState(199);
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
			setState(204);
			((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
			setState(205);
			((SwitchstmtContext)_localctx).expr = expr(0);
			setState(206);
			match(LLAVEIZQ);
			setState(208); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(207);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				((SwitchstmtContext)_localctx).casesvar.add(((SwitchstmtContext)_localctx).casestmt);
				}
				}
				setState(210); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE );
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(212);
				defaultstmt();
				}
			}

			setState(215);
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
			setState(218);
			((CasestmtContext)_localctx).CASE = match(CASE);
			setState(219);
			((CasestmtContext)_localctx).expr = expr(0);
			setState(220);
			match(COLON);
			setState(221);
			((CasestmtContext)_localctx).block = block();
			setState(223);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(222);
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
			setState(227);
			match(DEFAULT);
			setState(228);
			match(COLON);
			setState(229);
			((DefaultstmtContext)_localctx).block = block();
			setState(231);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(230);
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
			setState(235);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(236);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(237);
			match(LLAVEIZQ);
			setState(238);
			((WhilestmtContext)_localctx).block = block();
			setState(239);
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
			setState(262);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(243);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(244);
				match(IN);
				setState(245);
				((ForstmtContext)_localctx).STRING = match(STRING);
				setState(246);
				match(LLAVEIZQ);
				setState(247);
				((ForstmtContext)_localctx).block = block();
				setState(248);
				match(LLAVEDER);
				 
				    str = (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getText():null)
				    cadena = expressions.NewPrimitive((((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getLine():0), (((ForstmtContext)_localctx).STRING!=null?((ForstmtContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1], environment.STRING)
				    _localctx.forinstr = instructions.NewForIn((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), cadena, nil, nil, ((ForstmtContext)_localctx).block.blk) 

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(251);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(252);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(253);
				match(IN);
				setState(254);
				((ForstmtContext)_localctx).left = expr(0);
				setState(255);
				match(RANGEPTS);
				setState(256);
				((ForstmtContext)_localctx).right = expr(0);
				setState(257);
				match(LLAVEIZQ);
				setState(258);
				((ForstmtContext)_localctx).block = block();
				setState(259);
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
			setState(264);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(265);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(266);
			match(ELSE);
			setState(267);
			match(LLAVEIZQ);
			setState(268);
			((GuardstmtContext)_localctx).block = block();
			setState(269);
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
			setState(282);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(272);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				 _localctx.trns = instructions.NewBreak((((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getLine():0), (((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getCharPositionInLine():0)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(274);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				 _localctx.trns = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0)) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(276);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				 _localctx.trns = instructions.NewReturn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(278);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(279);
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
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SUB:
				{
				setState(285);
				((ExprContext)_localctx).op = match(SUB);
				setState(286);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetLine(), (((ExprContext)_localctx).right!=null?(((ExprContext)_localctx).right.start):null).GetColumn(), ((ExprContext)_localctx).right.e, "unario", expressions.NewPrimitive((((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getLine():0), (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getCharPositionInLine():0), -1, environment.INTEGER)) 
				}
				break;
			case PARIZQ:
				{
				setState(289);
				match(PARIZQ);
				setState(290);
				((ExprContext)_localctx).expr = expr(0);
				setState(291);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case NUMBER:
				{
				setState(294);
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
			case CHAR:
				{
				setState(296);
				((ExprContext)_localctx).CHAR = match(CHAR);
				 
				        str := (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getLine():0), (((ExprContext)_localctx).CHAR!=null?((ExprContext)_localctx).CHAR.getCharPositionInLine():0), str[1:len(str)-1], environment.CHAR) 
				    
				}
				break;
			case STRING:
				{
				setState(298);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRU:
				{
				setState(300);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FAL:
				{
				setState(302);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case ID:
				{
				setState(304);
				((ExprContext)_localctx).ID = match(ID);
				 _localctx.e = expressions.NewVar((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(345);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(343);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(308);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(309);
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
						setState(310);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(313);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(314);
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
						setState(315);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(318);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(319);
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
						setState(320);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(323);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(324);
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
						setState(325);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(328);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(329);
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
						setState(330);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(333);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(334);
						match(AND);
						setState(335);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(10);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(338);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(339);
						match(OR);
						setState(340);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(9);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(347);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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
		case 18:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 14);
		case 1:
			return precpred(_ctx, 13);
		case 2:
			return precpred(_ctx, 12);
		case 3:
			return precpred(_ctx, 11);
		case 4:
			return precpred(_ctx, 10);
		case 5:
			return precpred(_ctx, 9);
		case 6:
			return precpred(_ctx, 8);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3=\u015f\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\2\3\2\3\3\6\3.\n\3\r\3\16\3/\3\3\3\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4R\n\4\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\6\6g"+
		"\n\6\r\6\16\6h\3\6\3\6\5\6m\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\5\6z\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\5\t\u009e\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n"+
		"\u00aa\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\5\13\u00bd\n\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00cd\n\f\3\r\3\r\3\r\3\r\6\r\u00d3\n"+
		"\r\r\r\16\r\u00d4\3\r\5\r\u00d8\n\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3"+
		"\16\5\16\u00e2\n\16\3\16\3\16\3\17\3\17\3\17\3\17\5\17\u00ea\n\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\5\21\u0109\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u011d\n\23\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\5\24\u0135\n\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\7\24\u015a\n\24\f\24\16\24\u015d\13\24\3\24\2\3&\25\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&\2\7\3\2-/\3\2\60\61\4\2))++\4\2**,,\3\2#$\2"+
		"\u0178\2(\3\2\2\2\4-\3\2\2\2\6Q\3\2\2\2\bS\3\2\2\2\ny\3\2\2\2\f{\3\2\2"+
		"\2\16\u0083\3\2\2\2\20\u009d\3\2\2\2\22\u00a9\3\2\2\2\24\u00bc\3\2\2\2"+
		"\26\u00cc\3\2\2\2\30\u00ce\3\2\2\2\32\u00dc\3\2\2\2\34\u00e5\3\2\2\2\36"+
		"\u00ed\3\2\2\2 \u0108\3\2\2\2\"\u010a\3\2\2\2$\u011c\3\2\2\2&\u0134\3"+
		"\2\2\2()\5\4\3\2)*\7\2\2\3*+\b\2\1\2+\3\3\2\2\2,.\5\6\4\2-,\3\2\2\2./"+
		"\3\2\2\2/-\3\2\2\2/\60\3\2\2\2\60\61\3\2\2\2\61\62\b\3\1\2\62\5\3\2\2"+
		"\2\63\64\5\b\5\2\64\65\b\4\1\2\65R\3\2\2\2\66\67\5\n\6\2\678\b\4\1\28"+
		"R\3\2\2\29:\5\20\t\2:;\b\4\1\2;R\3\2\2\2<=\5\24\13\2=>\b\4\1\2>R\3\2\2"+
		"\2?@\5\26\f\2@A\b\4\1\2AR\3\2\2\2BC\5\30\r\2CD\b\4\1\2DR\3\2\2\2EF\5\36"+
		"\20\2FG\b\4\1\2GR\3\2\2\2HI\5 \21\2IJ\b\4\1\2JR\3\2\2\2KL\5\"\22\2LM\b"+
		"\4\1\2MR\3\2\2\2NO\5$\23\2OP\b\4\1\2PR\3\2\2\2Q\63\3\2\2\2Q\66\3\2\2\2"+
		"Q9\3\2\2\2Q<\3\2\2\2Q?\3\2\2\2QB\3\2\2\2QE\3\2\2\2QH\3\2\2\2QK\3\2\2\2"+
		"QN\3\2\2\2R\7\3\2\2\2ST\7\r\2\2TU\7\62\2\2UV\5&\24\2VW\7\63\2\2WX\b\5"+
		"\1\2X\t\3\2\2\2YZ\7\16\2\2Z[\5&\24\2[\\\7\64\2\2\\]\5\4\3\2]^\7\65\2\2"+
		"^_\b\6\1\2_z\3\2\2\2`a\7\16\2\2ab\5&\24\2bc\7\64\2\2cd\5\4\3\2df\7\65"+
		"\2\2eg\5\f\7\2fe\3\2\2\2gh\3\2\2\2hf\3\2\2\2hi\3\2\2\2ij\3\2\2\2jl\7\17"+
		"\2\2km\5\16\b\2lk\3\2\2\2lm\3\2\2\2mn\3\2\2\2no\b\6\1\2oz\3\2\2\2pq\7"+
		"\16\2\2qr\5&\24\2rs\7\64\2\2st\5\4\3\2tu\7\65\2\2uv\7\17\2\2vw\5\16\b"+
		"\2wx\b\6\1\2xz\3\2\2\2yY\3\2\2\2y`\3\2\2\2yp\3\2\2\2z\13\3\2\2\2{|\7\17"+
		"\2\2|}\7\16\2\2}~\5&\24\2~\177\7\64\2\2\177\u0080\5\4\3\2\u0080\u0081"+
		"\7\65\2\2\u0081\u0082\b\7\1\2\u0082\r\3\2\2\2\u0083\u0084\7\64\2\2\u0084"+
		"\u0085\5\4\3\2\u0085\u0086\7\65\2\2\u0086\u0087\b\b\1\2\u0087\17\3\2\2"+
		"\2\u0088\u0089\7\13\2\2\u0089\u008a\7\"\2\2\u008a\u008b\7:\2\2\u008b\u008c"+
		"\5\22\n\2\u008c\u008d\7(\2\2\u008d\u008e\5&\24\2\u008e\u008f\b\t\1\2\u008f"+
		"\u009e\3\2\2\2\u0090\u0091\7\13\2\2\u0091\u0092\7\"\2\2\u0092\u0093\7"+
		"(\2\2\u0093\u0094\5&\24\2\u0094\u0095\b\t\1\2\u0095\u009e\3\2\2\2\u0096"+
		"\u0097\7\13\2\2\u0097\u0098\7\"\2\2\u0098\u0099\7:\2\2\u0099\u009a\5\22"+
		"\n\2\u009a\u009b\7\66\2\2\u009b\u009c\b\t\1\2\u009c\u009e\3\2\2\2\u009d"+
		"\u0088\3\2\2\2\u009d\u0090\3\2\2\2\u009d\u0096\3\2\2\2\u009e\21\3\2\2"+
		"\2\u009f\u00a0\7\3\2\2\u00a0\u00aa\b\n\1\2\u00a1\u00a2\7\4\2\2\u00a2\u00aa"+
		"\b\n\1\2\u00a3\u00a4\7\6\2\2\u00a4\u00aa\b\n\1\2\u00a5\u00a6\7\5\2\2\u00a6"+
		"\u00aa\b\n\1\2\u00a7\u00a8\7\7\2\2\u00a8\u00aa\b\n\1\2\u00a9\u009f\3\2"+
		"\2\2\u00a9\u00a1\3\2\2\2\u00a9\u00a3\3\2\2\2\u00a9\u00a5\3\2\2\2\u00a9"+
		"\u00a7\3\2\2\2\u00aa\23\3\2\2\2\u00ab\u00ac\7\"\2\2\u00ac\u00ad\7(\2\2"+
		"\u00ad\u00ae\5&\24\2\u00ae\u00af\b\13\1\2\u00af\u00bd\3\2\2\2\u00b0\u00b1"+
		"\7\"\2\2\u00b1\u00b2\7\60\2\2\u00b2\u00b3\7(\2\2\u00b3\u00b4\5&\24\2\u00b4"+
		"\u00b5\b\13\1\2\u00b5\u00bd\3\2\2\2\u00b6\u00b7\7\"\2\2\u00b7\u00b8\7"+
		"\61\2\2\u00b8\u00b9\7(\2\2\u00b9\u00ba\5&\24\2\u00ba\u00bb\b\13\1\2\u00bb"+
		"\u00bd\3\2\2\2\u00bc\u00ab\3\2\2\2\u00bc\u00b0\3\2\2\2\u00bc\u00b6\3\2"+
		"\2\2\u00bd\25\3\2\2\2\u00be\u00bf\7\f\2\2\u00bf\u00c0\7\"\2\2\u00c0\u00c1"+
		"\7(\2\2\u00c1\u00c2\5&\24\2\u00c2\u00c3\b\f\1\2\u00c3\u00cd\3\2\2\2\u00c4"+
		"\u00c5\7\f\2\2\u00c5\u00c6\7\"\2\2\u00c6\u00c7\7:\2\2\u00c7\u00c8\5\22"+
		"\n\2\u00c8\u00c9\7(\2\2\u00c9\u00ca\5&\24\2\u00ca\u00cb\b\f\1\2\u00cb"+
		"\u00cd\3\2\2\2\u00cc\u00be\3\2\2\2\u00cc\u00c4\3\2\2\2\u00cd\27\3\2\2"+
		"\2\u00ce\u00cf\7\20\2\2\u00cf\u00d0\5&\24\2\u00d0\u00d2\7\64\2\2\u00d1"+
		"\u00d3\5\32\16\2\u00d2\u00d1\3\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d2\3"+
		"\2\2\2\u00d4\u00d5\3\2\2\2\u00d5\u00d7\3\2\2\2\u00d6\u00d8\5\34\17\2\u00d7"+
		"\u00d6\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00da\7\65"+
		"\2\2\u00da\u00db\b\r\1\2\u00db\31\3\2\2\2\u00dc\u00dd\7\21\2\2\u00dd\u00de"+
		"\5&\24\2\u00de\u00df\7:\2\2\u00df\u00e1\5\4\3\2\u00e0\u00e2\7\30\2\2\u00e1"+
		"\u00e0\3\2\2\2\u00e1\u00e2\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e4\b\16"+
		"\1\2\u00e4\33\3\2\2\2\u00e5\u00e6\7\22\2\2\u00e6\u00e7\7:\2\2\u00e7\u00e9"+
		"\5\4\3\2\u00e8\u00ea\7\30\2\2\u00e9\u00e8\3\2\2\2\u00e9\u00ea\3\2\2\2"+
		"\u00ea\u00eb\3\2\2\2\u00eb\u00ec\b\17\1\2\u00ec\35\3\2\2\2\u00ed\u00ee"+
		"\7\23\2\2\u00ee\u00ef\5&\24\2\u00ef\u00f0\7\64\2\2\u00f0\u00f1\5\4\3\2"+
		"\u00f1\u00f2\7\65\2\2\u00f2\u00f3\b\20\1\2\u00f3\37\3\2\2\2\u00f4\u00f5"+
		"\7\24\2\2\u00f5\u00f6\7\"\2\2\u00f6\u00f7\7\25\2\2\u00f7\u00f8\7!\2\2"+
		"\u00f8\u00f9\7\64\2\2\u00f9\u00fa\5\4\3\2\u00fa\u00fb\7\65\2\2\u00fb\u00fc"+
		"\b\21\1\2\u00fc\u0109\3\2\2\2\u00fd\u00fe\7\24\2\2\u00fe\u00ff\7\"\2\2"+
		"\u00ff\u0100\7\25\2\2\u0100\u0101\5&\24\2\u0101\u0102\7\26\2\2\u0102\u0103"+
		"\5&\24\2\u0103\u0104\7\64\2\2\u0104\u0105\5\4\3\2\u0105\u0106\7\65\2\2"+
		"\u0106\u0107\b\21\1\2\u0107\u0109\3\2\2\2\u0108\u00f4\3\2\2\2\u0108\u00fd"+
		"\3\2\2\2\u0109!\3\2\2\2\u010a\u010b\7\27\2\2\u010b\u010c\5&\24\2\u010c"+
		"\u010d\7\17\2\2\u010d\u010e\7\64\2\2\u010e\u010f\5\4\3\2\u010f\u0110\7"+
		"\65\2\2\u0110\u0111\b\22\1\2\u0111#\3\2\2\2\u0112\u0113\7\30\2\2\u0113"+
		"\u011d\b\23\1\2\u0114\u0115\7\31\2\2\u0115\u011d\b\23\1\2\u0116\u0117"+
		"\7\32\2\2\u0117\u011d\b\23\1\2\u0118\u0119\7\32\2\2\u0119\u011a\5&\24"+
		"\2\u011a\u011b\b\23\1\2\u011b\u011d\3\2\2\2\u011c\u0112\3\2\2\2\u011c"+
		"\u0114\3\2\2\2\u011c\u0116\3\2\2\2\u011c\u0118\3\2\2\2\u011d%\3\2\2\2"+
		"\u011e\u011f\b\24\1\2\u011f\u0120\7\61\2\2\u0120\u0121\5&\24\21\u0121"+
		"\u0122\b\24\1\2\u0122\u0135\3\2\2\2\u0123\u0124\7\62\2\2\u0124\u0125\5"+
		"&\24\2\u0125\u0126\7\63\2\2\u0126\u0127\b\24\1\2\u0127\u0135\3\2\2\2\u0128"+
		"\u0129\7\37\2\2\u0129\u0135\b\24\1\2\u012a\u012b\7 \2\2\u012b\u0135\b"+
		"\24\1\2\u012c\u012d\7!\2\2\u012d\u0135\b\24\1\2\u012e\u012f\7\b\2\2\u012f"+
		"\u0135\b\24\1\2\u0130\u0131\7\t\2\2\u0131\u0135\b\24\1\2\u0132\u0133\7"+
		"\"\2\2\u0133\u0135\b\24\1\2\u0134\u011e\3\2\2\2\u0134\u0123\3\2\2\2\u0134"+
		"\u0128\3\2\2\2\u0134\u012a\3\2\2\2\u0134\u012c\3\2\2\2\u0134\u012e\3\2"+
		"\2\2\u0134\u0130\3\2\2\2\u0134\u0132\3\2\2\2\u0135\u015b\3\2\2\2\u0136"+
		"\u0137\f\20\2\2\u0137\u0138\t\2\2\2\u0138\u0139\5&\24\21\u0139\u013a\b"+
		"\24\1\2\u013a\u015a\3\2\2\2\u013b\u013c\f\17\2\2\u013c\u013d\t\3\2\2\u013d"+
		"\u013e\5&\24\20\u013e\u013f\b\24\1\2\u013f\u015a\3\2\2\2\u0140\u0141\f"+
		"\16\2\2\u0141\u0142\t\4\2\2\u0142\u0143\5&\24\17\u0143\u0144\b\24\1\2"+
		"\u0144\u015a\3\2\2\2\u0145\u0146\f\r\2\2\u0146\u0147\t\5\2\2\u0147\u0148"+
		"\5&\24\16\u0148\u0149\b\24\1\2\u0149\u015a\3\2\2\2\u014a\u014b\f\f\2\2"+
		"\u014b\u014c\t\6\2\2\u014c\u014d\5&\24\r\u014d\u014e\b\24\1\2\u014e\u015a"+
		"\3\2\2\2\u014f\u0150\f\13\2\2\u0150\u0151\7\'\2\2\u0151\u0152\5&\24\f"+
		"\u0152\u0153\b\24\1\2\u0153\u015a\3\2\2\2\u0154\u0155\f\n\2\2\u0155\u0156"+
		"\7&\2\2\u0156\u0157\5&\24\13\u0157\u0158\b\24\1\2\u0158\u015a\3\2\2\2"+
		"\u0159\u0136\3\2\2\2\u0159\u013b\3\2\2\2\u0159\u0140\3\2\2\2\u0159\u0145"+
		"\3\2\2\2\u0159\u014a\3\2\2\2\u0159\u014f\3\2\2\2\u0159\u0154\3\2\2\2\u015a"+
		"\u015d\3\2\2\2\u015b\u0159\3\2\2\2\u015b\u015c\3\2\2\2\u015c\'\3\2\2\2"+
		"\u015d\u015b\3\2\2\2\24/Qhly\u009d\u00a9\u00bc\u00cc\u00d4\u00d7\u00e1"+
		"\u00e9\u0108\u011c\u0134\u0159\u015b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}