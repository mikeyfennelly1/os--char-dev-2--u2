package sysinfo

import (
	"fmt"
	"github.com/mikeyfennelly1/os--char-dev-2--u2/src/sysinfo"
	"github.com/stretchr/testify/require"
	"syscall"
	"testing"
)

func TestChangeIoctlToMemory(t *testing.T) {
	err := sysinfo.ChangeSysinfoDevMode(sysinfo.MEMORY_IOCTL)
	require.NoError(t, err)
}

func TestChangeIoctlToCPU(t *testing.T) {
	err := sysinfo.ChangeSysinfoDevMode(sysinfo.CPU_IOCTL)
	require.NoError(t, err)
}

func TestChangeIoctlToDisk(t *testing.T) {
	err := sysinfo.ChangeSysinfoDevMode(sysinfo.DISK_IOCTL)
	require.NoError(t, err)
}

func TestGetSysinfoStringDisk(t *testing.T) {
	syscall.Seteuid(1000)
	diskResult, err := sysinfo.GetSysinfoJSON(sysinfo.DISK_IOCTL)
	require.NoError(t, err)
	fmt.Println(*diskResult)
}

func TestGetSysinfoStringMemory(t *testing.T) {
	memoryResult, err := sysinfo.GetSysinfoJSON(sysinfo.MEMORY_IOCTL)
	require.NoError(t, err)
	fmt.Println(*memoryResult)
}

func TestGetSysinfoStringCPU(t *testing.T) {
	disk_result, err := sysinfo.GetSysinfoJSON(sysinfo.CPU_IOCTL)
	require.NoError(t, err)
	fmt.Println(*disk_result)
}

func TestStartServer(t *testing.T) {
	sysinfo.StartServer("8080", 10)
}
