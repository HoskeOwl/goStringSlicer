package goStringSlicer

import (
	"testing"
)

var exampleString string = "abcdefg"

func TestSlicingOk(t *testing.T) {
	var testVals map[[2]int]string = map[[2]int]string{
		{0, 0}: "",
		{0, 1}: "a",
		{6, 7}: "g",
		{7, 7}: "",
		{0, 7}: "abcdefg",
		{1, 3}: "bc",

		{-1, -1}: "",
		{-2, -1}: "f",
		{-6, -3}: "bcd",
		{-7, -6}: "a",
		{-7, -7}: "",
	}

	var res string
	for k, v := range testVals {
		res, _ = SliceString(exampleString, k[0], k[1])
		if res != v {
			t.Errorf("Wrong ansver. Input: %v, %v. Out: %v", k[0], k[1], res)
		}
	}
}

func TestSlicingErr(t *testing.T) {
	var testVals [][2]int = [][2]int{
		{5, 0},
		{2, 16},
		{12, 16},

		{-12, 2},
		{2, -12},
		{-3, -5},
	}

	var res string
	var err error
	for _, v := range testVals {
		res, err = SliceString(exampleString, v[0], v[1])
		if err == nil {
			t.Errorf("Error did not return. Input: %v, %v. Out: %v", v[0], v[1], res)
		}
		if res != "" {
			t.Errorf("Wrong ansver. Input: %v, %v. Out: %v. Err: %v", v[0], v[1], res, err)
		}
	}
}
