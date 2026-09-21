package compiler

import (
	"fmt"
	"strconv"

	"mj/parser"
)

func (l *CompilerWalk) ExitExpr(ctx *parser.ExprContext) {
	switch {
	case ctx.PLUS() != nil:
		Expretions := ctx.AllExpr()
		LeftType := l.ByteCode.ExprTypes[Expretions[0]]
		RightType := l.ByteCode.ExprTypes[Expretions[1]]
		fmt.Println("eee", LeftType, RightType)
		if LeftType != RightType {
			panic("")
		}
		switch LeftType {
		case "int":
			l.Add(IADD)
		}
		l.ByteCode.ExprTypes[ctx] = LeftType
	case ctx.MUL() != nil:
		Expretions := ctx.AllExpr()
		LeftType := l.ByteCode.ExprTypes[Expretions[0]]
		RightType := l.ByteCode.ExprTypes[Expretions[1]]

		if LeftType != RightType {
			panic("")
		}
		switch LeftType {
		case "int":
			l.Add(LMUL)
		}
		l.ByteCode.ExprTypes[ctx] = LeftType
	case ctx.Types_of_tokens().NUM() != nil:
		numberText := ctx.Types_of_tokens().NUM().GetText()
		number, err := strconv.Atoi(numberText)
		if err != nil {
			panic("Text cant be converted to int")
		}

		l.ByteCode.SmartAddIntReturn(number)
		l.ByteCode.ExprTypes[ctx] = "int"

	}

	//return JVMbytecode{}
}
