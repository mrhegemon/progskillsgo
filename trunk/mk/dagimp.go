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
	file, err := os.Open(s, os.O_RDONLY, 444)
	if err != nil { return _, err }

	//this function needs to read each non-tabbed line from the file
	//and add that line to slice of strings

	//then it passes everything to the target factory, which will
	//return the targets.

	//Return the name of the first target, and nil for the error.
	//If an error is encountered along the way, return _ for the string
	//and the error.
	return _, nil
}

func(s *SetImp) AddString(doc string, t TargetFactory) (string, os.Error) {
	//essentially the same as AddFile, except that instead of reading from
	//a document, we're reading from a string (that is basically the document)...
	
	//we may as well just read the entire document into a string in AddFile
	//and then just pass that into this function.
	
	return whatever, nil
}

func(s *SetImp) Add(lines []string, T TargetFactory) (string, os.Error) {
	//again, essentially the same as the line above, except that I think
	//this is the one where we'll actually do something.
	
	//we should parse each line, where the first value is the name of the new
	//target, and each following value is the name of a prerequisite.
	
	//we can then call the Put method for that target.
	
	//one thing to consider is that we may be putting targets in before their
	//prerequisites. I think that's what he means by merge, so that we put in
	//that prereq first with no dependencies, and then later, when we find it
	//in the document, we merge it with its existing target. (hence the Merge
	//method on Target)
	
	//btw, don't think i'm writing these comments in here to be a dick. it's
	//more for me than you. I know you know how to code. :)
}

func(s *SetImp) Put(t Target) (string, os.Error) {
	fromMap := s.Get(t.Name)
	if fromMap == nil {
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
}

func(t *TargImp) isDependent(depend string) bool {
	for _, y := range t.dependencies {
		if y == depend { return true }
	}
	return false
}

/*func(t *TargImp) ApplyPreq(a Action) os.Error {
	for _, y := range t.dependencies {
		//somehow get the target for y
		//send a to that target && store the resulting os.Error in err
		if err != nil { return err }
	}
	return nil
}*/

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