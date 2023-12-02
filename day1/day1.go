package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unicode"
)

var digitStrList = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func processEachLineEasy(line string) int64 {
	if len(line) == 0 {
		return 0
	}
	var firstDigit int
	var lastDigit int
	for _, b := range line {
		if unicode.IsDigit(b) {
			firstDigit = int(b - '0')
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			lastDigit = int(rune(line[i]) - '0')
			break
		}
	}
	sum := firstDigit*10 + lastDigit
	fmt.Printf("line:%s\tfirst digit:%d\tlast digit:%d\tsum:%d\n", line, firstDigit, lastDigit, sum)
	return int64(sum)
}

func processEachLineHard(line string) int64 {
	if len(line) == 0 {
		return 0
	}

	type digitStrWithIndex struct {
		v     int
		index int
	}
	// Check if there's any string rep of dig in the string.
	var firstDigitStrIndex = &digitStrWithIndex{index: math.MaxInt64, v: -1}
	var lastDigitStrIndex = &digitStrWithIndex{index: -1, v: -1}

	for i, digitStrList := range digitStrList {
		if idx := strings.Index(line, digitStrList); idx != -1 {
			if idx < firstDigitStrIndex.index {
				firstDigitStrIndex.v = i + 1
				firstDigitStrIndex.index = idx
			}
		}
		if lastIdx := strings.LastIndex(line, digitStrList); lastIdx != -1 {
			if lastIdx > lastDigitStrIndex.index {
				lastDigitStrIndex.v = i + 1
				lastDigitStrIndex.index = lastIdx
			}
		}
	}

	for idx, b := range line {
		if unicode.IsDigit(b) {
			if idx < firstDigitStrIndex.index {
				firstDigitStrIndex.v = int(b - '0')
				firstDigitStrIndex.index = idx
			}
			break
		}
	}

	for idx := len(line) - 1; idx >= 0; idx-- {
		if unicode.IsDigit(rune(line[idx])) {
			if idx > lastDigitStrIndex.index {
				lastDigitStrIndex.v = int(rune(line[idx]) - '0')
				lastDigitStrIndex.index = idx
			}
			break
		}
	}

	var sum int

	if firstDigitStrIndex.v != -1 {
		sum = firstDigitStrIndex.v*10 + lastDigitStrIndex.v
	}

	fmt.Printf("line:%s\tfirst digit:%d[%d]\tlast digit:%d[%d]\tsum:%d\n", line, firstDigitStrIndex.v, firstDigitStrIndex.index, lastDigitStrIndex.v, lastDigitStrIndex.index, sum)
	return int64(sum)
}

const (
	easyTestPath = "./day1/easytest.txt"
	hardTestPath = "./day1/hardtest.txt"
	inputPath    = "./day1/input.txt"
)

func Day1() {
	file, err := os.Open(inputPath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	cnt := 0
	var res int64
	for scanner.Scan() {
		lineStr := scanner.Text()
		res += processEachLineHard(strings.ToLower(lineStr))
		cnt++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("res:", res)
}
