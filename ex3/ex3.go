package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("../ex3/data.txt")

	if err != nil {
		log.Fatal(err)
	}
	inputArr := readAsString(file)

	getEdgeLinesAdjustemntNumbers(inputArr[0], inputArr[1])

	// for _, line := range inputArr {

	// }

	defer file.Close()
}

func getEdgeLinesAdjustemntNumbers(line string, nextLine string) {
	var resultArray []int

	numbers := getLineNumbers(line)

	for _, number := range numbers {
		index := strings.Index(line, number)
		if index == -1 {
			log.Fatal("Number not found")
		}
		length := len(number)

		diff := len(nextLine) - len(line)

		// check left side
		if index > 0 && (unicode.IsSymbol(rune(line[index-1])) || unicode.IsSymbol(rune(nextLine[index-1+diff]))) {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			resultArray = append(resultArray, numberInt)
			line = line[index+length+1:]
			continue
		}

		// check down/up

		for i, _ := range number {
		}
	}

}

func getLineNumbers(line string) []string {
	re := regexp.MustCompile(`\D`)
	onlyDigits := re.ReplaceAllString(line, " ")

	spaceRe := regexp.MustCompile(`\s+`)
	removedSpaces := spaceRe.ReplaceAllString(onlyDigits, " ")

	return strings.Split(strings.TrimSpace(removedSpaces), " ")
}

func readAsString(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	var resultArray []string

	for scanner.Scan() {
		resultArray = append(resultArray, scanner.Text())
	}

	return resultArray
}
