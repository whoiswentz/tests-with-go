package parallel

import (
	"testing"
	"time"
)

func TestSomethingA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestSomethingB(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
}

func TestSomethinC(t *testing.T) {
}

func TestSomethinD(t *testing.T) {
	t.Run("parallel subtest 1", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second)
	})

	t.Run("parallel subest 2", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second)
	})
}
