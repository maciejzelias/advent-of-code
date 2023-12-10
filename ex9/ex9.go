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

	part1Sum := 0
	part2Sum := 0

	for _, line := range data {
		leftNumber, rightNumber := processLine(line)
		part1Sum += rightNumber
		part2Sum += leftNumber
	}
	println(part1Sum)
	println(part2Sum)
}

func processLine(line []int) (int, int) {
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

	//part 1
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

	//part 2
	lineUnderlines[lastUnderlineIndex] = append([]int{0}, lineUnderlines[lastUnderlineIndex]...)
	index = len(lineUnderlines) - 2

	decrease := 0
	for index >= 0 {
		newElem := lineUnderlines[index][0] - decrease
		lineUnderlines[index] = append([]int{newElem}, lineUnderlines[index]...)
		decrease = newElem
		index--
	}

	return lineUnderlines[0][len(lineUnderlines[0])-1], lineUnderlines[0][0]
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
