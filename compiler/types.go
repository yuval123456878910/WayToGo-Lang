package compiler

import (
	"encoding/binary"

	"mj/parser"
)

func SmartAddInt(number int64) []byte {
	switch number {
	case 0:
		return []byte{ICONST_0}
	case 1:
		return []byte{ICONST_1}
	case 2:
		return []byte{ICONST_2}
	case 3:
		return []byte{ICONST_3}
	case 4:
		return []byte{ICONST_4}
	case 5:
		return []byte{ICONST_5}
	}
	if number <= 127 && number > -128 {
		return []byte{BIPUSH, byte(number)}
	}
}

func (l *CompilerWalk) EnterTypes_of_tokens(ctx *parser.Types_of_tokensContext) {
	switch {
	case ctx.NUM() != nil:
		l.Add()
	}
}

func TurnTypeNameToDescriptor_type(typeGiven parser.ITypesKeywordContext) []byte {
	switch {
	case typeGiven.INT_TYPE() != nil:
		return []byte{'I'}
	case typeGiven.STRING_TYPE() != nil:
		descriptor := "Ljava/lang/String;"
		bytes := []byte(descriptor)
		return bytes
	}
	panic("Uknowen type is refrenced")
}

func (j *JVMbytecode) MakeConstantUtf8(text string) {
	length := uint16(len(text))
	bytesReturn := make([]byte, 1+2+len(text))
	bytesReturn[0] = ConstantUtf8
	binary.BigEndian.PutUint16(bytesReturn[1:3], length)
	copy(bytesReturn[3:], []byte(text))
	j.CurrentLocPool++
	j.ContantPoolData[text] = j.CurrentLocPool
	j.ContantPool = append(j.ContantPool, bytesReturn...)
}
