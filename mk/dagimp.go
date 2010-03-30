package dag

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
