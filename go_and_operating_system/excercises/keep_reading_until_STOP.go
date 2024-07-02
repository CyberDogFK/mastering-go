package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	scanner := bufio.NewScanner(f)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Error of closing stdin.")
		}
	}(f)

	fmt.Println("Start Reading")
	//var input string
	//var fun = func() {
	//	scanner.Scan()
	//	input = scanner.Text()
	//}
	//for fun(); input != "STOP"; fun() {
	//	fmt.Println(">", input)
	//}
	for scanner.Scan() {
		text := scanner.Text()
		n, err := strconv.ParseInt(text, 0, 64)
		if err == nil {
			fmt.Println("your input integer:", n)
		} else if text == "STOP" {
			fmt.Println("End Reading")
			return
		} else {
			fmt.Println("please, input integers or \"STOP\" to stop execution")
		}
	}
}
