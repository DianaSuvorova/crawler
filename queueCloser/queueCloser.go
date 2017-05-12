package queueCloser

import (
  "fmt"
)

type QueueCloser struct {
  Quit chan bool
  Pause chan bool
  Resume chan bool
  todoQueue chan int
  todo int
  paused bool
}

func NewQueueCloser() *QueueCloser {
  q := new(QueueCloser)
  q.todoQueue = make(chan int)
  q.Quit = make(chan bool, 1)
  q.todo = 0

  go q.watch()
  return q
}

func (q *QueueCloser) Increment() {
  q.todoQueue <- 1
}

func (q *QueueCloser) Decrement() {
  q.todoQueue <- -1
}

func (q *QueueCloser) Todo() int {
  return q.todo;
}

func (q *QueueCloser) watch() {
  for i := range q.todoQueue {
    q.todo += i
    fmt.Println(q.todo)
    if (q.todo == 0) {
      q.Quit <- true
    } else if (q.todo > 100) {
      q.Pause <- true
      q.paused = true
    } else {
      if (q.paused) {
        q.Resume <- true
        q.paused = false
      }
    }
  }
}
