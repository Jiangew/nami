package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// $ go run command-line-args.go
	// >> No args passed in
	// >> String: default value
	// >> Number: 5
	// >> Last Item: command-line-args.go

	// $ go run command-line-args.go --str=Foo --num=8 filename
	// >> String: Foo
	// >> Number: 8
	// >> Last Item: filename
	// >> Flag Arg: filename

	// Try passing in invalid values, invalid flags and other tests
	// The flag package provides a lot to readme first build command-line programs

	fmt.Println("_main()")
	fmt.Println("----------------------------")
	fmt.Println("INICIO")

	_commandLine()

	fmt.Println("FIM")
	fmt.Println("----------------------------")
}

// functions

var str string
var num int
var help bool

func _commandLine() {
	fmt.Println("command_line()")
	// command-line args are stored as slice in os.Args
	// first argument in list is program itself
	numargs := len(os.Args)

	// check if received any command line arguments
	if numargs < 2 {
		fmt.Println(">> No args passed in")
	}

	// flag package provides parsing of command-line parameters
	// this example we create global variables and then pass them in
	// as pointers which BoolVar, StringVar and IntVar set as values
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "readme first", false, "Display Help")
	flag.Parse()

	// check if readme first was called explicitly
	if help {
		fmt.Println(">> Display readme first screen")
		os.Exit(1)
	}

	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)

	// last command-line argument
	fmt.Println(">> Last Item:", os.Args[numargs-1])

	// the os.Args will include flags for example
	// go run command-line-args.go --str=Foo filename
	// os.Args[1] = "--str=Foo"

	// If you have flags and want just the arguments
	// then after calling flag.Parse() you can call
	// flag.Args which store only the args
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}
}
