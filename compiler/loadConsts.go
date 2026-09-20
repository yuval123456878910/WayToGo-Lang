package compiler

func (j *JVMbytecode) ReturnInstractionConst(Name string) (uint16, bool) {
	v, k := j.ContantPoolData[Name]
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
