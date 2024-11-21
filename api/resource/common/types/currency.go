package types

import "strconv"

type Currency uint64

const (
	Cent Currency = 1
	Base Currency = 100
)

func CurrencyFromString(s string) Currency {
	c, _ := strconv.ParseUint(s, 10, 64)
	return Currency(c) / Base
}
