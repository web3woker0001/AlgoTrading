package main

import (
	"context"
	"fmt"
	"log"
	"os"

	binance "github.com/adshao/go-binance/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "bnbot"
	app.Usage = "bnbot"
	app.Version = "0.0.1"
	app.Commands = []*cli.Command{
		getAccount,
	}
	// run the app
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var getAccount = &cli.Command{
	Name:  "getAccount",
	Usage: "getAccount",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "apiKey",
			Value: "",
			Usage: "apiKey",
		},
		&cli.StringFlag{
			Name: "secretKey",
			Value: "",
			Usage: "secretKey",
		},
	}
	Action: func(c *cli.Context) error {


		if c.Is("apiKey") == "" {
			fmt.Println("apiKey is required")
			return nil
		}

		if c.Is("secretKey") == "" {
			fmt.Println("secretKey is required")
			return nil
		}

		client := binance.NewFuturesClient(c.String("apiKey"), c.string("secretKey"))

		// 示例：获取账户信息
		account, err := client.NewGetAccountService().Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Account Information: %+v\n", account)
		return nil
	},
}
