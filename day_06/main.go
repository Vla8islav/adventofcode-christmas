package day_06

import (
	"christmas-challenge/helpers"
	"fmt"
)
const FIRST_CYCLE_COUNTER = 8
const NTH_CYCLE_COUNTER = 6

type Fish struct{
	Day int
}

type FlockOfFish struct{
	State []Fish
	Day int
}

func (fl FlockOfFish) simulate(day int){
	for i:=0;i<=fl.Day;i++{
		fl.State = fl.nextDay()
	}
}

func (fl FlockOfFish) nextDay() []Fish{
	newState := make([]Fish, 0)
	for _, f := range fl.State{
		newFish := Fish{f.Day + 1}
		newState = append(newState, newFish)
		if f.getCycle() < newFish.getCycle(){
			newState = append(newState, Fish{0})
		}
	}
	return newState
}

func (f Fish) getCycle() int{
	nthCyclePart := f.Day - FIRST_CYCLE_COUNTER
	if nthCyclePart <= 0{
		return 0
	}else
	{
		return 1 + (nthCyclePart / (NTH_CYCLE_COUNTER + 1 + 1))
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

func RunDay06() {
	const sliceSize = 3
	rawDataFilename := "day_06.txt"
	helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/6/input")
}

func GetStateDayX(initialState []Fish, day int) []Fish{
	return initialState
}

