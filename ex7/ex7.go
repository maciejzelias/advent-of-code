package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("../ex7/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := mapDataToArray(file)

	cardsRanking := make(map[rune]int)
	cardsRanking['A'] = 14
	cardsRanking['K'] = 13
	cardsRanking['Q'] = 12
	cardsRanking['T'] = 10
	cardsRanking['9'] = 9
	cardsRanking['8'] = 8
	cardsRanking['7'] = 7
	cardsRanking['6'] = 6
	cardsRanking['5'] = 5
	cardsRanking['4'] = 4
	cardsRanking['3'] = 3
	cardsRanking['2'] = 2
	cardsRanking['J'] = 1

	var highCards []Card
	var onePairsCards []Card
	var twoPairsCards []Card
	var threeOfKindCards []Card
	var fullHouseCards []Card
	var fourOfKindCards []Card
	var fiveOfKindCards []Card

	for _, card := range data {

		count, differentChars := countRepetitions(card.hand)

		if count == 4 {
			fourOfKindCards = append(fourOfKindCards, card)
		} else if count == 3 && differentChars == 2 {
			fullHouseCards = append(fullHouseCards, card)
		} else if count == 3 && differentChars == 3 {
			threeOfKindCards = append(threeOfKindCards, card)
		} else if count == 2 && differentChars == 3 {
			twoPairsCards = append(twoPairsCards, card)
		} else if count == 2 && differentChars == 4 {
			onePairsCards = append(onePairsCards, card)
		} else {
			if count == 5 {
				fiveOfKindCards = append(fiveOfKindCards, card)
			} else {
				highCards = append(highCards, card)
			}
		}
	}

	cardDecks := [][]Card{
		highCards,
		onePairsCards,
		twoPairsCards,
		threeOfKindCards,
		fullHouseCards,
		fourOfKindCards,
		fiveOfKindCards,
	}

	rank := 1
	result := 0

	for _, deck := range cardDecks {
		sort.Slice(deck, func(i, j int) bool {
			return compareHands(deck[i].hand, deck[j].hand, cardsRanking)
		})
		for _, card := range deck {
			println(card.hand, card.bid, rank)
			result += card.bid * rank
			rank++
		}
	}
	println(result)
}

func compareHands(hand1, hand2 string, cardsRanking map[rune]int) bool {
	minLen := len(hand1)
	if len(hand2) < minLen {
		minLen = len(hand2)
	}

	for i := 0; i < minLen; i++ {
		rank1, rank2 := cardsRanking[rune(hand1[i])], cardsRanking[rune(hand2[i])]
		if rank1 != rank2 {
			return rank1 < rank2
		}
	}

	return len(hand1) > len(hand2)
}

type Card struct {
	hand string
	bid  int
}

func countRepetitions(str string) (int, int) {
	count := 0
	jokersCount := 0

	charCountMap := make(map[rune]int)

	for _, c := range str {
		if c == 'J' {
			jokersCount++
		} else {
			charCountMap[c]++
			if strings.Count(str, string(c)) > count {
				count = strings.Count(str, string(c))
			}
		}
	}

	return count + jokersCount, len(charCountMap)
}
func mapDataToArray(file *os.File) []Card {
	var resultArr []Card
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		bidNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		resultArr = append(resultArr, Card{
			hand: parts[0],
			bid:  bidNumber,
		})
	}

	// Check for errors during Scan.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return resultArr
}
