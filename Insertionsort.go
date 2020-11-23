Example 1:
package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1}
    insertionSort(sample)

    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    insertionSort(sample)
}

func insertionSort(arr []int) {
    len := len(arr)
    for i := 1; i < len; i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    
    fmt.Println("After Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}
Output:

After Sorting
1
2
3
4
5

After Sorting
-3
-1
1
2
3
4
5
7
8


Example 2:
package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
 
    slice := generateSlice(20)
    fmt.Println("\n--- Unsorted --- \n\n", slice)
    insertionsort(slice)
    fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
 
// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
 
    slice := make([]int, size, size)
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - rand.Intn(999)
    }
    return slice
}
  
func insertionsort(items []int) {
    var n = len(items)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if items[j-1] > items[j] {
                items[j-1], items[j] = items[j], items[j-1]
            }
            j = j - 1
        }
    }
}
C:\golang\time>go run sort.go

--- Unsorted ---

[503 -319 245 -537 -167 -804 469 -457 143 184 376 825 11 -309 -144 152 -493 684 165 29]

--- Sorted ---

[-804 -537 -493 -457 -319 -309 -167 -144 11 29 143 152 165 184 245 376 469 503 684 825]