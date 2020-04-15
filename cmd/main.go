package main

import (
	"flag"
	"fmt"
	"github.com/amitdavidson234/go-uniq/pkg/domain"
	"github.com/amitdavidson234/go-uniq/pkg/utils"
	"os"
)

func main() {
	shouldCountRepetitions := flag.Bool("c", false, "Precede each output line with the count of the number of times the line occurred in the input, followed by a single space.")
	outputRepeatedLines := flag.Bool("d", false, "Only output lines that are repeated in the input.")
	outputNonRepeatedLines := flag.Bool("u", false, "Only output lines that are not repeated in the input.")
	isCaseSensitive := flag.Bool("i", false, "Case insensitive comparison of lines.")
	ignoreFirstNumFields := flag.Int("f", 0, "Ignore the first num fields in each input line when doing comparisons.  A field is a string of non-blank characters separated from adjacent fields by blanks. Field numbers are one based, i.e., the first field is field one.")
	ignoreFirstChars := flag.Int("s", 0, " with the -f option, the first chars characters after the first num fields will be ignored.  Character numbers are one based, i.e., the first character is character one.")
	flag.Parse()

	if utils.Bool2int(*shouldCountRepetitions)+utils.Bool2int(*outputRepeatedLines)+utils.Bool2int(*outputNonRepeatedLines) > 1 {
		err := fmt.Errorf("provide either c, d or u flag. [c|d|u]")
		utils.ExitWithError(err)
	}
	inputFileName := flag.Arg(0)
	outputFileName := flag.Arg(1)
	input, err := utils.ReadFile(inputFileName)
	if err != nil {
		if os.IsNotExist(err) {
			err := fmt.Errorf("file %s doesn't exist", inputFileName)
			utils.ExitWithError(err)
		}
		utils.ExitWithError(err)
	}
	lines := domain.CreateLinesFromInput(input)

	// Manipulation on the text.
	cf := domain.CommandsFlags{}
	if *isCaseSensitive {
		cf.ToLower(lines)
	}

	if *ignoreFirstNumFields > 0 {
		cf.TrimLineByField(lines, *ignoreFirstNumFields)
	}

	if *ignoreFirstChars > 0 {
		cf.TrimLineByChar(lines, *ignoreFirstChars)
	}

	cf.CollapseLines(lines)

	if err != nil {
		utils.ExitWithError(err)
	}

	// Manipulation on the structure.
	if *outputRepeatedLines {
		lines = cf.GetRepeatedLines(lines)
	}

	if *outputNonRepeatedLines {
		lines = cf.GetNonRepeatedLines(lines)
	}

	linesToPrint := cf.GenerateLines(lines, *shouldCountRepetitions)
	if err != nil {
		utils.ExitWithError(err)
	}

	if outputFileName == "" {
		utils.WriteToStdout(linesToPrint)
	} else {
		f, err := utils.CreateFile(outputFileName)
		if err != nil {
			utils.ExitWithError(err)
		}
		err = utils.WriteFile(f, linesToPrint)
		if err != nil {
			utils.ExitWithError(err)
		}
		err = f.Close()
		if err != nil {
			utils.ExitWithError(err)
		}
	}
}
