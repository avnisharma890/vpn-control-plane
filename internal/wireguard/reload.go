package wireguard

import "os/exec"

func Reload() error {

	cmd := exec.Command(
		"bash",
		"-c",
		"wg syncconf wg0 <(wg-quick strip wg0)",
	)

	err := cmd.Run()
	return err
}