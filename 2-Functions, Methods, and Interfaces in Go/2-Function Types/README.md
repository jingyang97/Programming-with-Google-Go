# Function Types

## First-Class Values

### Variable as Functions

```go
var funcVar func(int) int
func incFn(x int) int {
   return x + 1
}
func main() {
   funcVar = incFn
   fmt.Print(funcVar(1))
}

```

### Functions as Argument

```go
func applyIt(aFunc func (int) int, val int) int {
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
