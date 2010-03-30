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

//Returns a new SetImp
func NewSet() Set {
	return nil
}
