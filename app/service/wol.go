package service

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"

	"github.com/illinois1204/WOLWEB-switch/app/constants"
)

func WakeUp(mac string, port uint16) error {
	mac = strings.ReplaceAll(mac, "-", "")
	macBytes, err := hex.DecodeString(mac)
	if err != nil {
		return err
	}

	var magicPacket [102]byte
	magicPacket[0] = 0xFF
	magicPacket[1] = 0xFF
	magicPacket[2] = 0xFF
	magicPacket[3] = 0xFF
	magicPacket[4] = 0xFF
	magicPacket[5] = 0xFF

	offset := 6
	for i := range 16 {
		magicPacket[6+(offset*i)] = macBytes[0]
		magicPacket[7+(offset*i)] = macBytes[1]
		magicPacket[8+(offset*i)] = macBytes[2]
		magicPacket[9+(offset*i)] = macBytes[3]
		magicPacket[10+(offset*i)] = macBytes[4]
		magicPacket[11+(offset*i)] = macBytes[5]
	}

	conn, err := net.Dial("udp", net.JoinHostPort(constants.AppEnv.Network, fmt.Sprintf("%d", port)))
	if err != nil {
		return err
	}

	defer conn.Close()

	if _, err = conn.Write(magicPacket[:]); err != nil {
		return err
	}

	return nil
}
