package main

import (
	"fmt"

	"github.com/Ryan-18-system/desafio-client-server-goexpert/client/service"
)

func main() {
	quoteResquest, err := service.GetCurrentQuote()
	if err != nil {
		fmt.Println(err)
		return
	}
	service.SaveCurrentQuote(*quoteResquest)
}
