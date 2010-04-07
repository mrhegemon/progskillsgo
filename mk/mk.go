/*
mk implements all of the mk.go interfaces to build files.
This relies on dag and dagimp to build the directed acyclic 
graph that governs the command order of a makefile.

Authors: William Broza, Tym Lipari
*/
package dag

import 	("os"
				"io/ioutil"
				"bytes"
				"strings"
)

//structure makeTarget
type makeTarget struct{
	Name string
	tgIm dagimp.TargImp
	//slice of strings
	dependencies []string
}

//functs for makeTarget

func(t *makeTarget) isDependent(depend string) bool {
	for _, y := range t.dependencies {
		if y == depend { return true }
	}
	return false
}

//MkTargetFact(s Set, str []string, t TargetFactory)
//target factory for mk implementation
//returns: error
func MkTargetFact(s Set, str []string, t TargetFactory) (Target, os.Error) {
		targ := new(TargImp)
		
		tokens := strings.Fields(str[0])
		
/*
		targ.name = tokens[0]
		targ.dependencies = make([]string,	20, 20)
		targ.dependlen = copy(targ.dependencies, tokens[1:])
		targ.dagset = s
		
		for y := 0; y < targ.dependlen; y++  {
			if s.Get(targ.dependencies[y]) == nil {
				temp := []string { targ.dependencies[y] }
				
				tempTarg, nerr := t(s, temp, t)
				if nerr != nil { return nil, nerr }
				
				_, nerr2 := s.Put(tempTarg)
				if nerr2 != nil { return tempTarg, nerr2 }
			}
		}
		return targ, nil
*/
}
