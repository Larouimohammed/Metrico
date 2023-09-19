package metrikoserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"metriko/db"
	"metriko/hardware"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

type Message struct {
	Type         string
	Cpupayload   hardware.CPU
	Ifacepayload []hardware.Iface
}

type Server struct {
	Addr    string
	Cpudb   db.CpuMetricStor
	Ifacedb db.IfaceMetricStor
}

func NewServer(Cpudb db.CpuMetricStor, Ifacedb db.IfaceMetricStor, addr string) *Server {
	return &Server{
		Cpudb:   Cpudb,
		Ifacedb: Ifacedb,
		Addr:    addr,
	}
}

func (s *Server) Start() {
	logrus.WithFields(logrus.Fields{"time": time.Now()}).Info("Metriko Server Starting on adress :" + s.Addr)
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	buffer := make([]byte, 1024)
	for {
		msg := Message{}
		msg.Type = "request"
		data, err := json.Marshal(msg)
		conn.Write(data)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("can't read from connection :%v\n", err)
			continue
		}
		err = json.Unmarshal(buffer[:n], &msg)
		if err != nil {
			fmt.Printf("can't unmarshal json :%v\n", err)
			continue
		}

		_, err = s.Cpudb.Insertcpu(context.Background(), msg.Cpupayload)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(msg)
		err = s.Ifacedb.InsertIface(context.Background(), msg.Ifacepayload)
		if err != nil {
			log.Fatal(err)
		}

		defer listener.Close()
		time.Sleep(1 * time.Second)
	}
}
