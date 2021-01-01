# Functions and Organization

## Call by Value, Reference

- Trade off: Data Encapsulation VS. Data Copying

Call by Reference

```go
// passing an unmutable variable
func foo(y *int) {
    *y = *y + 1
}

func main() {
    x := 2
    foo(&x)
    fmt.Print(x) // x = 3
}

// passing array (reference)
func foo(x *[3]int) int {
    (*x)[0] = (*x)[0] + 1
}

func main() {
    a := [3]int{1, 2, 3}
    foo(&a)
    fmt.Print(a)
}

// passing slice (slice contains a pointer)
func foo(sli int) int {
    sli[0] = sli[0] + 1
}

func main() {
    a := []int{1, 2, 3}
    foo(a)
    fmt.Print(a)
}
```
