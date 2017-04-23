package main

type queueCloser struct {
  queue chan processor
  todoQueue chan int
}

func newQueueCloser(queue chan processor) *queueCloser {
  q := new(queueCloser)
  q.todoQueue = make(chan int)
  q.queue = queue

  go q.watch()
  return q
}

func (q *queueCloser) increment() {
  q.todoQueue <- 1
}

func (q *queueCloser) decrement() {
  q.todoQueue <- -1
}

func (q *queueCloser) watch() {
  todo := 0
  for i:= range q.todoQueue {
    todo += i
    if (todo == 0) {
      close(q.queue)
    }
  }
}
