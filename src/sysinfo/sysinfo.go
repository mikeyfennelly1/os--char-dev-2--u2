// file to interact with the sysinfo device

package sysinfo

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
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

func getSysinfoString(infoType InfoType) (*string, error) {
	ioctlVal, err := GetSysinfoIoctlCMD(infoType)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	sysinfoDevice, err := os.OpenFile(sysinfoDeviceName, unix.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("Could not open device file: %d\n", sysinfoDeviceName)
	}
	defer sysinfoDevice.Close()

	err = changeSysinfoDevMode(sysinfoDevice, ioctlVal)
	if err != nil {
		return nil, fmt.Errorf("Could not perform ioctl on the device: %v\n", err)
	}

	buffer := make([]byte, 1024)

	_, err = sysinfoDevice.Read(buffer)
	if err != nil {
		return nil, err
	}

	json_str := string(buffer)

	return &json_str, nil
}

func changeSysinfoDevMode(fd *os.File, cmd IOCTLValue) error {
	err := unix.IoctlSetInt(int(fd.Fd()), uint(cmd), 0)
	if err != nil {
		return fmt.Errorf("Ioctl failed: %w", err)
	}

	return nil
}
