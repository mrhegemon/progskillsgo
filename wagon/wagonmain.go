/*
Authors: William Broza, Tym Lipari
Matrix Testing program

tests add and mult matrices

usage:
	matrix
*/
package main

import ("wagon"; "os"; "strconv"; "fmt")//; "bytes")

func main() {
	var row, col int
	var err os.Error

	if len(os.Args) != 3 {
		fmt.Println("USAGE: wagon [rows] [cols]")
		os.Exit(-1)
	} else {
			if row, err = strconv.Atoi(os.Args[1]); err != nil {
					fmt.Println(err.String())
					os.Exit(-1)
			}
			if col, err = strconv.Atoi(os.Args[2]); err != nil {
					fmt.Println(err.String())
					os.Exit(-1)
			}
	}
	//set up stty
	_, err = os.ForkExec("/bin/stty", []string{"stty", "cbreak"}, os.Environ(), "", nil)//exec.Run
	if err != nil {
		fmt.Println(err.String())
		os.Exit(-1)
	}
	
	wagon_game.NewGame(row, col)
	wagon_game.Print()
	fmt.Print("   a/A to add, u/U to move up, d/D to move down, l/L to move left, r/R to move right, q/Q to quit")
	//using ANSI, print enough lines so that there are "row" number of
	//lines on the screen
	fmt.Print("\033[" + strconv.Itoa(row + 1) + ";0H")

	inp := make([]byte, 4)
	os.Stdin.Read(inp)
	charString := string(inp)
	chars := []int(charString)
	
	for chars[0] != 'q' && chars[0] != 'Q' {
		wagon_game.Act(chars[0])
		wagon_game.Print()
		fmt.Print("   a/A to add, u/U to move up, d/D to move down, l/L to move left, r/R to move right, q/Q to quit")
		//using ANSI, print enough lines so that there are "row" number of
		//lines on the screen
		fmt.Print("\033[" + strconv.Itoa(row + 1) + ";0H")

		//refresh the input
		os.Stdin.Read(inp)
		charString = string(inp)
  		chars = []int(charString)
	}
	
	fmt.Print("\033c")
	//fix stty
	_, err = os.ForkExec("/bin/stty", []string{"stty", "sane"}, os.Environ(), "", nil)
	if err != nil {
		fmt.Println(err.String())
		os.Exit(-1)
	}

}


