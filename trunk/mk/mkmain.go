package main

import ("mk"; "dag"; "fmt"; "flag")
func main() { 
	flag.Parse()
	targetFact := mk.MkTargetFact
	action := mk.Act
	
	if err := dag.Main(targetFact, action); err != nil {
		fmt.Println("\n" + err.String())
	}
}
