package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
// ace		11	eight	8
// two		2	nine	9
// three	3	ten		10
// four		4	jack	10
// five		5	queen	10
// six		6	king	10
// seven	7	other	0
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
	var response int

	switch(card) {
		case "ace":
			response = 11
		case "two":
			response = 2
		case "three":
			response = 3
		case "eight":
			response = 8
		case "nine":
			response = 9
		case "ten":
			response = 10
		case "four":
			response = 4
		case "jack":
			response = 10
		case "five":
			response = 5
		case "queen":
			response = 10
		case "six":
			response = 6
		case "king":
			response = 10
		case "seven":
			response = 7
		default:
			response = 0
	}

	return response
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	// panic("Please implement the IsBlackjack function")

	if (card1 == "ace" && card2 == "ace") {
		return false
	} else if total := ParseCard(card1) + ParseCard(card2); total >= 21 {
		return true
	} else {
		return false
	}
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	// panic("Please implement the LargeHand function")

	if isBlackjack && dealerScore < 10 {
		return "W"
	} else if !isBlackjack {
		return "P"
	} else {
		return "S"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	// panic("Please implement the SmallHand function")

	if handScore >= 17 {
		return "S"
	} else if handScore <= 11 {
		return "H"
	} else {
		if dealerScore <= 6 {
			return "S"
		} else {
			return "H"
		}
	}
}
