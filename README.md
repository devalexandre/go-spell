A simple Go package for calculating the Levenshtein distance between two strings and finding words within a certain distance from a target word.

# Installation
To install the Levenshtein package, you can use the go get command:

```
go get github.com/devalexandre/go-spell
```

# Usage

Here's an example of how to use the Levenshtein package to find words within a certain distance from a target word:

```go
package main

import (
	"fmt"
	"github.com/devalexandre/go-spell/pkg"
)

func main() {
	// Example dictionary of words
	dictionary := []string{"apple", "banana", "cherry", "grape", "lemon", "lime", "orange", "strawberry"}

	// Target word we want to check
	targetWord := "apxle"

	// Define a maximum allowed distance
	maxDistance := 2

	// Find words within the specified distance from the target word
	words := levenshtein.Find(dictionary, maxDistance, targetWord)

	fmt.Println("Words:", words)
}
```

In this example, we import the levenshtein package and use it to find words in the dictionary that are within a maximum distance of 2 from the target word "apxle."

# Functions

**levenshteinDistance(s1, s2 string) int**

Calculates the Levenshtein distance between two strings s1 and s2 and returns an integer representing the distance.
**min(a, b int) int**

Helper function to calculate the minimum of two integers.
**Find(dictionary []string, maxDistance int, targetWord string) []string**

Finds words in the dictionary that are within a maximum maxDistance from the targetWord and returns them as a slice of strings.

# Contributing
Feel free to contribute to this project by opening issues or submitting pull requests 