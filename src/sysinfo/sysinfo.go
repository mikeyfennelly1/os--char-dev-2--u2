// file to interact with the sysinfo device

package sysinfo

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"sync"
)

const sysinfoDeviceName = "/dev/sysinfo"

type IOCTLValue int32

// integer values that correspond to ioctl
// functions on the character device
const (
	CPU_IOCTL    IOCTLValue = 1
	MEMORY_IOCTL IOCTLValue = 2
	DISK_IOCTL   IOCTLValue = 3
)

var mu sync.Mutex

// Get the sysinfo data in JSON string format from the device
//
// returns data read from device
func GetSysinfoJSON(ioctlVal IOCTLValue) (*string, error) {
	// lock mutex in critical section to prevent
	// conflictions with other request handling threads
	mu.Lock()

	// open the device via the open() syscall
	sysinfoDevice, err := os.OpenFile(sysinfoDeviceName, unix.O_RDWR, 0666)
	if err != nil {
		return nil, fmt.Errorf("Could not open device file: %s\n", sysinfoDeviceName)
	}
	defer sysinfoDevice.Close()

	// change the mode of the device
	err = ChangeSysinfoDevMode(ioctlVal)
	if err != nil {
		return nil, fmt.Errorf("Error performing ioctl on %v: %v\n", sysinfoDevice.Name(), err)
	}

	// read the device contents into buffer
	buffer := make([]byte, 1024)
	bytesRead, err := sysinfoDevice.Read(buffer)
	if err != nil {
		return nil, err
	}

	// unlock the mutex after critical read is done
	mu.Unlock()

	jsonStr := string(buffer[:bytesRead])

	return &jsonStr, nil
}

// Change the current_info_type of the sysinfo device
func ChangeSysinfoDevMode(cmd IOCTLValue) error {
	sysinfoDevice, err := os.OpenFile(sysinfoDeviceName, unix.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("ChangeSysinfoDevMode: Could not open device file: %d\n", sysinfoDeviceName)
	}
	defer sysinfoDevice.Close()

	err = unix.IoctlSetInt(int(sysinfoDevice.Fd()), uint(cmd), 0)
	if err != nil {
		return fmt.Errorf("Ioctl failed: %w", err)
	}

	return nil
}
