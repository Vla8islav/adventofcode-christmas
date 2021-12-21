package main

import (
	"bufio"
	"fmt"
	"os"
)

// var rawDataFilename string = "gamma_epsion.txt"
// var exampleString = "110100101110"
var rawDataFilename string = "gamma_epsion_sample.txt"
var exampleString = "00100"
var expected_string_length = len(exampleString)

func main() {
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/3/input")
	// const exampleString = "101011011110"
	numbersArray := extractBinaryNumbers(rawDataFilename, exampleString)

	// oxygenAnswer := getOxygen(numbersArray, 1, false)
	co2Answer := getOxygen(numbersArray, 0, true)
	// println("Oxygen answer:")
	// println(oxygenAnswer)
	println("Co2 answer:")
	println(co2Answer)
	// oxygenNumberInt := convertToInt(oxygenAnswer)
	// co2NumberInt := convertToInt(co2Answer)

	// fmt.Printf("%012s\t%012s\n", strconv.FormatInt(int64(oxygenNumberInt), 2), strconv.FormatInt(int64(co2NumberInt), 2))
	// fmt.Printf("Answer O2 in numerics %d co2 in numerics %d", oxygenNumberInt, co2NumberInt)

}

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

func calculateTendency(numbersArray map[string]bool, dominantValue int, chooseLeast bool) []int {
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
			continue
		}
		
		if chooseLeast {
			tendency[i] = flipNumber(tendency[i])
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

func flipTendency(tendency []int) {
	for i, _ := range tendency {
		if tendency[i] == 1 {
			tendency[i] = 0
		}
		if tendency[i] == 0 {
			tendency[i] = 1
		}
	}
}
func flipNumber(number int) int {
	if number == 1 {
		return 0
	}
	if number == 0 {
		return 1
	}
	return number
}

func getOxygen(numbersArray []string, dominantTendency int, chooseLeast bool) string {
	var oxygenAnswer string
	oxygenFlagMap := make(map[string]bool)
	for _, key := range numbersArray {
		oxygenFlagMap[key] = true
	}
	oxygenRemainingNumberCount := len(numbersArray)
	for position, _ := range numbersArray[0] {
		numbersTendency := calculateTendency(oxygenFlagMap, dominantTendency, chooseLeast)
		for _, binaryNumberString := range numbersArray {
			if oxygenFlagMap[binaryNumberString] {
				expectedNumber := numbersTendency[position]
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
		printIteration(position, oxygenFlagMap)
		if oxygenRemainingNumberCount == 1 {
			oxygenAnswer = getRemaining(oxygenFlagMap)[0]
			break
		}
	}
	return oxygenAnswer
}

func printIteration(position int, oxygenFlagMap map[string]bool) {
	println("Iteration")
	println(position)
	println("")
	printRemaining(oxygenFlagMap)
	println("")
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
