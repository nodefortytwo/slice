# slice

## Example

You can convert an existing string slice:

```go
import "github.com/nodefortytwo/slice"

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}  
    if slice.String(strs).Includes("apple") {
        return fmt.Println("Found an apple")
    }
}
```

or create a `slice.String` directly:

```go
import "github.com/nodefortytwo/slice"

func main() {
    var strs = slice.String{"peach", "apple", "pear", "plum"}  

    if strs.Includes("apple") {
        return fmt.Println("Found an apple")
    }
}
```

if you need a native slice back `Slice()` method