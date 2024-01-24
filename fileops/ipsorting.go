package fileops

import (
	"bufio"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ipaddressIPS []net.IP

func (ips ipaddressIPS) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}

func (ips ipaddressIPS) Len() int {
	return len(ips)
}

func (ips ipaddressIPS) Less(i, j int) bool {
	return lessIP(ips[i], ips[j])
}

// lessIP compares two net.IP addresses
func lessIP(ip1, ip2 net.IP) bool {
	return bytesLess(ip1.To4(), ip2.To4())
}

// bytesLess compares two byte slices lexicographically
func bytesLess(b1, b2 []byte) bool {
	for i := 0; i < len(b1) && i < len(b2); i++ {
		if b1[i] != b2[i] {
			return b1[i] < b2[i]
		}
	}
	return len(b1) < len(b2)
}

func ReadIPS(filePath string) []string {
	ips := []net.IP{}
	ipsText := []string{}
	if file, err := os.Open(filePath); err != nil {
		return nil
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ipText := scanner.Text()
			ips = append(ips, net.ParseIP(ipText))
		}
	}
	sort.Sort(ipaddressIPS(ips))
	fileToWrite, _ := os.Create("out.txt")

	for _, v := range ips {
		ipsText = append(ipsText, v.String())
		fileToWrite.WriteString(v.String())
	}

	return ipsText
}

func ReadIPS2(filePath string) []string {
	ipsText := []string{}
	if file, err := os.Open(filePath); err != nil {
		return nil
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ipText := scanner.Text()
			ipsText = append(ipsText, ipText)
		}
	}
	sort.Slice(ipsText, func(i, j int) bool {
		return compareIP(ipsText[i], ipsText[j]) < 0
	})

	fileToWrite, _ := os.Create("out.txt")
	defer fileToWrite.Close()
	for _, v := range ipsText {
		fileToWrite.WriteString(v)
		fileToWrite.WriteString("\n")
	}

	return ipsText
}

func compareIP(ip1, ip2 string) int {
	ipA := strings.Split(ip1, ".")
	ipB := strings.Split(ip2, ".")
	for counterI := 0; counterI < len(ipA); counterI++ {
		currentAInt, _ := strconv.Atoi(ipA[counterI])
		currentBInt, _ := strconv.Atoi(ipB[counterI])

		if currentAInt != currentBInt {
			return currentAInt - currentBInt
		}
	}
	return 0
}
