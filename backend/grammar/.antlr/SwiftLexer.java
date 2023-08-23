// Generated from c:\Users\TheJhonX\Documents\GitHub\OLC2_201900831\backend\grammar\SwiftLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, STR=4, CHARACTER=5, TRU=6, FAL=7, NIL=8, VAR=9, 
		LET=10, PRINT=11, IF=12, ELSE=13, SWITCH=14, CASE=15, DEFAULT=16, WHILE=17, 
		FOR=18, IN=19, GUARD=20, BREAK=21, CONTINUE=22, RETURN=23, STRUCT=24, 
		SELF=25, MUTATING=26, FUNC=27, NUMBER=28, CHAR=29, STRING=30, ID=31, DIF=32, 
		IG_IG=33, NOT=34, OR=35, AND=36, IG=37, MAY_IG=38, MEN_IG=39, MAYOR=40, 
		MENOR=41, MOD=42, MUL=43, DIV=44, ADD=45, SUB=46, PARIZQ=47, PARDER=48, 
		LLAVEIZQ=49, LLAVEDER=50, Q_MARK=51, ARROW=52, COMA=53, PUNTO=54, COLON=55, 
		WHITESPACE=56, COMMENT=57, LINE_COMMENT=58;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL", "VAR", 
			"LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", 
			"IN", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", "SELF", "MUTATING", 
			"FUNC", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", "MUL", "DIV", 
			"ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "Q_MARK", "ARROW", 
			"COMA", "PUNTO", "COLON", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'", 
			"'false'", "'nil'", "'var'", "'let'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'while'", "'for'", "'in'", "'guard'", "'break'", 
			"'continue'", "'return'", "'struct'", "'self'", "'mutating'", "'func'", 
			null, null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", 
			"'>='", "'<='", "'>'", "'<'", "'%'", "'*'", "'/'", "'+'", "'-'", "'('", 
			"')'", "'{'", "'}'", "'?'", "'->'", "','", "'.'", "':'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "CHARACTER", "TRU", "FAL", "NIL", 
			"VAR", "LET", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "IN", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", "SELF", 
			"MUTATING", "FUNC", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", 
			"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
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


	public SwiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2<\u0190\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\3\2\3"+
		"\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3"+
		"\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3"+
		"\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34"+
		"\3\34\3\34\3\34\3\35\6\35\u0117\n\35\r\35\16\35\u0118\3\35\3\35\6\35\u011d"+
		"\n\35\r\35\16\35\u011e\5\35\u0121\n\35\3\36\3\36\3\36\3\36\3\37\3\37\7"+
		"\37\u0129\n\37\f\37\16\37\u012c\13\37\3\37\3\37\3 \3 \7 \u0132\n \f \16"+
		" \u0135\13 \3!\3!\3!\3\"\3\"\3\"\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3\'\3\'"+
		"\3\'\3(\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61"+
		"\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\65\3\66\3\66\3\67\3\67"+
		"\38\38\39\69\u016f\n9\r9\169\u0170\39\39\3:\3:\3:\3:\7:\u0179\n:\f:\16"+
		":\u017c\13:\3:\3:\3:\3:\3:\3;\3;\3;\3;\7;\u0187\n;\f;\16;\u018a\13;\3"+
		";\3;\3<\3<\3<\3\u017a\2=\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f"+
		"\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63"+
		"\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62"+
		"c\63e\64g\65i\66k\67m8o9q:s;u<w\2\3\2\t\3\2\62;\3\2$$\5\2C\\aac|\6\2\62"+
		";C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u0196"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2"+
		"\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\3y\3\2\2\2\5}\3\2\2\2\7"+
		"\u0083\3\2\2\2\t\u0088\3\2\2\2\13\u008f\3\2\2\2\r\u0099\3\2\2\2\17\u009e"+
		"\3\2\2\2\21\u00a4\3\2\2\2\23\u00a8\3\2\2\2\25\u00ac\3\2\2\2\27\u00b0\3"+
		"\2\2\2\31\u00b6\3\2\2\2\33\u00b9\3\2\2\2\35\u00be\3\2\2\2\37\u00c5\3\2"+
		"\2\2!\u00ca\3\2\2\2#\u00d2\3\2\2\2%\u00d8\3\2\2\2\'\u00dc\3\2\2\2)\u00df"+
		"\3\2\2\2+\u00e5\3\2\2\2-\u00eb\3\2\2\2/\u00f4\3\2\2\2\61\u00fb\3\2\2\2"+
		"\63\u0102\3\2\2\2\65\u0107\3\2\2\2\67\u0110\3\2\2\29\u0116\3\2\2\2;\u0122"+
		"\3\2\2\2=\u0126\3\2\2\2?\u012f\3\2\2\2A\u0136\3\2\2\2C\u0139\3\2\2\2E"+
		"\u013c\3\2\2\2G\u013e\3\2\2\2I\u0141\3\2\2\2K\u0144\3\2\2\2M\u0146\3\2"+
		"\2\2O\u0149\3\2\2\2Q\u014c\3\2\2\2S\u014e\3\2\2\2U\u0150\3\2\2\2W\u0152"+
		"\3\2\2\2Y\u0154\3\2\2\2[\u0156\3\2\2\2]\u0158\3\2\2\2_\u015a\3\2\2\2a"+
		"\u015c\3\2\2\2c\u015e\3\2\2\2e\u0160\3\2\2\2g\u0162\3\2\2\2i\u0164\3\2"+
		"\2\2k\u0167\3\2\2\2m\u0169\3\2\2\2o\u016b\3\2\2\2q\u016e\3\2\2\2s\u0174"+
		"\3\2\2\2u\u0182\3\2\2\2w\u018d\3\2\2\2yz\7K\2\2z{\7p\2\2{|\7v\2\2|\4\3"+
		"\2\2\2}~\7H\2\2~\177\7n\2\2\177\u0080\7q\2\2\u0080\u0081\7c\2\2\u0081"+
		"\u0082\7v\2\2\u0082\6\3\2\2\2\u0083\u0084\7D\2\2\u0084\u0085\7q\2\2\u0085"+
		"\u0086\7q\2\2\u0086\u0087\7n\2\2\u0087\b\3\2\2\2\u0088\u0089\7U\2\2\u0089"+
		"\u008a\7v\2\2\u008a\u008b\7t\2\2\u008b\u008c\7k\2\2\u008c\u008d\7p\2\2"+
		"\u008d\u008e\7i\2\2\u008e\n\3\2\2\2\u008f\u0090\7E\2\2\u0090\u0091\7j"+
		"\2\2\u0091\u0092\7c\2\2\u0092\u0093\7t\2\2\u0093\u0094\7c\2\2\u0094\u0095"+
		"\7e\2\2\u0095\u0096\7v\2\2\u0096\u0097\7g\2\2\u0097\u0098\7t\2\2\u0098"+
		"\f\3\2\2\2\u0099\u009a\7v\2\2\u009a\u009b\7t\2\2\u009b\u009c\7w\2\2\u009c"+
		"\u009d\7g\2\2\u009d\16\3\2\2\2\u009e\u009f\7h\2\2\u009f\u00a0\7c\2\2\u00a0"+
		"\u00a1\7n\2\2\u00a1\u00a2\7u\2\2\u00a2\u00a3\7g\2\2\u00a3\20\3\2\2\2\u00a4"+
		"\u00a5\7p\2\2\u00a5\u00a6\7k\2\2\u00a6\u00a7\7n\2\2\u00a7\22\3\2\2\2\u00a8"+
		"\u00a9\7x\2\2\u00a9\u00aa\7c\2\2\u00aa\u00ab\7t\2\2\u00ab\24\3\2\2\2\u00ac"+
		"\u00ad\7n\2\2\u00ad\u00ae\7g\2\2\u00ae\u00af\7v\2\2\u00af\26\3\2\2\2\u00b0"+
		"\u00b1\7r\2\2\u00b1\u00b2\7t\2\2\u00b2\u00b3\7k\2\2\u00b3\u00b4\7p\2\2"+
		"\u00b4\u00b5\7v\2\2\u00b5\30\3\2\2\2\u00b6\u00b7\7k\2\2\u00b7\u00b8\7"+
		"h\2\2\u00b8\32\3\2\2\2\u00b9\u00ba\7g\2\2\u00ba\u00bb\7n\2\2\u00bb\u00bc"+
		"\7u\2\2\u00bc\u00bd\7g\2\2\u00bd\34\3\2\2\2\u00be\u00bf\7u\2\2\u00bf\u00c0"+
		"\7y\2\2\u00c0\u00c1\7k\2\2\u00c1\u00c2\7v\2\2\u00c2\u00c3\7e\2\2\u00c3"+
		"\u00c4\7j\2\2\u00c4\36\3\2\2\2\u00c5\u00c6\7e\2\2\u00c6\u00c7\7c\2\2\u00c7"+
		"\u00c8\7u\2\2\u00c8\u00c9\7g\2\2\u00c9 \3\2\2\2\u00ca\u00cb\7f\2\2\u00cb"+
		"\u00cc\7g\2\2\u00cc\u00cd\7h\2\2\u00cd\u00ce\7c\2\2\u00ce\u00cf\7w\2\2"+
		"\u00cf\u00d0\7n\2\2\u00d0\u00d1\7v\2\2\u00d1\"\3\2\2\2\u00d2\u00d3\7y"+
		"\2\2\u00d3\u00d4\7j\2\2\u00d4\u00d5\7k\2\2\u00d5\u00d6\7n\2\2\u00d6\u00d7"+
		"\7g\2\2\u00d7$\3\2\2\2\u00d8\u00d9\7h\2\2\u00d9\u00da\7q\2\2\u00da\u00db"+
		"\7t\2\2\u00db&\3\2\2\2\u00dc\u00dd\7k\2\2\u00dd\u00de\7p\2\2\u00de(\3"+
		"\2\2\2\u00df\u00e0\7i\2\2\u00e0\u00e1\7w\2\2\u00e1\u00e2\7c\2\2\u00e2"+
		"\u00e3\7t\2\2\u00e3\u00e4\7f\2\2\u00e4*\3\2\2\2\u00e5\u00e6\7d\2\2\u00e6"+
		"\u00e7\7t\2\2\u00e7\u00e8\7g\2\2\u00e8\u00e9\7c\2\2\u00e9\u00ea\7m\2\2"+
		"\u00ea,\3\2\2\2\u00eb\u00ec\7e\2\2\u00ec\u00ed\7q\2\2\u00ed\u00ee\7p\2"+
		"\2\u00ee\u00ef\7v\2\2\u00ef\u00f0\7k\2\2\u00f0\u00f1\7p\2\2\u00f1\u00f2"+
		"\7w\2\2\u00f2\u00f3\7g\2\2\u00f3.\3\2\2\2\u00f4\u00f5\7t\2\2\u00f5\u00f6"+
		"\7g\2\2\u00f6\u00f7\7v\2\2\u00f7\u00f8\7w\2\2\u00f8\u00f9\7t\2\2\u00f9"+
		"\u00fa\7p\2\2\u00fa\60\3\2\2\2\u00fb\u00fc\7u\2\2\u00fc\u00fd\7v\2\2\u00fd"+
		"\u00fe\7t\2\2\u00fe\u00ff\7w\2\2\u00ff\u0100\7e\2\2\u0100\u0101\7v\2\2"+
		"\u0101\62\3\2\2\2\u0102\u0103\7u\2\2\u0103\u0104\7g\2\2\u0104\u0105\7"+
		"n\2\2\u0105\u0106\7h\2\2\u0106\64\3\2\2\2\u0107\u0108\7o\2\2\u0108\u0109"+
		"\7w\2\2\u0109\u010a\7v\2\2\u010a\u010b\7c\2\2\u010b\u010c\7v\2\2\u010c"+
		"\u010d\7k\2\2\u010d\u010e\7p\2\2\u010e\u010f\7i\2\2\u010f\66\3\2\2\2\u0110"+
		"\u0111\7h\2\2\u0111\u0112\7w\2\2\u0112\u0113\7p\2\2\u0113\u0114\7e\2\2"+
		"\u01148\3\2\2\2\u0115\u0117\t\2\2\2\u0116\u0115\3\2\2\2\u0117\u0118\3"+
		"\2\2\2\u0118\u0116\3\2\2\2\u0118\u0119\3\2\2\2\u0119\u0120\3\2\2\2\u011a"+
		"\u011c\7\60\2\2\u011b\u011d\t\2\2\2\u011c\u011b\3\2\2\2\u011d\u011e\3"+
		"\2\2\2\u011e\u011c\3\2\2\2\u011e\u011f\3\2\2\2\u011f\u0121\3\2\2\2\u0120"+
		"\u011a\3\2\2\2\u0120\u0121\3\2\2\2\u0121:\3\2\2\2\u0122\u0123\7$\2\2\u0123"+
		"\u0124\13\2\2\2\u0124\u0125\7$\2\2\u0125<\3\2\2\2\u0126\u012a\7$\2\2\u0127"+
		"\u0129\n\3\2\2\u0128\u0127\3\2\2\2\u0129\u012c\3\2\2\2\u012a\u0128\3\2"+
		"\2\2\u012a\u012b\3\2\2\2\u012b\u012d\3\2\2\2\u012c\u012a\3\2\2\2\u012d"+
		"\u012e\7$\2\2\u012e>\3\2\2\2\u012f\u0133\t\4\2\2\u0130\u0132\t\5\2\2\u0131"+
		"\u0130\3\2\2\2\u0132\u0135\3\2\2\2\u0133\u0131\3\2\2\2\u0133\u0134\3\2"+
		"\2\2\u0134@\3\2\2\2\u0135\u0133\3\2\2\2\u0136\u0137\7#\2\2\u0137\u0138"+
		"\7?\2\2\u0138B\3\2\2\2\u0139\u013a\7?\2\2\u013a\u013b\7?\2\2\u013bD\3"+
		"\2\2\2\u013c\u013d\7#\2\2\u013dF\3\2\2\2\u013e\u013f\7~\2\2\u013f\u0140"+
		"\7~\2\2\u0140H\3\2\2\2\u0141\u0142\7(\2\2\u0142\u0143\7(\2\2\u0143J\3"+
		"\2\2\2\u0144\u0145\7?\2\2\u0145L\3\2\2\2\u0146\u0147\7@\2\2\u0147\u0148"+
		"\7?\2\2\u0148N\3\2\2\2\u0149\u014a\7>\2\2\u014a\u014b\7?\2\2\u014bP\3"+
		"\2\2\2\u014c\u014d\7@\2\2\u014dR\3\2\2\2\u014e\u014f\7>\2\2\u014fT\3\2"+
		"\2\2\u0150\u0151\7\'\2\2\u0151V\3\2\2\2\u0152\u0153\7,\2\2\u0153X\3\2"+
		"\2\2\u0154\u0155\7\61\2\2\u0155Z\3\2\2\2\u0156\u0157\7-\2\2\u0157\\\3"+
		"\2\2\2\u0158\u0159\7/\2\2\u0159^\3\2\2\2\u015a\u015b\7*\2\2\u015b`\3\2"+
		"\2\2\u015c\u015d\7+\2\2\u015db\3\2\2\2\u015e\u015f\7}\2\2\u015fd\3\2\2"+
		"\2\u0160\u0161\7\177\2\2\u0161f\3\2\2\2\u0162\u0163\7A\2\2\u0163h\3\2"+
		"\2\2\u0164\u0165\7/\2\2\u0165\u0166\7@\2\2\u0166j\3\2\2\2\u0167\u0168"+
		"\7.\2\2\u0168l\3\2\2\2\u0169\u016a\7\60\2\2\u016an\3\2\2\2\u016b\u016c"+
		"\7<\2\2\u016cp\3\2\2\2\u016d\u016f\t\6\2\2\u016e\u016d\3\2\2\2\u016f\u0170"+
		"\3\2\2\2\u0170\u016e\3\2\2\2\u0170\u0171\3\2\2\2\u0171\u0172\3\2\2\2\u0172"+
		"\u0173\b9\2\2\u0173r\3\2\2\2\u0174\u0175\7\61\2\2\u0175\u0176\7,\2\2\u0176"+
		"\u017a\3\2\2\2\u0177\u0179\13\2\2\2\u0178\u0177\3\2\2\2\u0179\u017c\3"+
		"\2\2\2\u017a\u017b\3\2\2\2\u017a\u0178\3\2\2\2\u017b\u017d\3\2\2\2\u017c"+
		"\u017a\3\2\2\2\u017d\u017e\7,\2\2\u017e\u017f\7\61\2\2\u017f\u0180\3\2"+
		"\2\2\u0180\u0181\b:\2\2\u0181t\3\2\2\2\u0182\u0183\7\61\2\2\u0183\u0184"+
		"\7\61\2\2\u0184\u0188\3\2\2\2\u0185\u0187\n\7\2\2\u0186\u0185\3\2\2\2"+
		"\u0187\u018a\3\2\2\2\u0188\u0186\3\2\2\2\u0188\u0189\3\2\2\2\u0189\u018b"+
		"\3\2\2\2\u018a\u0188\3\2\2\2\u018b\u018c\b;\2\2\u018cv\3\2\2\2\u018d\u018e"+
		"\7^\2\2\u018e\u018f\t\b\2\2\u018fx\3\2\2\2\13\2\u0118\u011e\u0120\u012a"+
		"\u0133\u0170\u017a\u0188\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}