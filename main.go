package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mikeyfennelly1/os--char-dev-2--u2/src/sysinfo"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		panic("Usage: go run main.go <port number>")
	}
	r := gin.Default()

	r.GET("/cpu", sysinfo.GetSysinfoJSON(sysinfo.CPU_IOCTL))
	r.GET("/memory", sysinfo.GetSysinfoJSON(sysinfo.CPU_IOCTL))
	r.GET("/disk", sysinfo.GetSysinfoJSON(sysinfo.CPU_IOCTL))

	portNum := os.Args[1]
	err := r.Run(fmt.Sprintf(":%d", portNum))
	if err != nil {
		fmt.Errorf("Unable to start server on port %d\n", portNum)
		return
	}
}
