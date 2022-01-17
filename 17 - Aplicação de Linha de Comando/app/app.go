package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servido na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			// Valor padrão caso não especifique a flag
			Value: "devbook.com.br",
		},
	}

	// Slice de comandos que a aplicação consegue ler e executar
	app.Commands = []cli.Command{
		{
			Name: "ip", 
			Usage: "Busca de IPS de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			// Nome do comando
			Name: "servidores",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
		host := c.String("host")

		servidores, erro := net.LookupNS(host) // name server
		if erro != nil {
			log.Fatal(erro)
		}

		for _, servidor := range servidores {
			fmt.Println(servidor.Host)
		}
	
}

