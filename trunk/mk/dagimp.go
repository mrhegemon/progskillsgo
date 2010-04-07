/*
dagimp implements all of the dag.go interfaces to build the directed acyclic 
graph that governs the command order of a makefile.

Authors: William Broza, Tym Lipari
*/

package dag

import 	("os"
		 "io/ioutil"
		 "bytes"
		 "strings"
)

//SetImp 
//Implements the map
type SetImp struct {
	setmap map[string]Target
}

//Get(name string)
//returns the name in setmap
//returns: Target
func(s *SetImp) Get(name string) Target {
	return s.setmap[name]
}

//Apply(t Target, a Action)
//applys action to target
//returns: error
func(s *SetImp) Apply(t Target, a Action) os.Error {
	if err := t.ApplyPreq(a); err != nil { return err }
	if err := t.Apply(a); err != nil { return err }
	return nil
}

//AddFile(fname string, t TargetFactory)
//adds file to target factory
//returns: string and error
func(s *SetImp) AddFile(fname string, t TargetFactory) (string, os.Error) {

	fileByteS, err := ioutil.ReadFile(fname) // may be ReadAll
	if err != nil { return  "" , err }
	
	str := bytes.NewBuffer(fileByteS).String()
	
	return s.AddString(str, t)
}

//AddString(doc string, t TargetFactory)
//adds string to target factory
//returns: string and error
func(s *SetImp) AddString(doc string, t TargetFactory) (string, os.Error) {
	split := strings.SplitAfter(doc, "\n", 0)
	return s.Add(split, t)
}

//Add(lines []string, t TargetFactory)
//adds elements to targ
//returns: string and error
func(s *SetImp) Add(lines []string, t TargetFactory) (string, os.Error) {
	/*var first string
	var begin int = -1
	
	for loc, val := range lines {
		//if the line is a target...
		if strings.Index(val, "\t") != 0 && strings.Index(val, " ") != 0 && strings.Index(val, "\n") != 0 && len(val) > 0 {
			if begin > -1 || loc == len(lines) - 1 {
				targ, err := t(s, lines[begin:loc], t)
				if err == nil {
					str, nerr := s.Put(targ)
					if nerr != nil { return "", nerr }
					
					if first == "" { first = str.Name() }
				} else { return "", err }
			}
			begin = loc
		}
	}
	return first, nil*/

	var first string
	var begin int = -1
	/*printer := func(s []string) {
		for y:= 0; y < len(s); y++ {
		}
	}*/
	for y := 0; y < len(lines); y++ {
		if strings.Index(lines[y], "\t") != 0 && strings.Index(lines[y], " ") != 0 && strings.Index(lines[y], "\n") != 0 {
			if begin > -1 || y == len(lines) - 1 {
				//printer(lines[begin:y])
				targ, err := t(s, lines[begin:y], t)
				if y == len(lines) - 1 { targ, err = t(s, lines[begin:y+1], t) }
				
				if first == "" { first = targ.Name() }
				
				if err == nil {
					str, nerr := s.Put(targ)
					if nerr != nil { return "", nerr 
					} else if (y == 0) {
						first = str.Name() 
					}
				} else { return "", err }
			}
			begin = y
		}
	}
	return first, nil
}

//Put(t Target)
//puts target in SetImp
//returns: target and error
func(s *SetImp) Put(t Target) (Target, os.Error) {
	fromMap := s.Get(t.Name())
	if fromMap != nil {
		fromMap.Merge(t)
	} else { s.setmap[t.Name()] = t }
	
	return s.Get(t.Name()), nil
}

//String()
//returns string from SetImp
//returns: string
func(s *SetImp) String() string {
	var toReturn string
	for f, g := range s.setmap {
		toReturn += "[" + f + "] = " + g.String() + "\n"
	}
	return toReturn
}		

//NewSet()
//Returns a new SetImp
func NewSet() Set {
	n := new(SetImp)
	n.setmap = make(map[string] Target)
	return n
}

//TargImp Struct
type TargImp struct {
	name string
	dependencies []string
	dagset Set
	dependlen int
	commandSent bool
	cyclic bool
}

//isDependent(depend string)
//tells if dependant
//returns: bool of dependance
func(t *TargImp) isDependent(depend string) bool {
	for _, y := range t.dependencies {
		if y == depend { return true }
	}
	return false
}

//ApplyPreq(a Action)
//Applys prereq to target
//returns: error
func(t *TargImp) ApplyPreq(a Action) os.Error {
	for y := 0; y < t.dependlen; y++ {
		var targ Target
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

//Name()
//returns name for the target
//returns: string Name
func(t *TargImp) Name() string {
	return t.name
}

//String()
//returns string for Target
//returns: string
func(t *TargImp) String() string {
	var toReturn string = t.Name() + ":\t"
	for _, y := range t.dependencies {
		toReturn += y + "  "
	}
	return toReturn
}

//Merge(other Target)
//merges targets
//returns: error
func(t *TargImp) Merge(other Target) (Target, os.Error) {
	x := other.(*TargImp)
	if x.Name() != t.Name() {
		return nil, os.NewError("cannot merge targets with different names")
	}
	
	for y := 0; y < x.dependlen; y++ {
		if !t.isDependent(x.dependencies[y]) {
			t.dependencies[t.dependlen] = x.dependencies[y]
			t.dependlen++
		}
	}
	
	return t, nil
}

//Apply(a Action)
//applys action to TargImp
//returns: error
func(t *TargImp) Apply(a Action) os.Error {

	if !t.commandSent {
		t.commandSent = true
		return a(t)
	}
	return nil
}

//DagTargetFact(s Set, str []string, t TargetFactory)
//target factory for dag implementation
//returns: error
func DagTargetFact(s Set, str []string, t TargetFactory) (Target, os.Error) {
		targ := new(TargImp)
		
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
		return targ, nil
}
