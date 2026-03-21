package wireguard

import (
	"bufio"
	"os"
	"strings"
)

func RemovePeer(publicKey string) error {

	input, err := os.Open("/etc/wireguard/wg0.conf")
	if err != nil {
		return err
	}
	defer input.Close()

	var outputLines []string
	scanner := bufio.NewScanner(input)

	var currentBlock []string
	inPeerBlock := false
	shouldSkip := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "[Peer]") {

			// flush previous block if needed
			if !shouldSkip {
				outputLines = append(outputLines, currentBlock...)
			}

			currentBlock = []string{line}
			inPeerBlock = true
			shouldSkip = false
			continue
		}

		if inPeerBlock {
			currentBlock = append(currentBlock, line)

			if strings.Contains(line, publicKey) {
				shouldSkip = true
			}

		} else {
			outputLines = append(outputLines, line)
		}
	}

	// flush last block
	if !shouldSkip {
		outputLines = append(outputLines, currentBlock...)
	}

	output := strings.Join(outputLines, "\n")

	err = os.WriteFile("/etc/wireguard/wg0.conf", []byte(output), 0644)
	return err
}