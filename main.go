// this file is a work in progress
package main

import (
	"fmt"
	"github.com/krischerven/goalias/src/util/files"
	"os"
	"strings"
)

func main() {
	os.Args = os.Args[1:]
	// help, --help are the same thing
	// both are mutually exclusive with any other arguments
	if len(os.Args) > 0 && (os.Args[0] == "help" || os.Args[0] == "--help") {
		os.Args = os.Args[0:1]
	}
	switch l := len(os.Args); l {
	case 0:
		fmt.Println("Error: no arguments provided (try 'goalias help')")
	case 1:
		switch strings.ToLower(os.Args[0]) {
		case "help", "--help":
			fmt.Printf("Usage:\n\n")
			fmt.Println("goalias [--]help")
			fmt.Println("goalias set [aliasname] [alias]")
			fmt.Println("goalias unset [aliasname]")
			fmt.Println("goalias check [aliasname]")
			fmt.Println()
			os.Exit(0)
		default:
			fmt.Println("Error: goalias requires at least two arguments")
		}
	case 2:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			file := files.Bin(os.Args[1])
			s, err := files.Read(file)
			if err != nil {
				fmt.Println(fmt.Sprintf("Go Error: '%s'", err))
			} else {
				fmt.Println(fmt.Sprintf("# %s\n", file) + s)
			}
		case "set":
			unimplemented(0)
		case "unset":
			unimplemented(1)
		default:
			fmt.Println(fmt.Sprintf("Error: unrecognized argument '%s'", os.Args[0]))
		}
	case 3:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			fmt.Println("Error: goalias check takes exactly one argument (2 provided)")
		case "set":
			unimplemented(2)
		case "unset":
			unimplemented(3)
		default:
			fmt.Println(fmt.Sprintf("Error: unrecognized argument '%s'", os.Args[0]))
		}
	default:
		fmt.Println(fmt.Sprintf("Error: bad number of arguments (%d provided)", l))
	}
}

func unimplemented(i uint) {
	panic(fmt.Sprintf("unimplemented (%d)", i))
}
