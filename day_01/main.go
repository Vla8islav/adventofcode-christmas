package day_01

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	rawDataFilename := "elevation.txt"
	elevationCounter := ReadElevationMapIntoSlice(rawDataFilename)
	fmt.Print(elevationCounter)
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

func GetRawDataFromWeb(rawDataFilename string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://adventofcode.com/2021/day/1/input", nil)
	req.Header.Set("cookie", "session=53616c7465645f5f6a5d552e712b707cdf60b99ae38583cfc51dbf519359b7597a3a35a925e91d934b334e8cbfc5b7af")
	result, err := client.Do(req)

	if err == nil {
		outFile, errFile := os.Create(rawDataFilename)
		if errFile == nil {
			defer outFile.Close()
			_, _ = io.Copy(outFile, result.Body)
		}
	}
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

func SumSlice(slice []int) int{
	sum := 0
	for _, val := range slice{
		sum += val
	}
	return sum
}