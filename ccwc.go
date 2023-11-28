package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getWordsInLine(line string) int {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)

	wordsCount := 0

	for scanner.Scan() {
		wordsCount++
	}

	return wordsCount
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}



func main () {
	bytes := flag.String("c", "", "Cannot provide the size of a file without a filename ")
	lines := flag.String("l", "", "Cannot provide a line count without a filename")
	words := flag.String("w", "", "Cannot provide a word count without a filename")
	letters := flag.String("m", "", "Cannot provide a letter count without a filename")


	// Parse command-line arguments
	var filename string
	var command string 

	if !isInputFromPipe() {
		flag.Parse()


		if len(*bytes) > 0 {
			filename = *bytes
		} else if len(*lines) > 0 {
			filename = *lines
		} else if len(*words) > 0 {
			filename = *words
		} else if len(*letters) > 0 {
			filename = *letters
		} else {
			if len(os.Args) > 1 { 
				filename = os.Args[1]
			} else {
				fmt.Println("Please provide a filename")
				os.Exit(1)
			}
		}
	} else {
		command = os.Args[1]
	}


	var file *os.File

	if filename != "" {
		var err error 
		file, err = os.Open(filename)


		if err != nil {
			fmt.Printf("Error opening file: %s", filename)
			os.Exit(1)
		}
	} else {
		file = os.Stdin
	}
		
	

	if len(*bytes) > 0 || command == "-c" {
		fileInfo, err := file.Stat()

		if err != nil {
			fmt.Printf("Error getting file information for %s", filename)
		}

		fileSize := fileInfo.Size()

		fmt.Printf("   %d %s", fileSize, filename)


		os.Exit(1)
	} 

	fileScanner := bufio.NewScanner(file)

	if len(*lines) > 0 || command == "-l" {
		linesCount := 0 
		for fileScanner.Scan() {
			linesCount++
		}

		fmt.Printf("   %d %s", linesCount, filename)
	} else if len(*words) > 0 || command == "-w" {
		wordsCount := 0

		fileScanner.Split(bufio.ScanWords)

		for fileScanner.Scan() {
			wordsCount++
		}

		fmt.Printf("   %d %s", wordsCount, filename)
	} else if len(*letters) > 0 || command == "-m" {
		bytesCount := 0

		fileScanner.Split(bufio.ScanRunes)
		
		for fileScanner.Scan() {
			bytesCount++ 
		}

		fmt.Printf("   %d %s", bytesCount, filename)
	} else {
		linesCount := 0 
		wordsCount := 0 
		bytesCount := 0 

		for fileScanner.Scan() {
			line := fileScanner.Text()

			linesCount++

			wordsCount += getWordsInLine(line)

			bytesCount += len(line) + 2
		}
		
		fmt.Printf("   %d   %d   %d %s", linesCount, wordsCount, bytesCount, filename)
	}
}