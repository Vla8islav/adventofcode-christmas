package day_04

import (
	"strconv"
	"strings"
	"christmas-challenge/helpers"
)

const crossedOutMarker = -1


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
	dataArray := helpers.ReadDataIntoStringArray(rawDataFilename)
	numbersString := extractNumbersString(dataArray)
	bingoBoardStringsArray := extractDashboardsStrings(dataArray)

	dashboardsRowsInt := dashboardStringsToNumbers(bingoBoardStringsArray)

	numbersToCross := numbersStringToIntArray(numbersString)

	winningBoards := make(map[int]int, 0)
	var winningBoardsOrder []int

	for _, numberToCross := range numbersToCross {
		for i := 0; i < len(dashboardsRowsInt); i++ {
			currentDashboard := dashboardsRowsInt[i]
			crossNumber(currentDashboard, numberToCross)
		}
		for i := 0; i < len(dashboardsRowsInt); i++ {
			if _, present := winningBoards[i]; present{
				continue
			}
			currentDashboard := dashboardsRowsInt[i]
			var sumOfUnmarkedNumbers int
			crossedOutRowFlag, sumOfUnmarkedNumbersRow := getSumOfUnmarkedNumbers(currentDashboard, numberToCross)
			dashboardsColsInt := transposeAll(dashboardsRowsInt)
			currentDashboardCols := dashboardsColsInt[i]
			var crossedOutColFlag bool
			crossedOutColFlag, sumOfUnmarkedNumbersCol := getSumOfUnmarkedNumbers(currentDashboardCols, numberToCross)
			if crossedOutRowFlag{
				sumOfUnmarkedNumbers = sumOfUnmarkedNumbersRow
			}
			if crossedOutColFlag{
				sumOfUnmarkedNumbers = sumOfUnmarkedNumbersCol
			}

			if crossedOutRowFlag || crossedOutColFlag {
				winningBoards[i] = sumOfUnmarkedNumbers * numberToCross
				winningBoardsOrder = append(winningBoardsOrder, i)
			}
		}
	}
	lastWinningBoardIndex := winningBoardsOrder[len(winningBoardsOrder)-1]
	result := winningBoards[lastWinningBoardIndex]

	println(numbersString)
	println(bingoBoardStringsArray)
	println(dashboardsRowsInt)
	println(numbersToCross)
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
