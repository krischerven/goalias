// this file is a work in progress
package main

import (
	"fmt"
	"github.com/krischerven/goalias/src/util/files"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

const (
	// the place where we keep track of our aliases
	registry = "/etc/goalias/registry.txt"
)

var (
	verbose = false
)

func main() {
	// first arg is filename
	os.Args = os.Args[1:]
	// if the second arg is --v, show all output)
	if len(os.Args) > 0 && os.Args[0] == "--v" {
		verbose = true
		os.Args = os.Args[1:]
	}
	// help, --help are the same thing
	// both are mutually exclusive with any other arguments
	if len(os.Args) > 0 && (os.Args[0] == "help" || os.Args[0] == "--help") {
		os.Args = os.Args[0:1]
	}
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
			handle(err, goerr)
		} else {
			fmt.Printf("Created registry at %s\n", registry)
		}
	} else if verbose {
		fmt.Printf("Registry already exists at %s\n", registry)
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
			fmt.Printf("Heap Size: %d KiB\n", ms.HeapAlloc/1024)
			fmt.Println("# Live Objects:", ms.HeapObjects)
			fmt.Println("# Allocs:", ms.Mallocs)
			fmt.Println("# Frees:", ms.Frees)
		default:
			fmt.Println("Usage: 'goalias [--v?] [help|debug]' or 'goalias [--v?] [set|unset|check] [aliasname] [alias?]'")
		}
	case 2:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			file := files.Bin(os.Args[1])
			s, err := files.Read(file)
			handle(err, goerr)
			fmt.Printf("# %s\n%s\n", file, s)
		case "set":
			fmt.Println("Error: goalias set takes exactly two arguments (1 provided)")
		case "unset":
			unregister(os.Args[1])
			unimplemented(1)
		default:
			fmt.Printf("Error: unrecognized argument '%s'\n", os.Args[0])
		}
	case 3:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			fmt.Println("Error: goalias check takes exactly one argument (2 provided)")
		case "set":
			register(os.Args[1], os.Args[2])
			unimplemented(2)
		case "unset":
			fmt.Println("Error: goalias unset takes exactly one argument (2 provided)")
		default:
			fmt.Printf("Error: unrecognized argument '%s'\n", os.Args[0])
		}
	default:
		fmt.Printf("Error: bad number of arguments (%d provided)\n", l)
	}
}

func unimplemented(i uint) {
	panic(fmt.Sprintf("unimplemented (%d)", i))
}

func goerr(e error) string {
	return fmt.Sprintf("Go Error: '%s'", e)
}

func handle(e error, handler func(error) string) {
	if e != nil {
		fmt.Println(handler(e))
		os.Exit(0)
	}
}

func register(name string, alias string) {
	r, err := files.Read(registry)
	handle(err, goerr)
	ioutil.WriteFile(registry, append([]byte(r), []byte(fmt.Sprintf("%s=%s\n", name, alias))...), 0755)
}

func unregister(string) {
	// TBI
}
