import ("dag"; "flag")

func main() {
	flag.Parse()
	//targetFact := [some code here]
	//action := func(t Target) os.Error {
		//fmt.Println(t.Name())
	//}

	dag.Main(targetFact, action)
}
