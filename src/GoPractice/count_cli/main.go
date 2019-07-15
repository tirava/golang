package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main()  {
	app:= cli.NewApp()
	//app.Name = "hello_cli"
	app.Usage = "Coun up or down"
	app.Commands = []cli.Command{
		{
			Name: "up", ShortName: "u",
			Usage: "Count up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "stop, s",
					Usage: "Value to Count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start < 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := 1; i<=start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name: "down", ShortName: "d",
			Usage: "Count down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name: "start, s",
					Usage: "Value to Count down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("Stop cannot be negative.")
				}
				for i := start; i>=0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}