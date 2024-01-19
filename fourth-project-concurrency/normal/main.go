package main

import (
	"fmt"
	user "fourth-project-concurrency/internal"
	"time"
)

func main() {
	//100 tane olsa nasıl yapılabilirdi. Paralel olmasa bile
	now := time.Now()
	var users []user.User
	for i := 0; i < 10; i++ {
		users = append(users, user.GetRandomUser())
		//fmt.Println(user.GetRandomUser())
	}
	fmt.Printf("Users generated in %s seconds", time.Since(now))
}
