package main

import (
	"fmt"
	user "fourth-project-concurrency/internal"
	"time"
)

func main() {
	// go routine değer yerine değer yerine kanal döndürsün.
	c := getRandomUsers()
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(10 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}

func fanIn(u1, u2 <-chan user.User) <-chan user.User {
	ch := make(chan user.User)
	go func() {
		for {
			select {
			case s := <-u1:
				ch <- s
			case s := <-u2:
				ch <- s
			}
		}
	}()
	return ch
}

func getRandomUsers() <-chan user.User {
	ch := make(chan user.User)

	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				ch <- user.GetRandomUser()
			}()
			//ch <- user.GetRandomUser()

		}
	}()

	return ch
}
