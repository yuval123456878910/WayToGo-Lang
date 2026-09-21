package compiler

import (
	"encoding/binary"

	"mj/parser"

	"github.com/antlr4-go/antlr/v4"
)

type CompilerWalk struct {
	*parser.BaseParserSeaListener
	ByteCode JVMbytecode
}

type Slot struct {
	Pos  uint16
	Type string
}

type JVMbytecode struct {
	MagicPool       []byte
	ContantPool     []byte
	ContantPoolData map[any]uint16
	CurrentLocPool  uint16
	Bytecode        []byte
	FuncDeclEnds    []uint32
	Slots           map[string]Slot
	NextSlot        uint16
	ExprTypes       map[antlr.ParserRuleContext]string
}

func (j JVMbytecode) ContantPoolFullWithCount() []byte {
	count := make([]byte, 2)
	binary.BigEndian.PutUint16(count[:2], uint16(len(j.ContantPoolData))+1)

	return append(count, j.ContantPool...)
}

func (j *JVMbytecode) PatchAll() []byte {
	return append(j.MagicPool, append(j.ContantPoolFullWithCount(), j.Bytecode...)...)
}

/*
func (s *parser.BaseParserSeaListener) EnterAssign(ctx *parser.AssignContext)
func (s *parser.BaseParserSeaListener) EnterBlock(ctx *parser.BlockContext)
func (l *CompilerWalk) EnterDecl(ctx *parser.ExprContext) JVMbytecode
func (s *parser.BaseParserSeaListener) EnterEveryRule(ctx antlr.ParserRuleContext)
func (s *parser.BaseParserSeaListener) EnterExpr(ctx *parser.ExprContext)
func (s *parser.BaseParserSeaListener) EnterFuncDecl(ctx *parser.FuncDeclContext)
func (s *parser.BaseParserSeaListener) EnterParam(ctx *parser.ParamContext)
func (s *parser.BaseParserSeaListener) EnterParamList(ctx *parser.ParamListContext)
func (s *parser.BaseParserSeaListener) EnterProg(ctx *parser.ProgContext)
func (s *parser.BaseParserSeaListener) EnterProgAbilities(ctx *parser.ProgAbilitiesContext)
func (s *parser.BaseParserSeaListener) EnterReturnList(ctx *parser.ReturnListContext)
func (s *parser.BaseParserSeaListener) EnterTypesKeyword(ctx *parser.TypesKeywordContext)
func (s *parser.BaseParserSeaListener) EnterTypes_of_tokens(ctx *parser.Types_of_tokensContext)
func (s *parser.BaseParserSeaListener) ExitAssign(ctx *parser.AssignContext)
func (s *parser.BaseParserSeaListener) ExitBlock(ctx *parser.BlockContext)
func (l *CompilerWalk) ExitDecl(ctx *parser.ExprContext) JVMbytecode
func (s *parser.BaseParserSeaListener) ExitEveryRule(ctx antlr.ParserRuleContext)
func (s *parser.BaseParserSeaListener) ExitExpr(ctx *parser.ExprContext)
func (s *parser.BaseParserSeaListener) ExitFuncDecl(ctx *parser.FuncDeclContext)
func (s *parser.BaseParserSeaListener) ExitParam(ctx *parser.ParamContext)
func (s *parser.BaseParserSeaListener) ExitParamList(ctx *parser.ParamListContext)
func (s *parser.BaseParserSeaListener) ExitProg(ctx *parser.ProgContext)
func (s *parser.BaseParserSeaListener) ExitProgAbilities(ctx *parser.ProgAbilitiesContext)
func (s *parser.BaseParserSeaListener) ExitReturnList(ctx *parser.ReturnListContext)
func (s *parser.BaseParserSeaListener) ExitTypesKeyword(ctx *parser.TypesKeywordContext)
func (s *parser.BaseParserSeaListener) ExitTypes_of_tokens(ctx *parser.Types_of_tokensContext)
func (s *parser.BaseParserSeaListener) VisitErrorNode(node antlr.ErrorNode)
func (s *parser.BaseParserSeaListener) VisitTerminal(node antlr.TerminalNode)
*/

func (l *CompilerWalk) Add(args ...byte) {
	l.ByteCode.Bytecode = append(l.ByteCode.Bytecode, args...)
}

func (l *CompilerWalk) EnterProg(ctx *parser.ProgContext) {
}

func (l *CompilerWalk) ExitProg(ctx *parser.ProgContext) {
}

func Compile(parsedCodeTree parser.IProgContext) JVMbytecode {
	listener := CompilerWalk{
		BaseParserSeaListener: &parser.BaseParserSeaListener{},
		ByteCode:              JVMbytecode{ContantPool: []byte{}, ContantPoolData: map[any]uint16{}, Bytecode: []byte{}, Slots: map[string]Slot{}, FuncDeclEnds: []uint32{}, ExprTypes: map[antlr.ParserRuleContext]string{}},
	}
	listener.ByteCode.MagicPool = append(MAGIC, append(MINOR_VERSION, MAJOR_VERSION...)...)

	antlr.ParseTreeWalkerDefault.Walk(&listener, parsedCodeTree)

	return listener.ByteCode
}
