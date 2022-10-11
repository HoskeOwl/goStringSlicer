# goStringSlicer
Fast and memory-frendly string slicer

## peculiarities
- O(n) algorithm for slicing string by rune indexes.
- Do not use additional memory (just for returning slice).
- Support negative indexes

## examples
```
import "goStringSlicer"

func main(){
    var sentense string = "Some Únicode"
    res := SliceString(sentense, 0, 3) // res = "Som"
}

```
```
import "goStringSlicer"

func main(){
    var sentense string = "Some Únicode"
    res := SliceString(sentense, -7, 8) // res = "Úni"
}

```