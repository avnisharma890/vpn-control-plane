package main

import (
	"fmt"

	"vpn-manager/internal/config"
	"vpn-manager/internal/ip"
	"vpn-manager/internal/wireguard"
)

func main() {

	privateKey, publicKey, err := wireguard.GenerateKeyPair()
	if err != nil {
		panic(err)
	}

	clientIP := ip.NextIP()

	serverPublicKey := "5ABgAyy7PLlR+dw971B2mwP4eiKIgdfKd+rfW7dmIlY="
	serverIP := "127.0.0.1"

	clientConfig := config.GenerateClientConfig(
		privateKey,
		serverPublicKey,
		serverIP,
		clientIP,
	)

	fmt.Println("New VPN Client Created")
	fmt.Println("-----------------------")

	fmt.Println("Client Public Key:", publicKey)
	fmt.Println("Assigned IP:", clientIP)

	fmt.Println("\nClient Config:\n")
	fmt.Println(clientConfig)
}