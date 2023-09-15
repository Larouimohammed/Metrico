package metric

import (
	"log"
	"metriko/hardware"

	"github.com/vishvananda/netlink"
)



func ListIface() []hardware.Iface {
	var Inface hardware.Iface
	var Infaces []hardware.Iface
	nt, _ := netlink.LinkList()

	for _, iface := range nt {
		Inface.Name = iface.Attrs().Name
		Inface.HardwareAddr = iface.Attrs().HardwareAddr.String()
		Inface.Type = iface.Type()
		Inface.Flags = iface.Attrs().Flags.String()
		routes, err := netlink.RouteList(iface, netlink.FAMILY_V4)

		if err != nil {
			log.Fatalf("failed to list routes: %s", err)
		}
		for _, rt := range routes {
			var route hardware.Route
			route.Dst = rt.Dst.String()
			route.Src = rt.Src.String()
			route.Gw = rt.Gw.String()
			Inface.Routes = append(Inface.Routes, route)

		}
		//fmt.Printf("%+v\n", Inface)
		Infaces = append(Infaces, Inface)

		//fmt.Println("-----------------------------------------------")
	}
	return Infaces
}
