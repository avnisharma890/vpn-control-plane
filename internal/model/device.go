package model

type Device struct {
	ID        int    `json:"id"`
	PublicKey string `json:"public_key"`
	VPNIP     string `json:"vpn_ip"`
	CreatedAt string `json:"created_at"`
}

