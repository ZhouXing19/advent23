package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/advent23/util"
)

const (
	easyTestPath = "./day3/easytest.txt"
	hardTestPath = "./day3/hardtest.txt"
	inputPath    = "./day3/input.txt"
)

func Day3() {
	res := hard(inputPath, -1)
	fmt.Printf("res: %d\n", res)
}

func getWidth(input []string) int {
	var width int
	for _, i := range input {
		if width == 0 {
			width = len(i)
			continue
		}
		if len(i) != width {
			panic(fmt.Sprintf("length of line %q is not %d", i, width))
		}
	}
	fmt.Printf("passed dim check, width is %d\n", width)
	return width
}

func markNumber(i int, j int, starLoc string, m [][]*matrixElem) {
	numRows := len(m)
	numCols := len(m[0])
	if i < 0 || i >= numRows || j < 0 || j >= numCols {
		fmt.Printf("(%d, %d) out of boundary\n", i, j)
		return
	}

	if m[i][j] != nil {
		m[i][j].marked = true
		if m[i][j].starLocs == nil {
			m[i][j].starLocs = make([]string, 0)
		}
		m[i][j].starLocs = append(m[i][j].starLocs, starLoc)
	}
}

func markSurrounding(i int, j int, m [][]*matrixElem) {
	for _, di := range []int{-1, 0, 1} {
		for _, dj := range []int{-1, 0, 1} {
			ni := i + di
			nj := j + dj
			markNumber(ni, nj, locToStr(i, j), m)
		}
	}
}

type matrixElem struct {
	marked   bool
	starLocs []string
}

var allSymbolsMap map[rune]int

func buildMatrix(lines []string) [][]*matrixElem {
	width := getWidth(lines)

	// make a recorder matrix to check if the number is marked.
	matrix := make([][]*matrixElem, len(lines))
	for i := range matrix {
		matrix[i] = make([]*matrixElem, width)
	}

	// Go through the input strings.
	for i, l := range lines {
		for j, c := range l {
			if unicode.IsNumber(c) {
				matrix[i][j] = &matrixElem{}
			} else {
				if _, ok := allSymbolsMap[c]; !ok {
					allSymbolsMap[c] = 1
				}
			}
		}
	}

	tmpSb := strings.Builder{}
	for k := range allSymbolsMap {
		tmpSb.WriteRune(k)
	}
	fmt.Printf("all symbols: %q\n", tmpSb.String())

	return matrix
}

func collectPartNums(lines []string, matrix [][]*matrixElem) []string {
	resStrSlice := make([]string, 0)
	nonPartNumSlice := make([]string, 0)

	sb := strings.Builder{}
	isPartNum := false
	for i, l := range lines {
		for j, c := range l {
			if !unicode.IsNumber(c) {
				sbStr := sb.String()
				if len(sbStr) != 0 {
					if isPartNum {
						resStrSlice = append(resStrSlice, sbStr)
					} else {
						nonPartNumSlice = append(nonPartNumSlice, fmt.Sprintf("(%d, %d)%q", i, j, sbStr))
					}

				}
				sb = strings.Builder{}
				isPartNum = false
				continue
			}
			if _, err := sb.WriteRune(c); err != nil {
				panic(err)
			}
			if matrix[i][j] != nil && matrix[i][j].marked == true {
				isPartNum = true
			}
		}
		sbStr := sb.String()
		if len(sbStr) != 0 {
			if isPartNum {
				resStrSlice = append(resStrSlice, sbStr)
			} else {
				nonPartNumSlice = append(nonPartNumSlice, fmt.Sprintf("(%d, %d)%q", i, len(matrix[0])-1, sbStr))
			}
		}
		sb = strings.Builder{}
		isPartNum = false
	}

	// fmt.Printf("part numbers: %s\n", resStrSlice)
	// fmt.Printf("non part numbers: %s\n", nonPartNumSlice)

	return resStrSlice
}

