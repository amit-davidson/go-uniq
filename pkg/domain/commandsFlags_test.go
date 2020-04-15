package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandsFlags_ToLower(t *testing.T) {
	var testCases = []struct {
		name string

		in *Line

		out *Line
	}{
		{"Multiple lines", CreateLines([]string{"LongLIne1", "LongLIne2"}), &Line{ComparisonText: "longline1", OriginalText: "LongLIne1", Next: &Line{ComparisonText: "longline2", OriginalText: "LongLIne2"}}},
		{"Regular", CreateLines([]string{"UpperCasedLine"}), &Line{ComparisonText: "uppercasedline", OriginalText: "UpperCasedLine"}},
		{"All lower cased", CreateLines([]string{"uppercasedline"}), CreateLines([]string{"uppercasedline"})},
		{"Multiple words", CreateLines([]string{"Upper cased Line"}), &Line{ComparisonText: "upper cased line", OriginalText: "Upper cased Line"}},
	}

	cf := CommandsFlags{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cf.ToLower(tc.in)
			assert.True(t, AreLinesEqual(tc.out, tc.in))
		})
	}
}

func BenchmarkCommandsFlags_ToLower(b *testing.B) {
	var benchmarks = []struct {
		name string

		in *Line
	}{
		{"Regular", CreateLines([]string{"UpperCasedLine"})},
		{"All lower cased", CreateLines([]string{"uppercasedline"})},
		{"Multiple words", CreateLines([]string{"Upper cased Line"})},
		{"Multiple lines", CreateLines([]string{"LongLIne1", "LongLIne2"})},
	}

	cf := CommandsFlags{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cf.ToLower(bm.in)
			}
		})
	}
}

func TestCommandsFlags_GetRepeatedLines(t *testing.T) {
	var testCases = []struct {
		name string

		in *Line

		out *Line
	}{
		{"Both repeated lines and non-repeated", &Line{Count: 2, Next: &Line{Count: 3, Next: &Line{Count: 1}}}, &Line{Count: 2, Next: &Line{Count: 3}}},
		{"Last is non repeated", &Line{Count: 2, Next: &Line{Count: 1}}, &Line{Count: 2}},
		{"First is non repeated", &Line{Count: 1, Next: &Line{Count: 3}}, &Line{Count: 3}},
		{"Only repeated lines", &Line{Count: 4, Next: &Line{Count: 5}}, &Line{Count: 4, Next: &Line{Count: 5}}},
		{"Only non-repeated lines", &Line{Count: 1}, nil},
		{"No lines", nil, nil},
	}

	cf := CommandsFlags{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := cf.GetRepeatedLines(tc.in)
			assert.True(t, AreLinesEqual(tc.out, res))
		})
	}
}

func BenchmarkCommandsFlags_GetRepeatedLines(b *testing.B) {
	var benchmarks = []struct {
		name string

		in *Line
	}{
		{"Both repeated lines and non-repeated", &Line{Count: 2, Next: &Line{Count: 3, Next: &Line{Count: 1}}}},
		{"Last is non repeated", &Line{Count: 2, Next: &Line{Count: 1}}},
		{"First is non repeated", &Line{Count: 1, Next: &Line{Count: 3}}},
		{"Only repeated lines", &Line{Count: 4, Next: &Line{Count: 5}}},
		{"Only non-repeated lines", &Line{Count: 1}},
		{"No lines", nil},
	}

	cf := CommandsFlags{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cf.GetRepeatedLines(bm.in)
			}
		})
	}
}

func TestCommandsFlags_GetNonRepeatedLines(t *testing.T) {
	var testCases = []struct {
		name string

		in *Line

		out *Line
	}{
		{"Both repeated lines and non-repeated", &Line{Count: 2, Next: &Line{Count: 3, Next: &Line{Count: 1}}}, &Line{Count: 1}},
		{"Last is non repeated", &Line{Count: 2, Next: &Line{Count: 1}}, &Line{Count: 1}},
		{"First is non repeated", &Line{Count: 1, Next: &Line{Count: 3}}, &Line{Count: 1}},
		{"Only repeated lines", &Line{Count: 4, Next: &Line{Count: 5}}, nil},
		{"Only non-repeated lines", &Line{Count: 1}, &Line{Count: 1}},
		{"No lines", nil, nil},
	}

	cf := CommandsFlags{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := cf.GetNonRepeatedLines(tc.in)
			assert.Equal(t, tc.out, res)
		})
	}
}

func BenchmarkCommandsFlags_GetNonRepeatedLines(b *testing.B) {
	var benchmarks = []struct {
		name string

		in *Line
	}{
		{"Both repeated lines and non-repeated", &Line{Count: 2, Next: &Line{Count: 3, Next: &Line{Count: 1}}}},
		{"Last is non repeated", &Line{Count: 2, Next: &Line{Count: 1}}},
		{"First is non repeated", &Line{Count: 1, Next: &Line{Count: 3}}},
		{"Only repeated lines", &Line{Count: 4, Next: &Line{Count: 5}}},
		{"Only non-repeated lines", &Line{Count: 1}},
		{"No lines", nil},
	}

	cf := CommandsFlags{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cf.GetNonRepeatedLines(bm.in)
			}
		})
	}
}

