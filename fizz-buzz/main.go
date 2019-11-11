package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		result := ""

		if i%3 == 0 {
			result = "fizz"
		}

		if i%5 == 0 {
			if result != "" {
				result += " "
			}
			result += "buzz"
		}

		if result == "" {
			result = fmt.Sprintf("%v", i)
		}

		fmt.Println(result)
	}
}
