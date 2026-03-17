package wireguard

import (
	"bytes"
	"os/exec"
	"strings"
)

func GenerateKeyPair() (string, string, error) {

	privateCmd := exec.Command("wg", "genkey")
	privateKeyBytes, err := privateCmd.Output()
	if err != nil {
		return "", "", err
	}

	privateKey := strings.TrimSpace(string(privateKeyBytes))

	publicCmd := exec.Command("wg", "pubkey")
	publicCmd.Stdin = bytes.NewBuffer(privateKeyBytes)

	publicKeyBytes, err := publicCmd.Output()
	if err != nil {
		return "", "", err
	}

	publicKey := strings.TrimSpace(string(publicKeyBytes))

	return privateKey, publicKey, nil
}