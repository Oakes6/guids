package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type volatileState struct {
	timestamp     int
	clockSequence int
	nodeID        int
}

// GUID structure w/ traditional fields
type GUID struct {
	timeLow                 uint32
	timeMid                 uint16
	timeHighAndVersion      uint16
	clockSeqHighAndReserved uint8
	clockSeqLow             uint8
	node                    [6]byte
}

var mutex *sync.Mutex

func main() {
	fmt.Println("hello guid world")
	// generate unoptimal alg first (read/write to shared storage per GUID, )
	fmt.Println("Printing Mac Address/es")
	_, err := getMacAddress()
	if err != nil {
		log.Fatal(err)
	}

}

// Returns the current UTC time as a 60-bit count of 100-nanosecond intervals since 00:00:00.00, 15 October 1582
func getTimestamp() uint64 {
	return 0
}

// returns the current system's MAC address
func getMacAddress() ([6]byte, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return [6]byte{}, err
	}

	// look for 'en0' interface
	var addr net.HardwareAddr
	for _, v := range interfaces {
		a := v.HardwareAddr.String()
		if a != "" && v.Name == "en0" {
			addr = v.HardwareAddr
		}
	}

	// return a single address, and nil if non-present
	if addr == nil {
		fmt.Println("en0 address not available")
		return [6]byte{}, nil
	}

	fmt.Println("The following mac address has been found: ", addr.String())
	macAddr := [6]byte{}
	copy(macAddr[:], addr[:6])
	return macAddr, nil
}

// generate GUID Version 1
// 0001 time-based
func getGUIDV1() {
	// Basic alg
	/* 1. Obtain a system-wide global lock
	 	   2. From a system-wide shared stable store (e.g., a file), read the
	 			UUID generator state: the values of the timestamp, clock sequence,
	 			and node ID used to generate the last UUID.
		   3. Get the current time as a 60-bit count of 100-nanosecond intervals
	 			since 00:00:00.00, 15 October 1582.
	 	   4. Get the current node ID.
	       5. If the state was unavailable (e.g., non-existent or corrupted), or
	 			the saved node ID is different than the current node ID, generate
	 			a random clock sequence value.
	 	   6. If the state was available, but the saved timestamp is later than
	 			the current timestamp, increment the clock sequence value.
	 	   7. Save the state (current timestamp, clock sequence, and node ID)
	 			back to the stable store.
	 	   8. Release the global lock.
	 	   9. Format a UUID from the current timestamp, clock sequence, and node
				 ID values according to the steps in Section 4.2.2.*/

	guid := GUID{}
	// 1) call mutex
	mutex.Lock()
	// 4) get current node ID (In GUID/UUID v1, this is the current system's MAC Address)
	var err error
	guid.node, err = getMacAddress()
	if err != nil {
		fmt.Println("Error with getMacAddress")
	}
	// 8) unlock mutex
	mutex.Unlock()

}

// generate GUID Version 3
// 0011 name-based (MD5 hash)
func getGUIDV3() {

}

// generate GUID Version 4
// 0100 randomly or pseudo-randomly generated
func getGUIDV4() {

}

// generate GUID Version 5
// 0101 name based version
func getGUIDV5() {

}
