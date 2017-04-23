package queueCloser

type QueueCloser struct {
  Quit chan bool
  todoQueue chan int
  todo int
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
    if (q.todo == 0) {
      q.Quit <- true
    }
  }
}
