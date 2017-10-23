package utils

import (
	"fmt"
)

func SendMail(typeemail string,email string,obj interface{}) {
	// TODO may be need pools limit concurrent nums
	go func() {
		fmt.Println("send email")
	}()
}
