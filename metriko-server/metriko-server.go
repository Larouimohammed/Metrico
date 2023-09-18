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
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			//fmt.Printf("can't read from connection :%v\n", err)
		}

		var msg Message

		err = json.Unmarshal(buffer[:n], &msg)
		if err != nil {

			// fmt.Printf("can't unmarshal json :%v\n", err)
			continue
		}

		//fmt.Printf("Message data Received: %+v\n", msg.Cpupayload.Version)

		s.InsertinDB(msg.Cpupayload, msg.Ifacepayload)

		time.Sleep(1 * time.Second)
	}
}
func (s *Server) InsertinDB(cpu hardware.CPU, inface []hardware.Iface) {

	_, err := s.Cpudb.Insertcpu(context.Background(), cpu)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Ifacedb.InsertIface(context.Background(), inface)
	if err != nil {
		log.Fatal(err)
	}

}
