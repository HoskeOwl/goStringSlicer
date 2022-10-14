package goStringSlicer

import "testing"

var exampleSmall string = "😸😼"
var exampleMedium string = "😸 Some very long string with nice unicode😼"
var exampleLarge string = "😸 Чёрная дыра́ — область пространства-времени, гравитационное притяжение которой настолько велико, что покинуть её не могут даже объекты, движущиеся со скоростью света, в том числе кванты самого света. Граница этой области называется горизонтом событий. В простейшем случае сферически симметричной чёрной дыры он представляет собой сферу с радиусом Шварцшильда, который считается характерным размером чёрной дыры."

const (
	smallBegin  int = 0
	smallEnd    int = 1
	mediumBegin int = 4
	mediumEnd   int = 18
	largeBegin  int = 10
	largeEnd    int = 120
)

func BenchmarkSmallSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleSmall)
		part := runes[smallBegin:smallEnd]
		_ = string(part)
	}
}

func BenchmarkSmallSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleSmall, smallBegin, smallEnd)
	}
}

func BenchmarkMediumSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleMedium)
		part := runes[mediumBegin:mediumEnd]
		_ = string(part)
	}
}

func BenchmarkMediumSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleMedium, mediumBegin, mediumEnd)
	}
}

func BenchmarkLargeSliceThroughRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runes := []rune(exampleLarge)
		part := runes[largeBegin:largeEnd]
		_ = string(part)
	}
}

func BenchmarkLargeSliceWithFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SliceString(exampleLarge, largeBegin, largeEnd)
	}
}
