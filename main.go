package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var StringsToDigits = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var part1Sum int
	var part2Sum int

	reader := bufio.NewScanner(file)

	reader.Split(bufio.ScanLines)

	for reader.Scan() {

		values := extractCalibrationVal(reader.Text())

		vals := ExtractPartTwo(reader.Text())

		if len(values) == 1 {
			digit, _ := strconv.Atoi(fmt.Sprintf("%d%d", values[0], values[0]))
			part1Sum += digit

		} else if len(values) >= 2 {
			digit, _ := strconv.Atoi(fmt.Sprintf("%d%d", values[0], values[len(values)-1]))
			part1Sum += digit

		}

		if len(vals) == 1 {
			digit, _ := strconv.Atoi(fmt.Sprintf("%d%d", vals[0], vals[0]))
			part2Sum += digit

		} else if len(vals) >= 2 {
			digit, _ := strconv.Atoi(fmt.Sprintf("%d%d", vals[0], vals[len(vals)-1]))
			part2Sum += digit

		}

	}

	fmt.Println("Part 1: ", part1Sum)
	fmt.Println("Part 2: ", part2Sum)
}

func extractCalibrationVal(line string) []int {
	values := []int{}

	for _, rune := range line {

		char := string(rune)
		if rune < '0' || rune > '9' {
		} else {
			digit, _ := strconv.Atoi(char)
			values = append(values, digit)
		}

	}
	return values

}

func ExtractPartTwo(line string) []int {

	// var calSum int

	var currentWords []int

	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for i := range line {

		found := re.FindString(line[i:])

		if found != "" {
			if _, err := strconv.Atoi(found); err == nil {
				digit, _ := strconv.Atoi(found)
				currentWords = append(currentWords, digit)
			} else if _, ok := StringsToDigits[found]; ok {
				currentWords = append(currentWords, StringsToDigits[found])
			}
		}
	}

	return currentWords

}
