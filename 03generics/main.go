package main

import "fmt"

func handleAddInt(a int, b int) int {
	return a + b
}

func handleAddFloat(a float32, b float32) float32 {
	return a + b
}

// funcatinality of of handleAddInt & handleAddFloat are same, generics can fix this issue

func handleAdd[T int | float32](a T, b T) T {
	return a + b
}

func main() {
	r1 := handleAddInt(2, 3)
	r2 := handleAddFloat(3, 4)
	r3 := handleAdd(4, 5)
	r4 := handleAdd[float32](4, 5)

	fmt.Println(r1, r2, r3, r4)
}
