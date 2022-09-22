// Adapted from https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
// Check https://pkg.go.dev/log for detail documentation of log package

package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	f, err := os.OpenFile("trace.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	// Multi outputs
	// multi := io.MultiWriter(f, os.Stdout)

	// ioutil.Discard can be used to disable logging
	Init(
		ioutil.Discard, // To nowhere
		f,              // To file
		os.Stdout,      // To stdout
		os.Stderr,      // To stderr
	)

	Trace.Println("You cannot see me!")
	Info.Println("Hello world!")
	Warning.Println("You might want to take a look at this?")
	Error.Println("Segmentation Fault :)?")
}
