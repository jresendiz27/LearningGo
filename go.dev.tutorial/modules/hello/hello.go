package main

import (
	"fmt"
	"log"
	"jresendiz27/sample"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names:= []string{"jresendiz", "jresendiz2", "jresendiz3"}

	messages, err:= sample.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}