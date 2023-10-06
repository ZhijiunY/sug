package main

import (
	"fmt"

	"github.com/ZhijiunY/sug/utils"
)

func main() {
	fmt.Println("這是一個面向對象的方式")
	// 現創建 utils.NewFamilyAccount() 的指針實例，再調用 MainMenu()
	utils.NewFamilyAccount().MainMenu() // 思路非常清晰
}
