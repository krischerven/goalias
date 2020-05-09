// this file is a work in progress
package main

import (
	"bytes"
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
	// for 'goalias set', we want to merge all args >= 2 into the second argument
	if len(os.Args) > 0 && os.Args[0] == "set" && len(os.Args) > 2 {
		for i := 3; i < len(os.Args); i++ {
			os.Args[2] += " "
			os.Args[2] += os.Args[i]
		}
		os.Args = os.Args[0:3]
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
			var ms runtime.MemStats
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
			b, err := files.Read(file)
			handle(err, goerr)
			fmt.Printf("# %s\n%s\n", file, b)
		case "set":
			fmt.Println("Error: goalias set takes exactly two arguments (1 provided)")
		case "unset":
			mustroot()
			if registered(os.Args[1]) {
				unregister(os.Args[1])
				handle(os.Remove(fmt.Sprintf("/usr/local/bin/%s", os.Args[1])), goerr)
			} else {
				fmt.Println("Error: this alias is not registered.")
			}
		default:
			fmt.Printf("Error: unrecognized argument '%s'\n", os.Args[0])
		}
	case 3:
		switch strings.ToLower(os.Args[0]) {
		case "check":
			fmt.Println("Error: goalias check takes exactly one argument (2 provided)")
		case "set":
			if registered(os.Args[1]) {
				fmt.Println("Error: an alias with this name already exists.")
			} else if files.UsrLocalBinExists(os.Args[1]) {
				fmt.Println("Error: a file with this name already exists.")
			} else {
				mustroot()
				register(os.Args[1], os.Args[2])
				handle(
					ioutil.WriteFile(
						fmt.Sprintf("/usr/local/bin/%s", os.Args[1]),
						[]byte(fmt.Sprintf("%s", os.Args[2])), 0755,
					),
					goerr,
				)
				if strings.Contains(os.Args[2], "cd ") {
					fmt.Println("It seems that your alias contains 'cd'.\n" +
						"Currently to change directory you have to run the alias with 'source aliasname'.")
				}
			}
		case "unset":
			fmt.Println("Error: goalias unset takes exactly one argument (2 provided)")
		default:
			fmt.Printf("Error: unrecognized argument '%s'\n", os.Args[0])
		}
	default:
		fmt.Printf("Error: bad number of arguments (%d provided)\n", l)
	}
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

func mustroot() {
	if os.Geteuid() != 0 {
		fmt.Println("Error: You must be running as the root user to run this command.")
		os.Exit(0)
	}
}

func unimplemented(i uint) {
	panic(fmt.Sprintf("unimplemented (%d)", i))
}

func register(name string, alias string) {
	b, err := files.Read(registry)
	handle(err, goerr)
	ioutil.WriteFile(
		registry,
		append(b, []byte(fmt.Sprintf("%s=%s\n", name, alias))...),
		0755,
	)
}

func registered(name string) bool {
	b, err := files.Read(registry)
	handle(err, goerr)
	return bytes.Contains(b, []byte(fmt.Sprintf("%s=", name)))
}

func unregister(name string) {
	b, err := files.Read(registry)
	handle(err, goerr)
	lines := bytes.Split(b, []byte("\n"))
	for i, line := range lines {
		if bytes.HasPrefix(line, []byte(fmt.Sprintf("%s=", name))) {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}
	ioutil.WriteFile(
		registry,
		bytes.Join(lines, []byte("\n")),
		0755,
	)
}
