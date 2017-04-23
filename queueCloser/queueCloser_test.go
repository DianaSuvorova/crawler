package queueCloser_test

import (
  "testing"
  "crawler/queueCloser"
)

func TestIncrement(t *testing.T) {
  closer := queueCloser.NewQueueCloser()
  closer.Increment();
  if (closer.Todo() != 1) {
      t.Error("expected  1")
  }
}

func TestDecrement(t *testing.T) {
  closer := queueCloser.NewQueueCloser()
  closer.Decrement();
  if (closer.Todo() != -1) {
      t.Error("expected  -1")
  }
}

func TestWatch(t *testing.T) {
  closer := queueCloser.NewQueueCloser()
  go func() {
    <- closer.Quit
  }()
  closer.Increment()
  closer.Decrement()

}
