package fileops

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindFamilyWithMemberCount() map[string][]string {
	familyWithCount := map[string][]string{}
	file, _ := os.Open("inputfiles/family.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		lastName := getFamilyName(name)
		members := familyWithCount[lastName]
		members = append(members, name)
		familyWithCount[lastName] = members
	}

	fmt.Println(familyWithCount)
	return familyWithCount
}

func getFamilyName(name string) string {
	fnameWithlname := strings.Fields(name)
	lastName := fnameWithlname[0]
	if len(fnameWithlname) > 1 {
		lastName = fnameWithlname[1]
	}

	switch {
	case strings.Contains(lastName, "Delgado"), strings.Contains(lastName, "Tucker"), strings.Contains(lastName, "Pritchett"):
		return "Pritchett"
	}

	return lastName
}
