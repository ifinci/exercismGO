package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
    switch card {
    case "ace":
        value = 11
    case "two":
        value = 2
    case "three":
        value = 3
    case "four":
        value = 4
    case "five":
        value = 5
    case "six":
        value = 6
    case "seven":
        value = 7
    case "eight":
        value = 8
    case "nine":
        value = 9
    case "ten":
        value = 10
    case "jack":
        value = 10
    case "queen":
        value = 10
    case "king":
        value = 10
    default:
        value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	const SPLIT = "P"
    const STAND = "S"
    const HIT = "H"
    const AUTOMATIC_WIN = "W"

    // if we have 2 Aces - we must split
    if card1 == "ace" && card2 == "ace" {
        return SPLIT
    }
    var cardsValue = ParseCard(card1) + ParseCard(card2)
    var dealerValue = ParseCard(dealerCard)

    switch {
        case cardsValue == 21:
        	if dealerValue < 10 {
        		return AUTOMATIC_WIN
        	} else {
          	  return STAND
        	}
        case cardsValue >= 17:
        	return STAND
        case cardsValue >= 12:
            if dealerValue >= 7 {
            	return HIT
        	} else {
            	return STAND
        	}
        default:
        	return HIT        	
    }
}
