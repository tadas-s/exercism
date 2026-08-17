package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cardValue := 0

	switch card {
	case "two":
		cardValue = 2
	case "three":
		cardValue = 3
	case "four":
		cardValue = 4
	case "five":
		cardValue = 5
	case "six":
		cardValue = 6
	case "seven":
		cardValue = 7
	case "eight":
		cardValue = 8
	case "nine":
		cardValue = 9
	case "ten", "jack", "queen", "king":
		cardValue = 10
	case "ace":
		cardValue = 11
	case "other":
		cardValue = 0
	}

	return cardValue
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case cardsValue == 21 && dealerValue >= 10:
		return "S"
	case cardsValue == 21:
		return "W"
	case cardsValue >= 17 && cardsValue <= 20:
		return "S"
	case cardsValue >= 12 && cardsValue <= 16 && dealerValue >= 7:
		return "H"
	case cardsValue >= 12 && cardsValue <= 16:
		return "S"
	default:
		return "H"
	}
}
