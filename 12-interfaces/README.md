# Interfaces

1. Use many, small interfaces
    - Simgle method interfaces are some of the most powerful and flexible
    - io.Writer, io.Reader, interface{}
2. Don't export interfaces for types that will be consumed
3. Do export interfaces for types that will be used by package
4. Design functions and methods to receive interfaces whenever possible
5. Type conversion

```go
var wc WriterCloser = NewBufferedWriterCloser()
bwc := wc.(*BufferedWriterCloser)
```
6. The empty interface and type switches

```go
var i interface {} = 0
switch i.(type) {
    case int:
        fmt.Println("i is an integer")
    case string:
        fmt.Println("i is a string")
    default:
        fmt.Println("i don't know what is")
}
```
