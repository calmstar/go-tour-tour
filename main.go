package main

import (
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}

	//s := "abcdefg 我"
	//for i, r := range s {
	//	fmt.Println(i, r)
	//}
}
