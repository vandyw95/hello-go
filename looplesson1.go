package main

import "fmt"

func tryLoop1() {
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	// Hashtable / map
	for key, char := range map[int]string{1: "A", 2: "B", 3: "C"} {
		fmt.Printf("%d => Char: %s\n", key, char)
	}

	// If you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}
}
