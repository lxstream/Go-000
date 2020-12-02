package main

import (
	"Go-000/Week02/service"
	"fmt"
)

func main() {
	user, err := service.GetUser(20)
	if err != nil {
		fmt.Println("记录日志：", err)
		return
	}
	fmt.Printf("user:%#v", user)
}
