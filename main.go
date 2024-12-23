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
	flag.StringVar(&fileName, "file", "default", "help message for flag int")
	flag.Parse()
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
	re := regexp.MustCompile(`#\s+`)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// Contain 1 or more "#" to denote a heading
		if strings.Contains(line, "#") {
			// fmt.Println("found line", strings.Replace(line, " ", "", -1))
			// only want to remove the first space after # or ##
			lineNospace := re.ReplaceAllString(line, "#")
			textToPrint = append(textToPrint, "[["+noteName+lineNospace+"|"+lineNospace+"]]")

		}
	}

	readFile.Close()
	for _, value := range textToPrint {
		fmt.Println(value)
	}

}
