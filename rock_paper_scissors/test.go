package main

import("rps"; "os")

func main() {
	view := rps.NewRPSView(os.Stdout, "A")
	view.Enable()
}