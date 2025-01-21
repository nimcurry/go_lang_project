package main

import (
	"fmt"
	"time"
)

func greet(msg string) {
	fmt.Println(msg)
}

func slowGreet(msg string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(msg)
	doneChan <- true
}

func main() {
	//go keyword or go routines helps in running the functions or statements in parallel aka concurrency
	//go greet("Hello how are you?")
	//go greet("Hello why not your?")
	done := make(chan bool)
	go slowGreet("Hchal nikal?", done)
	go greet("Hteri maa ki")
	<-done
	//fmt.Println(<-done)
}
