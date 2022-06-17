package main

import (
	"fmt"
	"github.com/idanieldrew/blog-golang/cmd"
	"github.com/idanieldrew/blog-golang/internal/config"
	"log"
)

var cfg = &config.Config{}

func init() {
	const path = "build/config/config.yml"
	if err := config.Read(path, cfg); err != nil {
		log.Fatalln("Error in parse file")
	}
}

func main() {
	if err := cmd.Serve(cfg); err != nil {
		log.Fatalln(err)
	}
	fmt.Print("success")
}
