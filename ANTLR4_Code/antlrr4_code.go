package antlr4code

import (
	"mj/parser"

	"github.com/antlr4-go/antlr/v4"
)

type ParserExecuter struct {
	Text   string
	lexer  *parser.ParserSeaLexer
	tokens *antlr.CommonTokenStream
	parsed *parser.ParserSeaParser
	tree   parser.IProgContext
}

func (p *ParserExecuter) InitLexer() {
	input := antlr.NewInputStream(p.Text)
	lexer := parser.NewParserSeaLexer(input)
	p.lexer = lexer
}

func (p *ParserExecuter) InitTokens() {
	p.tokens = antlr.NewCommonTokenStream(p.lexer, antlr.TokenDefaultChannel)
}

func (p *ParserExecuter) InitParser() {
	p.parsed = parser.NewParserSeaParser(p.tokens)
}

func (p *ParserExecuter) ExecuteTree() {
	p.tree = p.parsed.Prog()
}
func (p *ParserExecuter) AccessTree() (parser.IProgContext, *parser.ParserSeaParser) {
	return p.tree, p.parsed
}

func (p *ParserExecuter) ToString() string {
	t, pa := p.AccessTree()
	return t.ToStringTree(nil, pa)
}

func NewParserExecuter(text string) ParserExecuter {
	return ParserExecuter{Text: text}
}

func GenerateTheAST(text string) (parser.IProgContext, *parser.ParserSeaParser, ParserExecuter) {
	parser_executer := NewParserExecuter(text)
	parser_executer.InitLexer()
	parser_executer.InitTokens()
	parser_executer.InitParser()
	parser_executer.ExecuteTree()
	a, b := parser_executer.AccessTree()
	return a, b, parser_executer
}
