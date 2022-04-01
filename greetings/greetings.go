package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("hi, %v, Welcom!", name)
	return message
}
