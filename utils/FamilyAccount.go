package utils

import "fmt"

type FamilyAccount struct {
	// 聲明必須的字段

	key     string  //聲明一個變量，保存接收用戶輸入的選項
	loop    bool    //聲明一個變量，控制是否退出for循環
	balance float64 // 定義帳戶餘額
	money   float64 // 每次收支金額
	note    string  // 每次收支說明

	// 定義一個變量，紀錄是否有收支的行為
	// 預設初始化情況下，無收之行為
	flag bool

	// 收支的詳情用字符串紀錄
	// 當有收支時，只需要對details 進行拼接處理即可
	details string
}

// 編寫一個工廠模式的結構方法，返回一個 *FamilyAccount 實例，提高效率
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		// 放入聲明的初始化值
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t帳戶金額\t收支金額\t說明",
	}
}

// 將顯示明細寫成一個方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("------------當前收支明細------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("目前沒有收支，請來一筆！")
	}
}

// 將登記收入寫成一個方法，和 *FamilyAccount 綁定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金額：")
	fmt.Scan(&this.money)
	// 修改帳戶餘額
	this.balance += this.money
	fmt.Println("本次收入說明：")
	fmt.Scan(&this.note)
	// 將收入情況拼接到 details 變量
	// 收入  1100   1000  有人發紅包
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金額：")
	fmt.Scan(&this.money)
	// 條件判斷
	if this.money > this.balance {
		fmt.Println("餘額金額不足")
		// break 原先在 switch 需要使用
	}

	this.balance -= this.money
	fmt.Println("本次支出說明：")
	fmt.Scan(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true

}

func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

// 給該結構體 binding 相應的方法
// 顯示主菜單
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n------------家庭收支記帳------------")
		fmt.Println("		1 收支明細")
		fmt.Println("		2 登記收入")
		fmt.Println("		3 登記支出")
		fmt.Println("		4 退出條件")
		fmt.Println("		請選擇（1-4）:")

		fmt.Scan(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("請輸入正確的選項..")
		}

		if !this.loop {
			break
		}
	}
}
