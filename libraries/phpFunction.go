package libraries

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberFormat(number float64, decimals uint, point, separator string) string {
	var numbers []string
	var result string
	negative := false

	if number < 0 {
		number = -number
		negative = true
	}

	dec := int(decimals)
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)

	strplit  := strings.Split(str, ".")
	numeral  := strplit[0]
	numfloat := strplit[1]

	lnumeral := len(numeral)

	if lnumeral%3 != 0 {
		numbers = append(numbers, numeral[:(lnumeral%3)])
		numeral = numeral[(lnumeral%3):lnumeral]
	}

	for i, j := 0, 3; i < len(numeral); i++ {
		numbers = append(numbers, numeral[i:(j+i)])
		i += (3-1)
	}

	result  = strings.Join(numbers, separator)
	result += point + numfloat
	if negative {result = "-" + result}

	return result
}
