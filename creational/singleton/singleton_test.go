package singleton

import (
	"sync"
	"testing"
)

const count = 50

func TestPattern(t *testing.T) {
	s := GetInstance()
	s.getAddr()
}

func TestParallel(t *testing.T) {
	start := make(chan any)
	wg := sync.WaitGroup{}
	wg.Add(count)
	var instances [count]Singleton
	for i := 0; i < count; i++ {
		go func(i int) {
			<-start
			instance := GetInstance()
			instances[i] = instance
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < count; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
