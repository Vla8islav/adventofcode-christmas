package main
import "christmas-challenge/day_01"


func main(){
	rawDataFilename := "elevation.txt"
	day_01.GetRawDataFromWeb(rawDataFilename)
	elevationCount := day_01.CountElevations(rawDataFilename)
	println(elevationCount)

}