package domain

import (
	"strings"
)

type Line struct {
	Next *Line

	OriginalText   string
	ComparisonText string
	Count          int
}

func CompareLines(line1 *Line, line2 *Line) bool {
	return line1.OriginalText == line2.OriginalText && line1.ComparisonText == line2.ComparisonText && line1.Count == line2.Count
}

func CreateLines(lines []string) *Line {
	allocatedLines := make([]Line, len(lines)) // Allocate all lines at one attempt to save calls

	firstElem := &allocatedLines[0]
	firstElem.OriginalText = lines[0]
	firstElem.ComparisonText = lines[0]

	if len(lines) == 1 {
		return firstElem
	}
	elem := firstElem
	i := 1
	for range lines[i:] {
		elem.Next = &allocatedLines[i]
		elem.Next.OriginalText = lines[i]
		elem.Next.ComparisonText = lines[i]

		i++
		elem = elem.GetNext()
	}
	return firstElem
}

func AreLinesEqual(list1 *Line, list2 *Line) bool {
	elem1, elem2 := list1, list2
	for true {
		if elem1 == nil && elem2 != nil || elem1 != nil && elem2 == nil {
			return false
		} else if elem1 == nil && elem2 == nil {
			return true
		} else if !CompareLines(elem1, elem2) {
			return false
		} else {
			elem1, elem2 = elem1.GetNext(), elem2.GetNext()
		}
	}
	return true
}

func (node *Line) RemoveNext() {
	node.Next = node.Next.Next
}

func (node *Line) GetNext() *Line {
	return node.Next
}

func CreateLinesFromInput(content []byte) *Line {
	splitLines := strings.Split(string(content), "\n")
	if len(splitLines) == 1 {
		return nil
	}

	// Line is defined as a sequence of characters ending with \n. That's why we don't take the last item , as it's not
	//ending with /n. See: https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_206
	return CreateLines(splitLines[:len(splitLines)-1])
}
