package day_06

import (
	"christmas-challenge/helpers"
	"fmt"
	"strconv"
	"strings"
)

const FIRST_CYCLE_COUNTER = 8
const NTH_CYCLE_COUNTER = 6

var fishFlockCache map[int]Fish = make(map[int]Fish)

func getFishPointerFromCache(day int) *Fish {
	val, found := fishFlockCache[day]
	if !found {
		fishFlockCache[day] = Fish{day}
		val = fishFlockCache[day]
		// fmt.Printf("Flock cache size %d", len(fishFlockCache))
	}
	return &val
}

type Fish struct {
	Day int
}

type FlockOfFish struct {

	State map[Fish]int
}

func (fl FlockOfFish) Simulate(day int) FlockOfFish {
	currentState := fl.State
	for i := 0; i < day; i++ {
		currentState = nextDay(currentState)
		// fmt.Printf("Day %d\n", i)
	}
	return FlockOfFish{currentState}
}

func (fl FlockOfFish) GetFishNumber() int {
	retval:=0
	for _, val := range fl.State{
		retval += val
	}
	return retval
}

func (fl FlockOfFish) String() string {
	s := make([]int, 0)
	for f, number := range fl.State {
		for i:=0;i<number;i++{
			s = append(s, f.getTimer())
		}
	}
	// sort.Ints(s)
	retval := ""
	for i, f := range s {
		retval += fmt.Sprintf("%d", f)
		if i < len(s)-1 {
			retval += ","
		}
	}
	return retval
}

func nextDay(initialState map[Fish]int)map[Fish]int {
	var newState = make(map[Fish]int)

	for f, currentFishNumber := range initialState {
		newFish := Fish{f.Day + 1}
		currentNewStateFishNumber := newState[newFish]
		newState[newFish]=currentNewStateFishNumber+currentFishNumber

		if f.getCycle() < newFish.getCycle() {
			curNewFishVal := newState[Fish{0}]
			newState[Fish{0}]=curNewFishVal+currentFishNumber
		}
	}
	return newState
}

var fishCycleCahce map[int]int = make(map[int]int)

func (f Fish) getCycle() int {
	val, found := fishCycleCahce[f.Day]
	if !found {
		nthCyclePart := f.Day - FIRST_CYCLE_COUNTER - 1
		if nthCyclePart < 0 {
			val = 0
		} else {
			val = 1 + (nthCyclePart / (NTH_CYCLE_COUNTER + 1))
		}
		fishCycleCahce[f.Day] = val

	}
	return val
}

var fishTimerCahce map[int]int = make(map[int]int)
func (f Fish) getTimer() int {
	val, found := fishTimerCahce[f.Day]
	if !found {
		if f.getCycle() == 0 {
			val = FIRST_CYCLE_COUNTER - f.Day
		} else if f.getCycle() > 0 {
			nthCyclesElapsedDays := f.Day - FIRST_CYCLE_COUNTER - 1
			daysAfterLastCycle := nthCyclesElapsedDays % (NTH_CYCLE_COUNTER + 1)
			val = NTH_CYCLE_COUNTER - daysAfterLastCycle
		}
		fishTimerCahce[f.Day] = val
	}
	return val
}

func (f Fish) String() string {
	return fmt.Sprintf("day:%03d,cycle:%03d,timer:%d", f.Day, f.getCycle(), f.getTimer())
}

func ToFlockOfFish(fishTimeToGestationString string) FlockOfFish {
	// we assume that all fish in the zeroth cycle in the initial state
	retval := make(map[Fish]int, 0)
	for _, strNum := range strings.Split(fishTimeToGestationString, ",") {
		strNum = strings.TrimSpace(strNum)
		i, e := strconv.Atoi(strNum)
		if e == nil {
			val := retval[Fish{FIRST_CYCLE_COUNTER-i}]
			retval[Fish{FIRST_CYCLE_COUNTER-i}]=val+1
		}
	}
	return FlockOfFish{State: retval}
}

func RunDay06() {
	rawDataFilename := "day_06.txt"
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/6/input")
	dataArray := helpers.ReadDataIntoStringArray(rawDataFilename)
	inputString := dataArray[0]
	flock := ToFlockOfFish(inputString)
	const SIM_DAY = 256
	endResult := flock.Simulate(SIM_DAY)
	fmt.Printf("Number of fish after %d day is %d", SIM_DAY, endResult.GetFishNumber())
}

func GetStateDayX(initialState []Fish, day int) []Fish {
	return initialState
}
