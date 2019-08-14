package grep

import (
	//	"os"
	"fmt"
)

type parameters struct {
	printLineNumbers   bool
	printOnlyFileNames bool
	useCaseInsensitive bool
	invertSearch       bool
	matchEntireLine    bool
}

// Search takes a pattern, flags, and a list of files and returns a list of lines that match that pattern
func Search(pattern string, flags []string, files []string) []string {
	params := getParameters(flags)
	fmt.Printf("%v", params)
	//for file := range files {
	//}
	return []string{}
}

func getParameters(flags []string) parameters {
	p := parameters{}
	for _, flag := range flags {
		switch flag {
		case "-n":
			p.printLineNumbers = true
		case "-l":
			p.printOnlyFileNames = true
		case "-i":
			p.useCaseInsensitive = true
		case "-v":
			p.invertSearch = true
		case "-x":
			p.matchEntireLine = true
		}
	}
	return p
}
