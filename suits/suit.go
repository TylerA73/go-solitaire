package suits

type Class int

type Suit struct {
	Class Class
	Name  string
	IsRed bool
}

var (
	Hearts Suit = Suit{
		Class: 0,
		Name:  "Hearts",
		IsRed: true,
	}

	Diamonds Suit = Suit{
		Class: 1,
		Name:  "Diamonds",
		IsRed: true,
	}

	Clubs Suit = Suit{
		Class: 2,
		Name:  "Clubs",
		IsRed: false,
	}

	Spades Suit = Suit{
		Class: 3,
		Name:  "Spades",
		IsRed: false,
	}
)

func (s *Suit) IsEqual(suit *Suit) bool {
	return s.Class == suit.Class
}
