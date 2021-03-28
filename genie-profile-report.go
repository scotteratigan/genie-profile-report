package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const reportFilename = "report.csv"

func main() {
	logger := makeLogger()
	generateReport(logger)
}

func generateReport(logger dualLogger) {
	logger.info.Printf("Generating character listing report...\n\n")
	profileFileList := getAllProfiles()
	if len(profileFileList) == 0 {
		logger.error.Println("No xml profiles in this directory.")
		os.Exit(1)
	}

	var characters []charData
	for _, file := range profileFileList {
		content := readProfile(file, logger)
		if content == "" {
			continue
		}
		cd := extractProfileData(content)
		if len(cd.char) > 0 {
			characters = append(characters, cd)
		} else {
			logger.info.Println(fmt.Sprintf("Skipping %v due to error extracting profile data.\n", file.Name()))
		}
	}
	writeReportCSV(characters, logger)
}

func extractProfileData(profileText string) (cd charData) {
	profileRegexp := regexp.MustCompile("<Profile Account=\"([a-zA-Z0-9]+)\" Password=\".+\" Character=\"([a-zA-Z]+)\" Game=\"([a-zA-Z]+)\">")
	match := profileRegexp.FindAllStringSubmatch(profileText, -1)
	if len(match) > 0 && len(match[0]) == 4 {
		cd.acct = match[0][1]
		cd.char = match[0][2]
		cd.game = match[0][3]
	}
	return cd
}

func getAllProfiles() (profileFileNames []os.FileInfo) {
	// grab all files in directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	fileNameRegexp := regexp.MustCompile(".+.xml$")
	for _, file := range files {
		if fileNameRegexp.MatchString(file.Name()) {
			// then append files ending with '.xml' to slice
			profileFileNames = append(profileFileNames, file)
		}
	}
	return profileFileNames
}
