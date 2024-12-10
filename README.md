Some helpers to get user input from terminal. Thanks to [keyboard](https://github.com/atomicgo/keyboard) from atomic.go

### Example: select from multiple choice

```go
package main

import (
    "fmt"

    "github.com/Noblefel/vivi"
)

func main() {
    for {
        fmt.Println("==============")
        fmt.Println("What is 5 * 5 ?")

        answer := vivi.Choices(
            "[1] 25",
            "[2] 10",
            "[3] 50",
        )

        if answer == 0 {
            fmt.Println("✔️  Correct")
            break
        }

        fmt.Println("❌ Try Again")
    }
}

```
