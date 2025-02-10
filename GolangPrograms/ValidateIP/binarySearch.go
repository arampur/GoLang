package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Validating IP problem")
	s1 := validateIP("113.24.59.99")
	fmt.Println(s1)
}

func validateIP(s string) bool {
	str := strings.Split(s, ".")
	fmt.Println(str)

	if len(str) != 4 {
		return false
	}

	for i := 0; i < len(str); i++ {
		s1, err := strconv.Atoi(str[i])
		if err == nil {
			if s1 < 0 || s1 > 255 {
				return false
			}
		}
	}
	return true
}
