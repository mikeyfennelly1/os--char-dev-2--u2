// file to interact with the sysinfo device

package sysinfo

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

const port = 8989
const hostIP = "127.0.0.1"
const sysinfoDeviceName = "/dev/sysinfo_cdev"

type IOCTLValue int32

const (
	CPU_IOCTL    IOCTLValue = 1
	MEMORY_IOCTL IOCTLValue = 2
	DISK_IOCTL   IOCTLValue = 3
)

func GetSysinfoString(ioctlVal IOCTLValue) (*string, error) {
	// open the device via the open() syscall
	sysinfoDevice, err := os.OpenFile(sysinfoDeviceName, unix.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("Could not open device file: %d\n", sysinfoDeviceName)
	}
	defer sysinfoDevice.Close()

	// change the mode of the device
	err = ChangeSysinfoDevMode(ioctlVal)
	if err != nil {
		return nil, fmt.Errorf("Error performing ioctl on %v: %v\n", sysinfoDevice.Name(), err)
	}

	buffer := make([]byte, 1024)

	bytesRead, err := sysinfoDevice.Read(buffer)
	if err != nil {
		return nil, err
	}

	json_str := string(buffer[:bytesRead])

	return &json_str, nil
}

func ChangeSysinfoDevMode(cmd IOCTLValue) error {
	sysinfoDevice, err := os.OpenFile(sysinfoDeviceName, unix.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("ChangeSysinfoDevMode: Could not open device file: %d\n", sysinfoDeviceName)
	}
	defer sysinfoDevice.Close()
	fmt.Println("Opened device successfully\n")
	fmt.Printf("file descriptor: %d\n", int(sysinfoDevice.Fd()))
	fmt.Printf("command number: %d\n", uint(cmd))

	err = unix.IoctlSetInt(int(sysinfoDevice.Fd()), uint(cmd), 0)
	if err != nil {
		return fmt.Errorf("Ioctl failed: %w", err)
	}

	return nil
}
