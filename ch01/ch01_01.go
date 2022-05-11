package main

import (
	"fmt"
)

func main() {
	str := "パタトクカシーー"

	for i := 0; i < len(str); i += 6 {
		fmt.Println(str[i : i+3])
	}

}
