package converter

import "strings"

type converter struct {
	symbols []string
	values  []int
}

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

type Converter interface {
	IntToRoman(num int) string
}

func NewConverter(values []int, symbols []string) Converter {
	c := &converter{
		symbols: symbols,
		values:  values,
	}
	return c
}
