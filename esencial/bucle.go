package main

import (
	"fmt"
)

func main() {
	elefantes := 12

	for i := 0; i < elefantes; i++ {
		if i == 1 {
			fmt.Println(i, "elefante se balanceaba")
		} else {
			fmt.Println(i, "elefantes se balanceaban")
		}
	}
}
