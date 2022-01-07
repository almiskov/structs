package main

import (
	"fmt"

	"github.com/almiskov/structs/list"
)

func main() {
	sl := list.New[string]()
	sl.AddMany([]string{"hello", "priv", "zdarove"})
	sl.ForEach(func(v string) { fmt.Print(v, " ") })

	fmt.Println()

	il := list.New[int]()

	fmt.Println(il.Len())

	il.AddMany([]int{1, 2, 6, 323, 4})
	il.ForEach(func(v int) { fmt.Print(v, " ") })

	fmt.Println()

	il.Remove(func(v int) bool { return v == 323 })
	il.ForEach(func(v int) { fmt.Print(v, " ") })

	fmt.Println()
}
