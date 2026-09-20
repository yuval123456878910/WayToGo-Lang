// Generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class ParserSeaParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, RETURN=7, DEFINE=8, INT_TYPE=9, 
		STRING_TYPE=10, PUBLIC=11, PRIVATE=12, PLUS=13, MUL=14, ID=15, NUM=16, 
		COMMENT=17, STRING=18, WS=19;
	public static final int
		RULE_progAbilities = 0, RULE_prog = 1, RULE_decl = 2, RULE_block = 3, 
		RULE_assign = 4, RULE_flags = 5, RULE_types_of_tokens = 6, RULE_typesKeyword = 7, 
		RULE_param = 8, RULE_paramList = 9, RULE_returnList = 10, RULE_funcDecl = 11, 
		RULE_expr = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"progAbilities", "prog", "decl", "block", "assign", "flags", "types_of_tokens", 
			"typesKeyword", "param", "paramList", "returnList", "funcDecl", "expr"
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

	@Override
	public String getGrammarFileName() { return "ParserSea.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ParserSeaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgAbilitiesContext extends ParserRuleContext {
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public FuncDeclContext funcDecl() {
			return getRuleContext(FuncDeclContext.class,0);
		}
		public ProgAbilitiesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_progAbilities; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterProgAbilities(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitProgAbilities(this);
		}
	}

	public final ProgAbilitiesContext progAbilities() throws RecognitionException {
		ProgAbilitiesContext _localctx = new ProgAbilitiesContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_progAbilities);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(26);
				decl();
				}
				break;
			case 2:
				{
				setState(27);
				expr(0);
				}
				break;
			case 3:
				{
				setState(28);
				assign();
				}
				break;
			case 4:
				{
				setState(29);
				funcDecl();
				}
				break;
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
	public static class ProgContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(ParserSeaParser.EOF, 0); }
		public List<ProgAbilitiesContext> progAbilities() {
			return getRuleContexts(ProgAbilitiesContext.class);
		}
		public ProgAbilitiesContext progAbilities(int i) {
			return getRuleContext(ProgAbilitiesContext.class,i);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterProg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitProg(this);
		}
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(33); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(32);
				progAbilities();
				}
				}
				setState(35); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 368384L) != 0) );
			setState(37);
			match(EOF);
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
	public static class DeclContext extends ParserRuleContext {
		public TypesKeywordContext typesKeyword() {
			return getRuleContext(TypesKeywordContext.class,0);
		}
		public TerminalNode ID() { return getToken(ParserSeaParser.ID, 0); }
		public Types_of_tokensContext types_of_tokens() {
			return getRuleContext(Types_of_tokensContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterDecl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitDecl(this);
		}
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			typesKeyword();
			setState(40);
			match(ID);
			setState(41);
			match(T__0);
			setState(44);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(42);
				types_of_tokens();
				}
				break;
			case 2:
				{
				setState(43);
				expr(0);
				}
				break;
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
	public static class BlockContext extends ParserRuleContext {
		public List<ProgAbilitiesContext> progAbilities() {
			return getRuleContexts(ProgAbilitiesContext.class);
		}
		public ProgAbilitiesContext progAbilities(int i) {
			return getRuleContext(ProgAbilitiesContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(T__1);
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 368384L) != 0)) {
				{
				{
				setState(47);
				progAbilities();
				}
				}
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(53);
			match(T__2);
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
	public static class AssignContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ParserSeaParser.ID, 0); }
		public Types_of_tokensContext types_of_tokens() {
			return getRuleContext(Types_of_tokensContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterAssign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitAssign(this);
		}
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assign);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55);
			match(ID);
			setState(56);
			match(T__0);
			setState(59);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(57);
				types_of_tokens();
				}
				break;
			case 2:
				{
				setState(58);
				expr(0);
				}
				break;
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
	public static class FlagsContext extends ParserRuleContext {
		public TerminalNode PRIVATE() { return getToken(ParserSeaParser.PRIVATE, 0); }
		public TerminalNode PUBLIC() { return getToken(ParserSeaParser.PUBLIC, 0); }
		public FlagsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_flags; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterFlags(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitFlags(this);
		}
	}

	public final FlagsContext flags() throws RecognitionException {
		FlagsContext _localctx = new FlagsContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_flags);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			_la = _input.LA(1);
			if ( !(_la==PUBLIC || _la==PRIVATE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class Types_of_tokensContext extends ParserRuleContext {
		public TerminalNode NUM() { return getToken(ParserSeaParser.NUM, 0); }
		public TerminalNode STRING() { return getToken(ParserSeaParser.STRING, 0); }
		public Types_of_tokensContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types_of_tokens; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterTypes_of_tokens(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitTypes_of_tokens(this);
		}
	}

	public final Types_of_tokensContext types_of_tokens() throws RecognitionException {
		Types_of_tokensContext _localctx = new Types_of_tokensContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_types_of_tokens);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(63);
			_la = _input.LA(1);
			if ( !(_la==NUM || _la==STRING) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class TypesKeywordContext extends ParserRuleContext {
		public TerminalNode INT_TYPE() { return getToken(ParserSeaParser.INT_TYPE, 0); }
		public TerminalNode STRING_TYPE() { return getToken(ParserSeaParser.STRING_TYPE, 0); }
		public TypesKeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typesKeyword; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterTypesKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitTypesKeyword(this);
		}
	}

	public final TypesKeywordContext typesKeyword() throws RecognitionException {
		TypesKeywordContext _localctx = new TypesKeywordContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_typesKeyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			_la = _input.LA(1);
			if ( !(_la==INT_TYPE || _la==STRING_TYPE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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
	public static class ParamContext extends ParserRuleContext {
		public TypesKeywordContext typesKeyword() {
			return getRuleContext(TypesKeywordContext.class,0);
		}
		public TerminalNode ID() { return getToken(ParserSeaParser.ID, 0); }
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterParam(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitParam(this);
		}
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			typesKeyword();
			setState(68);
			match(ID);
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
	public static class ParamListContext extends ParserRuleContext {
		public List<ParamContext> param() {
			return getRuleContexts(ParamContext.class);
		}
		public ParamContext param(int i) {
			return getRuleContext(ParamContext.class,i);
		}
		public ParamListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterParamList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitParamList(this);
		}
	}

	public final ParamListContext paramList() throws RecognitionException {
		ParamListContext _localctx = new ParamListContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_paramList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			match(T__3);
			setState(79);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INT_TYPE || _la==STRING_TYPE) {
				{
				setState(71);
				param();
				setState(76);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__4) {
					{
					{
					setState(72);
					match(T__4);
					setState(73);
					param();
					}
					}
					setState(78);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(81);
			match(T__5);
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
	public static class ReturnListContext extends ParserRuleContext {
		public List<TypesKeywordContext> typesKeyword() {
			return getRuleContexts(TypesKeywordContext.class);
		}
		public TypesKeywordContext typesKeyword(int i) {
			return getRuleContext(TypesKeywordContext.class,i);
		}
		public ReturnListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterReturnList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitReturnList(this);
		}
	}

	public final ReturnListContext returnList() throws RecognitionException {
		ReturnListContext _localctx = new ReturnListContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_returnList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(83);
				match(T__3);
				setState(84);
				typesKeyword();
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__4) {
					{
					{
					setState(85);
					match(T__4);
					setState(86);
					typesKeyword();
					}
					}
					setState(91);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(92);
				match(T__5);
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
	public static class FuncDeclContext extends ParserRuleContext {
		public TerminalNode DEFINE() { return getToken(ParserSeaParser.DEFINE, 0); }
		public TerminalNode ID() { return getToken(ParserSeaParser.ID, 0); }
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public ReturnListContext returnList() {
			return getRuleContext(ReturnListContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public List<FlagsContext> flags() {
			return getRuleContexts(FlagsContext.class);
		}
		public FlagsContext flags(int i) {
			return getRuleContext(FlagsContext.class,i);
		}
		public FuncDeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDecl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterFuncDecl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitFuncDecl(this);
		}
	}

	public final FuncDeclContext funcDecl() throws RecognitionException {
		FuncDeclContext _localctx = new FuncDeclContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_funcDecl);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=1 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1+1 ) {
					{
					{
					setState(96);
					flags();
					}
					} 
				}
				setState(101);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			setState(102);
			match(DEFINE);
			setState(103);
			match(ID);
			setState(104);
			paramList();
			setState(105);
			returnList();
			setState(106);
			block();
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
		public TerminalNode ID() { return getToken(ParserSeaParser.ID, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public Types_of_tokensContext types_of_tokens() {
			return getRuleContext(Types_of_tokensContext.class,0);
		}
		public TerminalNode PLUS() { return getToken(ParserSeaParser.PLUS, 0); }
		public TerminalNode MUL() { return getToken(ParserSeaParser.MUL, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).enterExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ParserSeaListener ) ((ParserSeaListener)listener).exitExpr(this);
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
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(109);
				match(ID);
				setState(110);
				match(T__3);
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360448L) != 0)) {
					{
					setState(111);
					expr(0);
					setState(116);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__4) {
						{
						{
						setState(112);
						match(T__4);
						setState(113);
						expr(0);
						}
						}
						setState(118);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(121);
				match(T__5);
				}
				break;
			case 2:
				{
				setState(122);
				match(ID);
				}
				break;
			case 3:
				{
				setState(123);
				types_of_tokens();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(134);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(132);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(126);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(127);
						match(PLUS);
						setState(128);
						expr(5);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(129);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(130);
						match(MUL);
						setState(131);
						expr(4);
						}
						break;
					}
					} 
				}
				setState(136);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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
		case 12:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0013\u008a\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b"+
		"\u0002\f\u0007\f\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0003"+
		"\u0000\u001f\b\u0000\u0001\u0001\u0004\u0001\"\b\u0001\u000b\u0001\f\u0001"+
		"#\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002-\b\u0002\u0001\u0003\u0001\u0003\u0005\u0003"+
		"1\b\u0003\n\u0003\f\u00034\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004<\b\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0005\tK\b\t\n\t\f\tN\t\t\u0003"+
		"\tP\b\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0005\nX\b\n\n"+
		"\n\f\n[\t\n\u0001\n\u0001\n\u0003\n_\b\n\u0001\u000b\u0005\u000bb\b\u000b"+
		"\n\u000b\f\u000be\t\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0005\fs\b\f\n\f\f\fv\t\f\u0003\fx\b\f\u0001\f\u0001\f\u0001\f\u0003"+
		"\f}\b\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f\u0085"+
		"\b\f\n\f\f\f\u0088\t\f\u0001\f\u0001c\u0001\u0018\r\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u0000\u0003\u0001\u0000"+
		"\u000b\f\u0002\u0000\u0010\u0010\u0012\u0012\u0001\u0000\t\n\u008e\u0000"+
		"\u001e\u0001\u0000\u0000\u0000\u0002!\u0001\u0000\u0000\u0000\u0004\'"+
		"\u0001\u0000\u0000\u0000\u0006.\u0001\u0000\u0000\u0000\b7\u0001\u0000"+
		"\u0000\u0000\n=\u0001\u0000\u0000\u0000\f?\u0001\u0000\u0000\u0000\u000e"+
		"A\u0001\u0000\u0000\u0000\u0010C\u0001\u0000\u0000\u0000\u0012F\u0001"+
		"\u0000\u0000\u0000\u0014^\u0001\u0000\u0000\u0000\u0016c\u0001\u0000\u0000"+
		"\u0000\u0018|\u0001\u0000\u0000\u0000\u001a\u001f\u0003\u0004\u0002\u0000"+
		"\u001b\u001f\u0003\u0018\f\u0000\u001c\u001f\u0003\b\u0004\u0000\u001d"+
		"\u001f\u0003\u0016\u000b\u0000\u001e\u001a\u0001\u0000\u0000\u0000\u001e"+
		"\u001b\u0001\u0000\u0000\u0000\u001e\u001c\u0001\u0000\u0000\u0000\u001e"+
		"\u001d\u0001\u0000\u0000\u0000\u001f\u0001\u0001\u0000\u0000\u0000 \""+
		"\u0003\u0000\u0000\u0000! \u0001\u0000\u0000\u0000\"#\u0001\u0000\u0000"+
		"\u0000#!\u0001\u0000\u0000\u0000#$\u0001\u0000\u0000\u0000$%\u0001\u0000"+
		"\u0000\u0000%&\u0005\u0000\u0000\u0001&\u0003\u0001\u0000\u0000\u0000"+
		"\'(\u0003\u000e\u0007\u0000()\u0005\u000f\u0000\u0000),\u0005\u0001\u0000"+
		"\u0000*-\u0003\f\u0006\u0000+-\u0003\u0018\f\u0000,*\u0001\u0000\u0000"+
		"\u0000,+\u0001\u0000\u0000\u0000-\u0005\u0001\u0000\u0000\u0000.2\u0005"+
		"\u0002\u0000\u0000/1\u0003\u0000\u0000\u00000/\u0001\u0000\u0000\u0000"+
		"14\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u000023\u0001\u0000\u0000"+
		"\u000035\u0001\u0000\u0000\u000042\u0001\u0000\u0000\u000056\u0005\u0003"+
		"\u0000\u00006\u0007\u0001\u0000\u0000\u000078\u0005\u000f\u0000\u0000"+
		"8;\u0005\u0001\u0000\u00009<\u0003\f\u0006\u0000:<\u0003\u0018\f\u0000"+
		";9\u0001\u0000\u0000\u0000;:\u0001\u0000\u0000\u0000<\t\u0001\u0000\u0000"+
		"\u0000=>\u0007\u0000\u0000\u0000>\u000b\u0001\u0000\u0000\u0000?@\u0007"+
		"\u0001\u0000\u0000@\r\u0001\u0000\u0000\u0000AB\u0007\u0002\u0000\u0000"+
		"B\u000f\u0001\u0000\u0000\u0000CD\u0003\u000e\u0007\u0000DE\u0005\u000f"+
		"\u0000\u0000E\u0011\u0001\u0000\u0000\u0000FO\u0005\u0004\u0000\u0000"+
		"GL\u0003\u0010\b\u0000HI\u0005\u0005\u0000\u0000IK\u0003\u0010\b\u0000"+
		"JH\u0001\u0000\u0000\u0000KN\u0001\u0000\u0000\u0000LJ\u0001\u0000\u0000"+
		"\u0000LM\u0001\u0000\u0000\u0000MP\u0001\u0000\u0000\u0000NL\u0001\u0000"+
		"\u0000\u0000OG\u0001\u0000\u0000\u0000OP\u0001\u0000\u0000\u0000PQ\u0001"+
		"\u0000\u0000\u0000QR\u0005\u0006\u0000\u0000R\u0013\u0001\u0000\u0000"+
		"\u0000ST\u0005\u0004\u0000\u0000TY\u0003\u000e\u0007\u0000UV\u0005\u0005"+
		"\u0000\u0000VX\u0003\u000e\u0007\u0000WU\u0001\u0000\u0000\u0000X[\u0001"+
		"\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000"+
		"Z\\\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000\\]\u0005\u0006\u0000"+
		"\u0000]_\u0001\u0000\u0000\u0000^S\u0001\u0000\u0000\u0000^_\u0001\u0000"+
		"\u0000\u0000_\u0015\u0001\u0000\u0000\u0000`b\u0003\n\u0005\u0000a`\u0001"+
		"\u0000\u0000\u0000be\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000"+
		"ca\u0001\u0000\u0000\u0000df\u0001\u0000\u0000\u0000ec\u0001\u0000\u0000"+
		"\u0000fg\u0005\b\u0000\u0000gh\u0005\u000f\u0000\u0000hi\u0003\u0012\t"+
		"\u0000ij\u0003\u0014\n\u0000jk\u0003\u0006\u0003\u0000k\u0017\u0001\u0000"+
		"\u0000\u0000lm\u0006\f\uffff\uffff\u0000mn\u0005\u000f\u0000\u0000nw\u0005"+
		"\u0004\u0000\u0000ot\u0003\u0018\f\u0000pq\u0005\u0005\u0000\u0000qs\u0003"+
		"\u0018\f\u0000rp\u0001\u0000\u0000\u0000sv\u0001\u0000\u0000\u0000tr\u0001"+
		"\u0000\u0000\u0000tu\u0001\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000"+
		"vt\u0001\u0000\u0000\u0000wo\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000"+
		"\u0000xy\u0001\u0000\u0000\u0000y}\u0005\u0006\u0000\u0000z}\u0005\u000f"+
		"\u0000\u0000{}\u0003\f\u0006\u0000|l\u0001\u0000\u0000\u0000|z\u0001\u0000"+
		"\u0000\u0000|{\u0001\u0000\u0000\u0000}\u0086\u0001\u0000\u0000\u0000"+
		"~\u007f\n\u0004\u0000\u0000\u007f\u0080\u0005\r\u0000\u0000\u0080\u0085"+
		"\u0003\u0018\f\u0005\u0081\u0082\n\u0003\u0000\u0000\u0082\u0083\u0005"+
		"\u000e\u0000\u0000\u0083\u0085\u0003\u0018\f\u0004\u0084~\u0001\u0000"+
		"\u0000\u0000\u0084\u0081\u0001\u0000\u0000\u0000\u0085\u0088\u0001\u0000"+
		"\u0000\u0000\u0086\u0084\u0001\u0000\u0000\u0000\u0086\u0087\u0001\u0000"+
		"\u0000\u0000\u0087\u0019\u0001\u0000\u0000\u0000\u0088\u0086\u0001\u0000"+
		"\u0000\u0000\u000f\u001e#,2;LOY^ctw|\u0084\u0086";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}