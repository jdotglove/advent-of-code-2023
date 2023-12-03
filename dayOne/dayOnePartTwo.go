package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"unicode"
	"regexp"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("partTwoInput.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	numMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	var calValSum int = 0
	
	for fileScanner.Scan() {
		lineScanner := bufio.NewScanner(strings.NewReader(fileScanner.Text()))
		lineScanner.Split(bufio.ScanRunes)
		var firstDigit string = ""
		var lastDigit string = ""
		var numLetters string = ""

		for lineScanner.Scan() {

			for _, runeValue := range lineScanner.Text() {

				if (unicode.IsNumber(runeValue)) {
					if firstDigit == "" {
						firstDigit = string(runeValue)
					} else {
						lastDigit = string(runeValue)
					}
					
					numLetters = ""
				} else if (!unicode.IsNumber(runeValue)) {
					numLetters += string(runeValue)
					re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")
					digitCharArr := re.FindAllString(numLetters, -1)
					var digit string = "0"
					if len(digitCharArr) > 0 {
						digit = strconv.Itoa(numMap[digitCharArr[0]])
					}

					if digit != "0" {
						if firstDigit == "" {
							firstDigit = digit
						} else {
							lastDigit = digit
						}
						
						numLetters = string(runeValue)
					}
				}
			}
		}

		if (lastDigit == "") {
			lastDigit = firstDigit
		}
		lineSum, err := strconv.Atoi(firstDigit + lastDigit)
		check(err)

		calValSum += lineSum
	}
	fmt.Println(calValSum)
	readFile.Close()
}