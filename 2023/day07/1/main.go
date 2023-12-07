package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Card int

const (
	C2 Card = iota
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	T
	J
	Q
	K
	A
)

type Hand struct {
	handString string
	cards      []Card
	handType   HandType
	bid        int
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func parseInput(input []string) []Hand {
	hands := make([]Hand, 0)
	for _, line := range input {
		fields := strings.Fields(line)

		bid, _ := strconv.Atoi(fields[1])

		handString := fields[0]

		cardBytes := []byte(handString)
		cards := make([]Card, 0)
		cardMap := make(map[Card]int)
		for _, card := range cardBytes {
			switch card {
			case 'A':
				cardMap[A]++
				cards = append(cards, A)
			case 'K':
				cardMap[K]++
				cards = append(cards, K)
			case 'Q':
				cardMap[Q]++
				cards = append(cards, Q)
			case 'J':
				cardMap[J]++
				cards = append(cards, J)
			case 'T':
				cardMap[T]++
				cards = append(cards, T)
			case '9':
				cardMap[C9]++
				cards = append(cards, C9)
			case '8':
				cardMap[C8]++
				cards = append(cards, C8)
			case '7':
				cardMap[C7]++
				cards = append(cards, C7)
			case '6':
				cardMap[C6]++
				cards = append(cards, C6)
			case '5':
				cardMap[C5]++
				cards = append(cards, C5)
			case '4':
				cardMap[C4]++
				cards = append(cards, C4)
			case '3':
				cardMap[C3]++
				cards = append(cards, C3)
			case '2':
				cardMap[C2]++
				cards = append(cards, C2)
			}
		}

		handType := determineHandType(cardMap)

		hands = append(hands, Hand{
			handString,
			cards,
			handType,
			bid,
		})
	}

	return hands
}

func determineHandType(cardMap map[Card]int) HandType {
	if len(cardMap) == 1 {
		return FiveOfAKind
	} else if len(cardMap) == 2 {
		for _, count := range cardMap {
			if count == 1 || count == 4 {
				return FourOfAKind
			} else {
				return FullHouse
			}
		}
	} else if len(cardMap) == 3 {
		countMultiplication := 1
		for _, count := range cardMap {
			countMultiplication *= count
		}
		if countMultiplication == 3 {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	} else if len(cardMap) == 4 {
		return OnePair
	}
	return HighCard
}

// true if hand1 is bigger, else false
func compareHands(hand1, hand2 Hand) bool {
	if hand1.handType > hand2.handType {
		return true
	} else if hand1.handType < hand2.handType {
		return false
	}

	// hand1.handType == hand2.handType
	for i := 0; i < 5; i++ {
		if hand1.cards[i] > hand2.cards[i] {
			return true
		} else if hand1.cards[i] < hand2.cards[i] {
			return false
		}
	}

	// hand1 == hand2
	return true
}

func camelCards(input []string) {
	startTime := time.Now()

	hands := parseInput(input)

	for i := 0; i < len(hands); i++ {
		for j := i + 1; j < len(hands); j++ {
			if compareHands(hands[i], hands[j]) {
				hands[j], hands[i] = hands[i], hands[j]
			}
		}
	}

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}

	fmt.Println(totalWinnings)

	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func main() {
	file, err := os.Open("/home/bhavya5jain/projects/advent-of-code/2023/day07/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	camelCards(input)
}
