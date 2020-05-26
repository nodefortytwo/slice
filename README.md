# slice

## Gettring Started
```go
import "github.com/nodefortytwo/slice"

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}  
    if slice.String(strs).Includes("apple") {
        return fmt.Println("Found an apple")
    }
}
```