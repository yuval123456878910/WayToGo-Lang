package compiler

import (
	"mj/parser"
)

func (l *CompilerWalk) EnterDecl(ctx *parser.DeclContext) {
}

func SmartISTORE(n uint16) []byte {
	switch n {
	case 0:
		return []byte{ISTORE_0}
	case 1:
		return []byte{ISTORE_1}
	case 2:
		return []byte{ISTORE_2}
	case 3:
		return []byte{ISTORE_3}
	}
	return []byte{ISTORE, byte(n)}
}

func SmartASTORE(n uint16) []byte {
	switch n {
	case 0:
		return []byte{ASTORE_0}
	case 1:
		return []byte{ASTORE_1}
	case 2:
		return []byte{ASTORE_2}
	case 3:
		return []byte{ASTORE_3}
	}
	return []byte{ASTORE, byte(n)}
}

func (l *CompilerWalk) ExitDecl(ctx *parser.DeclContext) {
	Type := ctx.TypesKeyword().GetText()
	l.ByteCode.Slots[ctx.ID().GetText()] = Slot{Pos: l.ByteCode.NextSlot, Type: Type}
	if v := ctx.Types_of_tokens(); v != nil {
		if v.NUM() != nil && Type != "int" {
			panic("a token of not a value num cant be registered as an int")
		} else if v.STRING() != nil && Type != "string" {
			panic("a token of not a value string cant be registered as a string")
		}
	} else if v := ctx.Expr(); v != nil {
		typeGot := l.ByteCode.ExprTypes[v]
		if typeGot != Type {
			panic("Type mitchmach")
		}
	}
	switch Type {
	case "int":
		l.Add(SmartISTORE(l.ByteCode.NextSlot)...)
	case "string":
		l.Add(SmartASTORE(l.ByteCode.NextSlot)...)
	}
	l.ByteCode.NextSlot++
}
