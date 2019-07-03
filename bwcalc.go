package bwcalc

import "fmt"

var max uint32 = 6759999

func BwCalculate(n uint32) string {
	if n > max {
		panic("reach max size")
	}
	a := n / 10000
	return fmt.Sprintf("%s%s%04d", string('A'+a/26), string('A'+a%26), n%10000)
}
