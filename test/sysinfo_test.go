package sysinfo

import (
	"fmt"
	"github.com/mikeyfennelly1/os--char-dev-2--u2/src/sysinfo"
	"github.com/stretchr/testify/require"
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
	disk_result, err := sysinfo.GetSysinfoString(sysinfo.DISK_IOCTL)
	require.NoError(t, err)
	fmt.Println(*disk_result)
}

func TestGetSysinfoStringMemory(t *testing.T) {
	disk_result, err := sysinfo.GetSysinfoString(sysinfo.MEMORY_IOCTL)
	require.NoError(t, err)
	fmt.Println(*disk_result)
}

func TestGetSysinfoStringCPU(t *testing.T) {
	disk_result, err := sysinfo.GetSysinfoString(sysinfo.CPU_IOCTL)
	require.NoError(t, err)
	fmt.Println(*disk_result)
}
