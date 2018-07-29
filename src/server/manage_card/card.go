package manage_card

import (
	"errors"
	"server/msg"
)

type Card struct {
	PlayerId int64
	CardType msg.CardTypeId
}

func NewCard(playerId int64, cardType msg.CardTypeId) *Card {
	return &Card{
		PlayerId: playerId,
		CardType: cardType,
	}
}

func (m *Card) Duel(cardB *Card) (result int8, err error) {
	if !m.IsRightCardType() || !cardB.IsRightCardType() {
		err = errors.New("card type not right")
		return
	}
	return m.duel(m, cardB)
}

func (m *Card) IsRightCardType() bool {
	if m.CardType != msg.Card_Type_Jan || m.CardType != msg.Card_Type_Ken || m.CardType != msg.Card_Type_Po {
		return false
	}
	return true
}

func (m *Card) duel(cardTypeA, cardTypeB msg.CardTypeId) int8 {
	if cardTypeA == cardTypeB {
		return 0
	}
	switch cardTypeA {
	case msg.Card_Type_Jan:
		if cardTypeB == msg.Card_Type_Po {
			return 1
		} else {
			return -1
		}
	case msg.Card_Type_Ken:
		if cardTypeB == msg.Card_Type_Jan {
			return 1
		} else {
			return -1
		}
	case msg.Card_Type_Po:
		if cardTypeB == msg.Card_Type_Ken {
			return 1
		} else {
			return -1
		}
	}
}
