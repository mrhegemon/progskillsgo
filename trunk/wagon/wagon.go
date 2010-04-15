package wagon_game

import "list"

//----------WAGON----------//
type wagon struct {
	name string
	x, y int   //(x, y) location
}

func(this *wagon) getLocation() (int, int) {
	return this.x, this.y
}
func(this *wagon) move(r, c int) os.Error{
	if validMove(r,c) {
		this.x, this.y = r, c
		return nil
	}
	return os.NewError("WagonGame:  Invalid Location")
}

func newWagon(n string, r, c int) *wagon {
	temp := new(wagon)
	temp.name = n
	temp.x, temp.y = r, c
}

//----------GAME-----------//
var( row, col, char int
     list *LinkedList
)

func validMove(r, c int) bool {
	return r >= 0 && r < row && c >= 0 && c < col 
}

func NewGame(r, c int) {
	row, col = r, c
	list = new(LinkedList)

	Act('a')
	Act('A')
}

func Act(command int) os.Error {
	if command == 'a' {
		list.PushFront(newWagon(string([]int{char})), 0, 0)
		char++
	} else if command == 'A' {
		list.PushBack(newWagon(string[]int{char}), row-1, col-1)
		char++
	} else if command == 'u' {
		if err := list.ApplyToAllFront(up(true)); err != nil {
			return err
		}
	} else if command == 'U' {
		if err := list.ApplyToAllBack(up(false)); err != nil {
			return err
		}
	}

	return nil
}

func Print() {
	//iterate over all the values of the list, and print them,
	//using ANSI to align them properly. Each wagon shouldn't be
	//more than 1 character long, so we should be able to line them up
	//nicely.

	for y:= 0; y < list.Len(); y++ {
		val := list.At(y).(*wagon)
		
		r, c := val.getLocation()

		//print using ANSI
	}

	//using ANSI, print enough lines so that there are "row" number of
	//lines on the screen
}







