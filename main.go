package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/nutrun/lentil"
)

func main() {
	app := cli.NewApp()
	app.Name = "pinto"
	app.Usage = "pinto [command [args]]"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "host",
			Value:  "0.0.0.0:11300",
			Usage:  "Host for beanstalk",
			EnvVar: "BEANSTALK_HOST",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "list-tubes",
			Action: func(c *cli.Context) {
				listTubes(c.GlobalString("host"))
			},
			Aliases: []string{"l"},
			Usage: "Lists all available tubes",
		},
		{
			Name: "stats-tube",
			Action: func(c *cli.Context) {
                statsTube(c.GlobalString("host"), c.Args().First())
			},
			Aliases: []string{"p"},
			Usage: `Peeks at jobs in a specific tube (ex: pinto peek "tube-name")`,
		},
	}

	app.Run(os.Args)
}

func listTubes(host string) {
	conn, err := connect(host)
	if err != nil {
		fmt.Printf("error: Could not connect to Beanstalkd \n%v\n", err)
        os.Exit(1)
	}
    
    tubes, err := conn.ListTubes()
    if err != nil {
        fmt.Printf("error: Could not list tubes\n%v\n", err)
        os.Exit(1)
    }
    
    for _, tube := range tubes {
        fmt.Println(tube)
    }
}

func statsTube(host string, tube string) {
    conn, err := connect(host)
    if err != nil {
        fmt.Printf("error: Could not connect to Beanstalkd \n%v\n", err)
        os.Exit(1)
    }
    
    stats, err := conn.StatsTube(tube)
    if err != nil {
        fmt.Printf("error: Could not get tube stats\n%v\n", err)
        os.Exit(1)
    }
    
    for name, stat := range stats {
        fmt.Printf("%v: %v\n", name, stat)
    }
}

func connect(host string) (*lentil.Beanstalkd, error) {
	return lentil.Dial(host)
}
