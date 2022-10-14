package goStringSlicer

import "testing"

var example string = "😸 Some very long string with nice unicode😼"

func BenchmarkSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(example)
		part := runes[4:18]
		_ = string(part)
	}
}

func BenchmarkSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(example, 4, 18)
	}
}
