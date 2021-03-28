package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeReportCSV(characters []charData, logger dualLogger) {
	reportText := "Name,Account,Game\n"
	for _, char := range characters {
		reportText += char.toString()
	}
	logger.info.Println(reportText)
	outputFile, err := os.Create(reportFilename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			panic(err)
		}
	}()
	outputFile.WriteString(reportText)
	logger.info.Printf("Report saved to %v.\n", reportFilename)
}

func readProfile(file os.FileInfo, logger dualLogger) string {
	content, err := ioutil.ReadFile(file.Name())
	if err != nil {
		logger.error.Println(fmt.Sprintf("Error reading file %v\n%v\n", file.Name(), err))
	}
	return string(content)
}
