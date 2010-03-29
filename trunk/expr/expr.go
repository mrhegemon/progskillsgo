/*
Authors: William Broza, Tym Lipari
Arithmetic Expression Interpreter

expr evaluates expressions simular to the unix command shell

usage:
	expr [arg1] ... [argn] */
package main

import ("fmt"; "os"; "strconv"; "flag"; "container/vector")

//constants for interpreting
const( PLUS = "+"
       MINUS = "-"
       DIV = "/"
       TIMES = "*"
       MOD = "%"
       OPER = 0
       PAREN = 1
       NUM = 2
       BAD = -1
     )

//program main
func main(){
	var v *vector.Vector
	v = new(vector.Vector)
	var isInvalid bool
	var allInvalid bool = true
	
	for x:=0; x < flag.NArg(); x++ {
		isInvalid, _ = valid(flag.Arg(x))
	
		if !isInvalid && allInvalid {
			allInvalid = false
		} else if isInvalid && !allInvalid {
			fmt.Println("expr:  invalid argument")
			os.Exit(-1)
		}
		var subv vector.Vector
		subv.Push(flag.Arg(x))
		v.Push(&subv)
	}
	v, _ = regroup(v)
	v = evalVector(v)
	printVector(v)
}

/*printVector 
prints the final solution or error message.  A correctly formatted final 
solution is a vector of length one

usage:
	printVector(*vector) */
func printVector(v *vector.Vector) {
	for x:=0; x < v.Len(); x++ {
		var vx *vector.Vector = v.At(x).(*vector.Vector)
		if vx.Len() == 1 {
			fmt.Println(vx.At(0).(string))
		} else {
			fmt.Println("Error: [" + strconv.Itoa(x) + "] > ")
			printVector(vx)
		}
	}
}

/*valid
returns true if the value is an operator, parenthese or valid number,
false otherwise.

usage:
	valid(string)
returns: 
	0 for operator
	1 for parenthese
	2 for number
	-1 for invalid arguments*/
func valid(arg string) (bool, int) {
	if arg == PLUS || arg == MINUS || arg == DIV || arg == TIMES || arg == MOD {
		return true, OPER
	} else if arg == "(" || arg == ")"{
		return true, PAREN
	}
	var conv int
	conv, _ = strconv.Atoi(arg)
	if conv > 0 || arg == "0" {
		return true, NUM
	}
	return false, BAD
}

/* regroup
Regroups the given vector into terms and returns a new vector containing the
regrouped terms, and an int indicating the success state: 0 = good, -1 = bad
syntax 

usage:
	regroup(*vector)
returns:
	vector of nested vectors */
func regroup(toGroup *vector.Vector) (*vector.Vector, int) {
	for x:=0; x < toGroup.Len(); x++ {
		//Get the current subvector
		var subv *vector.Vector = toGroup.At(x).(*vector.Vector)
		//If the vector is of len 1, it's a token, so pull out the token
		//If not, it'll be "", so no problem
		var argu string
		if subv.Len() == 1 { 
			argu = subv.At(0).(string)
			//Set up temp vars for the inner loop, 
			//they'll be needed after the loop though
			var vec vector.Vector
			var y int = 0
			if argu == "(" {
				//parentheses count
				var pc = 1
				var moreParen bool = false
				for y=x+1; pc > 0; y++ {
					//Again, see if you're working with a
					//token this time, reset argu in case 
					//it's still a ( from before
					var vx *vector.Vector = 
					toGroup.At(y).(*vector.Vector)
					if vx.Len() == 1 {
						argu = vx.At(0).(string) 
					} else { 
						argu = "" 
					}
					//Augment the counts if necessary
					if argu == "(" {
						pc++
						moreParen = true
					} else if argu == ")" {
						pc--
						if pc == 0 {
							break
						}
					}
					//Push the subvector into another
					//subvector
					var tempV *vector.Vector =
					toGroup.At(y).(*vector.Vector)
					vec.Push(tempV)
				}
				var vNew *vector.Vector = &vec
				if moreParen {
					vNew, _ = regroup(&vec)
				}
				//Delete all subvectors from x-y, then put the
				//new subvector in its place
				for temp:=x; temp <= y; temp++{
					toGroup.Delete(x)
				}
				toGroup.Insert(x, vNew)
			}
		} else {
			toGroup.Delete(x)
			subv, _ = regroup(subv)
			toGroup.Insert(x, subv)
		}
	}
	return toGroup, 0
}

