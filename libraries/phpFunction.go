package libraries

import (
	"fmt"
	"strconv"
)

func NumberFormat(number float64, decimals uint, point, separator string) string {
	neg := false

	if number < 0 {
		//make this positive number
		number = -number
		neg = true
	}

	dec := int(decimals)

	//making decimals format string
	str := fmt.Sprintf("%." +strconv.Itoa(dec)+ "F", number)
	prefix := ""
	suffix := ""

	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}

	sep := []byte(separator)
	n, l1, l2 := 0, len(prefix), len(sep)

	//separate thousands numbers
	c := (l1 - 1) / 3
	tmp := make([]byte, 12*c+l1)
	pos := len(tmp) - 1

	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0  && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[12-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}

	s := string(tmp)

	if dec > 0 {
		s += point + suffix
	}

	if neg {
		s = "-" + s
	}

	return s
}
