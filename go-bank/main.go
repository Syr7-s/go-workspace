package main

import (
	"fmt"
	"math/rand"
)

var users = make(map[int]User)

func main() {
	fmt.Println("**************************** Welvome to Go Bank *****************************")
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
			var uName string
			var pwd string
			fmt.Println("Enter username : ")
			fmt.Scan(&uName)
			fmt.Println("Enter password : ")
			fmt.Scan(&pwd)
			isValidUser := login(uName, pwd)
			if isValidUser {
				fmt.Println("Valid User")

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
	userName string
	password string
	balance  int64
}

func registeredAccount() map[int]User {
	fmt.Println("Registered Account")
	//users := make(map[int]User)
	users[1] = User{userName: "tom", password: "t1m3+0n", balance: rand.Int63n(1000) * 10}
	users[2] = User{userName: "kate", password: "kt3*-", balance: rand.Int63n(1000) * 10}
	users[3] = User{userName: "chris", password: "chr1s*", balance: rand.Int63n(1000) * 10}

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
	users[len(users)+1] = User{userName: userName, password: password, balance: rand.Int63n(1000) * 10}

}

func login(userName string, password string) bool {
	for i, _ := range users {
		if users[i].userName == userName && users[i].password == password {
			return true
		}
	}
	return false
}

func atm() {
	var opt int
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Exit")
	fmt.Println("Enter your choice : ")
	opt, _ = fmt.Scan(&opt) // what is the return value of Scan function :  number of scanned items and error. Scan returns the number of items scanned. If that is less than the number of arguments, err will report why.
	// what is the & operator :  & is the address operator. It returns the address of a variable.
	switch opt {
	case 1:
		fmt.Println("Deposit")
	case 2:
		fmt.Println("Withdraw")
	case 3:
		fmt.Println("Check Balance")
	case 4:
		fmt.Println("Exit")
	}
}
