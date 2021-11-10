package main

import (
	"fmt"
	"github.com/bakhodur-nazriev/wallet/pkg/wallet"
)

func main() {
	svc := &wallet.Service{}
	account, err := svc.RegisterAccount("+99223123244")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = svc.Deposit(account.ID, 10)
	if err != nil {
		switch err {
		case wallet.ErrAmountMustBePositive:
			fmt.Println("Amount must be positive")
		case wallet.ErrAccountNotFound:
			fmt.Println("Account not found")
		}
		return
	}
	fmt.Println(account.Balance)
}
