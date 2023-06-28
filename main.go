package main

import (
	"fmt"

	"github.com/Bitspend01/lndclienttest/lndpay"
)

func main() {
	lndHost := "34.28.120.31:10009"
	tlsPath := "lnd_tls.cert"
	macDir := "./"
	network := "testnet" // testnet

	lndpayService, err := lndpay.NewInstance(lndHost, tlsPath, macDir, network)
	if err != nil {
		panic(err)
	}

	info, err := lndpayService.GetInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(info)
}
