package testUtils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"goUniq/pkg/utils"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

func RunCommand(command []string) ([]byte, error) {
	runCmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s", strings.Join(command, " ")))
	output, err := runCmd.CombinedOutput()
	return output, err
}

func RunCommandWithAssertion(t *testing.T, commandArgs []string) []byte {
	output, err := RunCommand(commandArgs)
	require.NoError(t, err)
	return output
}

func BuildCommand(executable string, flags []string, arguments []string) []string {
	packageCommand := []string{executable}
	if len(flags) > 0 {
		packageCommand = append(packageCommand, flags...)
	}
	if len(arguments) > 0 {
		packageCommand = append(packageCommand, "--")
		packageCommand = append(packageCommand, arguments...)
	}
	return packageCommand
}

func UpdateFile(t *testing.T, path string, data []byte, update bool) {
	if update {
		f, err := utils.CreateFile(path)
		require.NoError(t, err)
		err = f.Close()
		require.NoError(t, err)
		err = utils.WriteFile(f, data)
	}
}

func RunCommandWithGoldenFile(t *testing.T, testName string, binaryName string, update bool, flags []string) {
	fileOutputName := fmt.Sprintf("testdata/%sOutput.golden", testName)

	dirPath, err := os.Getwd()
	require.NoError(t, err)
	fileInputName := fmt.Sprintf("%s/testdata/input.txt", dirPath)
	executable := fmt.Sprintf("%s/%s", dirPath, binaryName)

	command := BuildCommand(executable, flags, []string{fileInputName})

	output := RunCommandWithAssertion(t, command)

	UpdateFile(t, fileOutputName, output, update)

	expected, err := utils.ReadFile(fileOutputName)
	require.NoError(t, err)

	assert.Equal(t, expected, output)
}

func RunCommandWithAgainstRealCommand(t *testing.T, binaryName string, flags []string) {
	dirPath, err := os.Getwd()
	require.NoError(t, err)
	executable := fmt.Sprintf("%s/%s", dirPath, binaryName)
	fileInputName := fmt.Sprintf("%s/testdata/input.txt", dirPath)

	uniqCommand := BuildCommand("uniq", flags, []string{fileInputName})

	packageCommand := BuildCommand(executable, flags, []string{fileInputName})

	outputUniq := RunCommandWithAssertion(t, uniqCommand)
	outputBinary := RunCommandWithAssertion(t, packageCommand)

	assert.Equal(t, outputUniq, outputBinary)
}

func GetModulePath() string {
	gp := os.Getenv("GOPATH")
	return path.Join(gp, "src/goUniq")
}

func GetBinary(binaryName string) (func() error, error) {
	modulePath := GetModulePath()
	mainPath := path.Join(modulePath, "cmd/main.go")

	buildCommand := BuildCommand("go", []string{"build", "-a", "-o", binaryName, mainPath}, []string{})
	_, err := RunCommand(buildCommand)
	if err != nil {
		return nil, err
	}

	return func() error {
		removeCmd := BuildCommand("rm", []string{binaryName}, []string{})
		_, err := RunCommand(removeCmd)
		return err
	}, nil

}
