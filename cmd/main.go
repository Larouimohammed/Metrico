package main

import (
	"context"
	"log"
	"metriko/db"
	metrikoagent "metriko/metriko-agent"
	metrikoserver "metriko/metriko-server"
	"net"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load("./.env.default"); err != nil {
		log.Fatal(err)
	}
	mongo_db_url := os.Getenv("MONGO_DB_URL")
	mongo_db_name := os.Getenv("Mongo_DB_Name")
	// addr := os.Getenv("HTTP_ADDR_Server_Listen")
	addrServer := os.Getenv("Addr_Server")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_db_url))
	if err != nil {
		log.Fatal(err)
	}
	var c = db.NewMongoCpuMetricStore(client, mongo_db_name)
	var i = db.NewMongoIfaceMetricStore(client, mongo_db_name)
	wg := &sync.WaitGroup{}
	//Run api
	// api := metrikoapi.NewApi(client, addr, mongo_db_name, c, i)
	// api.Run(wg)

	//run metriko agent

	agent := metrikoagent.NewAgent(net.IPAddr{IP: net.IPv4(192, 168, 1, 11)}, addrServer)
	agent.StartMetriko(wg)

	//run server
	server := metrikoserver.NewServer(c, i, addrServer)
	server.Start()
	wg.Wait()

}
