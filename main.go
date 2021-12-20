package main

import (
	"bufio"
	"fmt"
	"os"
)

var exampleString = "110100101110"
var expected_string_length = len(exampleString)

func calculateFrequency(numbersArray map[string]bool) []int {
	onesFrequency := make([]int, expected_string_length)
	for binaryNumberString, active := range numbersArray {
		if active {
			for i, letter := range binaryNumberString {
				if letter == '1' {
					onesFrequency[i]++
				}
			}
		}
	}
	return onesFrequency
}

func calculateTendency(numbersArray map[string]bool) []int {
	enabledNumberLength := len(numbersArray)
	for _, flag := range numbersArray{
		if !flag{
			enabledNumberLength--
		}
	}
	freqArray := calculateFrequency(numbersArray)
	tendency := make([]int, expected_string_length)
	for i, frequency := range freqArray {
		if frequency > enabledNumberLength/2 {
			tendency[i] = 1
		}
	}
	return tendency
}

func chooseOxygenFromTwoOrLess(possibleNumbers []string, comparePosition int) string{
	if len(possibleNumbers) == 1{
		return possibleNumbers[0]
	}
	if len(possibleNumbers) == 2{
		firstNumberIndicator := possibleNumbers[0][comparePosition]
		secondNumberIndicator := possibleNumbers[1][comparePosition]

	}

}

func main() {
	rawDataFilename := "gamma_epsion.txt"
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/3/input")
	// const exampleString = "101011011110"
	numbersArray := extractBinaryNumbers(rawDataFilename, exampleString)

	oxygenFlagMap := make(map[string]bool)
	for _, key := range numbersArray {
		oxygenFlagMap[key] = true
	}

	oxygenNumber := ""
	oxygenRemainingNumberCount := len(numbersArray)

	for position, _ := range numbersArray[0] {
		numbersTendency := calculateTendency(oxygenFlagMap)
		print(numbersTendency)
		for _, binaryNumberString := range numbersArray {
			if oxygenFlagMap[binaryNumberString] {
				expectedNumber := 0
				if numbersTendency[position] == '1'{
					expectedNumber = 1
				}
				actualNumber := 0
				if binaryNumberString[position] == '1'{
					actualNumber = 1
				}
				
				if expectedNumber != actualNumber {
					oxygenFlagMap[binaryNumberString] = false
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
