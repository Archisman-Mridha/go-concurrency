package main

import (
	"fmt"
	"sync"
)

func Print(data string, waitGroup *sync.WaitGroup) {
	if(waitGroup != nil) { defer waitGroup.Done( ) }

	fmt.Println(data)
}

func main( ) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go Print("count 1", &waitGroup)

	Print("count 2", nil)

	waitGroup.Wait( )
}