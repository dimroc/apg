package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcutil/base58"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "apg"
	app.Usage = "automatic password generator. Inspired by the old brew install apg."
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "hex",
			Usage: "prints hexadecimal instead of the default base58",
		},
	}
	app.Action = printGeneratedPassword
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printGeneratedPassword(c *cli.Context) error {
	uuid := uuid.NewV4()
	if c.Bool("hex") {
		fmt.Println(hex.EncodeToString(uuid.Bytes()))
		return nil
	}

	fmt.Println(base58.Encode(uuid.Bytes()))
	return nil
}
