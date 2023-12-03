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
	
	var gameNumSum int = 0
	for fileScanner.Scan() {
		maxColorNumMap := map[string]int {
			"red": 0,
			"blue": 0,
			"green": 0,
		}
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
				if colorAmt > maxColorNumMap[cubeInfoArr[1]] {
					maxColorNumMap[cubeInfoArr[1]] = colorAmt
				}
							
			}
			
		
		}

		var powerOfSet int = 1
		for _, amt := range maxColorNumMap {
			powerOfSet *= amt
		}
		
		gameNumSum += powerOfSet
	}
	
	fmt.Println(gameNumSum)
	readFile.Close()
}