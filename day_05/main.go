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

func RunDay05() {
	rawDataFilename := "day_05.txt"
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
	println(moreThanOneIntersectionCounter)
}

func getPointsInLine(line Line)[]Point {
	endMatchesStartX := line.end.x == line.start.x
	endMatchesStartY := line.end.y == line.start.y
	if !endMatchesStartX && !endMatchesStartY{
		return []Point{}
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
			minValue = line.start.y
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
			minValue = line.start.x
		}
		for i := minValue; i <= maxValue; i++{
			points = append(points, Point{i, line.end.y})
		} 
	}

	return points
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