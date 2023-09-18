package metrikoapi

import (
	"log"
	"metriko/db"

	"github.com/gin-gonic/gin"
)

type IfaceHandler struct {
	IfaceStore db.IfaceMetricStor
}

func NewIfaceHandler(ifaceStore db.IfaceMetricStor) *IfaceHandler {
	return &IfaceHandler{
		IfaceStore: ifaceStore,
	}
}

func (s *IfaceHandler) GetIfaces(c *gin.Context) {

	Infaces, err := s.IfaceStore.GetIfaces(c)
	if err != nil {
		log.Fatal(err)
	}

	c.Bind(&Infaces)
	c.JSON(200, Infaces)
}

func (s *IfaceHandler) GetIfacesByID(c *gin.Context) {
	id := c.Param("id")
	iface, err := s.IfaceStore.GetIfaceByID(c, id)
	if err != nil {
		log.Fatal(err)
	}

	c.Bind(&iface)
	c.JSON(200, iface)

}
