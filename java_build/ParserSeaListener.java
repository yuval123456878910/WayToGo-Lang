// Generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ParserSeaParser}.
 */
public interface ParserSeaListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#progAbilities}.
	 * @param ctx the parse tree
	 */
	void enterProgAbilities(ParserSeaParser.ProgAbilitiesContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#progAbilities}.
	 * @param ctx the parse tree
	 */
	void exitProgAbilities(ParserSeaParser.ProgAbilitiesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(ParserSeaParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(ParserSeaParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#decl}.
	 * @param ctx the parse tree
	 */
	void enterDecl(ParserSeaParser.DeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#decl}.
	 * @param ctx the parse tree
	 */
	void exitDecl(ParserSeaParser.DeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(ParserSeaParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(ParserSeaParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#assign}.
	 * @param ctx the parse tree
	 */
	void enterAssign(ParserSeaParser.AssignContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#assign}.
	 * @param ctx the parse tree
	 */
	void exitAssign(ParserSeaParser.AssignContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#flags}.
	 * @param ctx the parse tree
	 */
	void enterFlags(ParserSeaParser.FlagsContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#flags}.
	 * @param ctx the parse tree
	 */
	void exitFlags(ParserSeaParser.FlagsContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#types_of_tokens}.
	 * @param ctx the parse tree
	 */
	void enterTypes_of_tokens(ParserSeaParser.Types_of_tokensContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#types_of_tokens}.
	 * @param ctx the parse tree
	 */
	void exitTypes_of_tokens(ParserSeaParser.Types_of_tokensContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#typesKeyword}.
	 * @param ctx the parse tree
	 */
	void enterTypesKeyword(ParserSeaParser.TypesKeywordContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#typesKeyword}.
	 * @param ctx the parse tree
	 */
	void exitTypesKeyword(ParserSeaParser.TypesKeywordContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#param}.
	 * @param ctx the parse tree
	 */
	void enterParam(ParserSeaParser.ParamContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#param}.
	 * @param ctx the parse tree
	 */
	void exitParam(ParserSeaParser.ParamContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#paramList}.
	 * @param ctx the parse tree
	 */
	void enterParamList(ParserSeaParser.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#paramList}.
	 * @param ctx the parse tree
	 */
	void exitParamList(ParserSeaParser.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#returnList}.
	 * @param ctx the parse tree
	 */
	void enterReturnList(ParserSeaParser.ReturnListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#returnList}.
	 * @param ctx the parse tree
	 */
	void exitReturnList(ParserSeaParser.ReturnListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDecl(ParserSeaParser.FuncDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#funcDecl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDecl(ParserSeaParser.FuncDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link ParserSeaParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(ParserSeaParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link ParserSeaParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(ParserSeaParser.ExprContext ctx);
}