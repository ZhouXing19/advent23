package day2

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/advent23/util"
)

const (
	blue  = "blue"
	red   = "red"
	green = "green"
)

const gameIdPat = `Game (\d+)`
const attemptPat = `(\d+)\s([a-z]+)`

func splitRule(r rune) bool {
	return r == ':' || r == ';'
}

type result map[int64]map[string]int64

func (r result) Format() string {
	sb := strings.Builder{}
	for id, colorCnts := range r {
		sb.WriteString(fmt.Sprintf("id: %d", id))
		sb.WriteString("\t")
		for color, cnt := range colorCnts {
			sb.WriteString(fmt.Sprintf(" (%s: %d)", color, cnt))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func readInput(filePath string, cap int, rule func(exV int64, newV int64) (resV int64)) result {
	lines := util.GetLines(filePath)
	res := make(result)
	for i, l := range lines {
		if cap > 0 && i == cap {
			break
		}

		newMap := map[string]int64{
			blue:  0,
			red:   0,
			green: 0,
		}

		splitRes := strings.FieldsFunc(l, splitRule)
		if len(splitRes) < 2 {
			log.Fatalf("unknown fmt of line: %s", l)
		}

		// Get the game id.
		idStr := splitRes[0]
		gameIdReg := regexp.MustCompile(gameIdPat)
		gameIdList := gameIdReg.FindStringSubmatch(idStr)
		if len(gameIdList) < 2 {
			log.Fatalf("unknown fmt for game id str: %s", idStr)
		}
		gameId, err := strconv.ParseInt(gameIdList[1], 10, 64)
		if err != nil {
			panic(err)
		}
		res[gameId] = newMap

		for j := 1; j < len(splitRes); j++ {
			attpStr := splitRes[j]
			attpReg := regexp.MustCompile(attemptPat)

			attpMatch := attpReg.FindAllStringSubmatch(attpStr, -1)
			for _, cntColorGroup := range attpMatch {
				if len(cntColorGroup) != 3 {
					panic("group len not 3")
				}
				color := cntColorGroup[2]
				cnt, err := strconv.ParseInt(cntColorGroup[1], 10, 64)
				if err != nil {
					panic(err)
				}
				existingCnt, ok := newMap[color]
				if !ok {
					log.Fatalf("unknown color: %q", color)
				}
				newMap[color] = rule(existingCnt, cnt)
			}
		}
		res[gameId] = newMap
	}

	return res
}

func easy(f string, cap int) int64 {
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	var threshold = map[string]int64{
		red:   12,
		green: 13,
		blue:  14,
	}
	fmt.Printf("\nthreashold: (red:%d), (blue:%d), (green:%d)\n---\n", threshold[red], threshold[blue], threshold[green])
	games := readInput(f, cap, func(exV int64, newV int64) (resV int64) {
		if newV > exV {
			return newV
		}
		return exV
	})

	var idSum int64
	var validGame = make(result)
	for id, game := range games {
		if game[red] <= threshold[red] && game[green] <= threshold[green] && game[blue] <= threshold[blue] {
			idSum += id
			validGame[id] = game
		}
	}
	fmt.Println(validGame.Format())
	return idSum
}

func hard(f string, cap int) int64 {
	games := readInput(f, cap, func(exV int64, newV int64) (resV int64) {
		if newV > exV {
			return newV
		}
		return exV
	})

	var powerSum int64

	for _, game := range games {
		powerSum += game[red] * game[green] * game[blue]
	}

	return powerSum
}

const (
	easyTestPath = "./day2/easytest.txt"
	hardTestPath = "./day2/hardtest.txt"
	inputPath    = "./day2/input.txt"
)

func Day2() {
	// fmt.Printf("sum of id: %d\n", easy(inputPath, -1))
	fmt.Printf("powersum: %d\n", hard(inputPath, -1))
}
