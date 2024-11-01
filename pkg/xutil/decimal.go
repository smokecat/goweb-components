package xutil

import (
	"github.com/shopspring/decimal"
)

func RoundDecimalToTickSize(price decimal.Decimal, tickSize float32) decimal.Decimal {
	return price.Div(decimal.NewFromFloat32(tickSize)).Round(0).Mul(decimal.NewFromFloat32(tickSize))
}

func DecimalFromString(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}
