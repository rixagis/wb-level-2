package or

import (
	"testing"
	"time"
)

func TestOr(t *testing.T){

	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
	}()
	return c
	}
	
	start := time.Now()
	<-Or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(15*time.Second),
		sig(10*time.Second),
	)
	
	passed := time.Since(start)
	t.Logf("done after %v\n", passed)
	if passed < 1*time.Second - 100*time.Millisecond || passed > 1*time.Second + 100*time.Millisecond {
		t.Errorf("unexpected done time, expected: ~%v, got: %v", 1*time.Second, passed)
	}
	
}