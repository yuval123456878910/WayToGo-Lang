// Code generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ParserSea

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ParserSeaParser struct {
	*antlr.BaseParser
}

var ParserSeaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parserseaParserInit() {
	staticData := &ParserSeaParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'{'", "'}'", "'('", "','", "')'", "'return'", "'def'", "'int'",
		"'string'", "'float'", "'pub'", "'pri'", "'+'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "RETURN", "DEFINE", "INT_TYPE", "STRING_TYPE",
		"FLOAT_TYPE", "PUBLIC", "PRIVATE", "PLUS", "MUL", "ID", "FLOAT", "NUM",
		"COMMENT", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"progAbilities", "prog", "decl", "block", "assign", "flags", "types_of_tokens",
		"typesKeyword", "param", "paramList", "returnList", "funcDecl", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 138, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 31, 8, 0,
		1, 1, 4, 1, 34, 8, 1, 11, 1, 12, 1, 35, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 45, 8, 2, 1, 3, 1, 3, 5, 3, 49, 8, 3, 10, 3, 12, 3, 52,
		9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 60, 8, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 75,
		8, 9, 10, 9, 12, 9, 78, 9, 9, 3, 9, 80, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 5, 10, 88, 8, 10, 10, 10, 12, 10, 91, 9, 10, 1, 10, 1, 10,
		3, 10, 95, 8, 10, 1, 11, 5, 11, 98, 8, 11, 10, 11, 12, 11, 101, 9, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 5, 12, 115, 8, 12, 10, 12, 12, 12, 118, 9, 12, 3, 12, 120, 8,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 125, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 5, 12, 133, 8, 12, 10, 12, 12, 12, 136, 9, 12, 1, 12, 1,
		99, 1, 24, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 3, 1,
		0, 12, 13, 2, 0, 17, 18, 20, 20, 1, 0, 9, 11, 142, 0, 30, 1, 0, 0, 0, 2,
		33, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6, 46, 1, 0, 0, 0, 8, 55, 1, 0, 0, 0,
		10, 61, 1, 0, 0, 0, 12, 63, 1, 0, 0, 0, 14, 65, 1, 0, 0, 0, 16, 67, 1,
		0, 0, 0, 18, 70, 1, 0, 0, 0, 20, 94, 1, 0, 0, 0, 22, 99, 1, 0, 0, 0, 24,
		124, 1, 0, 0, 0, 26, 31, 3, 4, 2, 0, 27, 31, 3, 24, 12, 0, 28, 31, 3, 8,
		4, 0, 29, 31, 3, 22, 11, 0, 30, 26, 1, 0, 0, 0, 30, 27, 1, 0, 0, 0, 30,
		28, 1, 0, 0, 0, 30, 29, 1, 0, 0, 0, 31, 1, 1, 0, 0, 0, 32, 34, 3, 0, 0,
		0, 33, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 5, 0, 0, 1, 38, 3, 1, 0, 0, 0,
		39, 40, 3, 14, 7, 0, 40, 41, 5, 16, 0, 0, 41, 44, 5, 1, 0, 0, 42, 45, 3,
		12, 6, 0, 43, 45, 3, 24, 12, 0, 44, 42, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0,
		45, 5, 1, 0, 0, 0, 46, 50, 5, 2, 0, 0, 47, 49, 3, 0, 0, 0, 48, 47, 1, 0,
		0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 53,
		1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 3, 0, 0, 54, 7, 1, 0, 0, 0,
		55, 56, 5, 16, 0, 0, 56, 59, 5, 1, 0, 0, 57, 60, 3, 12, 6, 0, 58, 60, 3,
		24, 12, 0, 59, 57, 1, 0, 0, 0, 59, 58, 1, 0, 0, 0, 60, 9, 1, 0, 0, 0, 61,
		62, 7, 0, 0, 0, 62, 11, 1, 0, 0, 0, 63, 64, 7, 1, 0, 0, 64, 13, 1, 0, 0,
		0, 65, 66, 7, 2, 0, 0, 66, 15, 1, 0, 0, 0, 67, 68, 3, 14, 7, 0, 68, 69,
		5, 16, 0, 0, 69, 17, 1, 0, 0, 0, 70, 79, 5, 4, 0, 0, 71, 76, 3, 16, 8,
		0, 72, 73, 5, 5, 0, 0, 73, 75, 3, 16, 8, 0, 74, 72, 1, 0, 0, 0, 75, 78,
		1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0,
		78, 76, 1, 0, 0, 0, 79, 71, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 82, 5, 6, 0, 0, 82, 19, 1, 0, 0, 0, 83, 84, 5, 4, 0, 0, 84,
		89, 3, 14, 7, 0, 85, 86, 5, 5, 0, 0, 86, 88, 3, 14, 7, 0, 87, 85, 1, 0,
		0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 6, 0, 0, 93, 95, 1, 0, 0, 0,
		94, 83, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 21, 1, 0, 0, 0, 96, 98, 3,
		10, 5, 0, 97, 96, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0,
		99, 97, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 103,
		5, 8, 0, 0, 103, 104, 5, 16, 0, 0, 104, 105, 3, 18, 9, 0, 105, 106, 3,
		20, 10, 0, 106, 107, 3, 6, 3, 0, 107, 23, 1, 0, 0, 0, 108, 109, 6, 12,
		-1, 0, 109, 110, 5, 16, 0, 0, 110, 119, 5, 4, 0, 0, 111, 116, 3, 24, 12,
		0, 112, 113, 5, 5, 0, 0, 113, 115, 3, 24, 12, 0, 114, 112, 1, 0, 0, 0,
		115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 125, 5, 6, 0, 0, 122, 125, 5, 16,
		0, 0, 123, 125, 3, 12, 6, 0, 124, 108, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0,
		124, 123, 1, 0, 0, 0, 125, 134, 1, 0, 0, 0, 126, 127, 10, 4, 0, 0, 127,
		128, 5, 14, 0, 0, 128, 133, 3, 24, 12, 5, 129, 130, 10, 3, 0, 0, 130, 131,
		5, 15, 0, 0, 131, 133, 3, 24, 12, 4, 132, 126, 1, 0, 0, 0, 132, 129, 1,
		0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0,
		0, 135, 25, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 15, 30, 35, 44, 50, 59, 76,
		79, 89, 94, 99, 116, 119, 124, 132, 134,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ParserSeaParserInit initializes any static state used to implement ParserSeaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParserSeaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserSeaParserInit() {
	staticData := &ParserSeaParserStaticData
	staticData.once.Do(parserseaParserInit)
}

// NewParserSeaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParserSeaParser(input antlr.TokenStream) *ParserSeaParser {
	ParserSeaParserInit()
	this := new(ParserSeaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParserSeaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ParserSea.g4"

	return this
}

// ParserSeaParser tokens.
const (
	ParserSeaParserEOF         = antlr.TokenEOF
	ParserSeaParserT__0        = 1
	ParserSeaParserT__1        = 2
	ParserSeaParserT__2        = 3
	ParserSeaParserT__3        = 4
	ParserSeaParserT__4        = 5
	ParserSeaParserT__5        = 6
	ParserSeaParserRETURN      = 7
	ParserSeaParserDEFINE      = 8
	ParserSeaParserINT_TYPE    = 9
	ParserSeaParserSTRING_TYPE = 10
	ParserSeaParserFLOAT_TYPE  = 11
	ParserSeaParserPUBLIC      = 12
	ParserSeaParserPRIVATE     = 13
	ParserSeaParserPLUS        = 14
	ParserSeaParserMUL         = 15
	ParserSeaParserID          = 16
	ParserSeaParserFLOAT       = 17
	ParserSeaParserNUM         = 18
	ParserSeaParserCOMMENT     = 19
	ParserSeaParserSTRING      = 20
	ParserSeaParserWS          = 21
)

// ParserSeaParser rules.
const (
	ParserSeaParserRULE_progAbilities   = 0
	ParserSeaParserRULE_prog            = 1
	ParserSeaParserRULE_decl            = 2
	ParserSeaParserRULE_block           = 3
	ParserSeaParserRULE_assign          = 4
	ParserSeaParserRULE_flags           = 5
	ParserSeaParserRULE_types_of_tokens = 6
	ParserSeaParserRULE_typesKeyword    = 7
	ParserSeaParserRULE_param           = 8
	ParserSeaParserRULE_paramList       = 9
	ParserSeaParserRULE_returnList      = 10
	ParserSeaParserRULE_funcDecl        = 11
	ParserSeaParserRULE_expr            = 12
)

// IProgAbilitiesContext is an interface to support dynamic dispatch.
type IProgAbilitiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decl() IDeclContext
	Expr() IExprContext
	Assign() IAssignContext
	FuncDecl() IFuncDeclContext

	// IsProgAbilitiesContext differentiates from other interfaces.
	IsProgAbilitiesContext()
}

type ProgAbilitiesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgAbilitiesContext() *ProgAbilitiesContext {
	var p = new(ProgAbilitiesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_progAbilities
	return p
}

func InitEmptyProgAbilitiesContext(p *ProgAbilitiesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_progAbilities
}

func (*ProgAbilitiesContext) IsProgAbilitiesContext() {}

func NewProgAbilitiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgAbilitiesContext {
	var p = new(ProgAbilitiesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_progAbilities

	return p
}

func (s *ProgAbilitiesContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgAbilitiesContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ProgAbilitiesContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgAbilitiesContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ProgAbilitiesContext) FuncDecl() IFuncDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *ProgAbilitiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgAbilitiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgAbilitiesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterProgAbilities(s)
	}
}

func (s *ProgAbilitiesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitProgAbilities(s)
	}
}

func (p *ParserSeaParser) ProgAbilities() (localctx IProgAbilitiesContext) {
	localctx = NewProgAbilitiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserSeaParserRULE_progAbilities)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(26)
			p.Decl()
		}

	case 2:
		{
			p.SetState(27)
			p.expr(0)
		}

	case 3:
		{
			p.SetState(28)
			p.Assign()
		}

	case 4:
		{
			p.SetState(29)
			p.FuncDecl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllProgAbilities() []IProgAbilitiesContext
	ProgAbilities(i int) IProgAbilitiesContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserEOF, 0)
}

