package stat

import (
	"bytes"
	"os/exec"
	"strconv"
)

func GetDiskUsage() (float64, error) {
	command := []string{"bash", "-c", `df -h | grep -e "/$" | awk '{print $5}'`}
	cmd := exec.Command(command[0], command[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	val := out.String()
	n, err := strconv.Atoi(val[:len(val)-2])
	return float64(n), nil
}
