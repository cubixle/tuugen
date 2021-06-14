package tuugen

import "os/exec"

func runCommand(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	err := c.Start()
	if err != nil {
		return nil
	}
	err = c.Wait()
	if err != nil {
		return err
	}
	return nil
}
