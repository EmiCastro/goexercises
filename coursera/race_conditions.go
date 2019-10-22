package main

import (
	"fmt"
	"sync"
)

/* This method generates a rate condition because the two goroutines
* are accessing to the same variable concurrently so the output is not deterministic
* because sometimes one goroutine will print the variable before the other goroutine
* modifies it and sometimes will occur the opposite
 */
func main() {
	var wg sync.WaitGroup
	x := 2

	wg.Add(2)
	go func() {
		defer wg.Done()
		x += 5
	}()

	go func() {
		defer wg.Done()
		fmt.Println(x)
	}()

	wg.Wait()
}
