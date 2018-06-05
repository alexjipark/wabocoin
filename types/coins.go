package types

import "fmt"

type Coin struct {
	Denom  		string 		`json:"denom"`
	Amount 		int64 		`json:"amount"`
}

func (coin Coin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
}

type Coins 	[]Coin

func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}