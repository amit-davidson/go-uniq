package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateLinesFromInput(t *testing.T) {
	var testCases = []struct {
		name   string
		inData []byte

		out *Line
	}{
		{"empty file", []byte{}, nil},
		{"single newline", []byte{0x0a}, &Line{OriginalText: "", ComparisonText: "", Count: 0}},
		{"file not ending with new line", []byte{0x57, 0x0a, 0x57}, &Line{OriginalText: string(0x57), ComparisonText: string(0x57)}},
		{"regular text", []byte{0x57, 0x0a, 0x57, 0x0a}, &Line{OriginalText: string(0x57), ComparisonText: string(0x57), Next: &Line{OriginalText: string(0x57), ComparisonText: string(0x57)}}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := CreateLinesFromInput(tc.inData)
			assert.Equal(t, tc.out, res)
		})
	}
}

func BenchmarkCreateLinesFromInput(b *testing.B) {
	var benchmarks = []struct {
		name   string
		inData []byte
	}{
		{"empty file", []byte{}},
		{"single newline", []byte{0x0a}},
		{"file not ending with new line", []byte{0x57, 0x0a, 0x57}},
		{"regular text", []byte{0x57, 0x0a, 0x57, 0x0a}},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CreateLinesFromInput(bm.inData)
			}
		})
	}
}

func TestCreateLines(t *testing.T) {
	var testCases = []struct {
		name   string
		inData []string

		out *Line
	}{
		{"one line", []string{"W"}, &Line{OriginalText: "W", ComparisonText: "W", Count: 0}},
		{"multiple lines", []string{"W", "W"}, &Line{OriginalText: "W", ComparisonText: "W", Count: 0, Next: &Line{OriginalText: "W", ComparisonText: "W", Count: 0}}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := CreateLines(tc.inData)
			assert.Equal(t, tc.out, res)
		})
	}
}

func BenchmarkCreateLines(b *testing.B) {
	var benchmarks = []struct {
		name   string
		inData []string
	}{
		{"one line", []string{"W"}},
		{"multiple lines single character words", []string{"W", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T", "W", "T"}}, // when run with -test.benchmem number of allocs/op should be 1
		{"multiple lines multiple characters words", []string{"I", "Love", "EATING", "Pizza", "and", "hamburger", "GGGGG", "213OKO", "klakdad", "lplhgpgh", "dlakld", "mzcnxz", "T123132", "Wasdkad", "Taskdl", "asdasdasW", "l;lpl1023132T", "---0-23W", "Tasdad", "Wsqdas213", "T[p][]q[q]", "W.;pqdsl=123", "1=p;3lx-T", "W{}{T", "Tpaslpd12"}}, // when run with -test.benchmem number of allocs/op should be 1
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CreateLines(bm.inData)
			}
		})
	}
}
