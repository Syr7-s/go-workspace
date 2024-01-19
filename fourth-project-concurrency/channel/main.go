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
	go getRandomUsers(ch)      // go routine oluşturuluyor. 9 ncu satırda kanaldan okuyor.
	for i := 0; i < 10; i++ {
		users = append(users, <-ch) //channel dan okuma
	}
	fmt.Printf("Users generated in %s seconds", time.Since(now))
}

func getRandomUsers(ch chan user.User) {
	for i := 0; i < 10; i++ { //10 tane eleman kanala atılıyor.
		//	ch <- user.GetRandomUser() //getRandomUser fonksiyonu çağrılıyor.Imperative ilerliyor
		//}
		//	time.Sleep(1 * time.Second)
		go func() { //functional ilerliyor. 10 tane go routine oluşturuluyor.
			ch <- user.GetRandomUser()
		}() //fonksiyon çağrılıyor
	}
}
