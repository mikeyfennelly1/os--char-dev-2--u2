// file to interact with the sysinfo device

package sysinfo

import (
	"fmt"
)

const port = 8989
const hostIP = "127.0.0.1"
const sysinfoDeviceName = "/dev/sysinfo"

type InfoType string
type IOCTLValue int32

const (
	MEMORY       InfoType   = "MEMORY"
	DISK         InfoType   = "DISK"
	CPU          InfoType   = "CPU"
	MEMORY_IOCTL IOCTLValue = 1
	DISK_IOCTL   IOCTLValue = 2
	CPU_IOCTL    IOCTLValue = 3
)

func GetSysinfoIoctlCMD(infoType InfoType) (IOCTLValue, error) {
	switch infoType {
	case MEMORY:
		return MEMORY_IOCTL, nil
	case DISK:
		return DISK_IOCTL, nil
	case CPU:
		return CPU_IOCTL, nil
	default:
		return -1, fmt.Errorf("Invalid IOCTL type: %s\n", infoType)
	}
}

//func getSysinfoString(infoType InfoType) (*string, error) {
//	ioctlVal, err := getSysinfoIoctlCMD(infoType)
//	if err != nil {
//		return nil, fmt.Errorf(err.Error())
//	}
//
//}
