package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

var message string

func updateMessage(data string) {
	defer waitGroup.Done( )

	message= data
}

func main( ) {
	message= "initialized"

	waitGroup.Add(2)

	go updateMessage("count 1")
	go updateMessage("count 2")

	waitGroup.Wait( )

	fmt.Println(message)
}