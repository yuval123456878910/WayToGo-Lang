package compiler

import (
	"mj/parser"
)

func (l *CompilerWalk) ExitExpr(ctx *parser.ExprContext) {
	switch {
	case ctx.PLUS() != nil:
		Expretions := ctx.AllExpr()
		LeftType := l.ByteCode.ExprTypes[Expretions[0]]
		RightType := l.ByteCode.ExprTypes[Expretions[1]]

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
	}

	//return JVMbytecode{}
}
