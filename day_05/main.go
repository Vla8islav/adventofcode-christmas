package day_05

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct{
	x int
	y int
}

type Line struct{
	start Point
	end Point
}

var startPoint Point = Point{0, 0}

func RunDay05() {
	rawDataFilename := "day_05_example.txt"
	// helpers.GetRawDataFromWeb(rawDataFilename, "https://adventofcode.com/2021/day/5/input")
	linesList := readDataIntoLinesList(rawDataFilename)
	overallPointsList := make([]Point, 0)
	for _, line := range linesList{
		pointsList := getPointsInLine(line)
		overallPointsList = append(overallPointsList, pointsList...)
		// println(line.start.x)
	}
	pointsFrequency := make(map[Point]int, 0)
	for _, point := range overallPointsList{
		// fmt.Printf("%d, %d\n", point.x, point.y)
		pointValue, pointExist := pointsFrequency[point]
		if !pointExist{
			pointsFrequency[point] = 0
		}
		pointsFrequency[point] = pointValue + 1
	}

	var moreThanOneIntersectionCounter int
	for key, value := range pointsFrequency{
		if value > 1{
			fmt.Printf("%d, %d -> %d\n", key.x, key.y, value)
			moreThanOneIntersectionCounter++
		}
	}
	resultField := getResult(pointsFrequency)
	printResult(rawDataFilename + "_my_output.txt", resultField)
	println(moreThanOneIntersectionCounter)
}

func printResult(resultFileName string, field [][]int) {
	resultFile, err := os.OpenFile(resultFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if nil == err {
		writer := bufio.NewWriter(resultFile)
		for _, row := range field{
			for _, freq := range row{
				if freq == 0{
					fmt.Fprintf(writer, ".")

				}else{
					fmt.Fprintf(writer, "%d", freq)
				}
			}
			fmt.Fprintf(writer, "\n")
		}
		writer.Flush()
		resultFile.Close()
	}

}

func findLowerRightPoint(pointsFrequency map[Point]int)Point{
	var x, y int
	for key, _ := range pointsFrequency{
		if x < key.x{
			x = key.x
		}
		if y < key.y{
			y = key.y
		}
	}
	return Point{x, y}
}

func getResult(pointsFrequency map[Point]int)[][]int{
	lowerRightPoint := findLowerRightPoint(pointsFrequency)
	var field = make([][]int, lowerRightPoint.y + 1)
	for i:=0;i<=lowerRightPoint.y;i++{
		field[i] = make([]int, lowerRightPoint.x + 1)
	}
	for point, frequency := range pointsFrequency{
		field[point.y][point.x] = frequency
	}
	return field
}

func getPointsInLine(line Line)[]Point {
	endMatchesStartX := line.end.x == line.start.x
	endMatchesStartY := line.end.y == line.start.y
	if !endMatchesStartX && !endMatchesStartY{
		points := calculateDiagonalLinePoints(line)
		return points
//		panic("at least one of the coordinates must match because the line must be straight")
	}
	points := make([]Point, 0)
	if endMatchesStartX{
		minValue := line.start.y
		if minValue > line.end.y{
			minValue = line.end.y
		}

		maxValue := line.end.y
		if maxValue < line.start.y{
			maxValue = line.start.y
		}
		for i := minValue; i <= maxValue; i++{
			points = append(points, Point{line.start.x, i})
		} 
	}

	if endMatchesStartY{
		minValue := line.start.x
		if minValue > line.end.x{
			minValue = line.end.x
		}

		maxValue := line.end.x
		if maxValue < line.start.x{
			maxValue = line.start.x
		}
		for i := minValue; i <= maxValue; i++{
			points = append(points, Point{i, line.end.y})
		} 
	€€}

	return points
}


func calculateDiagonalLinePoints(line Line)[]Point {
	slope := float64(line.start.y - line.end.y) / float64(line.start.x - line.end.x)
	retval := make([]Point, 0)
	if line.start.x <= line.end.x && line.start.y <= line.end.y{
		// left-to-right, up-to-down
		for i := line.start.x; i <= line.end.x; i++{
			shift := i - line.start.x
			retval = append(retval, Point{line.start.x+shift, line.start.y+shift})
		}
	}

	if line.start.x > line.end.x && line.start.y > line.end.y{
		// right-to-left, down-to-up
		for i := line.end.x; i <= line.start.x; i++{
			shift := i - line.end.x
			retval = append(retval, Point{line.start.x-shift, line.start.y+shift})
		}

	}
	return retval

}

func readDataIntoLinesList(rawDataFilename string) []Line {
	arrayOfValues := make([]Line, 0)
	rawDataFile, err := os.Open(rawDataFilename)
	if nil == err {
		scanner := bufio.NewScanner(rawDataFile)
		for scanner.Scan() {
			var x1, y1, x2, y2 int
			currentText := scanner.Text()
			fmt.Sscanf(currentText, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
			arrayOfValues = append(arrayOfValues, Line{start: Point{x1, y1}, end: Point{x2, y2}})
		}
	}
	return arrayOfValues
}