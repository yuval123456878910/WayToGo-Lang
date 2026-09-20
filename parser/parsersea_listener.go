// Code generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ParserSea

import "github.com/antlr4-go/antlr/v4"

// ParserSeaListener is a complete listener for a parse tree produced by ParserSeaParser.
type ParserSeaListener interface {
	antlr.ParseTreeListener

	// EnterProgAbilities is called when entering the progAbilities production.
	EnterProgAbilities(c *ProgAbilitiesContext)

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterFlags is called when entering the flags production.
	EnterFlags(c *FlagsContext)

	// EnterTypes_of_tokens is called when entering the types_of_tokens production.
	EnterTypes_of_tokens(c *Types_of_tokensContext)

	// EnterTypesKeyword is called when entering the typesKeyword production.
	EnterTypesKeyword(c *TypesKeywordContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterReturnList is called when entering the returnList production.
	EnterReturnList(c *ReturnListContext)

	// EnterFuncDecl is called when entering the funcDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProgAbilities is called when exiting the progAbilities production.
	ExitProgAbilities(c *ProgAbilitiesContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitFlags is called when exiting the flags production.
	ExitFlags(c *FlagsContext)

	// ExitTypes_of_tokens is called when exiting the types_of_tokens production.
	ExitTypes_of_tokens(c *Types_of_tokensContext)

	// ExitTypesKeyword is called when exiting the typesKeyword production.
	ExitTypesKeyword(c *TypesKeywordContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitReturnList is called when exiting the returnList production.
	ExitReturnList(c *ReturnListContext)

	// ExitFuncDecl is called when exiting the funcDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
