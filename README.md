python sample
```py
import socket

def send_wol(mac_address: str, target_ip: str, port=9):
    mac_bytes = bytes.fromhex(mac_address.replace(":", "").replace("-", ""))
    packet = b"\xFF" * 6 + mac_bytes * 16

    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.sendto(packet, (target_ip, port))
    sock.close()

# mac = "7C-10-C9-9D-14-01"
mac = "7C:10:C9:9D:14:01"
ip = "192.168.1.255"

send_wol(mac, ip)
```

go sample
```go
package sender

import (
	"encoding/hex"
	"net"
	"strings"
)

func main() {
	mac := "7C-10-C9-9D-14-01"
	macBytes, err := hex.DecodeString(strings.ReplaceAll(mac, "-", ""))
	if err != nil {
		panic(err)
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

	conn, err := net.Dial("udp", "192.168.1.255:9")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if _, err = conn.Write(magicPacket[:]); err != nil {
		panic(err)
	}
}
```
