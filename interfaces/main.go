package main

import "fmt"

// NOTE: 5:50 - 54.Interfaces in Practice
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi, there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
