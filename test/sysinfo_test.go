package sysinfo

import (
	"github.com/mikeyfennelly1/os--char-dev-2--u2/src/sysinfo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSysinfoIoctlCMD(t *testing.T) {
	result, err := sysinfo.GetSysinfoIoctlCMD(sysinfo.MEMORY)
	require.NoError(t, err)
	require.Equal(t, sysinfo.MEMORY_IOCTL, result)
}
