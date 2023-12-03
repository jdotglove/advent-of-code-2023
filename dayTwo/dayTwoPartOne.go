package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func check (e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	maxColorNumMap := map[string]int {
		"red": 12,
		"blue": 14,
		"green": 13,
	}
	var gameNumSum int = 0
	for fileScanner.Scan() {
		var validGame bool = true
		splitArr := strings.Split(fileScanner.Text(), ": ")
		gameNumInfo := splitArr[0]
		cubeGameNumInfo := splitArr[1]
		cubeGameSetArr := strings.Split(cubeGameNumInfo, "; ")

		for _, cubeSet := range cubeGameSetArr {
			cubeCountArr := strings.Split(cubeSet, ", ")

			for _, cubeCount := range cubeCountArr {
				cubeInfoArr := strings.Split(cubeCount, " ")
				colorAmt, err1 := strconv.Atoi(cubeInfoArr[0])
				check(err1)
				maxColorAmt := maxColorNumMap[cubeInfoArr[1]]

				if colorAmt > maxColorAmt {
					validGame = false
					break
				}
			}
			
			if !validGame {
				break
			}
		}

		if validGame {
			gameInfoArr := strings.Split(gameNumInfo, " ")
			gameNum, err := strconv.Atoi(gameInfoArr[1])
			check(err)

			gameNumSum += gameNum
		}
	}
	
	fmt.Println(gameNumSum)
	readFile.Close()
}