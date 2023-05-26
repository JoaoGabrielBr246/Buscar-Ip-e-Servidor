package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a apliação de linha de comando pronta para ser executada
func Gerar() *cli.App { //retorno do pacote cli
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	} //Comandos para usar

	app.Commands = []cli.Command{
		{
			//Buscando IP
			Name:   "ip",
			Usage:  "Busca de IP de endereços na net",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome do servidores na net",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app

}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//funcao do pacote net
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

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
