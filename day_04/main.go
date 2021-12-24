package day_04

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const crossedOutMarker = -1

func readDataIntoStringArray(rawDataFilename string) []string {
	arrayOfValues := make([]string, 0)
	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		for scanner.Scan() {
			currentText := scanner.Text()
			arrayOfValues = append(arrayOfValues, currentText)
		}
	}
	return arrayOfValues
}

func extractNumbersString(dataArray []string) string {
	return dataArray[0]
}

func extractDashboardsStrings(dataArray []string) [][]string {
	var retval [][]string
	var currentBingoBoard []string
	for i := 2; i < len(dataArray); i++ {
		if len(dataArray[i]) == 0 {
			retval = append(retval, currentBingoBoard)
			currentBingoBoard = make([]string, 0)
			continue
		}
		currentBingoBoard = append(currentBingoBoard, dataArray[i])
	}
	return retval
}

func dashboardStringsToNumbers(dashboardStringsArray [][]string) [][][]int {
	var retval [][][]int
	for _, dashboadStrings := range dashboardStringsArray {
		var dashboad [][]int
		for _, dashboadRowString := range dashboadStrings {
			var dashboadRow []int
			for _, dashboardNumber := range strings.Split(dashboadRowString, " ") {
				numberToAppend, error := strconv.Atoi(dashboardNumber)
				if error == nil {
					dashboadRow = append(dashboadRow, numberToAppend)
				}
			}
			dashboad = append(dashboad, dashboadRow)
		}
		retval = append(retval, dashboad)
	}
	return retval
}

func transpose(bingoBoard [][]int) [][]int {
	colLen := len(bingoBoard)
	if colLen == 0 {
		return nil
	}
	rowLen := len(bingoBoard[0])

	retval := make([][]int, rowLen)
	for i := range retval {
		retval[i] = make([]int, colLen)
	}

	for i := 0; i < colLen; i++ {
		for j := 0; j < rowLen; j++ {
			retval[i][j] = bingoBoard[j][i]
		}
	}
	return retval
}

func transposeAll(bingoBoards [][][]int) [][][]int {
	boardsNumber := len(bingoBoards)
	retval := make([][][]int, boardsNumber)
	for i := range bingoBoards {
		retval[i] = transpose(bingoBoards[i])
	}
	return retval
}

func addNumberToRows(dashboardsInt [][][]int, selectedRows [][][]int) {

}

func numbersStringToIntArray(numbersString string) []int {
	var retval []int
	numberStrings := strings.Split(numbersString, ",")
	for _, numberString := range numberStrings {
		numberToAppend, error := strconv.Atoi(numberString)
		if error == nil {
			retval = append(retval, numberToAppend)
		}
	}
	return retval

}

func RunDay04() {
	rawDataFilename := "day_04.txt"
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/4/input")
	dataArray := readDataIntoStringArray(rawDataFilename)
	numbersString := extractNumbersString(dataArray)
	bingoBoardStringsArray := extractDashboardsStrings(dataArray)

	dashboardsRowsInt := dashboardStringsToNumbers(bingoBoardStringsArray)

	numbersToCross := numbersStringToIntArray(numbersString)

	var sumOfUnmarkedNumbers int
	var lastNumber int


	for _, numberToCross := range numbersToCross {
		lastNumber = numberToCross
		var crossedOutRowFlag bool
		for i := 0; i < len(dashboardsRowsInt); i++ {
			currentDashboard := dashboardsRowsInt[i]
			crossNumber(currentDashboard, numberToCross)
			crossedOutRowFlag, sumOfUnmarkedNumbers = getSumOfUnmarkedNumbers(currentDashboard, numberToCross)
			if crossedOutRowFlag{
				break
			}
			dashboardsColsInt := transposeAll(dashboardsRowsInt)
			currentDashboardCols := dashboardsColsInt[i]
			crossedOutRowFlag, sumOfUnmarkedNumbers = getSumOfUnmarkedNumbers(currentDashboardCols, numberToCross)
			if crossedOutRowFlag{
				break
			}
		}
		if crossedOutRowFlag{
			break
		}
	}
	result := sumOfUnmarkedNumbers * lastNumber

	println(numbersString)
	println(bingoBoardStringsArray)
	println(dashboardsRowsInt)
	println(numbersToCross)
	println(sumOfUnmarkedNumbers)
	println(result)
	
	// println(dashboardsColsInt)
}

func getSumOfUnmarkedNumbers(currentDashboard [][]int, numberToCross int) (bool, int) {
	if hasCrossedOutRow(currentDashboard) {
		println(numberToCross)
		sumOfUnmarkedNumbers := calculateSumOfUnmarkedNumbers(currentDashboard)
		return true, sumOfUnmarkedNumbers
	}
	return false, 0
}

func calculateSumOfUnmarkedNumbers(bingoBoard [][]int) int {
	var retval int
	for _, bingoRow := range bingoBoard {
		for _, number := range bingoRow {
			if number != crossedOutMarker {
				retval += number
			}
		}
	}
	return retval
}

func hasCrossedOutRow(bingoBoard [][]int) bool {
	for _, bingoRow := range bingoBoard {
		if isCrossedOut(bingoRow) {
			return true
		}
	}
	return false
}

func isCrossedOut(bingoRow []int) bool {
	for _, number := range bingoRow {
		if number != crossedOutMarker {
			return false
		}
	}
	return true
}

func crossNumber(bingoBoard [][]int, numberToCross int) {
	for i, bingoRow := range bingoBoard {
		for j, number := range bingoRow {
			if number == numberToCross {
				bingoBoard[i][j] = crossedOutMarker
			}
		}
	}
}
