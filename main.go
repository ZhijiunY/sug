package main

import "fmt"

func main() {

	//聲明一個變量，保存接收用戶輸入的選項
	key := ""

	//聲明一個變量，控制是否退出for循環
	loop := true

	// 定義帳戶餘額
	balance := 10000.0

	// 每次收支金額
	money := 0.0

	// 每次收支說明
	note := ""

	// 收支的詳情用字符串紀錄
	// 當有收支時，只需要對details 進行拼接處理即可
	details := "收支\t帳戶金額\t收支金額\t說明"

	// 顯示主菜單
	for {
		fmt.Println("\n------------家庭收支記帳------------")
		fmt.Println("		1 收支明細")
		fmt.Println("		2 登記收入")
		fmt.Println("		3 登記支出")
		fmt.Println("		4 退出條件")
		fmt.Println("		請選擇（1-4）:")

		fmt.Scan(&key)

		switch key {
		case "1":
			fmt.Println("------------當前收支明細------------")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金額：")
			fmt.Scan(&money)
			// 修改帳戶餘額
			balance += money
			fmt.Println("本次收入說明：")
			fmt.Scan(&note)
			// 將收入情況拼接到 details 變量
			// 收入  1100   1000  有人發紅包
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
		case "3":
			fmt.Println("本次支出金額：")
			fmt.Scan(&money)
			// 條件判斷
			if money > balance {
				fmt.Println("餘額金額不足")
				break
			}
			balance -= money
			fmt.Println("本次支出說明：")
			fmt.Scan(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)

		case "4":
			fmt.Println("是否確定要退出？（Y/N）")
			choice := ""
			for {
				fmt.Scan(&choice)
				if choice == "Y" || choice == "N" {
					break
				}
				fmt.Println("你的輸入有誤，請重新輸入 Y/N")
			}

			if choice == "Y" {
				loop = false
			}

		default:
			fmt.Println("請輸入正確的選項..")
		}

		if !loop {
			break
		}

	}
	fmt.Println("退出使用...")
}
