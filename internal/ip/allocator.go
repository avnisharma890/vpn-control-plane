package ip

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func NextIP() (string, error) {

	file, err := os.Open("/etc/wireguard/wg0.conf")
	if err != nil {
		return "", err
	}

	defer file.Close()

	max := 1

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "AllowedIPs") {

			parts := strings.Split(line, "=")
			ipPart := strings.TrimSpace(parts[1])

			ip := strings.Split(ipPart, "/")[0]

			octets := strings.Split(ip, ".")
			lastOctet := octets[3]

			num, _ := strconv.Atoi(lastOctet)

			if num > max {
				max = num
			}
		}
	}

	next := max + 1

	return "10.0.0." + strconv.Itoa(next), nil
}