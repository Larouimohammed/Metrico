package main

import (
	"log"
	metrikoagent "metriko/metriko-agent/metriko-plugin"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("./.env.default"); err != nil {
		log.Fatal(err)
	}

	//run metriko agent
	addrServer := os.Getenv("Addr_Server")

	agent := metrikoagent.NewAgent(net.IPAddr{IP: net.IPv4(192, 168, 1, 11)}, addrServer)
	agent.StartMetriko()

}
