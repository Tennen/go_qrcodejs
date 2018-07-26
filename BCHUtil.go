package BCHUtil

const (
	g15     = 1335
	g18     = 7973
	g15Mask = 21522
)

var g15BCHDigit = bchDigit(g15)
var g18BCHDigit = bchDigit(g18)

func BchTypeInfo(data int) int {
	d := data << 10
	for bchDigit(data) >= g15BCHDigit {
		d ^= g15 << (bchDigit(d) - g15BCHDigit)
	}
	return ((data << 10) | d) ^ g15Mask
}

func BchTypeNumber(data int) int {
	d := data << 12
	for bchDigit(d) >= g18BCHDigit {
		d ^= g18 << (bchDigit(d) - g18BCHDigit)
	}
	return (data << 12) | d
}

func bchDigit(data int) (digit uint) {
	d := uint(data)
	for d != 0 {
		digit += 1
		d <<= 1
	}
	return digit
}
