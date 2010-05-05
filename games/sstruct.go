package sstruct

type StringStruct struct {
	S string
}

func Make(s string) StringStruct {
	var temp StringStruct
	temp.S = s
	return temp
}
