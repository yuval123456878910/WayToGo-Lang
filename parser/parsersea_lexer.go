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
		"'string'", "'float'", "'pub'", "'pri'", "'+'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "RETURN", "DEFINE", "INT_TYPE", "STRING_TYPE",
		"FLOAT_TYPE", "PUBLIC", "PRIVATE", "PLUS", "MUL", "ID", "FLOAT", "NUM",
		"COMMENT", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "RETURN", "DEFINE",
		"INT_TYPE", "STRING_TYPE", "FLOAT_TYPE", "PUBLIC", "PRIVATE", "PLUS",
		"MUL", "ID", "FLOAT", "NUM", "COMMENT", "STRING", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 21, 163, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 98, 8, 15,
		10, 15, 12, 15, 101, 9, 15, 1, 16, 3, 16, 104, 8, 16, 1, 16, 4, 16, 107,
		8, 16, 11, 16, 12, 16, 108, 1, 16, 1, 16, 4, 16, 113, 8, 16, 11, 16, 12,
		16, 114, 1, 17, 1, 17, 3, 17, 119, 8, 17, 1, 17, 1, 17, 5, 17, 123, 8,
		17, 10, 17, 12, 17, 126, 9, 17, 3, 17, 128, 8, 17, 1, 18, 1, 18, 5, 18,
		132, 8, 18, 10, 18, 12, 18, 135, 9, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5,
		19, 141, 8, 19, 10, 19, 12, 19, 144, 9, 19, 1, 19, 1, 19, 1, 19, 5, 19,
		149, 8, 19, 10, 19, 12, 19, 152, 9, 19, 1, 19, 3, 19, 155, 8, 19, 1, 20,
		4, 20, 158, 8, 20, 11, 20, 12, 20, 159, 1, 20, 1, 20, 0, 0, 21, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12,
		25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21,
		1, 0, 8, 1, 0, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 49,
		57, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13, 3, 0, 10, 10, 13, 13, 34, 34, 3,
		0, 10, 10, 13, 13, 39, 39, 3, 0, 9, 10, 13, 13, 32, 32, 174, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0,
		33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0,
		0, 41, 1, 0, 0, 0, 1, 43, 1, 0, 0, 0, 3, 45, 1, 0, 0, 0, 5, 47, 1, 0, 0,
		0, 7, 49, 1, 0, 0, 0, 9, 51, 1, 0, 0, 0, 11, 53, 1, 0, 0, 0, 13, 55, 1,
		0, 0, 0, 15, 62, 1, 0, 0, 0, 17, 66, 1, 0, 0, 0, 19, 70, 1, 0, 0, 0, 21,
		77, 1, 0, 0, 0, 23, 83, 1, 0, 0, 0, 25, 87, 1, 0, 0, 0, 27, 91, 1, 0, 0,
		0, 29, 93, 1, 0, 0, 0, 31, 95, 1, 0, 0, 0, 33, 103, 1, 0, 0, 0, 35, 127,
		1, 0, 0, 0, 37, 129, 1, 0, 0, 0, 39, 154, 1, 0, 0, 0, 41, 157, 1, 0, 0,
		0, 43, 44, 5, 61, 0, 0, 44, 2, 1, 0, 0, 0, 45, 46, 5, 123, 0, 0, 46, 4,
		1, 0, 0, 0, 47, 48, 5, 125, 0, 0, 48, 6, 1, 0, 0, 0, 49, 50, 5, 40, 0,
		0, 50, 8, 1, 0, 0, 0, 51, 52, 5, 44, 0, 0, 52, 10, 1, 0, 0, 0, 53, 54,
		5, 41, 0, 0, 54, 12, 1, 0, 0, 0, 55, 56, 5, 114, 0, 0, 56, 57, 5, 101,
		0, 0, 57, 58, 5, 116, 0, 0, 58, 59, 5, 117, 0, 0, 59, 60, 5, 114, 0, 0,
		60, 61, 5, 110, 0, 0, 61, 14, 1, 0, 0, 0, 62, 63, 5, 100, 0, 0, 63, 64,
		5, 101, 0, 0, 64, 65, 5, 102, 0, 0, 65, 16, 1, 0, 0, 0, 66, 67, 5, 105,
		0, 0, 67, 68, 5, 110, 0, 0, 68, 69, 5, 116, 0, 0, 69, 18, 1, 0, 0, 0, 70,
		71, 5, 115, 0, 0, 71, 72, 5, 116, 0, 0, 72, 73, 5, 114, 0, 0, 73, 74, 5,
		105, 0, 0, 74, 75, 5, 110, 0, 0, 75, 76, 5, 103, 0, 0, 76, 20, 1, 0, 0,
		0, 77, 78, 5, 102, 0, 0, 78, 79, 5, 108, 0, 0, 79, 80, 5, 111, 0, 0, 80,
		81, 5, 97, 0, 0, 81, 82, 5, 116, 0, 0, 82, 22, 1, 0, 0, 0, 83, 84, 5, 112,
		0, 0, 84, 85, 5, 117, 0, 0, 85, 86, 5, 98, 0, 0, 86, 24, 1, 0, 0, 0, 87,
		88, 5, 112, 0, 0, 88, 89, 5, 114, 0, 0, 89, 90, 5, 105, 0, 0, 90, 26, 1,
		0, 0, 0, 91, 92, 5, 43, 0, 0, 92, 28, 1, 0, 0, 0, 93, 94, 5, 42, 0, 0,
		94, 30, 1, 0, 0, 0, 95, 99, 7, 0, 0, 0, 96, 98, 7, 1, 0, 0, 97, 96, 1,
		0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0,
		100, 32, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 104, 5, 45, 0, 0, 103, 102,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 107, 7, 2,
		0, 0, 106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0,
		108, 109, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 5, 46, 0, 0, 111,
		113, 7, 2, 0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 34, 1, 0, 0, 0, 116, 128, 5, 48,
		0, 0, 117, 119, 5, 45, 0, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0,
		119, 120, 1, 0, 0, 0, 120, 124, 7, 2, 0, 0, 121, 123, 7, 3, 0, 0, 122,
		121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125,
		1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 116, 1, 0,
		0, 0, 127, 118, 1, 0, 0, 0, 128, 36, 1, 0, 0, 0, 129, 133, 5, 35, 0, 0,
		130, 132, 8, 4, 0, 0, 131, 130, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133,
		131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135, 133,
		1, 0, 0, 0, 136, 137, 6, 18, 0, 0, 137, 38, 1, 0, 0, 0, 138, 142, 5, 34,
		0, 0, 139, 141, 8, 5, 0, 0, 140, 139, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0,
		142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 145, 155, 5, 34, 0, 0, 146, 150, 5, 39, 0, 0, 147, 149,
		8, 6, 0, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0,
		0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0,
		153, 155, 5, 39, 0, 0, 154, 138, 1, 0, 0, 0, 154, 146, 1, 0, 0, 0, 155,
		40, 1, 0, 0, 0, 156, 158, 7, 7, 0, 0, 157, 156, 1, 0, 0, 0, 158, 159, 1,
		0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0,
		0, 161, 162, 6, 20, 0, 0, 162, 42, 1, 0, 0, 0, 13, 0, 99, 103, 108, 114,
		118, 124, 127, 133, 142, 150, 154, 159, 1, 6, 0, 0,
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
	ParserSeaLexerFLOAT_TYPE  = 11
	ParserSeaLexerPUBLIC      = 12
	ParserSeaLexerPRIVATE     = 13
	ParserSeaLexerPLUS        = 14
	ParserSeaLexerMUL         = 15
	ParserSeaLexerID          = 16
	ParserSeaLexerFLOAT       = 17
	ParserSeaLexerNUM         = 18
	ParserSeaLexerCOMMENT     = 19
	ParserSeaLexerSTRING      = 20
	ParserSeaLexerWS          = 21
)
