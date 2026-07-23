# Testing in Go

## Steps for Writing Test Suites in Golang

- Create a file whose name ends with `_test.go`
- Import package testing by import `"testing"` command
- Write the test function of form `func TestXxx(*testing.T)` which uses any of Error, Fail, or related methods to signal failure.
- Put the file in any package.
- Run command `go test`
- Create go module
- Exported method names should be starting with capital case
- We should write all the clean-up codes in `t.Cleanup` method

**Note**: test file will be excluded in package build and will only get executed on `go test` command.

## Example: Basic Unit Test

```go
// src/app.go
package src

func SayHello() string {
    return "hello"
}

func Min(a, b int) int {
    if a < b { return a }
    return b
}

// test/app_test.go
package test

func TestReturnGeeks(t *testing.T) {
    actualString := src.SayHello()
    expectedString := "hello"
    if actualString != expectedString {
        t.Errorf("Expected String(%s) is not same as"+
            " actual string (%s)", expectedString, actualString)
    }
    t.Cleanup(func() { fmt.Println("This is cleanup function") })
}

func TestMinBasic(t *testing.T) {
    ans := src.Min(2, -2)
    if ans != -2 {
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}
```

## Table-Driven Tests

Writing tests can be repetitive — table-driven style lists inputs and expected outputs, then loops over them. `t.Run` enables subtests shown separately with `go test -v`.

```go
type MinTestCase struct {
    a, b int
    want int
}

func minTestCases() []MinTestCase {
    return []MinTestCase{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }
}

func Test_MinTableDriven(t *testing.T) {
    for _, tt := range minTestCases() {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := src.Min(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
```

## Benchmark Tests

Named beginning with `Benchmark`. The testing runner increases `b.N` until a precise measurement is collected.

```go
func BenchmarkMin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        src.Min(1, 2)
    }
}
```

Run with: `go test -bench=.`

## Testing Struct Methods

```go
// card.go
type Cards []string

func NewCards() Cards {
    cards := Cards{}
    cardType := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
    cardValues := []string{"Jack", "Nine", "Ace", "Ten", "King", "Queen", "Eight", "Seven"}
    for _, suit := range cardType {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)
        }
    }
    return cards
}

func (c *Cards) Shuffle() {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)
    cards := *c
    for i := range cards {
        newPosition := r.Intn(len(cards) - 1)
        cards[i], cards[newPosition] = cards[newPosition], cards[i]
    }
}

func (c *Cards) GiveDeck() ([][]string, [][]string) {
    c.Shuffle()
    allCards := []string(*c)
    firstDeck := make([][]string, 4)
    secondDeck := make([][]string, 4)
    for i := 0; i < 4; i++ {
        firstDeck[i] = allCards[i*4 : ((i+1)*4)-1]
        secondDeck[i] = allCards[16+i*4 : 16+((i+1)*4)-1]
    }
    return firstDeck, secondDeck
}

func (cards Cards) SaveToFile(filename string) error {
    return os.WriteFile(filename, []byte(strings.Join(cards, ",")), 0666)
}

// card_test.go
func TestNewCards(t *testing.T) {
    cards := c.NewCards()
    if len(cards) != 32 {
        t.Errorf("Expected deck length of 32, but got %v", len(cards))
    }
}

func TestGiveDeck(t *testing.T) {
    cards := c.NewCards()
    firstDeck, secondDeck := cards.GiveDeck()
    if len(firstDeck) != 4 || len(secondDeck) != 4 {
        t.Errorf("Expected 4 hands each, got %v %v", len(firstDeck), len(secondDeck))
    }
}
```

## Key Differences

| Method | Behavior |
|---|---|
| `t.Error` / `t.Errorf` | Reports failure, continues test |
| `t.Fatal` / `t.Fatalf` | Reports failure, stops immediately |
| `t.Cleanup(fn)` | Registers cleanup function, runs after test completes |
| `t.Log` / `t.Logf` | Logs info (shown with `-v` flag) |
| `t.Run(name, fn)` | Runs a subtest (shown separately with `-v`) |
