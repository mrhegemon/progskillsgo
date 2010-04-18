package wagon_game

import ("list"; "os")

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
	return temp
}

//----------GAME-----------//
var( row, col, char int
     my_list *list.LinkedList
)

func validMove(r, c int) bool {
	return r >= 0 && r < row && c >= 0 && c < col 
}

func NewGame(r, c int) {
	row, col = r, c
	my_list = new(list.LinkedList)

	Act('a')
	Act('A')
}

func Act(command int) os.Error {
	switch command {
		case 'a': {
			my_list.PushFront(newWagon(string([]int{char}), 0, 0))
			char++
		}
		case 'A': {
			my_list.PushBack(newWagon(string([]int{char}), 0, 0))
			char++
		}
		//move up or down
		case 'u': return my_list.ApplyToAllFromFront(upDown(true, true))
		case 'U': return my_list.ApplyToAllFromBack(upDown(false, true))
		case 'd': return my_list.ApplyToAllFromFront(upDown(true, false))
		case 'D': return my_list.ApplyToAllFromBack(upDown(false, false))
		
		//move left or right
		case 'l': return my_list.ApplyToAllFromFront(leftRight(true, true))
		case 'L': return my_list.ApplyToAllFromBack(leftRight(false, true))
		case 'r': return my_list.ApplyToAllFromFront(leftRight(true, false))
		case 'R': return my_list.ApplyToAllFromBack(leftRight(false, false))
	}
	return nil
}

func upDown(front, up bool) func(interface{}, int)os.Error {
	if front {
		return func(val interface{}, index int)os.Error {
			wag := val.(*wagon)
			if index == 0 {
				row, col := wag.getLocation()
				switch up {
					case true: return wag.move(row - 1, col)
					case false: return wag.move(row + 1, col)
				}
			} else {
				val, _ := my_list.At(index - 1)
				prev := val.(*wagon)
				return wag.move(prev.getLocation())
			}
			return nil
		}
	}
	return func(val interface{}, index int)os.Error {
		wag := val.(*wagon)
		if index == my_list.Len() - 1 {
			row, col := wag.getLocation()
			switch up{
				case true: return wag.move(row - 1, col)
				case false: return wag.move(row + 1, col)
			}
		} else {
			val, _ := my_list.At(index + 1)
			prev := val.(*wagon)
			return wag.move(prev.getLocation())
		}
		return nil
	}
}
func leftRight(front, left bool) func(interface{}, int)os.Error {
	if front {
		return func(val interface{}, index int)os.Error {
			wag := val.(*wagon)
			if index == 0 {
				row, col := wag.getLocation()
				switch left {
					case true: return wag.move(row, col - 1)
					case false: return wag.move(row, col - 1)
				}
			} else {
				val, _ := my_list.At(index - 1)
				prev := val.(*wagon)
				return wag.move(prev.getLocation())
			}
			return nil
		}
	}
	return func(val interface{}, index int)os.Error {
		wag := val.(*wagon)
		if index == my_list.Len() - 1 {
			row, col := wag.getLocation()
			switch left {
				case true: return wag.move(row, col - 1)
				case false: return wag.move(row, col + 1)
			}
		} else {
			val, _ := my_list.At(index + 1)
			prev := val.(*wagon)
			return wag.move(prev.getLocation())
		}
		return nil
	}
}

func Print() {
	//iterate over all the values of the list, and print them,
	//using ANSI to align them properly. Each wagon shouldn't be
	//more than 1 character long, so we should be able to line them up
	//nicely.

	for y:= 0; y < my_list.Len(); y++ {
		//val, _ := my_list.At(y)
		
		//r, c := val.(*wagon).getLocation()

		//print using ANSI
	}

	//using ANSI, print enough lines so that there are "row" number of
	//lines on the screen
}







