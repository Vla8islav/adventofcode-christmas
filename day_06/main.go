package day_06

import (
	"christmas-challenge/helpers"
	"fmt"
	"strconv"
	"strings"
)
const FIRST_CYCLE_COUNTER = 8
const NTH_CYCLE_COUNTER = 6

type Fish struct{
	Day int
}

type FlockOfFish struct{
	State []Fish
}

func (fl FlockOfFish) Simulate(day int)FlockOfFish{
	currentState := fl.State
	for i:=0;i<day;i++{
		currentState = nextDay(currentState)
	}
	return FlockOfFish{currentState}
}

func (fl FlockOfFish) String() string{
	s := make([]int, 0)
	for _, f := range fl.State{
		s = append(s, f.getTimer())
	}
	// sort.Ints(s)
	retval := ""
	for i, f := range s{
		retval += fmt.Sprintf("%d", f)
		if i < len(s) - 1{
			retval += ","
		}
	}
	return retval
}

func nextDay(initialState []Fish)[]Fish{
	newState := make([]Fish, 0)
	for _, f := range initialState{
		newFish := Fish{f.Day + 1}
		newState = append(newState, newFish)
		if f.getCycle() < newFish.getCycle(){
			newState = append(newState, Fish{0})
		}
	}
	return newState
}

func (f Fish) getCycle() int{
	nthCyclePart := f.Day - FIRST_CYCLE_COUNTER - 1
	if nthCyclePart < 0{
		return 0
	}else
	{
		return 1 + (nthCyclePart / (NTH_CYCLE_COUNTER + 1))
	}
}

func (f Fish) getTimer() int{
	if f.getCycle() == 0{
		return FIRST_CYCLE_COUNTER - f.Day
	}else if f.getCycle() > 0{
		nthCyclesElapsedDays := f.Day - FIRST_CYCLE_COUNTER - 1
		daysAfterLastCycle := nthCyclesElapsedDays % (NTH_CYCLE_COUNTER + 1)
		return NTH_CYCLE_COUNTER - daysAfterLastCycle
	}
	panic("negative getCycle value")
}

func (f Fish) String() string{
	return fmt.Sprintf("day:%03d,cycle:%03d,timer:%d", f.Day, f.getCycle(),f.getTimer())
}

func ToFlockOfFish(fishTimeToGestationString string)FlockOfFish{
    // we assume that all fish in the zeroth cycle in the initial state
    retval := make([]Fish, 0)
    for _, strNum := range strings.Split(fishTimeToGestationString, ","){
        strNum = strings.TrimSpace(strNum)
        i, e := strconv.Atoi(strNum)
        if e == nil {
            retval = append(retval, Fish{Day: FIRST_CYCLE_COUNTER - i})
        }
    }
    return FlockOfFish{State:retval}
}

func RunDay06() {
	const sliceSize = 3
	rawDataFilename := "day_06.txt"
	helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/6/input")
	dataArray := helpers.ReadDataIntoStringArray(rawDataFilename)
	inputString := dataArray[0]
	flock := ToFlockOfFish(inputString)
	const SIM_DAY = 80;
	endResult := flock.Simulate(SIM_DAY)
	fmt.Printf("Number of fish after %d day is %d", SIM_DAY, len(endResult.State))
}

func GetStateDayX(initialState []Fish, day int) []Fish{
	return initialState
}

