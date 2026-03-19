package main

import (
	"fmt"

	"vpn-manager/internal/config"
	"vpn-manager/internal/ip"
	"vpn-manager/internal/wireguard"
)

func main() {

	// Generate a new WireGuard key pair for the client
	privateKey, publicKey, err := wireguard.GenerateKeyPair()
	if err != nil {
		panic(err)
	}

	// Allocate the next available VPN IP by scanning wg0.conf
	clientIP, err := ip.NextIP()
	if err != nil {
		panic(err)
	}

	// Register the client peer in the WireGuard server config
	err = wireguard.AddPeer(publicKey, clientIP)
	if err != nil {
		panic(err)
	}

	// Reload the WireGuard interface to apply changes
	err = wireguard.Reload()
	if err != nil {
		panic(err)
	}

	// Public key of the WireGuard server (needed by the client config)
	serverPublicKey := "5ABgAyy7PLlR+dw971B2mwP4eiKIgdfKd+rfW7dmIlY="
	
	// Endpoint clients will connect to (VirtualBox port forwarding → localhost)
	serverIP := "127.0.0.1"

	// Generate the WireGuard client configuration file
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
