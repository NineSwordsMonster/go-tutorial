package main

import (
	"fmt"
	"github.com/NineSwordsMonster/go-tutorial/greetings"
)
import "rsc.io/quote"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(quote.Go())
	message := greetings.Hello("Wang")
	fmt.Println(message)
}
