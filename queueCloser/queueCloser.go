package queueCloser

type QueueCloser struct {
  Quit chan bool
  todoQueue chan int
  todo int
}

func NewQueueCloser() *QueueCloser {
  q := new(QueueCloser)
  q.todoQueue = make(chan int, 5000)
  q.Quit = make(chan bool)
  q.todo = 0

  return q
}

func (q *QueueCloser) Increment() {
  q.todoQueue <- 1
}

func (q *QueueCloser) Decrement() {
  <- q.todoQueue
  // if (len(q.todoQueue) == 0) {
  //   q.Quit <- true
  // }
}

func (q *QueueCloser) Todo() int {
  return len(q.todoQueue);
}
