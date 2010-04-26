/*

This package contains wagonmain.go, wagon.go, listTest.go and list.go

Authors: William Broza, Tym Lipari

list.go contains the package list for creating, adding, and removing
elements from a list.

listTest.go test the list's functionality

wagon.go contains the package list for creating, adding, and removing
wagons from a wagon train using a linkedList.

wagonmain.go contains the main for the wagon game.

to run listTest.go, use "make test" or "make test -v" for verbose testing.

To run wagonmain.go, run make using makefile.  Then run wagon with 2 arguments.
The first argument sets the row length for the game.
The second argument sets the column length for the game.

NOTE: The program does not properly support stty commands.  They must be
	manually set to function properly.  ForkExec calls are implemented but do not
	seem to function.

	Example: 	stty cbreak
						wagon 20 10
						stty sane

*/
package documentation
