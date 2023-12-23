package year2023

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type HandType int32

const (
	HighCard     HandType = 0
	OnePair      HandType = 1
	TwoPair      HandType = 2
	ThreeOfAKind HandType = 3
	FullHouse    HandType = 4
	FourOfAKind  HandType = 5
	FiveOfAKind  HandType = 6
)

type Hand struct {
	cards         string
	countBySymbol map[string]int
	bid           int
	handType      HandType
}

func RunDay7() {
	fmt.Println("Day 7")
	file, err := os.Open("inputs/2023_07.txt")
	if nil != err {
		panic("Failed to ope file")
	}

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, parseHand(line))
	}

	handsPart1 := make([]Hand, len(hands))
	copy(handsPart1, hands)

	sort.SliceStable(handsPart1, func(i, j int) bool {
		return !hasHigherRank(handsPart1[i], handsPart1[j], symbolScorePart1)
	})

	answerPart1 := 0
	for i, hand := range handsPart1 {
		answerPart1 += (i + 1) * hand.bid
	}

	fmt.Printf("Answer 1: %d\n", answerPart1)

	handsPart2 := make([]Hand, len(hands))
	copy(handsPart2, hands)

	for i, hand := range handsPart2 {
		maxHandType := maximizeHand(hand)
		if maxHandType != hand.handType {
			handsPart2[i].handType = maxHandType
		}
	}

	sort.SliceStable(handsPart2, func(i, j int) bool {
		return !hasHigherRank(handsPart2[i], handsPart2[j], symbolScorePart2)
	})

	answerPart2 := 0
	for i, hand := range handsPart2 {
		answerPart2 += (i + 1) * hand.bid
	}

	fmt.Printf("Answer 2: %d\n", answerPart2)
}

func maximizeHand(hand Hand) HandType {
	jayCount := hand.countBySymbol["J"]
	if 0 == jayCount || FiveOfAKind == hand.handType {
		return hand.handType
	}

	if 1 == jayCount {
		switch hand.handType {
		case HighCard: // five different cards
			return OnePair
		case OnePair: // one pair and three different cards
			return ThreeOfAKind
		case TwoPair: // Two different pairs and the Jay
			return FullHouse
		case ThreeOfAKind: //
			return FourOfAKind
		case FourOfAKind:
			return FiveOfAKind
		}
	} else if 2 == jayCount {
		switch hand.handType {
		case OnePair: // three different cards and two J
			return ThreeOfAKind
		case TwoPair: // One different pair, one J-Pair and a single card
			return FourOfAKind
		case FullHouse: // One J-Pair and three of the same symbol
			return FiveOfAKind
		}
	} else if 3 == jayCount {
		switch hand.handType {
		case FullHouse: // other is a pair
			return FiveOfAKind
		default: // two different cards
			return FourOfAKind
		}
	} else if 4 == jayCount {
		return FiveOfAKind
	}

	return hand.handType
}

func parseHand(line string) Hand {
	bid, _ := strconv.Atoi(line[6:])
	countBySymbol, handType := extractHandType(line)
	return Hand{line[0:5], countBySymbol, bid, handType}
}

func extractHandType(line string) (map[string]int, HandType) {
	countBySymbol := make(map[string]int)
	for _, ch := range line[0:5] {
		countBySymbol[string(ch)] += 1
	}

	numHandTypes := make(map[HandType]int)
	for _, v := range countBySymbol {
		switch v {
		case 0:
			continue
		case 1:
			numHandTypes[HighCard] += 1
		case 2:
			numHandTypes[OnePair] += 1
		case 3:
			numHandTypes[ThreeOfAKind] += 1
		case 4:
			numHandTypes[FourOfAKind] += 1
		case 5:
			numHandTypes[FiveOfAKind] += 1
		}
	}

	var handType = HighCard
	if 1 == numHandTypes[FiveOfAKind] {
		handType = FiveOfAKind
	} else if 1 == numHandTypes[FourOfAKind] {
		handType = FourOfAKind
	} else if 1 == numHandTypes[ThreeOfAKind] {
		if 1 == numHandTypes[OnePair] {
			handType = FullHouse
		} else {
			handType = ThreeOfAKind
		}
	} else if 0 < numHandTypes[OnePair] {
		if 1 == numHandTypes[OnePair] {
			handType = OnePair
		} else {
			handType = TwoPair
		}
	} else if 1 == numHandTypes[OnePair] {
		handType = OnePair
	}
	return countBySymbol, handType
}

// Returns `true` if the first hand has a higher rank then the second
func hasHigherRank(first Hand, second Hand, scoreFn func(symbol uint8) int) bool {
	if first.handType != second.handType {
		return first.handType > second.handType
	}

	for i := 0; i < len(first.cards); i++ {
		firstScore := scoreFn(first.cards[i])
		secondScore := scoreFn(second.cards[i])
		if firstScore == secondScore {
			continue
		} else {
			return firstScore > secondScore
		}
	}

	return true
}

func symbolScorePart1(symbol uint8) int {
	score := 0
	switch symbol {
	case 'A':
		score = 14
	case 'K':
		score = 13
	case 'Q':
		score = 12
	case 'J':
		score = 11
	case 'T':
		score = 10
	default:
		score, _ = strconv.Atoi(string(symbol))
	}

	return score
}

func symbolScorePart2(symbol uint8) int {
	score := 0
	switch symbol {
	case 'A':
		score = 14
	case 'K':
		score = 13
	case 'Q':
		score = 12
	case 'J':
		score = 0
	case 'T':
		score = 10
	default:
		score, _ = strconv.Atoi(string(symbol))
	}

	return score
}

func handTypeToStr(handType HandType) string {
	switch handType {
	case HighCard:
		return "HC"
	case OnePair:
		return "1P"
	case TwoPair:
		return "2P"
	case ThreeOfAKind:
		return "3K"
	case FullHouse:
		return "FH"
	case FourOfAKind:
		return "4K"
	case FiveOfAKind:
		return "5K"
	default:
		return "--"
	}
}
