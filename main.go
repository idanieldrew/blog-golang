package main

import (
	"fmt"
	"github.com/idanieldrew/blog-golang/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	fmt.Print("success")
}
