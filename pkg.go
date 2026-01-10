package main

import (
	_ "embed"

	"github.com/Ayushrao1970/todo/cmd"
)

func main() {

	cmd.Execute()

	// if len(Args) < 2 {
	// 	fmt.Printf("usage: todo [add|list|check|update|delete]")
	// 	return
	// }

	// switch Args[1] {
	// case "add":
	// 	t.Add(Args[2])
	// case "delete":
	// 	v, _ := strconv.Atoi(Args[2])
	// 	t.Delete(v - 1)
	// case "update":
	// 	v, _ := strconv.Atoi(Args[2])
	// 	t.Update(v-1, Args[3])
	// case "check":
	// 	v, _ := strconv.Atoi(Args[2])
	// 	t.Check(v - 1)
	// case "uncheck":
	// 	v, _ := strconv.Atoi(Args[2])
	// 	t.UnCheck(v - 1)
	// case "list":
	// 	t.Get()
	// }

}
