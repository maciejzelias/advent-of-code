package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	sequence := `LLRRLRRRLLRLRRLLLLRLRRLRRRLRLRRRLLRRLRRRLLRRLRRLRRLLRRRLRRLRRLRRRLRRLRLRLRRLRRLRRRLLRRLLLRRLRRRLRRRLRRRLRRLRRRLRLLRLRRRLRLRRLLRLRRRLRRRLRLRRRLRRRLRLRLRRLRRLRLRRLLRRRLRRRLRRRLLRRRLRLRLRLRLLRRRLRRRLRRLRRRLLRLRRLRRLRRRLRRRLRRLRLRLRRRLRRLRRLRRRLLRRLRLRLRRRLRLRLRRLRRLLRRLRRRLLRLLRLRLRRRR`

	file, err := os.Open("/Users/maciejzelias/Development/advent_of_code/ex8/data.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := readDataToMap(file)

	startingKeys := getStartingPoints(data)

	var restulNumbers []int

	for _, key := range startingKeys {
		counter := 0
		actualKey := key
		for actualKey[2] != 'Z' {
			if sequence[counter%len(sequence)] == 'L' {
				actualKey = data[actualKey].Left
			} else {
				actualKey = data[actualKey].Right
			}
			counter++
		}
		restulNumbers = append(restulNumbers, counter)
	}

	println(lcmArray(restulNumbers))
}

type Element struct {
	Left  string
	Right string
}

func getStartingPoints(data map[string]Element) []string {
	var startingPoints []string

	for key, _ := range data {
		if key[2] == 'A' {
			startingPoints = append(startingPoints, key)
		}
	}

	return startingPoints
}

func readDataToMap(file *os.File) map[string]Element {
	data := make(map[string]Element)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")

		key := strings.TrimSpace(parts[0])
		valuesStr := strings.TrimSpace(parts[1])
		trimmedValues := valuesStr[1 : len(valuesStr)-1]
		values := strings.Split(trimmedValues, ",")

		data[key] = Element{
			Left:  strings.TrimSpace(values[0]),
			Right: strings.TrimSpace(values[1]),
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmArray(arr []int) int {
	result := arr[0]
	for _, num := range arr[1:] {
		result = lcm(result, num)
	}
	return result
}
