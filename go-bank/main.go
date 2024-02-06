package main

import (
	"fmt"
	"math/rand"
	"time"
)

var users = make(map[int]User)

func main() {
	fmt.Println("**************************** Welcome to Go Bank *****************************")
	//go func() {
	//	fmt.Println("Go Routine get registered Account ")
	//	registeredAccount()
	//}()
	fmt.Println("1. Create Account")
	fmt.Println("2. Login")
	fmt.Println("3. Get Registered Account")
	fmt.Println("4. Delete Account")
	fmt.Println("5. Exit")

	for {
		fmt.Println("Enter your choice : ")
		var choice int
		scan, err := fmt.Scan(&choice)
		if err != nil {
			return
		}
		fmt.Println("scan = ", scan)
		switch choice {
		case 1:
			createAccount()

		case 2:
			fmt.Println("Welcome to Login")
			var uId int
			var pwd string
			fmt.Println("Enter userID : ")
			fmt.Scan(&uId)
			fmt.Println("Enter password : ")
			fmt.Scan(&pwd)
			fmt.Println("Entered Value : ")
			fmt.Println(uId)
			fmt.Println(pwd)
			fmt.Println("****************")
			isValidUser := login(uId, pwd)
			if isValidUser {
				fmt.Println("Valid User")
				atm(uId)

			} else {
				fmt.Println("Invalid User")
			}
		case 3:
			regUsers := users
			if len(regUsers) == 0 {
				fmt.Println("No registered users")
				// throw exception
				registeredAccount()
				//panic("No registered users. Please create an account")
				//What is the panic function? : panic function is used to create a runtime error. It stops the ordinary flow of the program and begins panicking.
				//When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
				//To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.
				//Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

			}
			for k, v := range regUsers {
				fmt.Printf("ID = %d Username = %s Password = %s Balance = %d \n", k, v.userName, v.password, v.balance)
			}
		case 4:
			deleteAccount()

		case 5:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid choice")

		}
	}

}

type User struct {
	userID   int
	userName string
	password string
	balance  int64
}

func registeredAccount() map[int]User {
	fmt.Println("Registered Account")
	//users := make(map[int]User)
	users[111] = User{userID: 111, userName: "tom", password: "t1m3+0n", balance: rand.Int63n(1000) * 10}
	users[112] = User{userID: 112, userName: "kate", password: "kt3*-", balance: rand.Int63n(1000) * 10}
	users[113] = User{userID: 113, userName: "chris", password: "chr1s*", balance: rand.Int63n(1000) * 10}

	return users
}

func deleteAccount() {
	fmt.Println("Delete Account")
	var deletedUser int
	fmt.Println("Enter user id : ")
	fmt.Scan(&deletedUser)
	if users[deletedUser].userName != "" {
		delete(users, deletedUser)
	} else {
		fmt.Println("Invalid user id")
	}
}

func createAccount() {
	fmt.Println("Create Account")
	var userName string
	var password string

	fmt.Println("Enter username : ")
	fmt.Scan(&userName)
	fmt.Println("Enter password : ")
	fmt.Scan(&password)
	userId := int(rand.Int63n(1000))
	users[userId] = User{userID: userId, userName: userName, password: password, balance: rand.Int63n(1000) * 10}

}

func login(userId int, password string) bool {
	fmt.Println("We are login")
	for i, _ := range users {
		if int(users[i].userID) == int(userId) && users[i].password == password {
			fmt.Println("We are here")
			return true
		}
		fmt.Println("We are not here, ", i)
	}
	return false
}

func atm(userId int) {

	fmt.Println("Welcome To  Your Account : ", users[userId].userName)
	for {
		var opt int
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Println("Enter your choice : ")
		selectedValue, err := fmt.Scan(&opt) // what is the return value of Scan function :  number of scanned items and error. Scan returns the number of items scanned. If that is less than the number of arguments, err will report why.
		// what is the & operator :  & is the address operator. It returns the address of a variable.
		if err != nil {
			return
		}
		fmt.Println("Opt : ", selectedValue)
		switch opt {
		case 1:
			fmt.Println("Deposit Money")
			depositToUserAccount(userId)

		case 2:
			fmt.Println("Withdraw Money")
			withDrawMoneyFromAccount(userId)

		case 3:
			fmt.Println("Check Balance")
			checkMoneyFromAccount(userId)

		case 4:
			fmt.Println("User is quiting from Login Screen")
			for i := 0; i < 5; i++ {
				time.Sleep(time.Second * 3)
				fmt.Print(".\t")
			}
			fmt.Println()
			return
		default:
			fmt.Println("Invalid choice")

		}

	}

}

func depositToUserAccount(userId int) {
	var amountDeposit int
	fmt.Println("Enter the amount deposit")
	fmt.Scan(&amountDeposit)
	fmt.Println("Your money is depositing")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 3)
		fmt.Print(".\t")
	}
	fmt.Println()
	if amountDeposit < 5 {
		fmt.Println("You can deposit money min 5 $")
		depositToUserAccount(userId)
	}
	if amountDeposit > 10000 {
		fmt.Println("You can deposit money max 10000$")
		depositToUserAccount(userId)
	}
	user := users[userId]
	user.balance += int64(amountDeposit)
	users[userId] = user
	fmt.Println("Your money was successfully deposit. Current Money : ", users[userId].balance)
}

func withDrawMoneyFromAccount(userId int) {
	var amountWithdrawMoney int
	fmt.Println("Enter the withdraw money")
	fmt.Scan(&amountWithdrawMoney)
	fmt.Println("Your Money is drawing")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 3)
		fmt.Print(".\t")
	}
	fmt.Println()
	user := users[userId]
	if user.balance < int64(amountWithdrawMoney) {
		fmt.Println("Not enough balance.Please enter again if you would withdraw money")
		withDrawMoneyFromAccount(userId)
	}
	if amountWithdrawMoney%5 != 0 {
		fmt.Println("You can only draw money multiples of 5.")
		withDrawMoneyFromAccount(userId)
	}
	user.balance -= int64(amountWithdrawMoney)
	users[userId] = user
	fmt.Println("Your money was successfully draw. Current Money : ", users[userId].balance)

}

func checkMoneyFromAccount(userId int) {
	user := users[userId]
	fmt.Printf("%s named user has %d $ money\n", user.userName, user.balance)
}
