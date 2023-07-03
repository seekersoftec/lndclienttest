package main

import (
	"log"

	"github.com/Bitspend01/lndclienttest/lndpay"
	"github.com/gin-gonic/gin"
)

func main() {
	lndHost := "34.28.120.31" //"34.0.192.244" // 10.128.0.7
	tlsPath := "./tls.cert"
	macDir := "./"
	network := "testnet" // testnet

	lndpayService, err := lndpay.NewInstance(lndHost, tlsPath, macDir, network)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "LND Service",
		})
	})

	router.GET("/info", func(c *gin.Context) {
		info, err := lndpayService.GetInfo()
		if err != nil {
			log.Fatalf("%v", err)
			panic(err)
		}

		//
		log.Printf("%v", info)
		c.JSON(200, gin.H{
			"message": info,
		})
	})

	router.Run(":8080")
}
