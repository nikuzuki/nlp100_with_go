package main

import "fmt"

func main() {
	str := "stressed"
	ans := ""
	for _, s := range str {
		ans = string(s) + ans
	}
	fmt.Println(ans)
}
