package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrint(testStateManager *testing.T) {

	reader, writer, _ := os.Pipe( )

	standardOutput := os.Stdout
	os.Stdout= writer

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go Print("test", &waitGroup)

	waitGroup.Wait( )

	writer.Close( )
	testFunctionLogs, _ := io.ReadAll(reader)

	if(! strings.Contains(string(testFunctionLogs), "test")) {
		testStateManager.Error("error in Print( )")
	}

	os.Stdout= standardOutput
}