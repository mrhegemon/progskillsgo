/*
mk implements all of the mk.go interfaces to build files.
This relies on dag and dagimp to build the directed acyclic 
graph that governs the command order of a makefile.

Authors: William Broza, Tym Lipari
*/
package mk

import 	("os"
		 "dag"
		 "strings"
)

//structure makeTarget
type MakeTarget struct{
	dag.TargImp
	//slice of strings
	commands []string
}

func (t *MakeTarget) Commands() string {
	temp := ""
	for _, y := range t.commands {
		temp += y + "\n"
	}
	return temp
}

//MkTargetFact(s Set, str []string, t TargetFactory)
//target factory for mk implementation
//returns: error
func MkTargetFact(s Set, str []string, t TargetFactory) (Target, os.Error) {
		targ := new(MakeTarget)
		
		tokens := strings.Fields(str[0])
		

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
		
		targ.commands = str[1:]
		return targ, nil
}

func Act(targ dag.Target) os.Error {
	t := targ.(MakeTarget)
	age := os.Dir.Stat(t.Name()).Mtime_ns
	
	printCommands := false
	for n, y := range t.dependencies {
		preqAge := os.Dir.Stat(y).Mtime_ns
		
		if age < preqAge { printCommands = true; break }
	}
	
	if printCommands {
		for _, y := range t.commands {
			fmt.Println(y)
		}
	}
	
	return nil
}
			
