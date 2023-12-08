package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/advent23/util"
)

const (
	easyTestPath    = "./day6/easytest.txt"
	hardTestPath    = "./day6/hardtest.txt"
	inputPath       = "./day6/input.txt"
	harderInputPath = "./day6/harderinput.txt"
)

func Day6() {
	records := getRecords(harderInputPath)
	var res int64 = 1
	for _, r := range records {
		res *= int64(getWinStrategyCnt(r))
	}
	fmt.Printf("res:%d\n", res)
}

type record struct {
	duration int64
	distance int64
}

func getRecords(fPath string) []*record {
	lines := util.GetLines(fPath)
	if len(lines) != 2 {
		panic("len lines not two")
	}
	durationStrSlice := strings.Fields(lines[0])[1:]
	distanceStrSlice := strings.Fields(lines[1])[1:]

	if len(durationStrSlice) != len(distanceStrSlice) {
		panic("duration cnt != distance cnt")
	}

	res := make([]*record, len(durationStrSlice))

	for i, durStr := range durationStrSlice {
		durationInt, err := strconv.ParseInt(durStr, 10, 64)
		if err != nil {
			panic(err)
		}
		distanceInt, err := strconv.ParseInt(distanceStrSlice[i], 10, 64)
		if err != nil {
			panic(err)
		}
		res[i] = &record{
			duration: durationInt,
			distance: distanceInt,
		}
	}

	return res
}

func getWinStrategyCnt(r *record) int {
	a := float64(r.duration)
	b := float64(r.distance)
	// (duration - hold_time) * hold_time > distance
	// (a - x) * x > b
	// x^2 - ax + b < 0, with x in [0, a]

	delta := math.Pow(a, 2) - 4*b
	rootDelta := math.Pow(delta, 0.5)
	lowerBound := 0.5 * (a - rootDelta)
	higherBound := 0.5 * (a + rootDelta)

	cnt := 0
	for i := int64(0); i <= r.duration; i++ {
		if float64(i) > lowerBound && float64(i) < higherBound {
			cnt++
		} else if float64(i) >= higherBound {
			break
		}
	}
	return cnt

	// delta = a**2 - 4b
	// root-delta = (a**2 - 4b)** 1/2
	// 1/2*(a +- root-delta)
}
