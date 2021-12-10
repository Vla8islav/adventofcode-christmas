package main
import "christmas-challenge/day_01"


func main(){
	rawDataFilename := "elevation.txt"
	elevationCount := day_01.CountElevations(rawDataFilename)
	println(elevationCount)

}