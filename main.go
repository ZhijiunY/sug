package main

import "fmt"

func main() {

	//聲明一個變量，保存接收用戶輸入的選項
	key := ""

	//聲明一個變量，控制是否退出for循環
	loop := true

	// 顯示主菜單
	for {
		fmt.Println("------------家庭收支記帳------------")
		fmt.Println("		1 收支明細")
		fmt.Println("		2 登記收入")
		fmt.Println("		3 登記支出")
		fmt.Println("		4 退出條件")
		fmt.Println("		請選擇（1-4）:")

		fmt.Scan(&key)

		switch key {
		case "1":
			fmt.Println("------------當前收支明細------------")
		case "2":
		case "3":
			fmt.Println("登記支出..")
		case "4":
			loop = false
		default:
			fmt.Println("請輸入正確的選項..")
		}

		if !loop {
			break
		}

	}
	fmt.Println("退出使用...")
}
