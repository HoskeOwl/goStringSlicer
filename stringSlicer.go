package goStringSlicer

import (
	"fmt"
	"unicode/utf8"
)

// slice string by runes. begin - start rune index for slice. end - end rune index to slice.
// end index will not be include.
func SliceString(data string, begin, end int) (string, error) {
	// end - symbol to (not included)
	runeLen := utf8.RuneCountInString(data)
	var byteBegin, byteEnd int = -1, -1

	// transform to plain sequence
	if begin < 0 {
		begin = runeLen + begin
	}
	if end < 0 {
		end = runeLen + end
	}

	// process error cases
	if begin > runeLen || begin < 0 {
		return "", fmt.Errorf("Index %v out of range: rune length = %v", begin, runeLen)
	}
	if end > runeLen || end < 0 {
		return "", fmt.Errorf("Index %v out of range: rune length = %v", end, runeLen)
	}
	if begin > end {
		return "", fmt.Errorf("Start of slice bigger than end")
	}

	// calc byte indexes
	i := 0
	for byteIdx := range data {
		if i == begin {
			byteBegin = byteIdx
		}
		if i == end {
			byteEnd = byteIdx
		}
		if i > end {
			break
		}
		i += 1
	}

	// check if some index did not set
	if byteBegin < 0 {
		return "", nil
	}
	if byteEnd < 0 {
		byteEnd = len(data)
	}

	if byteBegin > byteEnd {
		return "", fmt.Errorf("Internal error: begin byte index bugger than end byte index")
	}
	return data[byteBegin:byteEnd], nil
}
