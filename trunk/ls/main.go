package main

import ("ls"; "os"; "fmt"; "io/ioutil")

var (
	long bool = false
	recurse bool = false
)

func main() {
	var nodes []ls.Node = getNodes(".")

	for y := 0; y < len(nodes); y++ {
		var funcs []ls.NodeFunc
		if(nodes[y].IsFile()) {
			funcs = filefuncs()
		} //else {
		//	funcs = dirfuncs()
		//}

		if err := ls.Make(funcs)(&nodes[y]); err != nil {
			fmt.Fprintln(os.Stderr, err.String())
			os.Exit(-1)
		}
	}
}

func filefuncs() []ls.NodeFunc {
	if !long && !recurse {
		return []ls.NodeFunc{
			ls.PrintName,
			ls.PrintNewLine,
		}
	} else if(long && !recurse) {
		return []ls.NodeFunc {
			 ls.PrintPermissions,
			 ls.PrintSpace,
			 ls.PrintOwner,
			 ls.PrintSpace,
			 ls.PrintGroup,
			 ls.PrintSpace,
			 ls.PrintSize,
			 ls.PrintSpace,
			 ls.PrintModified,
			 ls.PrintSpace,
			 ls.PrintName,
			 ls.PrintNewLine,
		}
	}
	return nil
}

func getNodes(dir string) []ls.Node {
	var dirs []*os.Dir
	var err os.Error

	dirs, err = ioutil.ReadDir(dir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.String())
		os.Exit(-1)
	}

	var nodes []ls.Node = make([]ls.Node, len(dirs))

	for y := 0; y < len(dirs); y++ {
		nodes[y].Name = dirs[y].Name
	}

	return nodes
}
