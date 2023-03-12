package main

import (
	"fmt"
	b "multi/logic"
	"multi/logic/gross"
)

func main() {
	b.Basic = 50000
	fmt.Println(gross.GrossSalary())
}
