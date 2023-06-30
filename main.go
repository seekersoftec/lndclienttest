package main

import (
	"fmt"

	"log"

	"github.com/Bitspend01/lndclienttest/lndpay"
	"github.com/gin-gonic/gin"
)

func main() {
	lndHost := "10.128.0.7" // 10.128.0.7
	tlsPath := "./tls.cert"
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

	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		log.Println(fmt.Sprintf("%v", info))

		c.JSON(200, gin.H{
			"info": info,
		})
	})

	router.Run(":8080")

}
