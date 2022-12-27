package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	fmt.Println(nowTime) //current y-m-d with time

	fmt.Println(nowTime.Year())               //current year
	fmt.Println(nowTime.Month())              //current month
	fmt.Println(nowTime.Day())                //current date
	fmt.Println(nowTime.Date())               //separetely prints Y M D
	fmt.Println(nowTime.Format("2006-01-02")) //current date in YYYY-MM-DD
}
