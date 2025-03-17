package service

import (
	"fmt"
	"os"
)

func SaveCurrentQuote(value string) error {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)

	}
	_, err = f.Write([]byte(fmt.Sprintf("Dólar: %s", value)))
	if err != nil {
		return err
	}
	return nil
}
