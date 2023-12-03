package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
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
	
	var calValSum int = 0
	for fileScanner.Scan() {
		re := regexp.MustCompile("[0-9]")
		digitCharArr := re.FindAllString(fileScanner.Text(), -1)
		lineSum, err := strconv.Atoi(digitCharArr[0] + digitCharArr[len(digitCharArr) - 1])
		check(err)
		calValSum += lineSum
	}
	
	fmt.Println(calValSum)
	readFile.Close()
}