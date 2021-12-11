package helpers

import (
	"io"
	"net/http"
	"os"
)


func GetRawDataFromWeb(rawDataFilename string, url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", "session=53616c7465645f5f6a5d552e712b707cdf60b99ae38583cfc51dbf519359b7597a3a35a925e91d934b334e8cbfc5b7af")
	result, err := client.Do(req)

	if err == nil {
		outFile, errFile := os.Create(rawDataFilename)
		if errFile == nil {
			defer outFile.Close()
			_, _ = io.Copy(outFile, result.Body)
		}
	}
}