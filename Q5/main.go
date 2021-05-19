//What will be printed when the code below is executed?
//And fix the issue to assure that `len(m)` is printed as 10.

package main

import (
	"sync"
)

const N = 10

var (
	m  map[int]int
	mu sync.Mutex
)

func Set(wg *sync.WaitGroup, i int) {
	mu.Lock()
	m[i] = i
	mu.Unlock()
	wg.Done()
}
func main() {
	m = make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go Set(wg, i)
	}
	wg.Wait()
	println(len(m))
}
