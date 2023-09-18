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
func (a *Agent) StartMetriko(wg *sync.WaitGroup) {
	go func() {
		wg.Add(1)
		logrus.WithFields(logrus.Fields{"time": time.Now()}).Info("Metriko Agent Starting on adress :" + a.Addr)

		conn, err := net.Dial("tcp", a.Addr)
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}
		buffer := make([]byte, 1024)
		var msg Message

		for {
			n, err := conn.Read(buffer)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(buffer[:n], &msg)
			if msg.Type == "request" {
                fmt.Printf("--------------------%v\n",msg)
				msg.Type = "response"
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
			}
			defer wg.Done()
			defer conn.Close()
			time.Sleep(1 * time.Second)
		}

	}()
}
