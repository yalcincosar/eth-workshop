package main

import (
	"eth-workshop/wallet"
	"fmt"
)

func main() {
	privateKey, walletAddress := wallet.GenerateWallet()
	fmt.Println("Private key: " + privateKey)
	fmt.Println("Wallet address: " + walletAddress)
}
