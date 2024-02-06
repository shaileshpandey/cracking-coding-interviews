package main

import (
	"fmt"
	"main/proofpoint"
)

func main() {
	words := proofpoint.ReadWords("./proofpoint/aliceinwonderland.txt")
	fmt.Println(fmt.Sprintf("Number of words =>%d", len(words)))
}
