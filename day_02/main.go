package day_02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rawDataFilename := "elevation.txt"
	elevationCounter := ReadElevationMapIntoSlice(rawDataFilename)
	fmt.Print(elevationCounter)

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
