package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Now enough arguments!")
	}

	fmt.Println("Thanks for the arguments(s)!")
}
