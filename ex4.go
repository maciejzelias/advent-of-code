package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := ``

	winningNumbers, numbers, err := extractArrays(input)
	if err != nil {
		fmt.Println(err)
	}
	println(winningNumbers)

	println(numbers)

	score := 0

	for i, winningNumbers := range winningNumbers {
		score += getCardScore(winningNumbers, numbers[i])
	}

	print(score)
}

func extractArrays(input string) ([][]int, [][]int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var winningNumbers [][]int
	var numbers [][]int

	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("error spliting lines")
		}

		firstPart := strings.TrimSpace(parts[0][strings.Index(parts[0], ":")+1:])
		secondPart := strings.TrimSpace(parts[1])

		firsPartArr, err := parseIntoIntArray(firstPart)
		if err != nil {
			return nil, nil, err
		}
		secondPartArr, err := parseIntoIntArray(secondPart)
		if err != nil {
			return nil, nil, err
		}
		winningNumbers = append(winningNumbers, firsPartArr)
		numbers = append(numbers, secondPartArr)

		lineIndex++
	}
	return winningNumbers, numbers, nil
}

func parseIntoIntArray(strArr string) ([]int, error) {
	numStrings := strings.Fields(strArr)

	var numbers []int
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func isInArray(target int, arr []int) bool {
	for _, num := range arr {
		if target == num {
			return true
		}
	}
	return false
}

func getCardScore(winningNumbers []int, numbers []int) int {
	card_score := 1
	for i := 0; i < len(numbers); i++ {
		if isInArray(numbers[i], winningNumbers) {
			card_score *= 2
		}
	}

	if card_score == 1 {
		return 0
	}

	return card_score / 2
}
