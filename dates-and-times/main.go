package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------------- Date and Time Package ---------------")
	fmt.Println("The video recorded at ", time.Now())
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go Launched at ", t)
	fmt.Println(t.Format(time.ANSIC))
	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTiem %T\n", parsedTime)
	fmt.Println("Date : ", parsedTime.AddDate(0, 5, 20))

}
