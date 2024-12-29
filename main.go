package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseArgs() string {
	flag.StringVar(&fileName, "file", "default", "The file to be used")
	flag.Parse()
	if fileName == "default" {
		fmt.Println("Error: no file was provided.")
		os.Exit(1)
	}
	return fileName

}

func readFile(fileName string) (*os.File, *bufio.Scanner) {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	return readFile, fileScanner
}

func startCodeBlockFunc(line string, startCodeBlock *regexp.Regexp, codeBlock bool) bool {
	// ran when codeBlock == false
	if startCodeBlock.MatchString(line) && !codeBlock {
		fmt.Println("start of code block with language:", startCodeBlock.FindStringSubmatch(line)[1])
		return true
	}
	return codeBlock

}

func endCodeBlockFunc(line string, endCodeBlock *regexp.Regexp, codeBlock bool) bool {
	if endCodeBlock.MatchString(line) && codeBlock {
		fmt.Println("end of code block")
		// fmt.Println(line)
		return false
	} else {
		return codeBlock
	}
}

func extractHeadings(noteName string, line string, regex *regexp.Regexp) {
	lineNospace := regex.ReplaceAllString(line, "#")
	textToPrint = append(textToPrint, "[["+noteName+lineNospace+"|"+lineNospace+"]]")
}

func printContents(toc []string) {
	fmt.Println("==== Clickable Table of Contents below: ")
	for _, value := range textToPrint {
		fmt.Println(value)

	}
}

var (
	textToPrint []string
	noteName    string
	fileName    string
	codeBlock   bool
)

func main() {

	// Parse arguments
	fileName = parseArgs()
	// remove .md from MD files and use as the note name
	noteName := strings.Replace(fileName, ".md", "", -1)
	fmt.Println("Note Name:", noteName)

	// read file and bring backa  scanner
	file, fileScanner := readFile(fileName)
	defer file.Close()

	// split by line
	fileScanner.Split(bufio.ScanLines)

	headingRegex := "^#+\\s"
	startRegex := "^```(\\w+)$" // Matches the start with a language specifier
	endRegex := "^```$"         // Matches the end with just triple backticks

	headingCheck := regexp.MustCompile(headingRegex) // matches headings (only those with space afterwaterds)
	startCodeBlock := regexp.MustCompile(startRegex)
	endCodeBlock := regexp.MustCompile(endRegex)

	fmt.Println("scanner")
	for fileScanner.Scan() {
		// fmt.Println(codeBlock)
		line := fileScanner.Text()

		codeBlock = startCodeBlockFunc(line, startCodeBlock, codeBlock)
		codeBlock = endCodeBlockFunc(line, endCodeBlock, codeBlock)

		if !codeBlock && headingCheck.MatchString(line) {
			extractHeadings(noteName, line, headingCheck)
		}
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Now the slice is built, its printed
	printContents(textToPrint)

}
