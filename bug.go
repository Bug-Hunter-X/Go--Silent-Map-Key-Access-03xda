```go
package main

import "fmt"

func main() {

    // Create a map
    m := make(map[string]int)

    // Add some values
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3

    // Iterate over the map and print the values
    for k, v := range m {
        fmt.Println(k, v)
    }

    //Try to access a non-existent key
    fmt.Println(m["d"]) //this will print 0, not an error

    //Check if a key exists before accessing it
    value, ok := m["d"]
    if ok{
        fmt.Println("Key exists. Value:", value)
    }else{
        fmt.Println("Key does not exist")
    }
}
```