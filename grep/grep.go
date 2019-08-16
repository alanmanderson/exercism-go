package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type parameters struct {
	printLineNumbers   bool
	printOnlyFileNames bool
	useCaseInsensitive bool
	invertSearch       bool
	matchEntireLine    bool
}

type result struct {
	file       string
	lineNumber int
	text       string
}

// Search takes a pattern, flags, and a list of files and returns a list of lines that match that pattern
func Search(pattern string, flags []string, files []string) []string {
	params := getParameters(flags)
	fmt.Printf("%v", params)
	results := make([]result, 0)
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("Error encountered")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNumber := 0
		for scanner.Scan() {
			line := scanner.Text()
			if isMatch(line, pattern) {
				results = append(results, result{fileName, lineNumber, line})
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file")
		}
	}
	out := getResultStrings(results)
	return out
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

func isMatch(text string, pattern string) bool {
	isMatch, _ := regexp.Match(pattern, []byte(pattern))
	return isMatch
}

func getResultStrings(in []result) (out []string) {
	for _, res := range in {
		out = append(out, res.text)
	}
	return
}
