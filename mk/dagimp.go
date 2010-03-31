package dag

import 	("os"
				"io/ioutil"
				"bytes"
				"strings"
				//"strconv"
)

type SetImp struct {
	setmap map[string]Target
}

func(s *SetImp) Get(name string) Target {
	return s.setmap[name]
}

func(s *SetImp) Apply(t Target, a Action) os.Error {
	if err := t.ApplyPreq(a); err != nil { return err }
	if err := t.Apply(a); err != nil { return err }
	return nil
}

func(s *SetImp) AddFile(fname string, t TargetFactory) (string, os.Error) {

	//probably wrong
	//file, err := os.Open(fname, os.O_RDONLY, 444)
	
	fileByteS, err := ioutil.ReadFile(fname) // may be ReadAll
	if err != nil { return  "" , err }
	
	str := bytes.NewBuffer(fileByteS).String()
	
	return s.AddString(str, t)
}

func(s *SetImp) AddString(doc string, t TargetFactory) (string, os.Error) {
	split := strings.SplitAfter(doc, "\n", 0)
	
	return s.Add(split, t)
}

func(s *SetImp) Add(lines []string, t TargetFactory) (string, os.Error) {
	var first string
	for y := 0; y < len(lines); y++ {
		if strings.Index(lines[y], "\t") != 0 && len(lines[y]) > 0 {
			targ, err := t(s, lines[y:y+1], t)
			if err == nil {
				str, nerr := s.Put(targ)
				if nerr != nil { return "", nerr 
				} else if (y == 0) {
					first = str.Name() 
				}
			} else { return "", err }
		}
	}
	return first, nil
}

func(s *SetImp) Put(t Target) (Target, os.Error) {
	fromMap := s.Get(t.Name())
	if fromMap != nil {
		fromMap.Merge(t)
	} else { s.setmap[t.Name()] = t }
	
	return s.Get(t.Name()), nil
}

func(s *SetImp) String() string {
	var toReturn string
	for f, g := range s.setmap {
		toReturn += "[" + f + "] = " + g.String() + "\n"
	}
	return toReturn
}		

//Returns a new SetImp
func NewSet() Set {
	n := new(SetImp)
	n.setmap = make(map[string] Target)
	return n
}

type TargImp struct {
	name string
	dependencies []string
	dagset Set
}

func(t *TargImp) isDependent(depend string) bool {
	for _, y := range t.dependencies {
		if y == depend { return true }
	}
	return false
}

func(t *TargImp) ApplyPreq(a Action) os.Error {
	for _, y := range t.dependencies {
		var targ Target
		//if the target doesn't exist, send an error
		if targ = t.dagset.Get(y); targ == nil { 
			return os.NewError("non-existant Target:  " + y) 
		}
		
		//if the targets prereqs sent an error, send it on
		if err1 := targ.ApplyPreq(a); err1 != nil {
			println(err1.String())
			return err1
		}
		
		//if the target sent an error, send it on
		if err2 := targ.Apply(a); err2 != nil {
			println(err2.String())
			return err2
		}
	}
	return nil
}

func(t *TargImp) Name() string {
	return t.name
}

func(t *TargImp) String() string {
	var toReturn string = t.Name() + ":\t"
	for _, y := range t.dependencies {
		toReturn += y + "  "
	}
	return toReturn
}

func(t *TargImp) Merge(other Target) (Target, os.Error) {
	x := other.(*TargImp)
	println("CURRENT:  " + t.String())
	println("OTHER:  " + other.String())
	if x.Name() != t.Name() {
		return nil, os.NewError("cannot merge targets with different names")
	}
	
	for _, y := range other.(*TargImp).dependencies {
		if !t.isDependent(y) {
			t.dependencies[len(t.dependencies)] = y
		}
	}
	
	return t, nil
}

func(t *TargImp) Apply(a Action) os.Error {
	return a(t)
}

func DagTargetFact(s Set, str []string, t TargetFactory) (Target, os.Error) {
		targ := new(TargImp)
		
		tokens := strings.Fields(str[0])
		
		targ.name = tokens[0]
		targ.dependencies = make([]string, 20)
		copy(targ.dependencies, tokens[1:])
		targ.dagset = s
		
		for _, y := range targ.dependencies {
			if s.Get(y) == nil {
				temp := []string { y }
				
				for _, xyz := range temp { println(xyz) }
				tempTarg, nerr := t(s, temp, t)
				if nerr != nil { return nil, nerr }
				
				_, nerr2 := s.Put(tempTarg)
				if nerr2 != nil { return tempTarg, nerr2 }
			}
		}
		return targ, nil
}
