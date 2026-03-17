package config

import "fmt"

func GenerateClientConfig(privateKey, serverPublicKey, serverIP, clientIP string) string {

	config := fmt.Sprintf(`[Interface]
PrivateKey = %s
Address = %s/24
DNS = 1.1.1.1

[Peer]
PublicKey = %s
Endpoint = %s:51820
AllowedIPs = 0.0.0.0/0
PersistentKeepalive = 25
`,
		privateKey,
		clientIP,
		serverPublicKey,
		serverIP,
	)

	return config
}