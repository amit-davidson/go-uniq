package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SplitStringByCharIndex(t *testing.T) {
	var testCases = []struct {
		name string

		inString string
		inChar   rune
		inIndex  int

		outFirstPart  string
		outSecondPart string
	}{
		{"Regular", "This is a long string", 0x0020, 4, "This is a long", "string"},
		{"Index out of range before", "This is a long string", 0x0020, 0, "", "This is a long string"},
		{"Index out of range after", "This is a long string", 0x0020, 6, "This is a long string", ""},
		{"Char doesn't exist", "This is a long string", 0x0021, 2, "This is a long string", ""},
		{"Empty string", "", 0x0020, 2, "", ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			firstPart, secondPart := SplitStringByCharIndex(tc.inString, tc.inChar, tc.inIndex)
			assert.Equal(t, tc.outFirstPart, firstPart)
			assert.Equal(t, tc.outSecondPart, secondPart)
		})
	}
}

func Benchmark_SplitStringByCharIndex(b *testing.B) {
	var benchmarks = []struct {
		name string

		inString string
		inChar   rune
		inIndex  int
	}{
		{"Regular", "This is a long string", 0x0020, 4},
		{"Index out of range before", "This is a long string", 0x0020, 0},
		{"Index out of range after", "This is a long string", 0x0020, 6},
		{"Char doesn't exist", "This is a long string", 0x0021, 2},
		{"Empty string", "", 0x0020, 2},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SplitStringByCharIndex(bm.inString, bm.inChar, bm.inIndex)
			}
		})
	}
}
