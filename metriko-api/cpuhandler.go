package metrikoapi

import (
	"fmt"
	"metriko/db"

	"github.com/gin-gonic/gin"
)

type CpuHandler struct {
	cpuStore db.CpuMetricStor
}

func NewCpuHandler(cpuStore db.CpuMetricStor) *CpuHandler {
	return &CpuHandler{
		cpuStore: cpuStore,
	}
}
func (s *CpuHandler) GetCpuByID(c *gin.Context) {
	id := c.Param("id")

	cpu, err := s.cpuStore.GetCpuByID(c,id)
	if err != nil {
		fmt.Printf("id not found ")
	}
		c.Bind(&cpu)
		c.JSON(200,cpu)
	
	}


func (s *CpuHandler) GetCpu(c *gin.Context) {
	cpus, err := s.cpuStore.GetCpus(c)

	if err != nil {
		fmt.Printf("id not found ")
	}
	c.Bind(&cpus)
    c.JSON(200,cpus)

}
