package main

import (
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

const reportFilename = "report.csv"

func main() {
	fmt.Println("Generating character listing report...")
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fileNameRegexp := regexp.MustCompile(".+.xml$")
	profileRegexp := regexp.MustCompile("<Profile Account=\"([a-zA-Z0-9]+)\" Password=\".+\" Character=\"([a-zA-Z]+)\" Game=\"([a-zA-Z]+)\">")
	var characters []charData
	for _, file := range files {
		if fileNameRegexp.MatchString(file.Name()) {
			content, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			profileText := string(content)
			match := profileRegexp.FindAllStringSubmatch(profileText, -1)
			if len(match) > 0 && len(match[0]) == 4 {
				cd := charData{
					acct: match[0][1],
					char : match[0][2],
					game : match[0][3],
				}
				characters = append(characters, cd)
			}
		}
	}
	reportText := "Name,Account,Game\n"
	for _, char := range characters {
		reportText += fmt.Sprintf("%v,%v,%v\n", char.char, char.acct, char.game)
	}
	fmt.Println(reportText)
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
	fmt.Printf("Report saved to %v.\n", reportFilename)
}