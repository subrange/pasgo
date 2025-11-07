// https://www.tutorialspoint.com/pascal/pascal_program_structure.htm
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s <file>.pas\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

// Unless you need to deal with advanced grammar,
// I recommend just implementing a simple recursive
// descent parser and not bothering with formalisms.

// 17+6 5*(3+7) -> 5*(10) -> 50

// shunting-yard algorithm
// https://www.freepascal.org/docs-html/ref/refch12.html
// https://tiarkrompf.github.io/notes/?/just-write-the-parser/
func evaluateExpression(expression string) (int, error) {
	expression = strings.ReplaceAll(expression, " ", "")

	// var result int
	// switch operator {
	// case "div":
	// 	result = a / b
	// case "mod":
	// 	result
	// }

	return 0, nil
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Unable to open %s.\n", fileName)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0

	var programName string
	var block bool = false

	for scanner.Scan() {
		currentLine := scanner.Text()
		currentLine = strings.TrimSpace(currentLine)
		// fmt.Printf("Line %d: %s\n", i, currentLine)

		// line 1
		if strings.HasPrefix(currentLine, "program") {
			fields := strings.Fields(currentLine)
			if len(fields) == 2 {
				programName = strings.TrimSuffix(fields[1], ";")
				fmt.Println("ProgramName:", programName)
			} else {
				fmt.Printf("ProgramName: Error! %s\n", currentLine)
				return
			}
		} else if currentLine == "begin" {
			block = true
		} else if strings.Contains(currentLine, "writeln") {
			re := regexp.MustCompile(`\((.*?)\)`)
			match := re.FindStringSubmatch(currentLine)

			// TODO: evaluating the argument and ensuring that it handles things like variables, expressions, and newlines—is the next step.
			if len(match) > 1 {
				arg := match[1]
				// fmt.Println(arg) //[1 : len(arg)-1])

				result, err := evaluateExpression(arg)
				if err != nil {
					fmt.Println("Error evaluating expression %s\n", err)
					return
				}
				fmt.Println(result)
				// return
			} else {
				fmt.Printf("WriteLn: Error! %s\n", currentLine)
				return
			}
		} else if currentLine == "end." {
			if !block {
				fmt.Println("No begin: Error!")
				return
			}
			block = false
			break
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner: Error!")
	}
}
