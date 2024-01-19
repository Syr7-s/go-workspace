package main

import (
	"fmt"
	user "fourth-project-concurrency/internal"
	"time"
)

func main() {
	now := time.Now()
	var users []user.User
	ch := make(chan user.User) // channel tanım
	//go getRandomUsers(ch)      // go routine oluşturuluyor. 9 ncu satırda kanaldan okuyor.
	//ch := getRandomUsers()
	for i := 0; i < 10; i++ {
		users = append(users, <-ch) //channel dan okuma
	}
	fmt.Printf("Users generated in %s seconds", time.Since(now))
}

func getRandomUsers() <-chan user.User {
	ch := make(chan user.User)

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				ch <- user.GetRandomUser()
			}()
		}
	}()

	return ch
}
