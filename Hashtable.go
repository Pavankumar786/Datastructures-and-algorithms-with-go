Example :HashTable 
func main() {
m := make(map[string]int)
m["Apple"] = 40
m["Banana"] = 30
m["Mango"] = 50
for key, val := range m {
fmt.Print("[ ",key," -> ", val," ]")
}
fmt.Println("Apple price:", m["Apple"])
delete(m, "Apple")
fmt.Println("Apple price:", m["Apple"])
v, ok := m["Apple"]
fmt.Println("Apple price:", v, "Present:", ok)
v2, ok2 := m["Banana"]
fmt.Println("Banana price:", v2, "Present:", ok2)
}
Output:
[ Apple -> 40 ][ Banana -> 30 ][ Mango -> 50 ]
Apple price: 40
Apple price: 0
Apple price: 0 Present: false
Banana price: 30 Present:true