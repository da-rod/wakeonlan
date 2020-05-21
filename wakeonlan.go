package wakeonlan

import "net"

// MagicPacket is a broadcast frame containing 6 bytes of all 255 (FF FF FF FF FF FF in hexadecimal),
// followed by sixteen repetitions of the target computer's 48-bit MAC address, for a total of 102 bytes.
// https://wiki.wireshark.org/WakeOnLAN
type MagicPacket [102]byte

// NewMagicPacket returns the magic packet for the provided MAC address.
func NewMagicPacket(macAddress string) (MagicPacket, error) {
	var mp MagicPacket

	// Check MAC address
	hwAddr, err := net.ParseMAC(macAddress)
	if err != nil {
		return mp, err
	}

	// MAC address is valid, let's build this magic packet!
	syncStream := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	// Start with the 6 bytes Synchronization Stream:
	copy(mp[0:], syncStream)

	// After that, add the MAC address * 16:
	idx := len(hwAddr)
	for i := 0; i < 16; i++ {
		copy(mp[idx:], hwAddr)
		idx += len(hwAddr)
	}

	return mp, nil
}

// Send sends the magic packet via UDP to the broadcast address on port 9.
func (mp MagicPacket) Send() error {
	conn, err := net.Dial("udp", "255.255.255.255:9")
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(mp[:])
	return err
}
