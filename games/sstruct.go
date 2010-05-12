/*
 String structure wrapper. Used to send strings across a netchan.
*/

package sstruct

type StringStruct struct {
	S string
}

//Makes a new StringStruct
func Make(s string) StringStruct {
	var temp StringStruct
	temp.S = s
	return temp
}
