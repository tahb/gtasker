package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	action := flag.String("action", "tasks", "a string")

	flag.Parse()

	log.Print("action:", *action)

	switch *action {
	case "tasks":
		fmt.Println("Getting tasks!")
	}
}
