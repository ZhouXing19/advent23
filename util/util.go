package util

import (
	"bufio"
	"os"
)

func GetLines(fPath string) []string {
	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	res := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		res = append(res, lineStr)
	}
	return res
}
