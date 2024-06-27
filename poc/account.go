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
	Name:  "getAccountInfo",
	Usage: "getAccountInfo",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "apiKey",
			Value: "",
			Usage: "apiKey",
		},
		&cli.StringFlag{
			Name:  "secretKey",
			Value: "",
			Usage: "secretKey",
		},
	},
	Action: func(c *cli.Context) error {

		if c.String("apiKey") == "" {
			fmt.Println("apiKey is required")
			return nil
		}

		if c.String("secretKey") == "" {
			fmt.Println("secretKey is required")
			return nil
		}

		client := binance.NewFuturesClient(c.String("apiKey"), c.String("secretKey"))

		// 示例：获取账户信息
		account, err := client.NewGetAccountService().Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("Account Information: %+v\n", account)

		balance, err := client.NewGetBalanceService().Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("balance Information: %+v\n", balance)

		for _, v := range balance {
			fmt.Printf("balance: %+v\n", v)
		}

		// 查询所有订单
		orders, err := client.NewListOrdersService().Symbol("ORDIUSDT").Do(context.Background())
		for i := 0; i < len(orders); i++ {
			fmt.Printf("orders: %+v\n", orders[i])
		}
		// 挂单
		// order, err := client.NewCreateOrderService().TimeInForce("GTC").Quantity("0.5").Type("LIMIT").Side("BUY").Price("60").Symbol("ORDIUSDT").Do(context.Background())
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// print(order.OrderID)

		// // 撤单
		// cancelorder, err := client.NewCancelOrderService().Symbol("ORDIUSDT").OrderID(order.OrderID).Do(context.Background())
		// if err != nil {
		// 	fmt.Println(err)
		// 	return err
		// }
		// print(cancelorder)
		return nil
	},
}
