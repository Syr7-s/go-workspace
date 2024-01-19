package main

import (
	"fmt"
	user "fourth-project-concurrency/internal"
	"sync"
	"time"
)

func main() {
	// yapı olarak deltayı hemen ekle.

	now := time.Now()
	var users []user.User
	var ws sync.WaitGroup
	// boyutu biliniyor ise
	//ws.Add(2) verilebilir.
	for i := 0; i < 2; i++ {
		ws.Add(1)
		go getRandomUser(&ws, users)
	}
	//fmt.Println(len(users))
	ws.Wait()
	fmt.Printf("Users generated in %s seconds", time.Since(now))
}

func getRandomUser(wg *sync.WaitGroup, users []user.User) {
	defer wg.Done() // artırdığın count 1 azaltır.
	users = append(users, user.GetRandomUser())

}
