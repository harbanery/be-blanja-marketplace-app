package services

import (
	"os"
	"sync"

	"github.com/veritrans/go-midtrans"
)

var Client midtrans.Client
var midtransOnce sync.Once

func InitMidtrans() {
	midtransOnce.Do(func() {
		serverKey := os.Getenv("SERVER_KEY")
		clientKey := os.Getenv("CLIENT_KEY")

		Client = midtrans.Client{
			ServerKey: serverKey,
			ClientKey: clientKey,
		}

		Client.APIEnvType = midtrans.Sandbox
	})
}
