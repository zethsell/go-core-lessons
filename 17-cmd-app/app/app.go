package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "busca Ips e nomes de server na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ips de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: buscarServer,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServer(c *cli.Context) {
	host := c.String("host")
	servidores, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
