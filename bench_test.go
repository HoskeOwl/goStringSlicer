package goStringSlicer

import "testing"

var exampleSmall string = "😸😼"                                          // len = 2
var exampleMedium string = "😸 Some very long string with nice unicode😼" // len = 42
// len = 413
var exampleLarge string = "😸 Чёрная дыра́ — область пространства-времени, гравитационное притяжение которой настолько велико, что покинуть её не могут даже объекты, движущиеся со скоростью света, в том числе кванты самого света. Граница этой области называется горизонтом событий. В простейшем случае сферически симметричной чёрной дыры он представляет собой сферу с радиусом Шварцшильда, который считается характерным размером чёрной дыры."

const (
	smallBegin     int = 0
	smallEnd       int = 1
	smallNegBegin  int = -2
	smallNegEnd    int = -1
	mediumBegin    int = 4
	mediumEnd      int = 18
	mediumNegBegin int = -38
	mediumNegEnd   int = -20
	largeBegin     int = 10
	largeEnd       int = 120
	largeNegBegin  int = -400
	largeNegEnd    int = -380
	largeEqual     int = 410
)

func BenchmarkSmallSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleSmall)
		part := runes[smallBegin:smallEnd]
		_ = string(part)
	}
}

func BenchmarkPositiveSmallSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleSmall, smallBegin, smallEnd)
	}
}
func BenchmarkNegativeSmallSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleSmall, smallNegBegin, smallNegEnd)
	}
}

func BenchmarkMediumSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleMedium)
		part := runes[mediumBegin:mediumEnd]
		_ = string(part)
	}
}

func BenchmarkPositiveMediumSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleMedium, mediumBegin, mediumEnd)
	}
}

func BenchmarkNegativeMediumSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleMedium, mediumNegBegin, mediumNegEnd)
	}
}
func BenchmarkLargeSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleLarge)
		part := runes[largeBegin:largeEnd]
		_ = string(part)
	}
}

func BenchmarkPositiveLargeSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleLarge, largeBegin, largeEnd)
	}
}

func BenchmarkNegativeLargeSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleLarge, largeNegBegin, largeNegEnd)
	}
}
func BenchmarkPositiveLargeSliceWithFunctionEqualIndexes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleLarge, largeEqual, largeEqual)
	}
}
