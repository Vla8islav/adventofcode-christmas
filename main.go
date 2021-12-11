package main

import "christmas-challenge/day_01"

func main() {
	sliceSumIncrease := 0
	const sliceSize = 3
	rawDataFilename := "elevation.txt"
	elevationMap := day_01.ReadElevationMapIntoSlice(rawDataFilename)

	for i, _ := range elevationMap {

		firstSliceLeftIndex := 0 + i
		firstSliceRightIndex := sliceSize - 1 + i
		if firstSliceRightIndex < len(elevationMap) {
			firstSlice := elevationMap[firstSliceLeftIndex : firstSliceRightIndex+1]
			secondSlice := elevationMap[firstSliceLeftIndex+1 : firstSliceRightIndex+1+1]
			firstSliceSum := day_01.SumSlice(firstSlice)
			secondSliceSum := day_01.SumSlice(secondSlice)
			valToAdd := 0
			if firstSliceSum < secondSliceSum{
				valToAdd = 1
			}
			sliceSumIncrease += valToAdd
		}

	}

	println(sliceSumIncrease)
}
