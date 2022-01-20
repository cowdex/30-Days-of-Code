package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read input from STDIN. 
	fmt.Println("Hello, World.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputString := scanner.Text()
		// Print output to STDOUT
		fmt.Println(inputString)
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
