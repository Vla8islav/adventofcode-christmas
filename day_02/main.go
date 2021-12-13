package day_02

import (
	"bufio"
	"christmas-challenge/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawDataFilename := "directions.txt"
	helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/2/input")
	horPosition := 0
	depth := 0
	aim := 0

	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		for scanner.Scan() {
			currentText := scanner.Text()
			textSlices := strings.Split(currentText, " ")
			if len(textSlices) == 2 {
				direction := textSlices[0]
				value, conversionError := strconv.Atoi(textSlices[1])
				if conversionError == nil{
					switch direction{
					case "down":
						aim += value
					case "up":
						aim -= value
					case "forward":
						horPosition += value
						depth += aim*value
					default:
						println("incorrect direction " + direction)
					}

				}

			}

		}
	}
	fmt.Printf("Depth %d hor position %d answer to enter %d", depth, horPosition, depth*horPosition)
}