package safe

import (
	"sync"
	"testing"
)

func TestInt64(t *testing.T) {
	i := NewInt64()
	wg := &sync.WaitGroup{}
	for j := 0; j < 1000; j++ {
		wg.Add(1)
		go func() {
			i.Inc(1)
			wg.Done()
		}()
	}
	wg.Wait()

	if i.Get() != 1000 {
		t.Errorf("Expected %v, got %v", 1000, i.Get())
	}
}
