package main

import (
	"flag"
	"fmt"
	"strconv"
)

var (
	captureFlag = flag.String("capture", "", "capture data to solve")
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
func sum(digits []int) int {
	total := 0
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
	return total
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
	fmt.Printf("Digits = %v, Sum = %d\n", digits, sum(digits))
}
