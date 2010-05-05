/*This package contains games.go, rps.go, ttt.go, view.go

Authors: William Broza, Tym Lipari
(games.go authored by Axel Schreiner)

games.go defines the interface for making a game view

view.go implements that interface generically. it interacts with string channel
to determine what commands/printout need to be sent to the user's screen

ttt.go implements a referee for the tic-tac-toe game, using two game views

rps.go implements a referee for the rock paper scissors game, using two game 
views

net.go implements an interface for game views to interact through net channels

to run the rock paper scissors game, execute
	rps-stdin [terminal]

where terminal is an optional second terminal to print one of the players to. if excluded, it uses the same terminal for both players. the order of input is 
random for which player goes before the other.

to run the tic tac toe game, execute
	ttt-stdin [terminal]

where terminal is an optional second terminal to print one of the players to. if excluded, it uses the same terminal for both players. the order of input is 
random for which player goes before the other.


Valid moves for Rock Paper Scissors:
	p - paper
	r  - rock
	s - scissors

Note: the program does not currently error check. the output is undefined for 
any input other than the ones listed above.


Valid moves for Tic Tac Toe:
	c - center square
	nw - northwest square
	n - north square
	ne - northeast square
	e - east square
	se - southeast square
	s  - south square
	sw - southwest square
	w - west square

Note: the program does little error checking. the output is undefined for any 
input other than the ones listed above.

To quit either game, press Ctrl-C
*/
package documentation
