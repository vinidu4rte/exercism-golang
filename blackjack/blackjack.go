package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores 21 points or larger.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack {
		// Split (pair of aces)
		return "P"
	}
	if isBlackjack && dealerScore < 10 {
		// Automatic win
		return "W"
	}
	// Stand and wait
	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore int, dealerScore int) string {
	if handScore >= 17 {
		// Stand when cards are 17 or higher
		return "S"
	} else if handScore <= 11 {
		// Hit when cards are 11 or lower
		return "H"
	} else if dealerScore >= 7 {
		// Hit when cards are 12..16 and dealer is showing 7 or higher
		return "H"
	} else {
		// Stand when cards are 12..16 and dealer is showing 6 or lower
		return "S"
	}
}

// FirstTurn returns the semi-optimal decision for the first turn, given the cards of the player and the dealer.
// This function is already implemented and does not need to be edited. It pulls the other functions together in a
// complete decision tree for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	if handScore > 20 {
		return LargeHand(IsBlackjack(card1, card2), dealerScore)
	}

	return SmallHand(handScore, dealerScore)
}
