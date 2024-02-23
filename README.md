# b64

Simple Base64 Go Package

Example:

```go
  package main

  import (
    "fmt"
    "github.com/vxcute/base64/b64"
  )

  func main() {
    original := "AAAAAAAAhhhhhhhhhh912738712dskajdkxnchWWWHKHkashkdhakhkhsk"
    encoded := b64.Base64Encode(original)
    fmt.Println(encoded)
    decoded := b64.Base64Decode(encoded)
    fmt.Println(decoded)
    fmt.Println(original == decoded)
  }
```
