package main

import (
	"fmt"

	"github.com/Bitspend01/lndclienttest/lndpay"
)

func main() {
	lndHost := "10.128.0.7:10009"
	tlsPath := "lnd_tls.cert"
	macDir := "./nam"
	network := "testnet" // testnet

	lndpayService, err := lndpay.NewInstance(lndHost, tlsPath, macDir, network)
	if err != nil {
		panic(err)
	}

	info, err := lndpayService.GetInfo()
	if err != nil {
		panic(err)
	}

	fmt.Println(&info)
}
