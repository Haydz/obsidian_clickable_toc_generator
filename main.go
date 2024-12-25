package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "default", "The file to be used")
	flag.Parse()
	if fileName == "default" {
		fmt.Println("Error: no file was provided.")
		os.Exit(1)
	}
	// open file
	var textToPrint []string
	readFile, err := os.Open(fileName)
	// remove .md from MD files and use as the note name
	noteName := strings.Replace(fileName, ".md", "", -1)
	if err != nil {
		fmt.Println(err)
	}
	// read file line by int
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	// regex for identifying the space before #
	re := regexp.MustCompile("^#+\\s")

	codeBlock := false
	startRegex := "^```(\\w+)$" // Matches the start with a language specifier
	endRegex := "^```$"         // Matches the end with just triple backticks

	startCodeBlock := regexp.MustCompile(startRegex)
	endCodeBlock := regexp.MustCompile(endRegex)

	// fmt.Println(regexCode)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		// Contain 1 or more "#" to denote a heading
		//check for back ticks

		if startCodeBlock.MatchString(line) && !codeBlock {
			fmt.Println("start of code block with language:", startCodeBlock.FindStringSubmatch(line)[1])
			fmt.Println(line)
			codeBlock = true
			continue

		}

		if endCodeBlock.MatchString(line) && codeBlock {
			fmt.Println("end of code block")
			fmt.Println(line)
			codeBlock = false
			continue

		}
		if strings.Contains(line, "#") && !codeBlock {
			lineNospace := re.ReplaceAllString(line, "#")
			textToPrint = append(textToPrint, "[["+noteName+lineNospace+"|"+lineNospace+"]]")

		}
	}

	readFile.Close()
	for _, value := range textToPrint {
		fmt.Println(value)
	}

}
