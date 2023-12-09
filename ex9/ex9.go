package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../ex9/data.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data := readDataToArr(file)

	resultSum := 0

	for _, line := range data {
		resultSum += processLine(line)
	}
	println(resultSum)
}

func processLine(line []int) int {
	lineUnderlines := [][]int{line}
	for {
		underline, allZero := calculateDifferences(line)
		lineUnderlines = append(lineUnderlines, underline)
		line = underline
		// if all elements are 0, break
		if allZero {
			break
		}
	}
	lastUnderlineIndex := len(lineUnderlines) - 1
	lineUnderlines[lastUnderlineIndex] = append(lineUnderlines[lastUnderlineIndex], 0)
	index := len(lineUnderlines) - 2

	increase := 0
	for index >= 0 {
		newElem := lineUnderlines[index][len(lineUnderlines[index])-1] + increase
		lineUnderlines[index] = append(lineUnderlines[index], newElem)
		increase = newElem
		index--
	}
	return lineUnderlines[0][len(lineUnderlines[0])-1]
}

func calculateDifferences(line []int) ([]int, bool) {
	allZero := true
	var underline []int
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]
		underline = append(underline, diff)
		if diff != 0 {
			allZero = false
		}
	}
	return underline, allZero
}

func readDataToArr(file *os.File) [][]int {
	var resultArray [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splittedLine := strings.Split(scanner.Text(), " ")
		var lineInt []int
		for _, e := range splittedLine {
			number, err := strconv.Atoi(e)
			if err != nil {
				log.Println(err)
			}
			lineInt = append(lineInt, number)
		}
		resultArray = append(resultArray, lineInt)
	}
	return resultArray
}