func (s *ProgContext) AllProgAbilities() []IProgAbilitiesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgAbilitiesContext); ok {
			len++
		}
	}

	tst := make([]IProgAbilitiesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgAbilitiesContext); ok {
			tst[i] = t.(IProgAbilitiesContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) ProgAbilities(i int) IProgAbilitiesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgAbilitiesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgAbilitiesContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *ParserSeaParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParserSeaParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1523456) != 0) {
		{
			p.SetState(32)
			p.ProgAbilities()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(ParserSeaParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypesKeyword() ITypesKeywordContext
	ID() antlr.TerminalNode
	Types_of_tokens() ITypes_of_tokensContext
	Expr() IExprContext

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_decl
	return p
}

func InitEmptyDeclContext(p *DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_decl
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) TypesKeyword() ITypesKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesKeywordContext)
}

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserID, 0)
}

func (s *DeclContext) Types_of_tokens() ITypes_of_tokensContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypes_of_tokensContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypes_of_tokensContext)
}

func (s *DeclContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *ParserSeaParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserSeaParserRULE_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.TypesKeyword()
	}
	{
		p.SetState(40)
		p.Match(ParserSeaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(ParserSeaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(42)
			p.Types_of_tokens()
		}

	case 2:
		{
			p.SetState(43)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllProgAbilities() []IProgAbilitiesContext
	ProgAbilities(i int) IProgAbilitiesContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllProgAbilities() []IProgAbilitiesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IProgAbilitiesContext); ok {
			len++
		}
	}

	tst := make([]IProgAbilitiesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IProgAbilitiesContext); ok {
			tst[i] = t.(IProgAbilitiesContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) ProgAbilities(i int) IProgAbilitiesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgAbilitiesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgAbilitiesContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ParserSeaParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserSeaParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(ParserSeaParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1523456) != 0 {
		{
			p.SetState(47)
			p.ProgAbilities()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(ParserSeaParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Types_of_tokens() ITypes_of_tokensContext
	Expr() IExprContext

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserID, 0)
}

func (s *AssignContext) Types_of_tokens() ITypes_of_tokensContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypes_of_tokensContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypes_of_tokensContext)
}

func (s *AssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *ParserSeaParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParserSeaParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(ParserSeaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(ParserSeaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(57)
			p.Types_of_tokens()
		}

	case 2:
		{
			p.SetState(58)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFlagsContext is an interface to support dynamic dispatch.
type IFlagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIVATE() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode

	// IsFlagsContext differentiates from other interfaces.
	IsFlagsContext()
}

type FlagsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlagsContext() *FlagsContext {
	var p = new(FlagsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_flags
	return p
}

func InitEmptyFlagsContext(p *FlagsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_flags
}

func (*FlagsContext) IsFlagsContext() {}

func NewFlagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlagsContext {
	var p = new(FlagsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_flags

	return p
}

func (s *FlagsContext) GetParser() antlr.Parser { return s.parser }

func (s *FlagsContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserPRIVATE, 0)
}

func (s *FlagsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserPUBLIC, 0)
}

func (s *FlagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterFlags(s)
	}
}

func (s *FlagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitFlags(s)
	}
}

func (p *ParserSeaParser) Flags() (localctx IFlagsContext) {
	localctx = NewFlagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParserSeaParserRULE_flags)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParserSeaParserPUBLIC || _la == ParserSeaParserPRIVATE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypes_of_tokensContext is an interface to support dynamic dispatch.
type ITypes_of_tokensContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUM() antlr.TerminalNode
	STRING() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsTypes_of_tokensContext differentiates from other interfaces.
	IsTypes_of_tokensContext()
}

