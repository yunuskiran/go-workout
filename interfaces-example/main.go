package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type turkishBot struct{}

func main() {
	eb := englishBot{}
	tb := turkishBot{}

	printGreeting(eb)
	printGreeting(tb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (turkishBot) getGreeting() string {
	return "Merhaba"
}