func easy(f string, cap int) int64 {
	allSymbolsMap = make(map[rune]int)

	lines := util.GetLines(f)
	if len(lines) == 0 {
		panic("non input")
	}
	// make a recorder matrix to check if the number is marked.
	matrix := buildMatrix(lines)

	for i, l := range lines {
		for j, c := range l {
			if !unicode.IsNumber(c) && c != '.' {
				markSurrounding(i, j, matrix)
			}
		}
	}

	resStrSlice := collectPartNums(lines, matrix)
	return sumFromStrSlice(resStrSlice)
}

func sumFromStrSlice(strs []string) int64 {
	var res int64
	for _, s := range strs {
		if parsedInt, err := strconv.ParseInt(s, 10, 64); err != nil {
			panic(fmt.Sprintf("unable for parse %s to int", s))
		} else {
			res += parsedInt
		}
	}
	return res
}

func printStarAdjStrMap(starAdjStrMap map[string][]string) {
	fmt.Println("--------starAdjStrMap-------")
	for starLoc, v := range starAdjStrMap {
		fmt.Printf("(%s): %s\n", starLoc, v)
	}
	fmt.Println("---------------")
}

// starAdjStrMap: {location of star} -> [all digit strings surrounding it]
func collectGears(lines []string, matrix [][]*matrixElem, starAdjStrMap map[string][]string) {
	isPartNum := false

	sb := strings.Builder{}
	// Locations of all stars adjacent to the current digit string.
	starLocs := make([]string, 0)

	for i, l := range lines {
		for j, c := range l {
			// End of a digit string, summarize the info of this digit string.
			if !unicode.IsNumber(c) {
				sbStr := sb.String()
				// If this is a valid digit string, and is adjacent to a star.
				if len(sbStr) != 0 && isPartNum {
					for _, starLoc := range starLocs {
						starAdjStrMap[starLoc] = append(starAdjStrMap[starLoc], sbStr)
					}
				}
				sb = strings.Builder{}
				isPartNum = false
				starLocs = make([]string, 0)
				continue
			}
			if _, err := sb.WriteRune(c); err != nil {
				panic(err)
			}
			if matrix[i][j] != nil && matrix[i][j].marked == true {
				isPartNum = true
				starLocs = matrix[i][j].starLocs
			}
		}
		sbStr := sb.String()
		if len(sbStr) != 0 && isPartNum {
			for _, starLoc := range starLocs {
				starAdjStrMap[starLoc] = append(starAdjStrMap[starLoc], sbStr)
			}
		}
		sb = strings.Builder{}
		isPartNum = false
		starLocs = make([]string, 0)
	}
	printStarAdjStrMap(starAdjStrMap)
}

func locToStr(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func sumFromStarMap(starMap map[string][]string) int64 {
	var res int64
	for _, v := range starMap {
		if len(v) == 2 {
			// fmt.Printf("(%s): %s\n", starLoc, v)
			v1, err := strconv.ParseInt(v[0], 10, 64)
			if err != nil {
				panic(err)
			}
			v2, err := strconv.ParseInt(v[1], 10, 64)
			if err != nil {
				panic(err)
			}
			res += v1 * v2
		}
	}
	return res
}

func hard(f string, cap int) int64 {
	allSymbolsMap = make(map[rune]int)

	lines := util.GetLines(f)
	if len(lines) == 0 {
		panic("non input")
	}
	// make a recorder matrix to check if the number is marked.
	matrix := buildMatrix(lines)

	starAdjStrMap := make(map[string][]string)

	for i, l := range lines {
		for j, c := range l {
			if c == '*' {
				markSurrounding(i, j, matrix)
				starAdjStrMap[locToStr(i, j)] = make([]string, 0)
			}
		}
	}

	collectGears(lines, matrix, starAdjStrMap)

	return sumFromStarMap(starAdjStrMap)
}
