package metrikoagent

import (
	"encoding/json"
	"fmt"
	"log"
	"metriko/hardware"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Message struct {
	Type         string
	Cpupayload   hardware.CPU
	Ifacepayload []hardware.Iface
}

type Agent struct {
	Machine net.IPAddr
	Addr    string
}

func NewAgent(machine net.IPAddr, addr string) *Agent {
	return &Agent{
		Machine: machine,
		Addr:    addr,
	}
}
func (a *Agent) SendMetriko(wg *sync.WaitGroup) {
	logrus.WithFields(logrus.Fields{"time": time.Now()}).Info("Metriko Agent Starting on adress :" + a.Addr)

	conn, err := net.Dial("tcp", a.Addr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()
	defer wg.Done()

	var msg Message
	msg.Type = "json"
	msg.Cpupayload = a.GetCpu()
	msg.Ifacepayload = a.ListIface()

	data, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

}
