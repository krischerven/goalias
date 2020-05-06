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
	switch l := len(os.Args); l {
	case 0:
		fmt.Println("Error: no first argument (try 'set', 'unset', or 'check')")
	case 1:
		fmt.Println("Error: goalias requires at least two arguments")
	case 2:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			s, err := files.Read(os.Args[1])
			if err != nil {
				fmt.Println(fmt.Sprintf("Go Error: '%s'", err))
			} else {
				fmt.Println(fmt.Sprintf("# %s\n", os.Args[1]) + s)
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
