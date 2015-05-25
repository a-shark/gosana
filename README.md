# Gosana
Asana API client in Go

Example:
--------

```go
// example/example.go

package main

import (
  "fmt"
  "github.com/vanhalt/gosana"
)

func main() {
  client := gosana.NewClient("6asdasdf.asdasdasd")

  tasks := client.Tasks("111222333")
  for _, task := range tasks.Data {
    fmt.Println(task.Name)
  }
}
```
