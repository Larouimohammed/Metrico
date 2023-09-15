package metric

import (
	"bytes"
	"fmt"
	"metriko/hardware"
	"os/exec"

	"gopkg.in/xmlpath.v2"
)

func GetCpu() *hardware.CPU {

	var CPU hardware.CPU

	cmdOut, err := exec.Command("lshw", "-C", "processor", "-xml").Output()

	if err != nil {
		fmt.Printf("%s\n", err)
	}
	//output := string(cmdOut[:])
	//fmt.Println(output)
	root, err := xmlpath.Parse(bytes.NewReader(cmdOut))
	if err != nil {
		fmt.Println(err)
	}
	path := xmlpath.MustCompile("/list/node/product")
	if value, ok := path.String(root); ok {
		CPU.Product = value
	}

	path2 := xmlpath.MustCompile("/list/node/vendor")
	if value, ok := path2.String(root); ok {
		CPU.Vendor = value
	}
	path3 := xmlpath.MustCompile("/list/node/version")
	if value, ok := path3.String(root); ok {
		CPU.Version = value
	}
	path4 := xmlpath.MustCompile("/list/node/width")
	if value, ok := path4.String(root); ok {
		CPU.Width = value
	}

	return &CPU
}
