package manupilationtypes

import "encoding/binary"

func Uint16ToUint8s(uint16Num uint16) []byte {
	resultType := make([]byte, 2)
	binary.BigEndian.PutUint16(resultType[:], uint16Num)
	return resultType
}
