package main

import (
	"fmt"
	"vpn-manager/internal/config"
	"vpn-manager/internal/db"
	"vpn-manager/internal/ip"
	"vpn-manager/internal/wireguard"
)

func main() {	
	// Initialize database
	database, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	// defer database.Close()

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

	err = db.CreateDevice(database, publicKey, clientIP)
	if err != nil {
		panic(err)
	}

	// Register the client peer in the WireGuard server config
	err = wireguard.AddPeer(publicKey, clientIP)
	if err != nil {
		panic(err)
	}

	newPublicKey := "x4NegBYbp9yxWQlbYR2plL+1PlQIWa5OKB0Yi7Dz2Es="

	// Remove peer and delete device from database
	err = wireguard.RemovePeer(newPublicKey)
	if err != nil {
		panic(err)
	}

	err = db.DeleteDevice(database, publicKey)
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
