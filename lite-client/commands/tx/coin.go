package tx

import (
	"regexp"
	"github.com/pkg/errors"
	"strconv"
	"fmt"
	"strings"
	"sort"
)

type Coin struct {
	Denom	string 		`json:"denom"`
	Amount	int64		`json:"amount"`
}

func (coin Coin) String() string {
	return fmt.Sprintf("%v%v",coin.Amount, coin.Denom)
}

//regex code for extracting coins from string..
var reDenom = regexp.MustCompile("")
var reAmount = regexp.MustCompile("(\\d+)")
var reCoin = regexp.MustCompile("^([[:digit:]]+)[[:space:]]*([[:alpha:]]+)$")


func ParseCoin (str string) (Coin, error) {
	var coin Coin

	matches := reCoin.FindStringSubmatch(str)
	if matches == nil {
		return coin, errors.Errorf("%s is invalid coin definition", str)
	}

	// Parse the amount..
	amount, err := strconv.Atoi(matches[1])
	if err != nil {
		return coin, err
	}
	coin = Coin{ matches[2], int64(amount)}

	return coin, nil
}

// ------------------------
type Coins []Coin

func (coins Coins) Sort() {sort.Sort(coins)}	// Some functions need to be defined..

func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _,coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}

func ParseCoins(str string) (Coins, error) {
	if len(str) == 0 {
		return nil, nil
	}

	split := strings.Split(str, ",")
	var coins Coins

	for _, coin := range split {
		coin, err := ParseCoin(coin)
		if err != nil {
			return nil,err
		}
		coins = append(coins, coin)
	}

	coins.Sort()
	if !coins.IsValid() {
		return nil, errors.Errorf("Parsecoin Invalid..")
	}

}

func (coins Coins) IsValid() bool {
	return true
}
