package main

import (
	"fmt"

	"github.com/morf1lo/email-validator"
)

func main() {
	email := "blablabla@example.lol"
	if err := emailvalidator.Valid(email); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email is OK")
}
