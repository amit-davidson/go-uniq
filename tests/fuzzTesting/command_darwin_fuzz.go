// +build darwin

package fuzzTesting

import (
	"bytes"
	"github.com/amitdavidson234/go-uniq/pkg/utils"
	"github.com/amitdavidson234/go-uniq/tests/testUtils"
	"log"
	"path"
)

func FuzzRegularCommandFlagsRunAgainstUniq(data []byte) int {
	modulePath := testUtils.GetProjectRootPath()
	binaryName := path.Join(modulePath, "cmdCommand")

	fileInputName := path.Join(modulePath, "tests", "fuzzTesting", "corpus", randSeq(10)+".txt")
	writeToFuzzFile(fileInputName, data)
	defer utils.RemoveFile(fileInputName)

	// execution part
	var flags []string
	arguments := []string{fileInputName}
	command := testUtils.BuildCommand(binaryName, flags, arguments)
	resCommand, err := testUtils.RunCommand(command)
	if err != nil {
		if resCommand != nil {
			log.Print(err.Error())
			return -1
		}
		log.Print(err.Error())
		return -1
	}

	uniqCommand := testUtils.BuildCommand("uniq", flags, []string{fileInputName})
	resUniq, err := testUtils.RunCommand(uniqCommand)
	if err != nil {
		if resUniq != nil {
			log.Print(err.Error())
			panic("uniq return error and result")
		}
		panic("Nice job, you failed real uniq")
	}
	if !bytes.Equal(resCommand, resUniq) {
		panic("uniq and command don't match")
	}

	return 0
}
