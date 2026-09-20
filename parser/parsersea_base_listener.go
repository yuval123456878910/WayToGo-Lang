// Code generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ParserSea

import "github.com/antlr4-go/antlr/v4"

// BaseParserSeaListener is a complete listener for a parse tree produced by ParserSeaParser.
type BaseParserSeaListener struct{}

var _ ParserSeaListener = &BaseParserSeaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParserSeaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParserSeaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParserSeaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParserSeaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgAbilities is called when production progAbilities is entered.
func (s *BaseParserSeaListener) EnterProgAbilities(ctx *ProgAbilitiesContext) {}

// ExitProgAbilities is called when production progAbilities is exited.
func (s *BaseParserSeaListener) ExitProgAbilities(ctx *ProgAbilitiesContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseParserSeaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseParserSeaListener) ExitProg(ctx *ProgContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseParserSeaListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseParserSeaListener) ExitDecl(ctx *DeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseParserSeaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseParserSeaListener) ExitBlock(ctx *BlockContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseParserSeaListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseParserSeaListener) ExitAssign(ctx *AssignContext) {}

// EnterFlags is called when production flags is entered.
func (s *BaseParserSeaListener) EnterFlags(ctx *FlagsContext) {}

// ExitFlags is called when production flags is exited.
func (s *BaseParserSeaListener) ExitFlags(ctx *FlagsContext) {}

// EnterTypes_of_tokens is called when production types_of_tokens is entered.
func (s *BaseParserSeaListener) EnterTypes_of_tokens(ctx *Types_of_tokensContext) {}

// ExitTypes_of_tokens is called when production types_of_tokens is exited.
func (s *BaseParserSeaListener) ExitTypes_of_tokens(ctx *Types_of_tokensContext) {}

// EnterTypesKeyword is called when production typesKeyword is entered.
func (s *BaseParserSeaListener) EnterTypesKeyword(ctx *TypesKeywordContext) {}

// ExitTypesKeyword is called when production typesKeyword is exited.
func (s *BaseParserSeaListener) ExitTypesKeyword(ctx *TypesKeywordContext) {}

// EnterParam is called when production param is entered.
func (s *BaseParserSeaListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseParserSeaListener) ExitParam(ctx *ParamContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseParserSeaListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseParserSeaListener) ExitParamList(ctx *ParamListContext) {}

// EnterReturnList is called when production returnList is entered.
func (s *BaseParserSeaListener) EnterReturnList(ctx *ReturnListContext) {}

// ExitReturnList is called when production returnList is exited.
func (s *BaseParserSeaListener) ExitReturnList(ctx *ReturnListContext) {}

// EnterFuncDecl is called when production funcDecl is entered.
func (s *BaseParserSeaListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production funcDecl is exited.
func (s *BaseParserSeaListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseParserSeaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseParserSeaListener) ExitExpr(ctx *ExprContext) {}
