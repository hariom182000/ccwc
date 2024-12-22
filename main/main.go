package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	lineFlag := flag.Bool("l", false, "count line")
	wordFlag := flag.Bool("w", false, "word count")
	byteFlag := flag.Bool("c", false, "byte count")
	charFlag := flag.Bool("m", false, "char count")

	var lineCount int64 = 0
	var byteCount int64 = 0
	var charCount int64 = 0
	var wordCount int64 = 0

	flag.Parse()

	filePath := ""
	if len(flag.Args()) > 0 {
		filePath = flag.Args()[0]
	}
	fmt.Printf("file path is %s \n", filePath)

	if *lineFlag == false && *wordFlag == false && *charFlag == false && *byteFlag == false {
		*lineFlag = true
		*byteFlag = true
		*wordFlag = true

	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lastSpace := false
	for {
		c, size, err := reader.ReadRune()
		if err != nil {
			break
		}
		byteCount += int64(size)
		charCount++
		if c == '\n' {
			if lastSpace == false {
				wordCount++
			}
			lineCount++
			continue
		}
		if unicode.IsSpace(c) {
			if lastSpace == false {
				wordCount++
			}
			lastSpace = true
		} else {
			lastSpace = false
		}
	}

	var cols []string
	if *lineFlag {
		cols = append(cols, strconv.FormatInt(lineCount, 10))
	}

	if *wordFlag {
		cols = append(cols, strconv.FormatInt(wordCount, 10))
	}

	if *charFlag {
		cols = append(cols, strconv.FormatInt(charCount, 10))
	}

	if *byteFlag {
		cols = append(cols, strconv.FormatInt(byteCount, 10))
	}

	cols = append(cols, filePath)
	fmt.Println(strings.Join(cols, " "))

}
