package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myDate string

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n",
			filepath.Base(os.Args[0]))
		return
	}
	myDate = os.Args[1]

	d, err := time.Parse("02 January 2006", myDate)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}
}
