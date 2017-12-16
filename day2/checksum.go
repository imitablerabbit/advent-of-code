package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	spreadsheetFlag = flag.String("spreadsheet", "", "`file` of spreadsheet data to compute checksum for")
	typeFlag        = flag.Int("type", 1, "type of checksum to calculate, 1 is lowest and highest diff, 2 is even division")
)

// type constants
const (
	DIFF = 1 + iota
	EVENDIVISION
)

func init() {
	flag.Parse()
}

func computeRowDiff(row []int) int {
	lowest := row[0]
	highest := row[0]
	for _, num := range row {
		if num < lowest {
			lowest = num
		}
		if num > highest {
			highest = num
		}
	}
	return highest - lowest
}

func computeEvenDivision(row []int) int {
	for aIndex, a := range row {
		for bIndex, b := range row {
			if aIndex == bIndex {
				continue
			}
			if a%b == 0 {
				return a / b
			}
		}
	}
	return 0 // Should not get here
}

func computeDiffChecksum(spreadsheet [][]int) (checksum int) {
	for _, row := range spreadsheet {
		diff := computeRowDiff(row)
		checksum = checksum + diff
	}
	return
}

func computeEvenDivisionChecksum(spreadsheet [][]int) (checksum int) {
	for _, row := range spreadsheet {
		division := computeEvenDivision(row)
		checksum = checksum + division
	}
	return
}

func parseSpreadsheetString(spreadsheet string) (data [][]int, err error) {
	trimmedSpreadsheet := strings.Trim(spreadsheet, "\n")
	rowStrings := strings.Split(trimmedSpreadsheet, "\n")
	for _, rowString := range rowStrings {
		whitespaceRegex := regexp.MustCompile("\\s+") // Any whitespace
		numberStrings := whitespaceRegex.Split(rowString, -1)
		var rowDigits []int
		for _, numberString := range numberStrings {
			digit, err := strconv.Atoi(numberString)
			if err != nil {
				return data, err
			}
			rowDigits = append(rowDigits, digit)
		}
		data = append(data, rowDigits)
	}
	return
}

func main() {
	if *spreadsheetFlag == "" {
		flag.PrintDefaults()
		return
	}
	spreadsheetBytes, err := ioutil.ReadFile(*spreadsheetFlag)
	if err != nil {
		fmt.Printf("Error: read spreadsheet file: %v\n", err)
		return
	}
	spreadsheetData, err := parseSpreadsheetString(string(spreadsheetBytes))
	if err != nil {
		fmt.Printf("Error: could not parse spreadsheet data: %v\n%v", err, string(spreadsheetBytes))
		return
	}
	var checksum int
	switch *typeFlag {
	case DIFF:
		checksum = computeDiffChecksum(spreadsheetData)
		break
	case EVENDIVISION:
		checksum = computeEvenDivisionChecksum(spreadsheetData)
		break
	}
	fmt.Printf("Spreadsheet = %v\nChecksum = %d\n", spreadsheetData, checksum)
}