func TestCommandsFlags_TrimLineByField(t *testing.T) {
	var testCases = []struct {
		name string

		inLines                  *Line
		inNumberOfFieldsToIgnore int

		out *Line
	}{
		{"All ignored", CreateLines([]string{"UpperCasedLine"}), 2, &Line{ComparisonText: "", OriginalText: "UpperCasedLine"}},
		{"Partially ignored", CreateLines([]string{"UpperCasedLine", "Partially ignored"}), 1, &Line{ComparisonText: "", OriginalText: "UpperCasedLine", Next: &Line{ComparisonText: "ignored", OriginalText: "Partially ignored"}}},
		{"Nothing ignored", CreateLines([]string{"UpperCasedLine", "Not ignored"}), 0, &Line{ComparisonText: "UpperCasedLine", OriginalText: "UpperCasedLine", Next: &Line{ComparisonText: "Not ignored", OriginalText: "Not ignored"}}},
	}

	cf := CommandsFlags{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cf.TrimLineByField(tc.inLines, tc.inNumberOfFieldsToIgnore)
			assert.True(t, AreLinesEqual(tc.out, tc.inLines))
		})
	}
}

func BenchmarkCommandsFlags_TrimLineByField(b *testing.B) {
	var benchmarks = []struct {
		name string

		inLines                  *Line
		inNumberOfFieldsToIgnore int
	}{
		{"All ignored", CreateLines([]string{"UpperCasedLine"}), 2},
		{"Partially ignored", CreateLines([]string{"UpperCasedLine", "Partially ignored"}), 1},
		{"Nothing ignored", CreateLines([]string{"UpperCasedLine", "Not ignored"}), 0},
	}

	cf := CommandsFlags{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cf.TrimLineByField(bm.inLines, bm.inNumberOfFieldsToIgnore)
			}
		})
	}
}

//
func TestCommandsFlags_TrimLineByChar(t *testing.T) {
	var testCases = []struct {
		name string

		inLines                 *Line
		inNumberOfCharsToIgnore int

		out *Line
	}{
		{"All ignored", CreateLines([]string{"UpperCasedLine", "Completely ignored"}), 5, &Line{ComparisonText: "CasedLine", OriginalText: "UpperCasedLine", Next: &Line{ComparisonText: "etely ignored", OriginalText: "Completely ignored"}}},
		{"input is shorter than ignore number", CreateLines([]string{"Uppe", "Comp"}), 5, &Line{OriginalText: "Uppe", Next: &Line{OriginalText: "Comp"}}},
		{"Not ignoring anything", CreateLines([]string{"UpperCasedLine", "Not ignored"}), 0, &Line{ComparisonText: "UpperCasedLine", OriginalText: "UpperCasedLine", Next: &Line{ComparisonText: "Not ignored", OriginalText: "Not ignored"}}},
	}

	cf := CommandsFlags{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cf.TrimLineByChar(tc.inLines, tc.inNumberOfCharsToIgnore)
			assert.True(t, AreLinesEqual(tc.out, tc.inLines))
		})
	}
}

func BenchmarkCommandsFlags_TrimLineByChar(b *testing.B) {
	var benchmarks = []struct {
		name string

		inLines                 *Line
		inNumberOfCharsToIgnore int
	}{
		{"All ignored", CreateLines([]string{"UpperCasedLine", "Completely ignored"}), 5},
		{"input is shorter than ignore number", CreateLines([]string{"Uppe", "Comp"}), 5},
		{"Not ignoring anything", CreateLines([]string{"UpperCasedLine", "Not ignored"}), 0},
	}

	cf := CommandsFlags{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cf.TrimLineByChar(bm.inLines, bm.inNumberOfCharsToIgnore)
			}
		})
	}
}

func TestCommandsFlags_CollapseLines(t *testing.T) {
	var testCases = []struct {
		name string

		in *Line

		out *Line
	}{
		{"No collapse", CreateLines([]string{"UpperCasedLine", "Completely ignored"}), &Line{ComparisonText: "CasedLine", OriginalText: "UpperCasedLine", Count: 1, Next: &Line{ComparisonText: "etely ignored", OriginalText: "Completely ignored", Count: 1}}},
		{"Collapse", CreateLines([]string{"UpperCasedLine", "UpperCasedLine"}), &Line{ComparisonText: "CasedLine", OriginalText: "UpperCasedLine", Count: 2}},
		{"Multiple collapse", CreateLines([]string{"UpperCasedLine", "UpperCasedLine", "UpperCasedLine"}), &Line{ComparisonText: "CasedLine", OriginalText: "UpperCasedLine", Count: 3}},
	}

	cf := CommandsFlags{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cf.CollapseLines(tc.in)
			assert.True(t, AreLinesEqual(tc.out, tc.out))
		})
	}
}

func BenchmarkCommandsFlags_CollapseLines(b *testing.B) {
	var benchmarks = []struct {
		name string

		in *Line
	}{
		{"No collapse", CreateLines([]string{"UpperCasedLine", "Completely ignored"})},
		{"Collapse", CreateLines([]string{"UpperCasedLine", "UpperCasedLine"})},
		{"Multiple collapse", CreateLines([]string{"UpperCasedLine", "UpperCasedLine", "UpperCasedLine"})},
	}

	cf := CommandsFlags{}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cf.CollapseLines(bm.in)
			}
		})
	}
}
