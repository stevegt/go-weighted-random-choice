# Golang Weighted Random Choice
This is a simple package for creating Weighted Random Choices.

Distantly derived from "github.com/kontoulis/go-weighted-random-choice":
- switched weights to float64
- simplified the code
- added test cases

### Usage

```go
import Wrc "github.com/stevegt/go-weighted-random-choice"

func main() {
    wrc := Wrc.New()
    wrc.AddElement("common", 59)
    wrc.AddElement("epic", 1)
    wrc.AddElement("rare", 10)
    wrc.AddElement("green", 30)
    choice := wrc.GetRandomChoice()
    fmt.Println(choice)
}


``` 

You can also add multiple elements by passing a `map[string]int` in `AddElements` 

```go
wrc.AddElements(map[string]int{
    "common"    : 59,
    "epic"      : 1,
    "rare"      : 10,
    "green"     : 30,
})
```
