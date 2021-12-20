package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// var rawDataFilename string = "gamma_epsion.txt"
// var exampleString = "110100101110"
var rawDataFilename string = "gamma_epsion_sample.txt"
var exampleString = "00100"
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

func calculateTendency(numbersArray map[string]bool, dominantValue int) []int {
	enabledNumberLength := len(numbersArray)
	for _, flag := range numbersArray {
		if !flag {
			enabledNumberLength--
		}
	}
	freqArray := calculateFrequency(numbersArray)
	tendency := make([]int, expected_string_length)
	for i, frequency := range freqArray {
		numberOfOnes := frequency
		numberOfZeros := enabledNumberLength - frequency
		if numberOfOnes > numberOfZeros {
			tendency[i] = 1
		} else if numberOfOnes == numberOfZeros {
			tendency[i] = dominantValue
		}
	}
	return tendency
}

func getRemaining(numbersArray map[string]bool) []string {
	retval := make([]string, 0)
	for k, v := range numbersArray {
		if v {
			retval = append(retval, k)
		}
	}
	return retval
}

func printRemaining(numbersArray map[string]bool) {
	remaining := getRemaining(numbersArray)
	for _, v := range remaining {
		println(v)
	}
}

func main() {
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/3/input")
	// const exampleString = "101011011110"
	numbersArray := extractBinaryNumbers(rawDataFilename, exampleString)

	oxygenAnswer := getOxygen(numbersArray, 1)
	println("Answer:")
	println(oxygenAnswer)
	oxygenNumberInt := convertToInt(oxygenAnswer)

	fmt.Printf("%012s\n", strconv.FormatInt(int64(oxygenNumberInt), 2))
	fmt.Printf("Answer O2 in numerics %d", oxygenNumberInt)

}

func convertToInt(oxygenAnswer string) uint16 {
	number := uint16(0)

	leftShift := expected_string_length - 1
	for i, value := range oxygenAnswer {
		if value == '1' {
			number = number | ((0x0001 << leftShift) >> i)
		}
	}
	return number
}

func getOxygen(numbersArray []string, dominantTendency int) string {
	var oxygenAnswer string
	oxygenFlagMap := make(map[string]bool)
	for _, key := range numbersArray {
		oxygenFlagMap[key] = true
	}
	oxygenRemainingNumberCount := len(numbersArray)
	for position, _ := range numbersArray[0] {
		numbersTendency := calculateTendency(oxygenFlagMap, dominantTendency)
		for _, binaryNumberString := range numbersArray {
			if oxygenFlagMap[binaryNumberString] {
				expectedNumber := 0
				if numbersTendency[position] == 1 {
					expectedNumber = 1
				}
				actualNumber := 0
				if binaryNumberString[position] == '1' {
					actualNumber = 1
				}

				if expectedNumber != actualNumber {
					oxygenFlagMap[binaryNumberString] = false
					oxygenRemainingNumberCount--
				}
			}
		}
		println("Iteration")
		println(position)
		println("")
		printRemaining(oxygenFlagMap)
		println("")
		if oxygenRemainingNumberCount == 1 {
			oxygenAnswer = getRemaining(oxygenFlagMap)[0]

		}
	}
	return oxygenAnswer
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
