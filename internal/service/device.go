package service

import (
	"database/sql"
	"vpn-manager/internal/config"
	"vpn-manager/internal/db"
	"vpn-manager/internal/ip"
	"vpn-manager/internal/wireguard"
)

type DeviceResponse struct {
	PublicKey   string `json:"public_key"`
	VPNIP       string `json:"vpn_ip"`
	ClientConfig string `json:"client_config"`
}

func CreateDevice(database *sql.DB, serverPublicKey, serverIP string) (*DeviceResponse, error) {
	privateKey, publicKey, err := wireguard.GenerateKeyPair()
	if err != nil {
		return nil, err
	}

	clientIP, err := ip.NextIP()
	if err != nil {
		return nil, err
	}

	err = db.CreateDevice(database, publicKey, clientIP)
	if err != nil {
		return nil, err
	}

	err = wireguard.AddPeer(publicKey, clientIP)
	if err != nil {
		return nil, err
	}

	err = wireguard.Reload()
	if err != nil {
		return nil, err
	}

	clientConfig := config.GenerateClientConfig(
		privateKey,
		serverPublicKey,
		serverIP,
		clientIP,
	)

	return &DeviceResponse{
		PublicKey:   publicKey,
		VPNIP:       clientIP,
		ClientConfig: clientConfig,
	}, nil
}

func DeleteDevice(database *sql.DB, id int) error {

	publicKey, err := db.GetDeviceByID(database, id)
	if err != nil {
		return err
	}

	err = wireguard.RemovePeer(publicKey)
	if err != nil {
		return err
	}

	err = db.DeleteDevice(database, publicKey)
	if err != nil {
		return err
	}

	err = wireguard.Reload()
	if err != nil {
		return err
	}

	return nil
}