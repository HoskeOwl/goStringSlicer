# goStringSlicer
Fast and memory-frendly string slicer

## Description
- indexing by letters
- O(n) 
- Do not use additional memory (just for result)
- Support negative indexes

## examples
```
import "goStringSlicer"

func main(){
    var sentense string = "Some Únicode"
    res, _ := SliceString(sentense, 0, 3) // res = "Som"
}

```
```
import "goStringSlicer"

func main(){
    var sentense string = "Some Únicode"
    res, _ := SliceString(sentense, -7, 8) // res = "Úni"
}

```
Also function can return some errors:
- one of index or both out of bound
- end index bigger than begin

```
import "goStringSlicer"

func main(){
    var sentense string = "Some Únicode"
    res, err := SliceString(sentense, -7, 8) // res = "Úni"
    if err == nil{
        // do something
    }
}

```