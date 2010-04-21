/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	matrix
*/
package main

import ("wagon"; "os"; "strconv")//; "bytes")

func main() {
	var row, col int
	var err os.Error

	if len(os.Args) != 3 {
		println("USAGE:   wagon [rows] [cols]")
		os.Exit(-1)
	} else {
			if row, err = strconv.Atoi(os.Args[1]); err != nil {
					println(err.String())
					os.Exit(-1)
			}
			if col, err = strconv.Atoi(os.Args[2]); err != nil {
					println(err.String())
					os.Exit(-1)
			}
	}
	

	//set up stty
	_, err = os.ForkExec("/bin/stty", []string{"cbreak"}, os.Environ(), "", nil)
	if err != nil {
		println(err.String())
		os.Exit(-1)
	}

	inp := make([]byte, 4)
	os.Stdin.Read(inp)
	charString := string(inp)
  chars := []int(charString)

	wagon_game.NewGame(row, col)
	wagon_game.Print()
	for chars[0] != 'q' && chars[0] != 'Q' {
		wagon_game.Act(chars[0])
		wagon_game.Print()


		//refresh the input
		os.Stdin.Read(inp)
		charString = string(inp)
  	chars = []int(charString)
	}
/*
	println("\033[7m Linux OS! Best OS!! \033[0m")
	println("\033[31m I am in Red")
	println("\033[5;10H Hello")
	println("\033[7m Linux OS! Best OS!! \033[0m")
	println("\033[31m I am in Red")
	println("\033c")
*/
}


