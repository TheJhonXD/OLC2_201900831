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
		FOR=18, IN=19, RANGEPTS=20, GUARD=21, BREAK=22, CONTINUE=23, RETURN=24, 
		STRUCT=25, SELF=26, MUTATING=27, FUNC=28, APPEND=29, REMOVELAST=30, REMOVE=31, 
		EMPTY=32, COUNT=33, INOUT=34, CAME=35, NUMBER=36, CHAR=37, STRING=38, 
		ID=39, DIF=40, IG_IG=41, NOT=42, OR=43, AND=44, IG=45, MAY_IG=46, MEN_IG=47, 
		MAYOR=48, MENOR=49, MOD=50, MUL=51, DIV=52, ADD=53, SUB=54, PARIZQ=55, 
		PARDER=56, LLAVEIZQ=57, LLAVEDER=58, CORCHIZQ=59, CORCHDER=60, Q_MARK=61, 
		ARROW=62, COMA=63, PUNTO=64, COLON=65, AMP=66, WHITESPACE=67, COMMENT=68, 
		LINE_COMMENT=69;
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
			"IN", "RANGEPTS", "GUARD", "BREAK", "CONTINUE", "RETURN", "STRUCT", "SELF", 
			"MUTATING", "FUNC", "APPEND", "REMOVELAST", "REMOVE", "EMPTY", "COUNT", 
			"INOUT", "CAME", "NUMBER", "CHAR", "STRING", "ID", "DIF", "IG_IG", "NOT", 
			"OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MOD", "MUL", 
			"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORCHIZQ", 
			"CORCHDER", "Q_MARK", "ARROW", "COMA", "PUNTO", "COLON", "AMP", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT", "ESC_SEQ"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2G\u01df\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\3\2\3\2\3"+
		"\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23"+
		"\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3"+
		" \3!\3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3"+
		"$\3$\3%\6%\u0160\n%\r%\16%\u0161\3%\3%\6%\u0166\n%\r%\16%\u0167\5%\u016a"+
		"\n%\3&\3&\3&\3&\3\'\3\'\7\'\u0172\n\'\f\'\16\'\u0175\13\'\3\'\3\'\3(\3"+
		"(\7(\u017b\n(\f(\16(\u017e\13(\3)\3)\3)\3*\3*\3*\3+\3+\3,\3,\3,\3-\3-"+
		"\3-\3.\3.\3/\3/\3/\3\60\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3"+
		"\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3="+
		"\3>\3>\3?\3?\3?\3@\3@\3A\3A\3B\3B\3C\3C\3D\6D\u01be\nD\rD\16D\u01bf\3"+
		"D\3D\3E\3E\3E\3E\7E\u01c8\nE\fE\16E\u01cb\13E\3E\3E\3E\3E\3E\3F\3F\3F"+
		"\3F\7F\u01d6\nF\fF\16F\u01d9\13F\3F\3F\3G\3G\3G\3\u01c9\2H\3\3\5\4\7\5"+
		"\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{"+
		"?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008d\2\3\2\t\3\2\62"+
		";\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t"+
		"\2\"#%%--/\60<<BB]_\2\u01e5\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3"+
		"\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2"+
		"\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37"+
		"\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3"+
		"\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2"+
		"\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C"+
		"\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2"+
		"\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2"+
		"\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i"+
		"\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2"+
		"\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081"+
		"\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3\2\2\2\2\u0089\3\2\2"+
		"\2\2\u008b\3\2\2\2\3\u008f\3\2\2\2\5\u0093\3\2\2\2\7\u0099\3\2\2\2\t\u009e"+
		"\3\2\2\2\13\u00a5\3\2\2\2\r\u00af\3\2\2\2\17\u00b4\3\2\2\2\21\u00ba\3"+
		"\2\2\2\23\u00be\3\2\2\2\25\u00c2\3\2\2\2\27\u00c6\3\2\2\2\31\u00cc\3\2"+
		"\2\2\33\u00cf\3\2\2\2\35\u00d4\3\2\2\2\37\u00db\3\2\2\2!\u00e0\3\2\2\2"+
		"#\u00e8\3\2\2\2%\u00ee\3\2\2\2\'\u00f2\3\2\2\2)\u00f5\3\2\2\2+\u00f9\3"+
		"\2\2\2-\u00ff\3\2\2\2/\u0105\3\2\2\2\61\u010e\3\2\2\2\63\u0115\3\2\2\2"+
		"\65\u011c\3\2\2\2\67\u0121\3\2\2\29\u012a\3\2\2\2;\u012f\3\2\2\2=\u0136"+
		"\3\2\2\2?\u0141\3\2\2\2A\u0148\3\2\2\2C\u0150\3\2\2\2E\u0156\3\2\2\2G"+
		"\u015c\3\2\2\2I\u015f\3\2\2\2K\u016b\3\2\2\2M\u016f\3\2\2\2O\u0178\3\2"+
		"\2\2Q\u017f\3\2\2\2S\u0182\3\2\2\2U\u0185\3\2\2\2W\u0187\3\2\2\2Y\u018a"+
		"\3\2\2\2[\u018d\3\2\2\2]\u018f\3\2\2\2_\u0192\3\2\2\2a\u0195\3\2\2\2c"+
		"\u0197\3\2\2\2e\u0199\3\2\2\2g\u019b\3\2\2\2i\u019d\3\2\2\2k\u019f\3\2"+
		"\2\2m\u01a1\3\2\2\2o\u01a3\3\2\2\2q\u01a5\3\2\2\2s\u01a7\3\2\2\2u\u01a9"+
		"\3\2\2\2w\u01ab\3\2\2\2y\u01ad\3\2\2\2{\u01af\3\2\2\2}\u01b1\3\2\2\2\177"+
		"\u01b4\3\2\2\2\u0081\u01b6\3\2\2\2\u0083\u01b8\3\2\2\2\u0085\u01ba\3\2"+
		"\2\2\u0087\u01bd\3\2\2\2\u0089\u01c3\3\2\2\2\u008b\u01d1\3\2\2\2\u008d"+
		"\u01dc\3\2\2\2\u008f\u0090\7K\2\2\u0090\u0091\7p\2\2\u0091\u0092\7v\2"+
		"\2\u0092\4\3\2\2\2\u0093\u0094\7H\2\2\u0094\u0095\7n\2\2\u0095\u0096\7"+
		"q\2\2\u0096\u0097\7c\2\2\u0097\u0098\7v\2\2\u0098\6\3\2\2\2\u0099\u009a"+
		"\7D\2\2\u009a\u009b\7q\2\2\u009b\u009c\7q\2\2\u009c\u009d\7n\2\2\u009d"+
		"\b\3\2\2\2\u009e\u009f\7U\2\2\u009f\u00a0\7v\2\2\u00a0\u00a1\7t\2\2\u00a1"+
		"\u00a2\7k\2\2\u00a2\u00a3\7p\2\2\u00a3\u00a4\7i\2\2\u00a4\n\3\2\2\2\u00a5"+
		"\u00a6\7E\2\2\u00a6\u00a7\7j\2\2\u00a7\u00a8\7c\2\2\u00a8\u00a9\7t\2\2"+
		"\u00a9\u00aa\7c\2\2\u00aa\u00ab\7e\2\2\u00ab\u00ac\7v\2\2\u00ac\u00ad"+
		"\7g\2\2\u00ad\u00ae\7t\2\2\u00ae\f\3\2\2\2\u00af\u00b0\7v\2\2\u00b0\u00b1"+
		"\7t\2\2\u00b1\u00b2\7w\2\2\u00b2\u00b3\7g\2\2\u00b3\16\3\2\2\2\u00b4\u00b5"+
		"\7h\2\2\u00b5\u00b6\7c\2\2\u00b6\u00b7\7n\2\2\u00b7\u00b8\7u\2\2\u00b8"+
		"\u00b9\7g\2\2\u00b9\20\3\2\2\2\u00ba\u00bb\7p\2\2\u00bb\u00bc\7k\2\2\u00bc"+
		"\u00bd\7n\2\2\u00bd\22\3\2\2\2\u00be\u00bf\7x\2\2\u00bf\u00c0\7c\2\2\u00c0"+
		"\u00c1\7t\2\2\u00c1\24\3\2\2\2\u00c2\u00c3\7n\2\2\u00c3\u00c4\7g\2\2\u00c4"+
		"\u00c5\7v\2\2\u00c5\26\3\2\2\2\u00c6\u00c7\7r\2\2\u00c7\u00c8\7t\2\2\u00c8"+
		"\u00c9\7k\2\2\u00c9\u00ca\7p\2\2\u00ca\u00cb\7v\2\2\u00cb\30\3\2\2\2\u00cc"+
		"\u00cd\7k\2\2\u00cd\u00ce\7h\2\2\u00ce\32\3\2\2\2\u00cf\u00d0\7g\2\2\u00d0"+
		"\u00d1\7n\2\2\u00d1\u00d2\7u\2\2\u00d2\u00d3\7g\2\2\u00d3\34\3\2\2\2\u00d4"+
		"\u00d5\7u\2\2\u00d5\u00d6\7y\2\2\u00d6\u00d7\7k\2\2\u00d7\u00d8\7v\2\2"+
		"\u00d8\u00d9\7e\2\2\u00d9\u00da\7j\2\2\u00da\36\3\2\2\2\u00db\u00dc\7"+
		"e\2\2\u00dc\u00dd\7c\2\2\u00dd\u00de\7u\2\2\u00de\u00df\7g\2\2\u00df "+
		"\3\2\2\2\u00e0\u00e1\7f\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7h\2\2\u00e3"+
		"\u00e4\7c\2\2\u00e4\u00e5\7w\2\2\u00e5\u00e6\7n\2\2\u00e6\u00e7\7v\2\2"+
		"\u00e7\"\3\2\2\2\u00e8\u00e9\7y\2\2\u00e9\u00ea\7j\2\2\u00ea\u00eb\7k"+
		"\2\2\u00eb\u00ec\7n\2\2\u00ec\u00ed\7g\2\2\u00ed$\3\2\2\2\u00ee\u00ef"+
		"\7h\2\2\u00ef\u00f0\7q\2\2\u00f0\u00f1\7t\2\2\u00f1&\3\2\2\2\u00f2\u00f3"+
		"\7k\2\2\u00f3\u00f4\7p\2\2\u00f4(\3\2\2\2\u00f5\u00f6\7\60\2\2\u00f6\u00f7"+
		"\7\60\2\2\u00f7\u00f8\7\60\2\2\u00f8*\3\2\2\2\u00f9\u00fa\7i\2\2\u00fa"+
		"\u00fb\7w\2\2\u00fb\u00fc\7c\2\2\u00fc\u00fd\7t\2\2\u00fd\u00fe\7f\2\2"+
		"\u00fe,\3\2\2\2\u00ff\u0100\7d\2\2\u0100\u0101\7t\2\2\u0101\u0102\7g\2"+
		"\2\u0102\u0103\7c\2\2\u0103\u0104\7m\2\2\u0104.\3\2\2\2\u0105\u0106\7"+
		"e\2\2\u0106\u0107\7q\2\2\u0107\u0108\7p\2\2\u0108\u0109\7v\2\2\u0109\u010a"+
		"\7k\2\2\u010a\u010b\7p\2\2\u010b\u010c\7w\2\2\u010c\u010d\7g\2\2\u010d"+
		"\60\3\2\2\2\u010e\u010f\7t\2\2\u010f\u0110\7g\2\2\u0110\u0111\7v\2\2\u0111"+
		"\u0112\7w\2\2\u0112\u0113\7t\2\2\u0113\u0114\7p\2\2\u0114\62\3\2\2\2\u0115"+
		"\u0116\7u\2\2\u0116\u0117\7v\2\2\u0117\u0118\7t\2\2\u0118\u0119\7w\2\2"+
		"\u0119\u011a\7e\2\2\u011a\u011b\7v\2\2\u011b\64\3\2\2\2\u011c\u011d\7"+
		"u\2\2\u011d\u011e\7g\2\2\u011e\u011f\7n\2\2\u011f\u0120\7h\2\2\u0120\66"+
		"\3\2\2\2\u0121\u0122\7o\2\2\u0122\u0123\7w\2\2\u0123\u0124\7v\2\2\u0124"+
		"\u0125\7c\2\2\u0125\u0126\7v\2\2\u0126\u0127\7k\2\2\u0127\u0128\7p\2\2"+
		"\u0128\u0129\7i\2\2\u01298\3\2\2\2\u012a\u012b\7h\2\2\u012b\u012c\7w\2"+
		"\2\u012c\u012d\7p\2\2\u012d\u012e\7e\2\2\u012e:\3\2\2\2\u012f\u0130\7"+
		"c\2\2\u0130\u0131\7r\2\2\u0131\u0132\7r\2\2\u0132\u0133\7g\2\2\u0133\u0134"+
		"\7p\2\2\u0134\u0135\7f\2\2\u0135<\3\2\2\2\u0136\u0137\7t\2\2\u0137\u0138"+
		"\7g\2\2\u0138\u0139\7o\2\2\u0139\u013a\7q\2\2\u013a\u013b\7x\2\2\u013b"+
		"\u013c\7g\2\2\u013c\u013d\7N\2\2\u013d\u013e\7c\2\2\u013e\u013f\7u\2\2"+
		"\u013f\u0140\7v\2\2\u0140>\3\2\2\2\u0141\u0142\7t\2\2\u0142\u0143\7g\2"+
		"\2\u0143\u0144\7o\2\2\u0144\u0145\7q\2\2\u0145\u0146\7x\2\2\u0146\u0147"+
		"\7g\2\2\u0147@\3\2\2\2\u0148\u0149\7k\2\2\u0149\u014a\7u\2\2\u014a\u014b"+
		"\7G\2\2\u014b\u014c\7o\2\2\u014c\u014d\7r\2\2\u014d\u014e\7v\2\2\u014e"+
		"\u014f\7{\2\2\u014fB\3\2\2\2\u0150\u0151\7e\2\2\u0151\u0152\7q\2\2\u0152"+
		"\u0153\7w\2\2\u0153\u0154\7p\2\2\u0154\u0155\7v\2\2\u0155D\3\2\2\2\u0156"+
		"\u0157\7k\2\2\u0157\u0158\7p\2\2\u0158\u0159\7q\2\2\u0159\u015a\7w\2\2"+
		"\u015a\u015b\7v\2\2\u015bF\3\2\2\2\u015c\u015d\7a\2\2\u015dH\3\2\2\2\u015e"+
		"\u0160\t\2\2\2\u015f\u015e\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u015f\3\2"+
		"\2\2\u0161\u0162\3\2\2\2\u0162\u0169\3\2\2\2\u0163\u0165\7\60\2\2\u0164"+
		"\u0166\t\2\2\2\u0165\u0164\3\2\2\2\u0166\u0167\3\2\2\2\u0167\u0165\3\2"+
		"\2\2\u0167\u0168\3\2\2\2\u0168\u016a\3\2\2\2\u0169\u0163\3\2\2\2\u0169"+
		"\u016a\3\2\2\2\u016aJ\3\2\2\2\u016b\u016c\7$\2\2\u016c\u016d\13\2\2\2"+
		"\u016d\u016e\7$\2\2\u016eL\3\2\2\2\u016f\u0173\7$\2\2\u0170\u0172\n\3"+
		"\2\2\u0171\u0170\3\2\2\2\u0172\u0175\3\2\2\2\u0173\u0171\3\2\2\2\u0173"+
		"\u0174\3\2\2\2\u0174\u0176\3\2\2\2\u0175\u0173\3\2\2\2\u0176\u0177\7$"+
		"\2\2\u0177N\3\2\2\2\u0178\u017c\t\4\2\2\u0179\u017b\t\5\2\2\u017a\u0179"+
		"\3\2\2\2\u017b\u017e\3\2\2\2\u017c\u017a\3\2\2\2\u017c\u017d\3\2\2\2\u017d"+
		"P\3\2\2\2\u017e\u017c\3\2\2\2\u017f\u0180\7#\2\2\u0180\u0181\7?\2\2\u0181"+
		"R\3\2\2\2\u0182\u0183\7?\2\2\u0183\u0184\7?\2\2\u0184T\3\2\2\2\u0185\u0186"+
		"\7#\2\2\u0186V\3\2\2\2\u0187\u0188\7~\2\2\u0188\u0189\7~\2\2\u0189X\3"+
		"\2\2\2\u018a\u018b\7(\2\2\u018b\u018c\7(\2\2\u018cZ\3\2\2\2\u018d\u018e"+
		"\7?\2\2\u018e\\\3\2\2\2\u018f\u0190\7@\2\2\u0190\u0191\7?\2\2\u0191^\3"+
		"\2\2\2\u0192\u0193\7>\2\2\u0193\u0194\7?\2\2\u0194`\3\2\2\2\u0195\u0196"+
		"\7@\2\2\u0196b\3\2\2\2\u0197\u0198\7>\2\2\u0198d\3\2\2\2\u0199\u019a\7"+
		"\'\2\2\u019af\3\2\2\2\u019b\u019c\7,\2\2\u019ch\3\2\2\2\u019d\u019e\7"+
		"\61\2\2\u019ej\3\2\2\2\u019f\u01a0\7-\2\2\u01a0l\3\2\2\2\u01a1\u01a2\7"+
		"/\2\2\u01a2n\3\2\2\2\u01a3\u01a4\7*\2\2\u01a4p\3\2\2\2\u01a5\u01a6\7+"+
		"\2\2\u01a6r\3\2\2\2\u01a7\u01a8\7}\2\2\u01a8t\3\2\2\2\u01a9\u01aa\7\177"+
		"\2\2\u01aav\3\2\2\2\u01ab\u01ac\7]\2\2\u01acx\3\2\2\2\u01ad\u01ae\7_\2"+
		"\2\u01aez\3\2\2\2\u01af\u01b0\7A\2\2\u01b0|\3\2\2\2\u01b1\u01b2\7/\2\2"+
		"\u01b2\u01b3\7@\2\2\u01b3~\3\2\2\2\u01b4\u01b5\7.\2\2\u01b5\u0080\3\2"+
		"\2\2\u01b6\u01b7\7\60\2\2\u01b7\u0082\3\2\2\2\u01b8\u01b9\7<\2\2\u01b9"+
		"\u0084\3\2\2\2\u01ba\u01bb\7(\2\2\u01bb\u0086\3\2\2\2\u01bc\u01be\t\6"+
		"\2\2\u01bd\u01bc\3\2\2\2\u01be\u01bf\3\2\2\2\u01bf\u01bd\3\2\2\2\u01bf"+
		"\u01c0\3\2\2\2\u01c0\u01c1\3\2\2\2\u01c1\u01c2\bD\2\2\u01c2\u0088\3\2"+
		"\2\2\u01c3\u01c4\7\61\2\2\u01c4\u01c5\7,\2\2\u01c5\u01c9\3\2\2\2\u01c6"+
		"\u01c8\13\2\2\2\u01c7\u01c6\3\2\2\2\u01c8\u01cb\3\2\2\2\u01c9\u01ca\3"+
		"\2\2\2\u01c9\u01c7\3\2\2\2\u01ca\u01cc\3\2\2\2\u01cb\u01c9\3\2\2\2\u01cc"+
		"\u01cd\7,\2\2\u01cd\u01ce\7\61\2\2\u01ce\u01cf\3\2\2\2\u01cf\u01d0\bE"+
		"\2\2\u01d0\u008a\3\2\2\2\u01d1\u01d2\7\61\2\2\u01d2\u01d3\7\61\2\2\u01d3"+
		"\u01d7\3\2\2\2\u01d4\u01d6\n\7\2\2\u01d5\u01d4\3\2\2\2\u01d6\u01d9\3\2"+
		"\2\2\u01d7\u01d5\3\2\2\2\u01d7\u01d8\3\2\2\2\u01d8\u01da\3\2\2\2\u01d9"+
		"\u01d7\3\2\2\2\u01da\u01db\bF\2\2\u01db\u008c\3\2\2\2\u01dc\u01dd\7^\2"+
		"\2\u01dd\u01de\t\b\2\2\u01de\u008e\3\2\2\2\13\2\u0161\u0167\u0169\u0173"+
		"\u017c\u01bf\u01c9\u01d7\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}