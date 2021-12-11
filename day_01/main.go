package day_01

import "christmas-challenge/helpers"

import (
	"bufio"
	"os"
	"strconv"
)

func Main() {
	sliceSumIncrease := 0
	const sliceSize = 3
	rawDataFilename := "elevation.txt"
	helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/1/input")
	elevationMap := ReadElevationMapIntoSlice(rawDataFilename)

	for i, _ := range elevationMap {
		firstSliceLeftIndex := 0 + i
		firstSliceRightIndex := sliceSize - 1 + i
		if firstSliceRightIndex < len(elevationMap) {
			firstSlice := elevationMap[firstSliceLeftIndex : firstSliceRightIndex+1]
			secondSlice := elevationMap[firstSliceLeftIndex+1 : firstSliceRightIndex+1+1]
			firstSliceSum := SumSlice(firstSlice)
			secondSliceSum := SumSlice(secondSlice)
			if firstSliceSum < secondSliceSum {
				sliceSumIncrease++
			}
		}

	}

	println(sliceSumIncrease)
}

func CountElevations(rawDataFilename string) int {
	elevationCounter := 0
	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		scanner.Scan()

		previousValue, conversionError := strconv.Atoi(scanner.Text())
		if conversionError == nil {
			for scanner.Scan() {
				nextValue, conversionError := strconv.Atoi(scanner.Text())
				if nil == conversionError {
					valToAdd := 0
					if previousValue < nextValue {
						valToAdd = 1
					}
					elevationCounter += valToAdd
					previousValue = nextValue
				}
			}
		}

	}
	return elevationCounter
}


func ReadElevationMapIntoSlice(rawDataFilename string) []int {
	elevationMap := []int{}
	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		for scanner.Scan() {
			value, conversionError := strconv.Atoi(scanner.Text())
			if nil == conversionError {
				elevationMap = append(elevationMap, value)
			}
		}
	}
	return elevationMap
}

func SumSlice(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}
