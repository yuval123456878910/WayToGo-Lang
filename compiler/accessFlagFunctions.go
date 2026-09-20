package compiler

import (
	"encoding/binary"
	"mj/parser"
)

func AccessFlagOR(Flags ...uint16) []byte {
	var ORresult uint16 = 0
	for _, flag := range Flags {
		ORresult = ORresult | flag
	}
	ReturnByte := make([]byte, 2)
	binary.BigEndian.PutUint16(ReturnByte[:], ORresult)
	return ReturnByte
}

func FlagReturnValue(FlagsNames ...string) []uint16 {
	ReturnUint16 := make([]uint16, len(FlagsNames))
	for _, name := range FlagsNames {
		V := uint16(0)
		switch name {
		case "pri":
			V = ACC_PRIVATE
		case "pub":
			V = ACC_PUBLIC
		default:
			panic("Flag value doesnt exist!")
		}
		ReturnUint16 = append(ReturnUint16, V)
	}
	return ReturnUint16
}

func GetFlagsFrom(flags []parser.IFlagsContext) []string {
	returnString := []string{}
	for _, f := range flags {
		switch {
		case f.PRIVATE() != nil:
			returnString = append(returnString, "pri")
		case f.PUBLIC() != nil:
			returnString = append(returnString, "pub")
		}
	}
	return returnString
}
