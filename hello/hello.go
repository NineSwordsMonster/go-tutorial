package main

import (
	"fmt"
	"github.com/NineSwordsMonster/go-tutorial/greetings"
	"log"
)
import "rsc.io/quote"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings:")
	log.SetFlags(0)
	message, err := greetings.Hello("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
