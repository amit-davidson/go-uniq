package fuzzTesting

import (
	"goUniq/pkg/utils"
	"goUniq/tests/testUtils"
	"path"
)

func FuzzRegularCommandFlags(data []byte) int {
	// initPart
	modulePath := testUtils.GetModulePath()
	binaryName := path.Join(modulePath, "cmdCommand")

	fileInputName := path.Join(modulePath, "tests", "fuzzTesting", "corpus", randSeq(10) + ".txt")
	writeToFuzzFile(fileInputName, data)
	defer utils.RemoveFile(fileInputName)

	// execution part
	var flags []string
	arguments := []string{fileInputName}
	command := testUtils.BuildCommand(binaryName, flags, arguments)
	_, err := testUtils.RunCommand(command)
	if err != nil {
		panic(err.Error())
	}
	return 0
}
