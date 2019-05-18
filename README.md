# levenshtein-distance
LevenshteinDistance returns a metric used to determine the difference between two strings using the algorithm from the wikipedia: https://en.wikipedia.org/wiki/Levenshtein_distance

"kitten" and "sitting" has an edit distance of 3.
"Sunday" and "Saturday" has an edit distance of 3.

# Quickstart
```bash
$ go get github.com/hueyjj/levenshtein-distance
```

# Example usage
```golang
package main

import (
    "fmt"
    "github.com/hueyjj/levenshtein_distance"
)

func main() {
    distance := levenshtein.Distance("kitten", "sitting")
    fmt.Printf("levenshtein.Distance(kitten, sitting)=%d\n", distance)
}
```

# Test
```bash
$ go test
PASS
ok      github.com/hueyjj/levenshtein-distance  0.284s
```