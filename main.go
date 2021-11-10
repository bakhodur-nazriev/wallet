package main

import "github.com/bakhodur-nazriev/wallet/pkg/wallet"

func main() {
	svc := &wallet.Service{}
	svc.RegisterAccount("+992985305255")
}