type Types_of_tokensContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypes_of_tokensContext() *Types_of_tokensContext {
	var p = new(Types_of_tokensContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_types_of_tokens
	return p
}

func InitEmptyTypes_of_tokensContext(p *Types_of_tokensContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_types_of_tokens
}

func (*Types_of_tokensContext) IsTypes_of_tokensContext() {}

func NewTypes_of_tokensContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Types_of_tokensContext {
	var p = new(Types_of_tokensContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_types_of_tokens

	return p
}

func (s *Types_of_tokensContext) GetParser() antlr.Parser { return s.parser }

func (s *Types_of_tokensContext) NUM() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserNUM, 0)
}

func (s *Types_of_tokensContext) STRING() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserSTRING, 0)
}

func (s *Types_of_tokensContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserFLOAT, 0)
}

func (s *Types_of_tokensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Types_of_tokensContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Types_of_tokensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterTypes_of_tokens(s)
	}
}

func (s *Types_of_tokensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitTypes_of_tokens(s)
	}
}

func (p *ParserSeaParser) Types_of_tokens() (localctx ITypes_of_tokensContext) {
	localctx = NewTypes_of_tokensContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserSeaParserRULE_types_of_tokens)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1441792) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesKeywordContext is an interface to support dynamic dispatch.
type ITypesKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_TYPE() antlr.TerminalNode
	STRING_TYPE() antlr.TerminalNode
	FLOAT_TYPE() antlr.TerminalNode

	// IsTypesKeywordContext differentiates from other interfaces.
	IsTypesKeywordContext()
}

type TypesKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesKeywordContext() *TypesKeywordContext {
	var p = new(TypesKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_typesKeyword
	return p
}

func InitEmptyTypesKeywordContext(p *TypesKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_typesKeyword
}

func (*TypesKeywordContext) IsTypesKeywordContext() {}

func NewTypesKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesKeywordContext {
	var p = new(TypesKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_typesKeyword

	return p
}

func (s *TypesKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesKeywordContext) INT_TYPE() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserINT_TYPE, 0)
}

func (s *TypesKeywordContext) STRING_TYPE() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserSTRING_TYPE, 0)
}

func (s *TypesKeywordContext) FLOAT_TYPE() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserFLOAT_TYPE, 0)
}

func (s *TypesKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterTypesKeyword(s)
	}
}

func (s *TypesKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitTypesKeyword(s)
	}
}

