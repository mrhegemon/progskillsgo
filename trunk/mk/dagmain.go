/*
Main function

Authors: William Broza, Tym Lipari
*/

package main

import ("dag"; "flag"; "os"; "fmt")

func main() {
	flag.Parse()
	targetFact := dag.DagTargetFact
	action := func(t dag.Target) os.Error {
		_, err := fmt.Println(t.Name())
		return err
	}

	if err := dag.Main(targetFact, action); err != nil {
		fmt.Println("\n" + err.String())
	}
}
