// Code generated from ANTLR4_Code/ParserSea.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ParserSeaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ParserSeaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parsersealexerLexerInit() {
	staticData := &ParserSeaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'{'", "'}'", "'('", "','", "')'", "'return'", "'def'", "'int'",
		"'string'", "'pub'", "'pri'", "'+'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "RETURN", "DEFINE", "INT_TYPE", "STRING_TYPE",
		"PUBLIC", "PRIVATE", "PLUS", "MUL", "ID", "NUM", "COMMENT", "STRING",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "RETURN", "DEFINE",
		"INT_TYPE", "STRING_TYPE", "PUBLIC", "PRIVATE", "PLUS", "MUL", "ID",
		"NUM", "COMMENT", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 139, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 88,
		8, 14, 10, 14, 12, 14, 91, 9, 14, 1, 15, 1, 15, 3, 15, 95, 8, 15, 1, 15,
		1, 15, 5, 15, 99, 8, 15, 10, 15, 12, 15, 102, 9, 15, 3, 15, 104, 8, 15,
		1, 16, 1, 16, 5, 16, 108, 8, 16, 10, 16, 12, 16, 111, 9, 16, 1, 16, 1,
		16, 1, 17, 1, 17, 5, 17, 117, 8, 17, 10, 17, 12, 17, 120, 9, 17, 1, 17,
		1, 17, 1, 17, 5, 17, 125, 8, 17, 10, 17, 12, 17, 128, 9, 17, 1, 17, 3,
		17, 131, 8, 17, 1, 18, 4, 18, 134, 8, 18, 11, 18, 12, 18, 135, 1, 18, 1,
		18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18,
		37, 19, 1, 0, 8, 1, 0, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3, 0, 10, 10, 13, 13,
		34, 34, 3, 0, 10, 10, 13, 13, 39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 147,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39,
		1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 43, 1, 0, 0, 0, 7, 45, 1, 0, 0, 0, 9,
		47, 1, 0, 0, 0, 11, 49, 1, 0, 0, 0, 13, 51, 1, 0, 0, 0, 15, 58, 1, 0, 0,
		0, 17, 62, 1, 0, 0, 0, 19, 66, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 77,
		1, 0, 0, 0, 25, 81, 1, 0, 0, 0, 27, 83, 1, 0, 0, 0, 29, 85, 1, 0, 0, 0,
		31, 103, 1, 0, 0, 0, 33, 105, 1, 0, 0, 0, 35, 130, 1, 0, 0, 0, 37, 133,
		1, 0, 0, 0, 39, 40, 5, 61, 0, 0, 40, 2, 1, 0, 0, 0, 41, 42, 5, 123, 0,
		0, 42, 4, 1, 0, 0, 0, 43, 44, 5, 125, 0, 0, 44, 6, 1, 0, 0, 0, 45, 46,
		5, 40, 0, 0, 46, 8, 1, 0, 0, 0, 47, 48, 5, 44, 0, 0, 48, 10, 1, 0, 0, 0,
		49, 50, 5, 41, 0, 0, 50, 12, 1, 0, 0, 0, 51, 52, 5, 114, 0, 0, 52, 53,
		5, 101, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5, 117, 0, 0, 55, 56, 5, 114,
		0, 0, 56, 57, 5, 110, 0, 0, 57, 14, 1, 0, 0, 0, 58, 59, 5, 100, 0, 0, 59,
		60, 5, 101, 0, 0, 60, 61, 5, 102, 0, 0, 61, 16, 1, 0, 0, 0, 62, 63, 5,
		105, 0, 0, 63, 64, 5, 110, 0, 0, 64, 65, 5, 116, 0, 0, 65, 18, 1, 0, 0,
		0, 66, 67, 5, 115, 0, 0, 67, 68, 5, 116, 0, 0, 68, 69, 5, 114, 0, 0, 69,
		70, 5, 105, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5, 103, 0, 0, 72, 20, 1,
		0, 0, 0, 73, 74, 5, 112, 0, 0, 74, 75, 5, 117, 0, 0, 75, 76, 5, 98, 0,
		0, 76, 22, 1, 0, 0, 0, 77, 78, 5, 112, 0, 0, 78, 79, 5, 114, 0, 0, 79,
		80, 5, 105, 0, 0, 80, 24, 1, 0, 0, 0, 81, 82, 5, 43, 0, 0, 82, 26, 1, 0,
		0, 0, 83, 84, 5, 42, 0, 0, 84, 28, 1, 0, 0, 0, 85, 89, 7, 0, 0, 0, 86,
		88, 7, 1, 0, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0,
		0, 89, 90, 1, 0, 0, 0, 90, 30, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 104,
		5, 48, 0, 0, 93, 95, 5, 45, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0,
		0, 95, 96, 1, 0, 0, 0, 96, 100, 7, 2, 0, 0, 97, 99, 7, 3, 0, 0, 98, 97,
		1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0,
		0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 92, 1, 0, 0, 0, 103,
		94, 1, 0, 0, 0, 104, 32, 1, 0, 0, 0, 105, 109, 5, 35, 0, 0, 106, 108, 8,
		4, 0, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0,
		0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112,
		113, 6, 16, 0, 0, 113, 34, 1, 0, 0, 0, 114, 118, 5, 34, 0, 0, 115, 117,
		8, 5, 0, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0,
		121, 131, 5, 34, 0, 0, 122, 126, 5, 39, 0, 0, 123, 125, 8, 6, 0, 0, 124,
		123, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127,
		1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 131, 5, 39,
		0, 0, 130, 114, 1, 0, 0, 0, 130, 122, 1, 0, 0, 0, 131, 36, 1, 0, 0, 0,
		132, 134, 7, 7, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138,
		6, 18, 0, 0, 138, 38, 1, 0, 0, 0, 10, 0, 89, 94, 100, 103, 109, 118, 126,
		130, 135, 1, 6, 0, 0,
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

// ParserSeaLexerInit initializes any static state used to implement ParserSeaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewParserSeaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserSeaLexerInit() {
	staticData := &ParserSeaLexerLexerStaticData
	staticData.once.Do(parsersealexerLexerInit)
}

// NewParserSeaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewParserSeaLexer(input antlr.CharStream) *ParserSeaLexer {
	ParserSeaLexerInit()
	l := new(ParserSeaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ParserSeaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "ParserSea.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ParserSeaLexer tokens.
const (
	ParserSeaLexerT__0        = 1
	ParserSeaLexerT__1        = 2
	ParserSeaLexerT__2        = 3
	ParserSeaLexerT__3        = 4
	ParserSeaLexerT__4        = 5
	ParserSeaLexerT__5        = 6
	ParserSeaLexerRETURN      = 7
	ParserSeaLexerDEFINE      = 8
	ParserSeaLexerINT_TYPE    = 9
	ParserSeaLexerSTRING_TYPE = 10
	ParserSeaLexerPUBLIC      = 11
	ParserSeaLexerPRIVATE     = 12
	ParserSeaLexerPLUS        = 13
	ParserSeaLexerMUL         = 14
	ParserSeaLexerID          = 15
	ParserSeaLexerNUM         = 16
	ParserSeaLexerCOMMENT     = 17
	ParserSeaLexerSTRING      = 18
	ParserSeaLexerWS          = 19
)
