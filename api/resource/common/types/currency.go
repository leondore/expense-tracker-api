package types

import "strconv"

type Currency uint64

const (
	Cent Currency = 1
	Base Currency = 100
)

func CurrencyFromString(s string) (Currency, error) {
	c, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return Currency(c) / Base, nil
}
