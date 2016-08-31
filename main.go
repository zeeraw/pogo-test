package main

import (
	"os"

	"github.com/pogodevorg/pgoapi-go/cli"
	"github.com/zeeraw/encrypt"
)

func main() {
	crypto := encrypt.NewCrypto()
	cli.Run(crypto, os.Args)
}
