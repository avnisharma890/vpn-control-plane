package wireguard

import (
	"fmt"
	"os"
)

func AddPeer(publicKey string, ip string) error {

	peerConfig := fmt.Sprintf(`
[Peer]
PublicKey = %s
AllowedIPs = %s/32
`, publicKey, ip)

	file, err := os.OpenFile("/etc/wireguard/wg0.conf",
		os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(peerConfig)

	return err
}