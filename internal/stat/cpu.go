package stat

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func GetCpuUsage() (float64,error) {
	command := []string{"grep", "cpu ", "/proc/stat"}
	cmd := exec.Command(command[0], command[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0,err
	}
	arr := strings.Split(out.String(), " ")
	n1, err := strconv.Atoi(arr[2])
	n2, err := strconv.Atoi(arr[4])
	n3, err := strconv.Atoi(arr[5])
	return float64(n1+n2) * 100 / float64(n1+n2+n3), nil
}
