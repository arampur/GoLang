package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "123.235.153.425"
	fmt.Println("result:", validateIP(str))
}

func validateIP(ip string) bool {
	//split the address based on .

	splitStr := strings.Split(ip, ".")

	re := regexp.MustCompile(`^[0-9]+$`)

	if len(splitStr) != 4 {
		return false
	}

	//check if each of them is number and between 0 - 255

	for i := 0; i < len(splitStr); i++ {
		strNum, _ := strconv.Atoi(splitStr[i])
		if !(re.MatchString(splitStr[i]) && strNum >= 0 && strNum <= 255) {
			return false
		}
	}

	//its a valid IP
	return true
}
