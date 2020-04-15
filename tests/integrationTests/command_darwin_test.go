// +build darwin

package integrationTests

import (
	"goUniq/tests/testUtils"
	"testing"
)

func TestCommandFlagsAgainstCommand(t *testing.T) {
	var testCases = []struct {
		name string

		inFlags []string
	}{
		{"Regular", []string{}},
		{"CaseInsensitive", []string{"-i"}},
		//{"CountLines", []string{"-c"}}, // lines are prefixed with spaces
		{"Output only repeated lines", []string{"-d"}},
		{"Skip fields", []string{"-f 1"}},
		{"Skip chars", []string{"-s 5"}},
		{"Skip fields and chars", []string{"-f 2", "-s 5"}},
		{"Zero fields and chars", []string{"-f 0", "-s 5"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testUtils.RunCommandWithAgainstRealCommand(t, binaryName, tc.inFlags)
		})
	}
}


