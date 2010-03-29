package ls

import ("fmt"; "os")

type NodeFunc func(*Node) os.Error

type Node struct{
	Name string
}

func (n *Node) getDir() *os.Dir {
	var dir *os.Dir
	var err os.Error
	dir, err = os.Stat(n.Name)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.String())
		os.Exit(-1)
	}

	return dir
}

func (n *Node) IsDirectory() bool {
	return n.getDir().IsDirectory()
}

func (n *Node) IsFile() bool {
	return !n.IsDirectory()
}


func PrintColon(n *Node) os.Error {
	_, x := fmt.Print(":")
	return x
}

func PrintSpace(n *Node) os.Error {
	_, x := fmt.Print(" ")
	return x
}

func PrintTab(n *Node) os.Error {
	_, x := fmt.Print("\t")
	return x
}

func PrintName(n *Node) os.Error {
	_, x := fmt.Print(n.Name + "  ")
	return x
}

func PrintNewLine(n *Node) os.Error {
	_, x := fmt.Print("\n")
	return x
}

//Prints the owner of the file
//(as indicated by the Node)
func PrintOwner(n *Node) os.Error {
	_, x := fmt.Print(n.getDir().Uid)
	return x
}

//Prints the group ownership of the file
//(as indicated by the Node)
func PrintGroup(n *Node) os.Error {
	_, x := fmt.Print(n.getDir().Gid)
	return x
}

//Prints the permissions of the file (as
//indicated by the Node)
func PrintPermissions(n *Node) os.Error {
	_, x := fmt.Println(n.getDir().Permission())
	return x
}

//Prints the size (in bytes) of the file
//(as indicated by the Node)
func PrintSize(n *Node) os.Error {
	_, x := fmt.Println(n.getDir().Size)
	return x
}

//Prints the time the file (as indicated by the Node)
//was last modified
func PrintModified(n *Node) os.Error {
	_, x := fmt.Println(n.getDir().Mtime_ns)
	return x
}



//func convTime(time uint64) string

//func prettyUser(uid int) string

//func prettyPermissions(perm int) string

func Make(body []NodeFunc) NodeFunc {
	return func(n *Node) os.Error {
		for f := 0; f < len(body); f++ {
			x := body[f](n)
			if x != nil {
				return x
			}
		}
		return nil;
	}
}
