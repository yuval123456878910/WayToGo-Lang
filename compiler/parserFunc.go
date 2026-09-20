package compiler

import (
	"encoding/binary"
	"fmt"
	mty "mj/manupilationTypes"
	"mj/parser"
)

func GetTheParamsDescription(paramList parser.IParamListContext) []byte {
	returnData := []byte{}
	for _, token := range paramList.AllParam() {
		returnData = append(returnData, TurnTypeNameToDescriptor_type(token.TypesKeyword())...)
	}
	return returnData
}

func GetTheReturnDescription(ReturnList parser.IReturnListContext) []byte {
	returnData := []byte{}
	for _, token := range ReturnList.AllTypesKeyword() {
		returnData = append(returnData, TurnTypeNameToDescriptor_type(token)...)
	}
	return returnData
}

func (l *CompilerWalk) EnterFuncDecl(ctx *parser.FuncDeclContext) {
	fmt.Println("enter func decleration")
	for _, i := range ctx.ParamList().AllParam() {
		i.TypesKeyword()
	}
	FlagNames := GetFlagsFrom(ctx.AllFlags())
	FlagUint16s := FlagReturnValue(FlagNames...)
	FinalFlag := AccessFlagOR(FlagUint16s...)
	l.Add(FinalFlag...)
	loc := l.ByteCode.POSTorGETnameConst(ctx.ID().GetText())
	l.Add(mty.Uint16ToUint8s(loc)...)
	descriptor := "(" +
		string(GetTheParamsDescription(ctx.ParamList())) +
		")" +
		string(GetTheReturnDescription(ctx.ReturnList()))
	index := l.ByteCode.POSTorGETnameConst(descriptor)
	l.Add(mty.Uint16ToUint8s(index)...)
	l.Add(0, 0, 0, 0)
	l.ByteCode.FuncDeclEnds = append(
		l.ByteCode.FuncDeclEnds,
		uint32(len(l.ByteCode.Bytecode)-4),
	)
}

func (l *CompilerWalk) ExitFuncDecl(ctx *parser.FuncDeclContext) {
	FuncDeclEndLength := len(l.ByteCode.FuncDeclEnds)

	code_length_place_holder := l.ByteCode.FuncDeclEnds[FuncDeclEndLength-1]
	code_start := code_length_place_holder + 4 // code len buff
	code_end := uint32(len(l.ByteCode.Bytecode))

	l.ByteCode.FuncDeclEnds = l.ByteCode.FuncDeclEnds[:FuncDeclEndLength]

	sliceTemp := l.ByteCode.Bytecode[code_length_place_holder : code_length_place_holder+4]
	binary.BigEndian.PutUint32(sliceTemp[:], code_end-code_start)
}
