package domain

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/amitdavidson234/go-uniq/pkg/utils"
)

type CommandsFlags struct{}

func (CommandsFlags) ToLower(lines *Line) {
	var elem *Line
	for elem = lines; elem != nil; elem = elem.GetNext() {
		elem.ComparisonText = strings.ToLower(elem.ComparisonText)
	}
}

func (CommandsFlags) GetRepeatedLines(lines *Line) (modifiedList *Line){
	if lines == nil { // No lines
		return
	}
	elem := lines
	for true { // We iterate until there's a first line with count >= 2. If we found we break, if we didn't and got to nil, we exit
		if elem.Count >= 2 {
			modifiedList = elem
			break
		}
		elem = elem.GetNext()
		if elem == nil {
			return
		}
	}

	for true {
		if elem.GetNext() == nil {
			return modifiedList
		} else if elem.GetNext().Count < 2 {
			elem.RemoveNext()
		} else {
			elem = elem.GetNext()
		}
	}
	return modifiedList
}

func (CommandsFlags) GetNonRepeatedLines(lines *Line) (modifiedList *Line) {
	if lines == nil { // No lines
		return
	}
	elem := lines
	for true { // We iterate until there's a first line with count >= 2. If we found we break, if we didn't and got to nil, we exit
		if elem.Count < 2 {
			modifiedList = elem
			break
		}
		elem = elem.GetNext()
		if elem == nil {
			return
		}
	}

	for true {
		if elem.GetNext() == nil {
			return modifiedList
		} else if elem.GetNext().Count >= 2 {
			elem.RemoveNext()
		} else {
			elem = elem.GetNext()
		}
	}
	return modifiedList
}

func (CommandsFlags) TrimLineByField(lines *Line, ignoreFirstNumFields int) {
	var elem *Line
	for elem = lines; elem != nil; elem = elem.GetNext() {
		_, formattedLine := utils.SplitStringByCharIndex(elem.ComparisonText, 0x0020, ignoreFirstNumFields)
		elem.ComparisonText = formattedLine
	}
}

func (CommandsFlags) TrimLineByChar(lines *Line, ignoreFirstChars int) {
	var elem *Line
	for elem = lines; elem != nil; elem = elem.GetNext() {
		if len(elem.ComparisonText) > ignoreFirstChars {
			elem.ComparisonText = elem.ComparisonText[ignoreFirstChars:]
		} else {
			elem.ComparisonText = ""
		}
	}
}

func (cf CommandsFlags) CollapseLines(lines *Line) {
	if lines == nil {
		return
	}

	elem := lines
	for true {
		nextElement := elem.GetNext()
		elem.Count++
		if nextElement == nil {
			return
		}
		if elem.ComparisonText != nextElement.ComparisonText {
			elem = elem.GetNext()
		} else {
			elem.RemoveNext()
		}
	}
}

func (CommandsFlags) GenerateLines(lines *Line, shouldCountRepetitions bool) []byte {
	var buf bytes.Buffer

	if lines == nil {
		return []byte{}
	}

	var elem *Line
	for elem = lines; elem != nil; elem = elem.GetNext() {
		if shouldCountRepetitions {
			buf.Write([]byte(fmt.Sprintf("%d %s\n", elem.Count, elem.OriginalText)))
		} else {
			buf.Write([]byte(fmt.Sprintf("%s\n", elem.OriginalText)))
		}
	}
	return buf.Bytes()
}
