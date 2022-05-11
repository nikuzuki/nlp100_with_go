// 「パトカー」＋「タクシー」の文字を先頭から交互に連結して
// 文字列「パタトクカシーー」を得よ．
package main

import (
	"fmt"
)

func main() {
	str1 := "パトカー"
	str2 := "タクシー"
	ans := ""

	for i := 0; i < len(str1); i += 3 {
		ans = ans + str1[i:i+3] + str2[i:i+3]
	}
	fmt.Println(ans)
}
