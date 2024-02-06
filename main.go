package main

import (
	"fmt"

	"main/interfaces"
)

func main() {
	// ipsWithWords := companies.FindBackLists("./companies/ip.txt")
	// companies.Display(ipsWithWords)
	var s interfaces.Shape = interfaces.Square{Length: 5}
	var r interfaces.Shape = interfaces.Rectangle{Length: 5, Width: 6}

	fmt.Printf("Square Area is %d\n", s.Area())
	fmt.Printf("Rectangle Area is %d\n", r.Area())
}
