Example 1:Queue in golang
package main
import "fmt"

func enqueue(queue[] int, element int) []int {
  queue = append(queue, element); // Simply append to enqueue.
  fmt.Println("Enqueued:", element);
  return queue
}

func dequeue(queue[] int) ([]int) {
  element := queue[0]; // The first element is the one to be dequeued.
  fmt.Println("Dequeued:", element)
  return queue[1:]; // Slice off the element once it is dequeued.
}

func main() {
  var queue[] int; // Make a queue of ints.

  queue = enqueue(queue, 10);
  queue = enqueue(queue, 20);
  queue = enqueue(queue, 30);

  fmt.Println("Queue:", queue);

  queue = dequeue(queue);
  fmt.Println("Queue:", queue);

  queue = enqueue(queue, 40);
  fmt.Println("Queue:", queue);
}

Output:
$go run queueingo.go
Enqueued: 10
Enqueued: 20
Enqueued: 30
Queue: [10 20 30]
Dequeued: 10
Queue: [20 30]
Enqueued: 40
Queue: [20 30 40]

Example2:
package main
import "container/list"
import "fmt"

func main() {
    // new linked list
    queue := list.New()

    // Simply append to enqueue.
    queue.PushBack(10)
    queue.PushBack(20)
    queue.PushBack(30)

    // Dequeue
    front:=queue.Front()
    fmt.Println(front.Value)
    // This will remove the allocated memory and avoid memory leaks
    queue.Remove(front)
}

Output:
$go run queueingo.go
10