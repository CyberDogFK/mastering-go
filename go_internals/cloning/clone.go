package main

import "fmt"

type Original struct {
	clone *To
}

type To struct {
	i int8
}

func main() {
	var to To
	origin := Original{
		clone: &to,
	}
	changingOfOriginOfTo(origin)
	fmt.Println("Origin.clone:", origin.clone)
	fmt.Println("To:", to)
}

func changingOfOriginOfTo(origin Original) {
	origin.clone.i = 10
	fmt.Println("Origin.clone in func:", origin.clone)
}
