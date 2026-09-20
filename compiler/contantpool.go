package compiler

const (
	ConstantUtf8    = byte(1)
	ConstantInteger = byte(2 + iota)
	ConstantFloat
	ConstantLong
	ConstantDouble
	ConstantClass
	ConstantString
	ConstantFieldref
	ConstantMethodref
	ConstantInterfaceMethodref
	ConstantNameAndType
)