func (p *ParserSeaParser) TypesKeyword() (localctx ITypesKeywordContext) {
	localctx = NewTypesKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserSeaParserRULE_typesKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3584) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TypesKeyword() ITypesKeywordContext
	ID() antlr.TerminalNode

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) TypesKeyword() ITypesKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesKeywordContext)
}

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserID, 0)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *ParserSeaParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserSeaParserRULE_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.TypesKeyword()
	}
	{
		p.SetState(68)
		p.Match(ParserSeaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *ParserSeaParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParserSeaParserRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(ParserSeaParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3584) != 0 {
		{
			p.SetState(71)
			p.Param()
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParserSeaParserT__4 {
			{
				p.SetState(72)
				p.Match(ParserSeaParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(73)
				p.Param()
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(81)
		p.Match(ParserSeaParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnListContext is an interface to support dynamic dispatch.
type IReturnListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTypesKeyword() []ITypesKeywordContext
	TypesKeyword(i int) ITypesKeywordContext

	// IsReturnListContext differentiates from other interfaces.
	IsReturnListContext()
}

type ReturnListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnListContext() *ReturnListContext {
	var p = new(ReturnListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_returnList
	return p
}

func InitEmptyReturnListContext(p *ReturnListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_returnList
}

func (*ReturnListContext) IsReturnListContext() {}

func NewReturnListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnListContext {
	var p = new(ReturnListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_returnList

	return p
}

func (s *ReturnListContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnListContext) AllTypesKeyword() []ITypesKeywordContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypesKeywordContext); ok {
			len++
		}
	}

	tst := make([]ITypesKeywordContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypesKeywordContext); ok {
			tst[i] = t.(ITypesKeywordContext)
			i++
		}
	}

	return tst
}

func (s *ReturnListContext) TypesKeyword(i int) ITypesKeywordContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesKeywordContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesKeywordContext)
}

func (s *ReturnListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterReturnList(s)
	}
}

func (s *ReturnListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitReturnList(s)
	}
}

func (p *ParserSeaParser) ReturnList() (localctx IReturnListContext) {
	localctx = NewReturnListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserSeaParserRULE_returnList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserSeaParserT__3 {
		{
			p.SetState(83)
			p.Match(ParserSeaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.TypesKeyword()
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ParserSeaParserT__4 {
			{
				p.SetState(85)
				p.Match(ParserSeaParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(86)
				p.TypesKeyword()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(92)
			p.Match(ParserSeaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFINE() antlr.TerminalNode
	ID() antlr.TerminalNode
	ParamList() IParamListContext
	ReturnList() IReturnListContext
	Block() IBlockContext
	AllFlags() []IFlagsContext
	Flags(i int) IFlagsContext

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserDEFINE, 0)
}

func (s *FuncDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserID, 0)
}

func (s *FuncDeclContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FuncDeclContext) ReturnList() IReturnListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnListContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) AllFlags() []IFlagsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFlagsContext); ok {
			len++
		}
	}

	tst := make([]IFlagsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFlagsContext); ok {
			tst[i] = t.(IFlagsContext)
			i++
		}
	}

	return tst
}

func (s *FuncDeclContext) Flags(i int) IFlagsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlagsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlagsContext)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterFuncDecl(s)
	}
}

func (s *FuncDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitFuncDecl(s)
	}
}

func (p *ParserSeaParser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParserSeaParserRULE_funcDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(96)
				p.Flags()
			}

		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(ParserSeaParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(ParserSeaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.ParamList()
	}
	{
		p.SetState(105)
		p.ReturnList()
	}
	{
		p.SetState(106)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Types_of_tokens() ITypes_of_tokensContext
	PLUS() antlr.TerminalNode
	MUL() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserSeaParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserSeaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserID, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) Types_of_tokens() ITypes_of_tokensContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypes_of_tokensContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypes_of_tokensContext)
}

func (s *ExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserPLUS, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(ParserSeaParserMUL, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserSeaListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *ParserSeaParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ParserSeaParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ParserSeaParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(109)
			p.Match(ParserSeaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(ParserSeaParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1507328) != 0 {
			{
				p.SetState(111)
				p.expr(0)
			}
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == ParserSeaParserT__4 {
				{
					p.SetState(112)
					p.Match(ParserSeaParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(113)
					p.expr(0)
				}

				p.SetState(118)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(121)
			p.Match(ParserSeaParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(122)
			p.Match(ParserSeaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(123)
			p.Types_of_tokens()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ParserSeaParserRULE_expr)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(127)
					p.Match(ParserSeaParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(128)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ParserSeaParserRULE_expr)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(130)
					p.Match(ParserSeaParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(131)
					p.expr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ParserSeaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParserSeaParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