/*evalVector
evaluates the expression stored in the vector recursively

usage:
	evalVector(*vector)
returs:
	*vector */
func evalVector(v *vector.Vector) *vector.Vector {
	//to preserve order of operations, two loops are used
	//mult, divide, mod loop
	for x:=0; x < v.Len(); x++ {
		var vx *vector.Vector = v.At(x).(*vector.Vector)
		if vx.Len() == 1 {
			var arg1 string = vx.At(0).(string)
			var operType int
			_, operType = valid(arg1)
			//checks if there is two consecutive number in the input
			if operType == NUM && ( x > 0 ) {
				var errChkV int
				_,errChkV = valid(v.At(x-1).(*vector.Vector).At(0).(string))
				if errChkV == NUM {
					fmt.Println("expr: syntax error")
					os.Exit(-2)
				}
			}
			if (operType == OPER && arg1 == TIMES || arg1 == DIV || arg1 == MOD) && v.At(x-1).(*vector.Vector).Len() == 1 && v.At(x+1).(*vector.Vector).Len() == 1 {
				var tempVec *vector.Vector
				tempVec = evaluate(v.At(x-1).(*vector.Vector), vx, v.At(x+1).(*vector.Vector))
				v.Delete(x-1)
				v.Delete(x-1)
				v.Delete(x-1)
				v.Insert(x-1, tempVec)
			}
		} else {		
			v.Delete(x)
			v.Insert(x, evalVector(vx).At(0))
		}
	}
	//add, sub loop
	for x:=0; x < v.Len(); x++ {
		var vx *vector.Vector = v.At(x).(*vector.Vector)
		if vx.Len() == 1 {
			var arg1 string = vx.At(0).(string)			
			var operType int
			_, operType = valid(arg1)
			//checks if there is two consecutive number in the input
			if operType == NUM && ( x > 0 ) {
				var errChkV int
				_,errChkV = valid(v.At(x-1).(*vector.Vector).At(0).(string))
				if errChkV == NUM {
					fmt.Println("expr: syntax error")
					os.Exit(-2)
				}
			}
			if (operType == OPER && arg1 == PLUS || arg1 == MINUS) && v.At(x-1).(*vector.Vector).Len() == 1 && v.At(x+1).(*vector.Vector).Len() == 1 {
				var tempVec *vector.Vector
				tempVec = evaluate(v.At(x-1).(*vector.Vector), vx, v.At(x+1).(*vector.Vector))
				v.Delete(x-1)
				v.Delete(x-1)
				v.Delete(x-1)
				v.Insert(x-1, tempVec)
			} 
		} else {
			v.Delete(x)
			v.Insert(x, evalVector(vx).At(0))
		}
	}
	if v.Len() != 1 {
		v = evalVector(v)
	}
	return v
}

/*evaluate
computes the result from the expression values

usage:
	evaluate(*vector, *vector, *vector)
returns:
	*vector */
func evaluate(arg1 *vector.Vector, oper *vector.Vector, arg2 *vector.Vector) *vector.Vector {
	//Store the operator in a temp string, to save typing it out
	var operS string
	operS = oper.At(0).(string)
	var val1, val2 int 
	var err1, err2 os.Error
	val1, err1 = strconv.Atoi(arg1.At(0).(string))
	val2, err2 = strconv.Atoi(arg2.At(0).(string))
	//screens for consecutive operators
	if(err1 != nil || err2 != nil){
		fmt.Println("expr: syntax error")
		os.Exit(-2)
	}
	var result int = -1
	//Evaluate based on the operator
	if operS == "+" {
		result = val1 + val2
	} else if operS == "-" {
		result = val1 - val2
	} else if operS == "/" {
		result = val1 / val2
	} else if operS == "*" {
		result = val1 * val2
	} else if operS == "%" {
		result = val1 % val2
	}
	//Clear the arg1 vector and add the result to it, then return
	//(saves memory by not creating a new vector)
	arg1.Cut(0, arg1.Len())
	arg1.Push(strconv.Itoa(result))
	return arg1
}
