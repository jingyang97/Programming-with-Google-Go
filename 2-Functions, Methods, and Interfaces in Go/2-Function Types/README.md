# Function Types

## First-Class Values

### Variable as Functions

```go
var funcVar func(int) int
```

### Functions as Argument

```go
func applyIt(aFunc fuct (int) int, val int) int {
    return aFunc(val)
}
```

### Anonymous Functions

```go
func main() {
    v := applyIt(func (x int) int {return x + 1}, 2)
    fmt.Println(v)
}
```

## Returning Functions

## Variadic and Deferred
