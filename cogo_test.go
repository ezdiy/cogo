package cogo

import (
	"sync"
	"testing"
)

func Test_cogo(t *testing.T) {
	if WantsPreempt() {
		t.Fatal("Wants to preempt for no reason (bad offset?)")
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for !WantsPreempt() {
		}
		wg.Done()
	}()
	var x *int
	for y := 0; y < 1000000; y++ {
		x = new(int)
	}
	_ = x
	wg.Wait()
}
