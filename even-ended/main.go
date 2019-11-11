package main

import "fmt"

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			number := i * j

			s := fmt.Sprintf("%d", number)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}
	fmt.Println(count)
}
