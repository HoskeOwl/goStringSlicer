package goStringSlicer

import (
	"fmt"
	"unicode/utf8"
)

// slice string by runes. begin - start rune index for slice. end - end rune index to slice.
// end index will not be include.
func SliceStringNegative(data string, begin, end int) (string, error) {
	if begin == end {
		return "", nil
	}
	// end - symbol to (not included)
	runeLen := utf8.RuneCountInString(data)
	var byteBegin, byteEnd int = -1, -1

	// transform to plain sequence
	// transform to plain sequence
	if begin < 0 {
		begin = runeLen + begin
	}
	if end < 0 {
		end = runeLen + end
	}
	if begin == end {
		return "", nil
	}

	// process error cases
	if begin < 0 || begin > runeLen {
		return "", fmt.Errorf("Index %v out of range: rune length = %v", begin, runeLen)
	}
	if end < 0 || end > runeLen {
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

	if byteEnd < 0 {
		byteEnd = len(data)
	}

	if byteBegin > byteEnd {
		return "", fmt.Errorf("Internal error: begin byte index bugger than end byte index")
	}
	return data[byteBegin:byteEnd], nil
}

// slice string by runes. begin - start rune index for slice. end - end rune index to slice.
// end index will not be include.
func SliceStringPositive(data string, begin, end int) (string, error) {
	if begin > end {
		return "", fmt.Errorf("Start of slice bigger than end")
	}
	if begin == end {
		return "", nil
	}
	var byteBegin, byteEnd int = -1, -1

	// calc byte indexes
	runeLen := 0
	for byteIdx := range data {
		if runeLen == begin {
			byteBegin = byteIdx
		}
		if runeLen == end {
			byteEnd = byteIdx
		}
		if runeLen > end {
			break
		}
		runeLen += 1
	}

	// process error cases
	if begin > runeLen {
		return "", fmt.Errorf("Index %v out of range: rune length = %v", begin, runeLen)
	}
	if end > runeLen {
		return "", fmt.Errorf("Index %v out of range: rune length = %v", end, runeLen)
	}

	if byteEnd < 0 {
		byteEnd = len(data)
	}

	return data[byteBegin:byteEnd], nil
}

func SliceString(data string, begin, end int) (string, error) {
	if begin > 0 && end > 0 {
		return SliceStringPositive(data, begin, end)
	}
	return SliceStringNegative(data, begin, end)
}
