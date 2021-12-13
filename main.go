package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rawDataFilename := "gamma_epsion.txt"
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/3/input")
	// const exampleString = "101011011110"
	const exampleString = "110100101110"
	expected_string_length := len(exampleString)
	number_count := 0
	gamma_one_frequency := make([]int, expected_string_length)

	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		for scanner.Scan() {
			currentText := scanner.Text()
			if len(currentText) == 0 {continue}
			if len(currentText) != expected_string_length {
				fmt.Errorf("string '%s' is not valid expected something like '%s'", currentText, exampleString)
			}
			number_count++
			for i := 0; i < expected_string_length; i++{
				if currentText[i] == '1'{
					gamma_one_frequency[i]++
				} else if currentText[i] != '0'{
					panic(fmt.Errorf("unexpected symbol %s in string %s", string(currentText[i]), currentText))
				}
			}
		}
	}

	println(gamma_one_frequency)
	gamma := uint16(0)
	epsilon := uint16(0)
	leftShift := expected_string_length - 1
	for i, frequency := range gamma_one_frequency{
		if frequency > (number_count / 2){
			gamma = gamma | ((0x0001 << leftShift) >> i)
		} else{
			epsilon = epsilon | ((0x0001 << leftShift) >> i)
		}
	}

	fmt.Printf("%05s\n", strconv.FormatInt(int64(gamma), 2))
	fmt.Printf("%05s\n", strconv.FormatInt(int64(epsilon), 2))
	fmt.Printf("gamma %d epsilon %d", gamma, epsilon)

	println(gamma*epsilon)

}

// 01101
// 10010


//101111010000
//010000101111

