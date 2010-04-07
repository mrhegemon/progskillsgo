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
		 "fmt"
		 //"strconv"
)

//structure makeTarget
type MakeTarget struct{
	//we kept getting conversion errors for missing things,
	//so we had to reimplement all of our methods... not sure why...
	dag.TargImp
	
	name string
	dependencies []string
	dagset dag.Set
	dependlen int
	commandSent bool
	cyclic bool
	commands []string
}

func (t *MakeTarget) Commands() string {
	temp := ""
	for _, y := range t.commands {
		if strings.Index(y, "\n") != 0 {
			temp += strings.TrimSpace(y) + "\n"
		}
	}
	return temp
}

//MkTargetFact(s Set, str []string, t TargetFactory)
//target factory for mk implementation
//returns: error
func MkTargetFact(s dag.Set, str []string, t dag.TargetFactory) (dag.Target, os.Error) {
		/*printer := func(s []string) {
			for x, y := range s {
				println("[" + strconv.Itoa(x) + "]  =  " + y)
			}
		}*/
		
		//printer(str)

		//println("TARGFACT 0:  " + strconv.Itoa(len(str[0])) + "   \"" + str[0] + "\"")
		targ := new(MakeTarget)
		
		tokens := strings.Fields(str[0])
		
		//printer(tokens)
		

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
		//println("TARG RETURNED:  \"" + targ.Name() + "\"")
		return targ, nil
}
//Merge(other Target)
//merges targets
//returns: error
func(t *MakeTarget) Merge(other dag.Target) (dag.Target, os.Error) {
	x := other.(*MakeTarget)
	if x.Name() != t.Name() {
		return nil, os.NewError("cannot merge targets with different names")
	}
	
	for y := 0; y < x.dependlen; y++ {
		if !t.isDependent(x.dependencies[y]) {
			t.dependencies[t.dependlen] = x.dependencies[y]
			t.dependlen++
		}
	}
	
	t.commands = other.(*MakeTarget).commands
	return t, nil
}
//isDependent(depend string)
//tells if dependant
//returns: bool of dependance
func(t *MakeTarget) isDependent(depend string) bool {
	for _, y := range t.dependencies {
		if y == depend { return true }
	}
	return false
}
func(t *MakeTarget) Name() string {
	return t.name
}
//Apply(a Action)
//applys action to TargImp
//returns: error
func(t *MakeTarget) Apply(a dag.Action) os.Error {

	if !t.commandSent {
		t.commandSent = true
		return a(t)
	}
	return nil
}

//ApplyPreq(a Action)
//Applys prereq to target
//returns: error
func(t *MakeTarget) ApplyPreq(a dag.Action) os.Error {
	for y := 0; y < t.dependlen; y++ {
		var targ dag.Target
		//if the target doesn't exist, send an error
		if targ = t.dagset.Get(t.dependencies[y]); targ == nil { 
			return os.NewError("non-existant Target:  " + t.dependencies[y]) 
		}
		
		if t.cyclic { return os.NewError("Target " + t.Name() + ":  cyclic") }

		t.cyclic = true
		//if the targets prereqs sent an error, send it on
		if err1 := targ.ApplyPreq(a); err1 != nil {
			return err1
		}
		t.cyclic = false
		
		//if the target sent an error, send it on
		if err2 := targ.Apply(a); err2 != nil {
			return err2
		}
	}
	return nil
}
func Act(targ dag.Target) os.Error {
	t := targ.(*MakeTarget)
	
	printCommands := false
	
	dir, err := os.Stat(targ.Name())
	//file doesn't exist..
	if err != nil { 
		printCommands = true 
	} else {
		reg := dir.IsRegular()
	
		//command, just run it.
		if !reg { 
			printCommands = true 
		} else {
			age := dir.Mtime_ns
	
			for y := 0; y < t.dependlen; y++{
				preqDir, err := os.Stat(targ.(*MakeTarget).dependencies[y])
				if err != nil { return err }
				preqAge := preqDir.Mtime_ns
		
				if age < preqAge { printCommands = true; break }
			}
		}
	}
	
	if printCommands { fmt.Println(targ.(*MakeTarget).Commands()) }
	
	return nil
}
			
