// this file is a work in progress
package main

import (
	"fmt"
	"github.com/krischerven/goalias/src/util/files"
	"os"
	"runtime"
	"strings"
)

const (
	// the place where we keep track of our aliases
	registry = "/etc/goalias/registry.txt"
)

func main() {
	// first arg is filename
	os.Args = os.Args[1:]
	// create registry files
	if err := os.MkdirAll(files.Dir(registry), 0755); err != nil {
		fmt.Print(goerr(err))
		if strings.HasSuffix(err.Error(), "permission denied") {
			fmt.Print(" (try running goalias with sudo the first time)")
		}
		fmt.Println()
		os.Exit(0)
	} else if !files.Exists(registry) {
		if _, err := os.Create(registry); err != nil {
			fmt.Println(goerr(err))
			os.Exit(0)
		} else {
			fmt.Println(fmt.Sprintf("Created registry at %s", registry))
		}
	} else {
		fmt.Println(fmt.Sprintf("Registry already exists at %s", registry))
	}
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
		case "debug":
			ms := runtime.MemStats{}
			runtime.ReadMemStats(&ms)
			fmt.Printf("Debugger:\n\n")
			fmt.Println(fmt.Sprintf("Heap Size: %d KiB", ms.HeapAlloc/1024))
			fmt.Println("# Live Objects:", ms.HeapObjects)
			fmt.Println("# Allocs:", ms.Mallocs)
			fmt.Println("# Frees:", ms.Frees)
		default:
			fmt.Println("Usage: 'goalias [help|debug]' or 'goalias [set|unset|check] [aliasname] [alias?]'")
		}
	case 2:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			file := files.Bin(os.Args[1])
			s, err := files.Read(file)
			if err != nil {
				fmt.Println(goerr(err))
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

func goerr(e error) string {
	return fmt.Sprintf("Go Error: '%s'", e)
}
