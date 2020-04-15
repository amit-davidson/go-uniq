package integrationTests

import (
	"fmt"
	"github.com/amitdavidson234/go-uniq/pkg/utils"
	"github.com/amitdavidson234/go-uniq/tests/testUtils"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"testing"
)

func TestCommandFlagsWithGoldenFiles(t *testing.T) {
	var testCases = []struct {
		name string

		inFlags []string
	}{
		{"Regular", []string{}},
		{"CaseInsensitive", []string{"-i"}},
		{"CountLines", []string{"-c"}},
		{"OutputOnlyRepeatedLines", []string{"-d"}},
		{"SkipFields", []string{"-f 1"}},
		{"SkipChars", []string{"-s 5"}},
		{"SkipFieldsAndChars", []string{"-f 2", "-s 5"}},
		{"ZeroFieldsAndChars", []string{"-f 0", "-s 5"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testUtils.RunCommandWithGoldenFile(t, tc.name, binaryName, *update, tc.inFlags)
		})
	}
}

func TestCommandFlagsWithFileOutput(t *testing.T) {
	fileOutputName := fmt.Sprintf("%v", rand.Int())
	goldenFileOutputName := fmt.Sprintf("testdata/%sOutput.golden", t.Name())

	dirPath, err := os.Getwd()
	require.NoError(t, err)
	executable := fmt.Sprintf("%s/%s", dirPath, binaryName)
	fileInputName := fmt.Sprintf("%s/testdata/input.txt", dirPath)

	command := testUtils.BuildCommand(executable, []string{}, []string{fileInputName, fileOutputName})
	testUtils.RunCommandWithAssertion(t, command)
	defer func() {
		err = utils.RemoveFile(fileOutputName)
		require.NoError(t, err)
	}()

	data, err := utils.ReadFile(fmt.Sprintf("%s/%s", dirPath, fileOutputName))
	testUtils.UpdateFile(t, goldenFileOutputName, data, *update)

	goldenData, err := utils.ReadFile(goldenFileOutputName)
	require.NoError(t, err)

	require.Equal(t, goldenData, data)
}
