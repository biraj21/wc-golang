package main

import (
	"fmt"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s [OPTION]... FILE...", os.Args[0])
	}

	options, filenames := parseArgs(os.Args[1:])

	for _, filename := range filenames {
		numLines, numWords, numBytes, numChars := GetStats(filename)

		if len(options) == 0 {
			fmt.Print(numLines, numWords, numBytes, " ")
		} else {
			if options["-l"] {
				fmt.Print(numLines, " ")
			}

			if options["-w"] {
				fmt.Print(numWords, " ")
			}

			if options["-m"] {
				fmt.Print(numChars, " ")
			}

			if options["-c"] {
				fmt.Print(numBytes, " ")
			}
		}

		fmt.Println(filename)
	}

}

func parseArgs(args []string) (map[string]bool, []string) {
	options := make(map[string]bool)
	filenames := []string{}

	for _, arg := range args {
		if arg[0] == '-' {
			options[arg] = true
		} else {
			filenames = append(filenames, arg)
		}
	}

	return options, filenames
}

func GetStats(filename string) (int64, int64, int64, int64) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var numLines, numWords, numBytes, numChars int64 = 0, 0, 0, int64(utf8.RuneCount(bytes))
	inWord := false

	for _, b := range bytes {
		if b == '\n' {
			numLines++
		}

		if unicode.IsSpace(rune(b)) {
			if inWord {
				numWords++
			}

			inWord = false
		} else {
			inWord = true
		}

		numBytes++
	}

	return numLines, numWords, numBytes, numChars
}
