package main

import "fmt"

type account struct {
	balance float64
}

func (a *account) Checkbalance() {
	fmt.Printf("Balance: %.2f\n", a.balance)
}

func (a *account) withdraw(money float64) {
	if money >= 100 && a.balance >= money {
		a.balance -= money
		fmt.Println("Money successfully withdrawn")
		fmt.Printf("New Balance: %.2f\n", a.balance)
	} else {
		fmt.Println("The minimum amount of withdraw is 100")
	}
}

func (a *account) deposit(money float64) {
	if money >= 100 {
		a.balance += money
		fmt.Println("Money successfully deposited")
		fmt.Printf("New Balance: %.2f\n", a.balance)
	} else {
		fmt.Println("The minimum amount of deposit is 100")
	}

}

func list() {
	fmt.Println("Select the choice to proceed")
	fmt.Println("-----1. Check Balance-----")
	fmt.Println("-----2. Withdraw-----")
	fmt.Println("-----3. Deposit-----")
	fmt.Println("-----4. Quit-----")
}

func main() {
	a := account{
		balance: 100000.90,
	}
	fmt.Println("Hello welcome to my bank")
	for {
		list()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			a.Checkbalance()
		case 2:
			fmt.Println("Enter the amount to withdraw:")
			var w float64
			fmt.Scanln(&w)
			a.withdraw(w)

		case 3:
			fmt.Println("Enter the amount to deposit:")
			var w float64
			fmt.Scanln(&w)
			a.deposit(w)
		case 4:
			fmt.Println("Bye thank you")
			return
		default:
			fmt.Println("Enter the correct choice to continue")
		}
	}
}
