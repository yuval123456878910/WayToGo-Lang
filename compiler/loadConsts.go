package compiler

func (j *JVMbytecode) ReturnInstractionConst(Name string) (uint16, bool) {
	v, k := j.ContantPoolData[Name]
	return v, k
}

func (j *JVMbytecode) ReturnInstractionConstInt(num int) (uint16, bool) {
	v, k := j.ContantPoolData[num]
	return v, k
}

func (j *JVMbytecode) ReturnInstractionConstFloat(num float32) (uint16, bool) {
	v, k := j.ContantPoolData[num]
	return v, k
}

func (j *JVMbytecode) POSTorGETnameConst(Name string) uint16 {
	v, ok := j.ReturnInstractionConst(Name)
	if ok {
		return v
	}
	j.MakeConstantUtf8(Name)
	v2, _ := j.ReturnInstractionConst(Name)
	return v2
}

func (j *JVMbytecode) POSTorGETintConst(num int) uint16 {
	v, ok := j.ReturnInstractionConstInt(num)
	if ok {
		return v
	}
	j.MakeConstantInt(num)
	v2, _ := j.ReturnInstractionConstInt(num)
	return v2
}

func (j *JVMbytecode) POSTorGETfloatConst(num float32) uint16 {
	v, ok := j.ReturnInstractionConstFloat(num)
	if ok {
		return v
	}
	j.MakeConstantFloat(num)
	v2, _ := j.ReturnInstractionConstFloat(num)
	return v2
}
