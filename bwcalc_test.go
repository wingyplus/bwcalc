package bwcalc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBwCalc(t *testing.T) {
	var tcs = []struct {
		input    uint32
		expected string
	}{
		{1, "AA0001"},
		{1259, "AA1259"},
		{11259, "AB1259"},
		{251259, "AZ1259"},
		{261259, "BA1259"},
		{511259, "BZ1259"},
		{521259, "CA1259"},
		{6751259, "ZZ1259"},
		{6750000, "ZZ0000"},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.expected, BwCalculate(tc.input))
	}
}
