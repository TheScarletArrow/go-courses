package converter

import "strings"

type converter struct {
	symbols []string
	values  []int
}
type Converter interface {
	IntToRoman(num int) string
}

// IntToRoman converts integer to roman
func (c *converter) IntToRoman(num int) string {
	var roman strings.Builder

	for i, val := range c.values {
		for num >= val {
			roman.WriteString(c.symbols[i])
			num -= val
		}
	}

	return roman.String()
}
func NewConverter(values []int, symbols []string) Converter {
	return &converter{
		symbols: symbols,
		values:  values,
	}
}
