// 2.13.14.15=spamhaus-xbl
// 12.13.14.15=sorbs2
// 68.180.93.22=spamcannibal
// 12.102.40.70=uce_protect2
// 72.88.9.5=uce_protect1
// 68.180.93.22=SORBS1
// 72.88.9.5=spamhaus-xbl
// 33.190.26.41=spamhaus-xbl
// 12.102.40.70=sorbs1
// 68.101.8.130=uce_protect3
// 72.88.9.5=uce_protect2
// 12.13.14.15=uce_protect1
// 72.88.9.5=spamcannibal
// 68.180.93.22=spamhaus-sbl
// 12.102.40.70=Spamhaus-xbl
// 68.180.93.22=sorbs1
// 72.88.9.5=spamhaus-sbl
// 9.4.22.1=spamhaus-xbl
// 12.102.40.70=spamcannibal
// 9.4.22.1=uce_protect2
// 68.180.93.22=uce_protect1
// 12.102.40.70=psbl
// 68.101.8.130=spamhaus-xbl
// 33.190.26.41=spamhaus-sbl
// 9.4.22.1=uce_protect2
// Write a simple program that loads in a list of ip addresses and blocklists names
// and outputs each ip address and how many blocklists that ip is listed on.
// 12.13.14.15 = 2

package proofpoint

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func FindBackLists(filePath string) map[string][]string {
	ipsWithWords := map[string][]string{}

	file, _ := os.Open(filePath)
	//fmt.Println(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		ipWithWord := scanner.Text()
		ipWithWordArray := strings.Split(ipWithWord, "=")
		array := ipsWithWords[ipWithWordArray[0]]
		nextWord := strings.ToLower(ipWithWordArray[1])
		if !slices.Contains(array, nextWord) {
			ipsWithWords[ipWithWordArray[0]] = append(array, nextWord)
		}
	}
	return ipsWithWords
}

func Display(ips map[string][]string) {
	for key, value := range ips {
		fmt.Println(fmt.Sprintf("%s:%d", key, len(value)))
	}
}
