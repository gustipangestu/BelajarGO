package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	// symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	// fmt.Println(RMB, symbol[RMB])
	tesArray()
}

func tesArray() {
	r := [...]int{99: -6}
	fmt.Println("array index ke-", (len(r)), " adalah ", r[len(r)-1])
	fmt.Println("array index ke-", (len(r) - 10), " adalah ", r[len(r)-9])
}
