# Go: Silent Map Key Access

This example demonstrates the potential issue in Go where accessing a non-existent key in a map returns the zero value for the data type instead of an error or exception. This silent behavior can lead to difficult-to-debug issues if not handled carefully.

## The Bug
The `main` function showcases how accessing a key that does not exist returns 0 without any error messages. This behavior can make it harder to detect errors during program execution.

## Solution
The solution is to always check the existence of a key using the `value, ok := map[key]` idiom before accessing the value. This approach ensures robust error handling.