package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	captureFlag = flag.String("capture", "", "capture data to solve")
	typeFlag    = flag.Int("type", 0, "the type of capture calculation, 0 is next digit, 1 is halfway around")
)

// type constants for how the sum is calculated
const (
	NEXT = iota
	HALFWAY
)

func init() {
	flag.Parse()
}

// splitString will take in the capture data as a string and turn it into a
// slice of integers
func splitString(data string) (digits []int, err error) {
	for _, c := range data {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			return digits, err
		}
		digits = append(digits, num)
	}
	return
}

// Loop through digits and sum them if they are matching next to each other
// Make sure to get the last digit summed with the first.
func sumNext(digits []int) (total int) {
	lastIndex := len(digits) - 1
	for index, a := range digits[:lastIndex] {
		b := digits[index+1]
		if a == b {
			total = total + a
		}
	}
	first := digits[0]
	last := digits[lastIndex]
	if first == last {
		total = total + last
	}
	return
}

// Loop through digits and sum them if they match the digit halfway around the
// slice.
func sumHalfway(digits []int) (total int) {
	offset := len(digits) / 2
	for index, a := range digits {
		var halfwayIndex int
		if index >= offset { // Need to wrap halfway index
			halfwayIndex = index - offset
		} else {
			halfwayIndex = index + offset
		}
		b := digits[halfwayIndex]
		if a == b {
			total = total + a
		}
	}
	return
}

func main() {
	if *captureFlag == "" {
		flag.PrintDefaults()
		return
	}
	digits, err := splitString(*captureFlag)
	if err != nil {
		fmt.Printf("Error: could not parse the capture data: %v\n", err)
		return
	}
	var sum int
	switch *typeFlag {
	case NEXT:
		sum = sumNext(digits)
		break
	case HALFWAY:
		sum = sumHalfway(digits)
		break
	}
	fmt.Printf("Digits = %v, Sum = %d\n", digits, sum)
}
