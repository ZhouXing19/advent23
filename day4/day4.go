package day4

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/advent23/util"
)

const (
	easyTestPath = "./day4/easytest.txt"
	hardTestPath = "./day4/hardtest.txt"
	inputPath    = "./day4/input.txt"
)

func Day4() {
	ps := getCardPairs(inputPath)
	res := getStretchNumber(ps)
	fmt.Println("res: ", res)
}

type winV struct {
	cnt     int
	matched bool
}

type cardPair struct {
	cardIdx     int64
	winningNums map[int64]*winV
	myNums      []int64
	matchedCnt  int
	stretchCnt  int64
}

func splitRule(r rune) bool {
	return r == ':' || r == '|'
}

func getCardPairs(fPath string) []*cardPair {
	lines := util.GetLines(fPath)
	res := make([]*cardPair, len(lines))
	for i, l := range lines {
		splitLine := strings.FieldsFunc(l, splitRule)
		if len(splitLine) < 3 {
			log.Fatalf("unknown fmt of line: %q", l)
		}

		curCardPair := &cardPair{
			winningNums: make(map[int64]*winV),
			myNums:      make([]int64, 0),
			stretchCnt:  1,
		}
		// Card index.
		cardIdxStr := splitLine[0]
		cardIdxStrSlice := strings.Fields(cardIdxStr)
		if len(cardIdxStrSlice) < 2 {
			panic("unknown fmt of cardIdxStr")
		}
		cardIdx, err := strconv.ParseInt(cardIdxStrSlice[1], 10, 16)
		if err != nil {
			panic("cannot parse for card index")
		}

		curCardPair.cardIdx = cardIdx

		// Winning numbers.
		winningNumsStr := splitLine[1]
		winNumsStrSlice := strings.Fields(winningNumsStr)
		for _, numStr := range winNumsStrSlice {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				panic(err)
			}
			if v, ok := curCardPair.winningNums[num]; !ok {
				curCardPair.winningNums[num] = &winV{
					cnt: 1,
				}
			} else {
				v.cnt++
			}
		}

		// My Numbers.
		myNumsStr := splitLine[2]
		myNumsStrSlice := strings.Fields(myNumsStr)
		for _, numStr := range myNumsStrSlice {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				panic(err)
			}
			curCardPair.myNums = append(curCardPair.myNums, num)
			if v, ok := curCardPair.winningNums[num]; ok {
				v.matched = true
				curCardPair.matchedCnt++
			}
		}

		sb := strings.Builder{}
		sb.WriteString("matched:( ")
		for k, v := range curCardPair.winningNums {
			if v.matched {
				sb.WriteString(strconv.FormatInt(k, 10))
				sb.WriteString(" ")
			}
		}
		sb.WriteString(")\n")
		fmt.Printf("[%s]%s", cardIdxStr, sb.String())
		res[i] = curCardPair
	}
	return res
}

func getPoints(ps []*cardPair) int64 {
	var res float64
	for _, p := range ps {
		if p.matchedCnt > 0 {
			res += math.Pow(float64(2), float64(p.matchedCnt-1))
		}
	}
	return int64(res)
}

func getStretchNumber(ps []*cardPair) int64 {
	curCardIdx := 0
	var res int64
	for curCardIdx < len(ps) {
		res += ps[curCardIdx].stretchCnt
		// Start from card 1.
		// Get the number of matches.
		curMatch := ps[curCardIdx].matchedCnt
		for i := 1; i < curMatch+1; i++ {
			copyCardIdx := curCardIdx + i
			if copyCardIdx < len(ps) {
				ps[copyCardIdx].stretchCnt += ps[curCardIdx].stretchCnt
			}
		}
		curCardIdx++
	}

	return res
}
