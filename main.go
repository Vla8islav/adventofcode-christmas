package main

import (
	"bufio"
	"fmt"
	"os"
)

var exampleString = "110100101110"
var expected_string_length = len(exampleString)

func calculateFrequency(numbersArray []string) []int {
	onesFrequency := make([]int, expected_string_length)
	for _, binaryNumberString := range numbersArray {
		for i, letter := range binaryNumberString {
			if letter == '1' {
				onesFrequency[i]++
			}
		}
	}
	return onesFrequency
}

func calculateTendency(numbersArray []string) []int {
	freqArray := calculateFrequency(numbersArray)
	tendency := make([]int, expected_string_length)
	for i, frequency := range freqArray {
		if frequency > len(numbersArray)/2 {
			tendency[i] = 1
		}
	}
	return tendency
}

func main() {
	rawDataFilename := "gamma_epsion.txt"
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/3/input")
	// const exampleString = "101011011110"
	numbersArray := extractBinaryNumbers(rawDataFilename, exampleString)
	numbersArrayLen := len(numbersArray)
	numbersTendency := calculateTendency(numbersArray)

	oxygenFlagArray := make([]bool, numbersArrayLen)
	for i, _ := range oxygenFlagArray {
		oxygenFlagArray[i] = true
	}

	oxygenRemainingNumberCount := len(numbersArray)

	for position, _ := range numbersArray[0] {
		for i, binaryNumberString := range numbersArray {
			if oxygenFlagArray[i] {
				if !(binaryNumberString[position] == '1' && numbersTendency[position] == 1) {
					oxygenFlagArray[i] = false
					oxygenRemainingNumberCount--
				}
			}
		}
	}

}

func extractBinaryNumbers(rawDataFilename string, exampleString string) []string {
	arrayOfValues := make([]string, 0)
	expected_string_length := len(exampleString)
	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		for scanner.Scan() {
			currentText := scanner.Text()
			if len(currentText) == 0 {
				continue
			}
			if len(currentText) != expected_string_length {
				fmt.Errorf("string '%s' is not valid expected something like '%s'", currentText, exampleString)
			}
			arrayOfValues = append(arrayOfValues, currentText)
		}
	}
	return arrayOfValues
}

// 01101
// 10010

//101111010000
//010000101111
