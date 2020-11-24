Example 1:stack using linked list.
package main
import "fmt"

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	stack.Push("this")
	stack.Push("is")
	stack.Push("pavan")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}

Output:
pavan
is
this

Example 2:
package main
import "fmt"
func main () {
var array = [6] int{10,20,30,40,50,60}
var slice = array[1:4]
fmt.Printf("slice: cap %v, len %v, %v\n", cap(slice), len(slice), slice)
fmt.Printf("array: cap %v, len %v, %v\n", cap(array), len(array), array)
}

Output:
$ go run main.go
slice: cap 5, len 3, [20 30 40]
array: cap 6, len 6, [10 20 30 40 50 60]


Example 3:
package main
import "fmt"
func main () {
var array = [6] int{100,200,300,400,500,600}
var slice = array[1:4]
fmt.Printf("slice: cap %v, len %v, %v\n", cap(slice), len(slice), slice)
fmt.Printf("array: cap %v, len %v, %v\n", cap(array), len(array), array)
}

Output:
$ go run main.go
slice: cap 5, len 3, [200 300 400]
array: cap 6, len 6, [100 200 300 400 500 600]

