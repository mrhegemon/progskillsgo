package dag

import "os"

type SetImp struct {
	setmap map[string]Target
}

func(s *SetImp) Get(name string) Target {
	return setmap[name]
}

func(s *SetImp) Apply(t Target, a Action) os.Error {
	if err := t.ApplyPreq(a); err != nil { return err }
	if err := t.Apply(a); err != nil { return err }
	return nil
}

func(s *SetImp) AddFile(fname string, t TargetFactory) (string, os.Error) {
	file, err := os.Open(fname, os.O_RDONLY, 444)
	if err != nil { return _, err }

		file

	return AddString(str, t)
}

func(s *SetImp) AddString(doc string, t TargetFactory) (string, os.Error) {
	split := strings.SplitAfter(doc, "\n", 0)
	
	return s.Add(split, t)
}

func(s *SetImp) Add(lines []string, t TargetFactory) (string, os.Error) {
	var first string
	for y := 0; y < len(lines); y++ {
		if strings.Index(lines[y], "\t" != 0 {
			targ, err := t(s, lines[y:y+1], t)
			if err == nil {
				str, nerr := s.Put(t)
				if nerr != nil { return _, nerr }
				else if y == 0 { first = str }
			} else { return _, err }
		}
	}
	return first, _
}

func(s *SetImp) Put(t Target) (string, os.Error) {
	fromMap := s.Get(t.Name)
	if fromMap != nil {
		fromMap.Merge(t)
	} else { s.setmap[t.Name] = t }
	
	return s.Get(t.Name), _
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
	var name string
	var dependencies []string
	var dagset Set
}

func(t *TargImp) isDependent(depend string) bool {
	for _, y := range t.dependencies {
		if y == depend { return true }
	}
	return false
}

func(t *TargImp) ApplyPreq(a Action) os.Error {
	for _, y := range t.dependencies {
		targ := dagset.Get(y)
		//if the target doesn't exist, send an error
		if targ:= dagset.Get(y)targ == nil { 
			return os.NewError("non-existant Target:  " + y) 
		}
		
		//if the targets prereqs sent an error, send it on
		if err1 := targ.ApplyPreq(a); err1 != nil {
			return err1
		}
		
		//if the target sent an error, send it on
		if err2 := targ.Apply(a); err2 != nil {
			return err2
		}
	}
	return nil
}

func(t *TargImp) Name() string {
	return t.Name
}

func(t *TargImp) String() string {
	var toReturn string = t.Name() + ":\t"
	for _, y := dependencies {
		toReturn += y + "  "
	}
	return toReturn
}

func(t *TargImp) Merge(other Target) (Target, os.Error) {
	x := other.(*TargImp)
	if x.Name() != t.Name() {
		return os.NewError("cannot merge targets with different names")
	}
	
	for _, y := range other.dependencies {
		if !t.isDependent(y) {
			t.dependencies[len(t.dependencies)] = y
		}
	}
	
	return t, _
}

func(t *TargImp) Apply(a Action) os.Error {
	return a(t)
}