package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

type Stats struct {
	numLines, numWords, numBytes, numChars int64
}

func main() {
	// "usage: %s [OPTION]... FILE...", os.Args[0])

	options, filenames := parseArgs(os.Args[1:])
	if len(filenames) == 0 {
		bytes := readBytesFromStdin()
		printStats(getStatsFromBytes(bytes), "", options)
		return
	}

	for _, filename := range filenames {
		printStats(GetStats(filename), filename, options)
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

func readBytesFromStdin() []byte {
	bytes := []byte{}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytes = append(bytes, line...)
		if err == io.EOF {
			break
		}
	}

	return bytes
}

func printStats(stats Stats, filename string, options map[string]bool) {
	if len(options) == 0 {
		fmt.Print(stats.numLines, stats.numWords, stats.numBytes, " ")
	} else {
		if options["-l"] {
			fmt.Print(stats.numLines, " ")
		}

		if options["-w"] {
			fmt.Print(stats.numWords, " ")
		}

		if options["-m"] {
			fmt.Print(stats.numChars, " ")
		}

		if options["-c"] {
			fmt.Print(stats.numBytes, " ")
		}
	}

	fmt.Println(filename)
}

func GetStats(filename string) Stats {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return getStatsFromBytes(bytes)
}

func getStatsFromBytes(bytes []byte) Stats {
	stats := Stats{0, 0, 0, int64(utf8.RuneCount(bytes))}
	inWord := false

	for _, b := range bytes {
		if b == '\n' {
			stats.numLines++
		}

		if unicode.IsSpace(rune(b)) {
			if inWord {
				stats.numWords++
			}

			inWord = false
		} else {
			inWord = true
		}

		stats.numBytes++
	}

	if inWord {
		stats.numWords++
	}

	return stats
}
