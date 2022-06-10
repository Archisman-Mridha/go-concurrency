# Concurrency in GoLang

Concurrency is the ability of an operating system to execute multiple instructions at the same instant. A process contains multiple lines of instructions. These instructions are grouped and multiple groups are created for a process. These groups can be called subprocesses or threads which can be executed without any dependency on each other.

Parent process
|
|- Subprocess A (Thread A)
|
|- Subprocess B (Thread B)

These threads are then executed at the same time by the cores of a multi-core processor.

## GoRoutines

In GoLang, a lightweight thread is called a go-routine. A group of go-routines is called co-routines. They are all managed by the go-scheduler. The main function in GoLang is a go-routine itself.

Let's see how we can create a go-routine :

```go
func print(data string) { fmt.Println(data) }

func main( ) {

    // creates a new go-routine
    go print("count 1")

    print("count 2")
}
```

We use the "go" keyword to create a go-routine. A thread is made using whatever code is there after the "go" keyword in that line. Now, if we execute this program, the output will be :

```log
count 2
```

But shouldn't "count 1" be also printed. Turns out that the main( ) function exited before our go-routine was created, assigned to the go-schedular and executed. The solution to this, is a concept called wait groups.

## Wait Groups

The job of a wait-group will be to stop the main( ) function from exiting, until all go-routines are finished executing.

```go
func Print(data string, waitGroup *sync.WaitGroup) {
	if(waitGroup != nil) {
    
        // mark the go-routine finished while exitting this function
        defer waitGroup.Done( )
    }

	fmt.Println(data)
}

func main( ) {

    // creating a wait-group
	var waitGroup sync.WaitGroup

    // registering the go-routine to the wait group
	waitGroup.Add(1)
	go Print("count 1", &waitGroup)

	Print("count 2", nil)

    // wait untill all the go-routines are executed
	waitGroup.Wait( )
}
```

WARNING - Don't call waitGroup.Done( ) if there are no active go-routines. This will give error.

## Concurrency for Tests in GoLang

Suppose we want to write a test for the Print( ) function defined above.

```go
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
```

The concept here is we will run the function which we want to test as a go-routine.

## Race Conditions