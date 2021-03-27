package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type charData struct {
	acct string
	char string
	game string
}

func (cd charData) toString() string {
	return fmt.Sprintf("%v,%v,%v\n", cd.char, cd.acct, cd.game)
}

const reportFilename = "report.csv"

type dualLogger struct {
	info *log.Logger
	error *log.Logger
}

func main() {
	logger := dualLogger{
		info: log.New(os.Stdout, "", 0),
		error: log.New(os.Stderr, "", 0),
	}
	logger.info.Printf("Generating character listing report...\n\n")
	profileFileList := getAllProfiles()
	var characters []charData
	if len(profileFileList) == 0 {
		logger.error.Println("No xml profiles in this directory.")
		os.Exit(1)
	}
	for _, file := range profileFileList {
		content, err := ioutil.ReadFile(file.Name())
		if err != nil {
			logger.error.Println(fmt.Sprintf("Error reading file %v\n%v\n", file.Name(), err))
			continue
		}
		cd, profileErr := extractProfileData(string(content))
		if profileErr != nil {
			logger.info.Println(fmt.Sprintf("Skipping %v due to error extracting profile data.\n", file.Name()))
		} else {
			characters = append(characters, cd)
		}
	}
	writeReportCSV(characters, logger)
}

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

func extractProfileData(profileText string) (cd charData, e error) {
	profileRegexp := regexp.MustCompile("<Profile Account=\"([a-zA-Z0-9]+)\" Password=\".+\" Character=\"([a-zA-Z]+)\" Game=\"([a-zA-Z]+)\">")
	match := profileRegexp.FindAllStringSubmatch(profileText, -1)
	if len(match) > 0 && len(match[0]) == 4 {
			cd.acct = match[0][1]
			cd.char = match[0][2]
			cd.game = match[0][3]
	} else {
		e = errors.New("no profile data found")
	}
	return cd, e
}

func getAllProfiles() (profileFileNames []os.FileInfo) {
	fileNameRegexp := regexp.MustCompile(".+.xml$")
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if fileNameRegexp.MatchString(file.Name()) {
			profileFileNames = append(profileFileNames, file)
		}
	}
	return profileFileNames
}