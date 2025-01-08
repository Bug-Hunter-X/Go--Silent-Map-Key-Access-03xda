```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3

    for k, v := range m {
        fmt.Println(k, v)
    }

    //Correct way to access a map key
    value, ok := m["d"]
    if ok {
        fmt.Println("Key exists. Value:", value)
    } else {
        fmt.Println("Key does not exist")
    }
}
```