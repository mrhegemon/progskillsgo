package main

import ("dag"; "flag"; "os"; "fmt")

func main() {
	flag.Parse()
	targetFact := dag.DagTargetFact
	action := func(t dag.Target) os.Error {
		_, err := fmt.Println(t.Name())
		return err
	}

	dag.Main(targetFact, action)
}
