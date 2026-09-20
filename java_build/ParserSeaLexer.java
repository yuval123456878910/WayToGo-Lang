// Generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class ParserSeaLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, RETURN=7, DEFINE=8, INT_TYPE=9, 
		STRING_TYPE=10, PUBLIC=11, PRIVATE=12, PLUS=13, MUL=14, ID=15, NUM=16, 
		COMMENT=17, STRING=18, WS=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "RETURN", "DEFINE", "INT_TYPE", 
			"STRING_TYPE", "PUBLIC", "PRIVATE", "PLUS", "MUL", "ID", "NUM", "COMMENT", 
			"STRING", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'{'", "'}'", "'('", "','", "')'", "'return'", "'def'", 
			"'int'", "'string'", "'pub'", "'pri'", "'+'", "'*'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "RETURN", "DEFINE", "INT_TYPE", 
			"STRING_TYPE", "PUBLIC", "PRIVATE", "PLUS", "MUL", "ID", "NUM", "COMMENT", 
			"STRING", "WS"
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


	public ParserSeaLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ParserSea.g4"; }

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
		"\u0004\u0000\u0013\u008b\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001"+
		"\u000e\u0001\u000e\u0005\u000eX\b\u000e\n\u000e\f\u000e[\t\u000e\u0001"+
		"\u000f\u0001\u000f\u0003\u000f_\b\u000f\u0001\u000f\u0001\u000f\u0005"+
		"\u000fc\b\u000f\n\u000f\f\u000ff\t\u000f\u0003\u000fh\b\u000f\u0001\u0010"+
		"\u0001\u0010\u0005\u0010l\b\u0010\n\u0010\f\u0010o\t\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0011\u0001\u0011\u0005\u0011u\b\u0011\n\u0011\f\u0011"+
		"x\t\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0005\u0011}\b\u0011\n\u0011"+
		"\f\u0011\u0080\t\u0011\u0001\u0011\u0003\u0011\u0083\b\u0011\u0001\u0012"+
		"\u0004\u0012\u0086\b\u0012\u000b\u0012\f\u0012\u0087\u0001\u0012\u0001"+
		"\u0012\u0000\u0000\u0013\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004"+
		"\t\u0005\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017"+
		"\f\u0019\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011#\u0012%\u0013\u0001"+
		"\u0000\b\u0001\u0000az\u0004\u000009AZ__az\u0001\u000019\u0001\u00000"+
		"9\u0002\u0000\n\n\r\r\u0003\u0000\n\n\r\r\"\"\u0003\u0000\n\n\r\r\'\'"+
		"\u0003\u0000\t\n\r\r  \u0093\u0000\u0001\u0001\u0000\u0000\u0000\u0000"+
		"\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000"+
		"\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b"+
		"\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001"+
		"\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001"+
		"\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001"+
		"\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001"+
		"\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001"+
		"\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000"+
		"\u0000\u0000%\u0001\u0000\u0000\u0000\u0001\'\u0001\u0000\u0000\u0000"+
		"\u0003)\u0001\u0000\u0000\u0000\u0005+\u0001\u0000\u0000\u0000\u0007-"+
		"\u0001\u0000\u0000\u0000\t/\u0001\u0000\u0000\u0000\u000b1\u0001\u0000"+
		"\u0000\u0000\r3\u0001\u0000\u0000\u0000\u000f:\u0001\u0000\u0000\u0000"+
		"\u0011>\u0001\u0000\u0000\u0000\u0013B\u0001\u0000\u0000\u0000\u0015I"+
		"\u0001\u0000\u0000\u0000\u0017M\u0001\u0000\u0000\u0000\u0019Q\u0001\u0000"+
		"\u0000\u0000\u001bS\u0001\u0000\u0000\u0000\u001dU\u0001\u0000\u0000\u0000"+
		"\u001fg\u0001\u0000\u0000\u0000!i\u0001\u0000\u0000\u0000#\u0082\u0001"+
		"\u0000\u0000\u0000%\u0085\u0001\u0000\u0000\u0000\'(\u0005=\u0000\u0000"+
		"(\u0002\u0001\u0000\u0000\u0000)*\u0005{\u0000\u0000*\u0004\u0001\u0000"+
		"\u0000\u0000+,\u0005}\u0000\u0000,\u0006\u0001\u0000\u0000\u0000-.\u0005"+
		"(\u0000\u0000.\b\u0001\u0000\u0000\u0000/0\u0005,\u0000\u00000\n\u0001"+
		"\u0000\u0000\u000012\u0005)\u0000\u00002\f\u0001\u0000\u0000\u000034\u0005"+
		"r\u0000\u000045\u0005e\u0000\u000056\u0005t\u0000\u000067\u0005u\u0000"+
		"\u000078\u0005r\u0000\u000089\u0005n\u0000\u00009\u000e\u0001\u0000\u0000"+
		"\u0000:;\u0005d\u0000\u0000;<\u0005e\u0000\u0000<=\u0005f\u0000\u0000"+
		"=\u0010\u0001\u0000\u0000\u0000>?\u0005i\u0000\u0000?@\u0005n\u0000\u0000"+
		"@A\u0005t\u0000\u0000A\u0012\u0001\u0000\u0000\u0000BC\u0005s\u0000\u0000"+
		"CD\u0005t\u0000\u0000DE\u0005r\u0000\u0000EF\u0005i\u0000\u0000FG\u0005"+
		"n\u0000\u0000GH\u0005g\u0000\u0000H\u0014\u0001\u0000\u0000\u0000IJ\u0005"+
		"p\u0000\u0000JK\u0005u\u0000\u0000KL\u0005b\u0000\u0000L\u0016\u0001\u0000"+
		"\u0000\u0000MN\u0005p\u0000\u0000NO\u0005r\u0000\u0000OP\u0005i\u0000"+
		"\u0000P\u0018\u0001\u0000\u0000\u0000QR\u0005+\u0000\u0000R\u001a\u0001"+
		"\u0000\u0000\u0000ST\u0005*\u0000\u0000T\u001c\u0001\u0000\u0000\u0000"+
		"UY\u0007\u0000\u0000\u0000VX\u0007\u0001\u0000\u0000WV\u0001\u0000\u0000"+
		"\u0000X[\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000YZ\u0001\u0000"+
		"\u0000\u0000Z\u001e\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000"+
		"\\h\u00050\u0000\u0000]_\u0005-\u0000\u0000^]\u0001\u0000\u0000\u0000"+
		"^_\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`d\u0007\u0002\u0000"+
		"\u0000ac\u0007\u0003\u0000\u0000ba\u0001\u0000\u0000\u0000cf\u0001\u0000"+
		"\u0000\u0000db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000eh\u0001"+
		"\u0000\u0000\u0000fd\u0001\u0000\u0000\u0000g\\\u0001\u0000\u0000\u0000"+
		"g^\u0001\u0000\u0000\u0000h \u0001\u0000\u0000\u0000im\u0005#\u0000\u0000"+
		"jl\b\u0004\u0000\u0000kj\u0001\u0000\u0000\u0000lo\u0001\u0000\u0000\u0000"+
		"mk\u0001\u0000\u0000\u0000mn\u0001\u0000\u0000\u0000np\u0001\u0000\u0000"+
		"\u0000om\u0001\u0000\u0000\u0000pq\u0006\u0010\u0000\u0000q\"\u0001\u0000"+
		"\u0000\u0000rv\u0005\"\u0000\u0000su\b\u0005\u0000\u0000ts\u0001\u0000"+
		"\u0000\u0000ux\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000\u0000vw\u0001"+
		"\u0000\u0000\u0000wy\u0001\u0000\u0000\u0000xv\u0001\u0000\u0000\u0000"+
		"y\u0083\u0005\"\u0000\u0000z~\u0005\'\u0000\u0000{}\b\u0006\u0000\u0000"+
		"|{\u0001\u0000\u0000\u0000}\u0080\u0001\u0000\u0000\u0000~|\u0001\u0000"+
		"\u0000\u0000~\u007f\u0001\u0000\u0000\u0000\u007f\u0081\u0001\u0000\u0000"+
		"\u0000\u0080~\u0001\u0000\u0000\u0000\u0081\u0083\u0005\'\u0000\u0000"+
		"\u0082r\u0001\u0000\u0000\u0000\u0082z\u0001\u0000\u0000\u0000\u0083$"+
		"\u0001\u0000\u0000\u0000\u0084\u0086\u0007\u0007\u0000\u0000\u0085\u0084"+
		"\u0001\u0000\u0000\u0000\u0086\u0087\u0001\u0000\u0000\u0000\u0087\u0085"+
		"\u0001\u0000\u0000\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\u0089"+
		"\u0001\u0000\u0000\u0000\u0089\u008a\u0006\u0012\u0000\u0000\u008a&\u0001"+
		"\u0000\u0000\u0000\n\u0000Y^dgmv~\u0082\u0087\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}