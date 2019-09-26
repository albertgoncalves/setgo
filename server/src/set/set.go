package set

import (
    "fmt"
    "math/rand"
    "time"
)

const (
    K = 3
    L = 2 // K - 1
)

func combinations(n int) [][K]int {
    result := make([][K]int, 0)
    if 0 < n {
        var state [K]int
        var closure func(int, int)
        closure = func(j, k int) {
            for i := k; i < n; i++ {
                state[j] = i
                if j == L {
                    var combination [K]int = state
                    result = append(result, combination)
                } else {
                    closure(j+1, i+1)
                }
            }
            return
        }
        closure(0, 0)
    }
    return result
}

type Token struct {
    Id        string `json:"id"`
    Shape     string `json:"shape"`
    Color     string `json:"color"`
    Fill      string `json:"fill"`
    Frequency uint8  `json:"frequency"`
}

var SHAPES = [K]string{
    "square",
    "circle",
    "triangle",
}

var FILLS = [K]string{
    "solid",
    "transparent",
    "empty",
}

var COLORS = [K]string{
    "green",
    "red",
    "blue",
}

var IDS = [9]string{
    "0,0",
    "0,1",
    "0,2",
    "1,0",
    "1,1",
    "1,2",
    "2,0",
    "2,1",
    "2,2",
}

func AllTokens() []*Token {
    tokens := make([]*Token, 0, 81)
    for _, shape := range SHAPES {
        for _, fill := range FILLS {
            for _, color := range COLORS {
                for frequency := 1; frequency < 4; frequency++ {
                    tokens = append(tokens, &Token{
                        Shape:     shape,
                        Fill:      fill,
                        Color:     color,
                        Frequency: uint8(frequency),
                    })
                }
            }
        }
    }
    return tokens
}

var ALL_TOKENS []*Token = AllTokens()

func ShuffleTokens() {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(ALL_TOKENS), func(i, j int) {
        ALL_TOKENS[i], ALL_TOKENS[j] = ALL_TOKENS[j], ALL_TOKENS[i]
    })
}

func Pop() (*Token, error) {
    if 0 < len(ALL_TOKENS) {
        var token *Token
        token, ALL_TOKENS = ALL_TOKENS[0], ALL_TOKENS[1:]
        return token, nil
    }
    return nil, fmt.Errorf("Pop()")
}

func Init() ([]*Token, error) {
    if 9 <= len(ALL_TOKENS) {
        var tokens []*Token
        tokens, ALL_TOKENS = ALL_TOKENS[:9], ALL_TOKENS[9:]
        for i, id := range IDS {
            tokens[i].Id = id
        }
        return tokens, nil
    }
    return nil, fmt.Errorf("Init()")
}

func compareString(a, b, c string) bool {
    return ((a == b) && (b == c)) || ((a != b) && (b != c) && (c != a))
}

func compareUint8(a, b, c uint8) bool {
    return ((a == b) && (b == c)) || ((a != b) && (b != c) && (c != a))
}

func Validate(tokens []*Token) bool {
    if len(tokens) != 3 {
        return false
    }
    first := tokens[0]
    second := tokens[1]
    third := tokens[2]
    if (first == nil) || (second == nil) || (third == nil) {
        return false
    }
    return compareString(first.Shape, second.Shape, third.Shape) &&
        compareString(first.Fill, second.Fill, third.Fill) &&
        compareString(first.Color, second.Color, third.Color) &&
        compareUint8(first.Frequency, second.Frequency, third.Frequency)
}

func AnySolution(tokens []*Token) bool {
    for _, indices := range combinations(len(tokens)) {
        if Validate([]*Token{
            tokens[indices[0]],
            tokens[indices[1]],
            tokens[indices[2]],
        }) {
            fmt.Println(
                tokens[indices[0]],
                tokens[indices[1]],
                tokens[indices[2]],
            )
            return true
        }
    }
    return false
}

func Remove(tokens []*Token, i int) []*Token {
    newTokens := make([]*Token, 0, len(tokens)-1)
    newTokens = append(newTokens, tokens[:i]...)
    return append(newTokens, tokens[i+1:]...)
}

func Lookup(index string) (int, error) {
    switch index {
    case "0,0":
        return 0, nil
    case "0,1":
        return 1, nil
    case "0,2":
        return 2, nil
    case "1,0":
        return 3, nil
    case "1,1":
        return 4, nil
    case "1,2":
        return 5, nil
    case "2,0":
        return 6, nil
    case "2,1":
        return 7, nil
    case "2,2":
        return 8, nil
    }
    return 0, fmt.Errorf("Lookup(%s)", index)
}
