package main

import (
	"context"
	"log"
	"metriko/db"
	"metriko/handlers"
	"metriko/metric-agent"
	"os"

	"github.com/gin-gonic/gin"
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
	addr := os.Getenv("HTTP_ADDR_Server_Listen")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_db_url))
	if err != nil {
		log.Fatal(err)
	}
	var c = db.NewMongoCpuMetricStore(client, mongo_db_name)
	var i = db.NewMongoIfaceMetricStore(client, mongo_db_name)
	//get metric
	cpu := metric.GetCpu()
	Infaces := metric.ListIface()
	//insert metric in db
	_, err = c.Insertcpu(context.Background(), cpu)
	if err != nil {
		log.Fatal(err)
	}
	err = i.InsertIface(context.Background(), Infaces)
	if err != nil {
		log.Fatal(err)
	}

	cpuhandler := handlers.NewCpuHandler(c)
	ifacehandler := handlers.NewIfaceHandler(i)
	r := gin.Default()
	//CPUhandlers
	r.GET("/cpu", cpuhandler.GetCpu)
	r.GET("/cpu/:id", cpuhandler.GetCpuByID)
	//interface handlers
	r.GET("/iface", ifacehandler.GetIfaces)
	r.GET("/iface/:id", ifacehandler.GetIfacesByID)
	err = r.Run(addr)
	if err != nil {
		log.Fatal(err)

	}

}
