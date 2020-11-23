package main
import "fmt"
 func binarySearch(needle int, haystack []int) bool {
  low := 0
    high := len(haystack) - 1
 
    for low <= high{
        median := (low + high) / 2
 
        if haystack[median] < needle {
            low = median + 1
        }else{
            high = median - 1
        }
    }
 
    if low == len(haystack) || haystack[low] != needle {
        return false
    }
 
    return true
}
 
 
func main(){
    items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
    fmt.Println(binarySearch(63, items))
}

Output:
C:\golang\time>go run sort.go

--- Unsorted ---

[-129 755 38 354 248 6 -160 212 -184 336 -85 222 776 587 490 -503 -420 -54 -502 -341]

--- Sorted ---

[-503 -502 -420 -341 -184 -160 -129 -85 -54 6 38 212 222 248 336 354 490 587 755 776]