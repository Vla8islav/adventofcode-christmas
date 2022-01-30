package helpers

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func ReadDataIntoStringArray(rawDataFilename string) []string {
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

func GetRawDataFromWeb(rawDataFilename string, url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", "session=53616c7465645f5f3e6a237738cf9405524959c89e3f151d5b03440ac23c3c58b680b58d4ba93aafdd8a1ea27dbfb4fe")
	result, err := client.Do(req)

	if err == nil {
		outFile, errFile := os.Create(rawDataFilename)
		if errFile == nil {
			defer outFile.Close()
			_, _ = io.Copy(outFile, result.Body)
		}
	}
}