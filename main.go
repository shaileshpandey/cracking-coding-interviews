package main

import "main/proofpoint"

func main() {
	ipsWithWords := proofpoint.FindBackLists("./proofpoint/ip.txt")
	proofpoint.Display(ipsWithWords)
}
