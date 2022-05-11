// “Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.”
//という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	tmp := strings.Split(str, " ")
	ans := []int{}

	for _, word := range tmp {
		replaced_tmp := strings.Replace(word, ",", "", 1)
		replaced_tmp = strings.Replace(replaced_tmp, ".", "", 1)
		ans = append(ans, len(replaced_tmp))
	}
	fmt.Println(ans)
}
