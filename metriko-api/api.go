package metrikoapi

import (
	"log"
	"metriko/db"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	client        *mongo.Client
	addr          string
	db_name       string
	db_cpu_stor   db.CpuMetricStor
	db_iface_stor db.IfaceMetricStor
}

func NewApi(client *mongo.Client, addr string, db_name string, db_cpu_stor db.CpuMetricStor,
	db_iface_stor db.IfaceMetricStor) *Api {
	return &Api{
		client:        client,
		addr:          addr,
		db_name:       db_name,
		db_cpu_stor:   db_cpu_stor,
		db_iface_stor: db_iface_stor,
	}
}
func (a *Api) Run(wg *sync.WaitGroup, cond *sync.Cond) {
	go func(*sync.WaitGroup, *sync.Cond) {
		wg.Add(1)
		cond.L.Lock()
		cond.Wait()
		logrus.WithFields(logrus.Fields{"time": time.Now()}).Info("API starting on adress :" + a.addr)

		cpuhandler := NewCpuHandler(a.db_cpu_stor)
		ifacehandler := NewIfaceHandler(a.db_iface_stor)
		r := gin.Default()
		//CPUhandlers
		r.GET("/cpu", cpuhandler.GetCpu)
		r.GET("/cpu/:id", cpuhandler.GetCpuByID)
		//interface handlers
		r.GET("/iface", ifacehandler.GetIfaces)
		r.GET("/iface/:id", ifacehandler.GetIfacesByID)
		err := r.Run(a.addr)
		if err != nil {
			log.Fatal(err)

		}

		defer wg.Done()
	}(wg, cond)
}
