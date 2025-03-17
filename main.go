package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mikeyfennelly1/os--char-dev-2--u2/src/sysinfo"
	"net/http"
	"syscall"
)

// Define a struct for a simple resource
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{
	{ID: "1", Name: "Item One"},
	{ID: "2", Name: "Item Two"},
}

func main() {
	syscall.Seteuid(1000)
	r := gin.Default()
	r.GET("/", getCPUData)

	r.GET("/cpu", getCPUData)
	r.GET("/memory", getMemoryData)
	r.GET("/disk", getDiskData)

	// Run the server
	r.Run(":8080")
}

// Handlers
func getCPUData(c *gin.Context) {
	cpuJSON, err := sysinfo.GetSysinfoJSON(sysinfo.CPU_IOCTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not get CPU data from device")
		return
	}
	c.JSON(http.StatusOK, cpuJSON)
}

func getMemoryData(c *gin.Context) {
	memJSON, err := sysinfo.GetSysinfoJSON(sysinfo.MEMORY_IOCTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not get memory data from device")
		return
	}
	c.JSON(http.StatusOK, memJSON)
}

func getDiskData(c *gin.Context) {
	diskJSON, err := sysinfo.GetSysinfoJSON(sysinfo.DISK_IOCTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not get disk data from device")
		return
	}
	c.JSON(http.StatusOK, diskJSON)
}
