package stat

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func GetRAMUsage()(float64,error){
	command := []string{"bash","-c","free | grep Mem"}
	cmd := exec.Command(command[0], command[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0,err
	}
	arr := strings.Split(out.String(), " ")
    var x,y int = -1, -1
    for i,s := range(arr){
        if s!="" && i>0{
            if x==-1{
                x = i
            } else if y==-1{
                y = i
            } else {
                break
            }
        }
    }
	n1, err := strconv.Atoi(arr[x])
	n2, err := strconv.Atoi(arr[y])
	return float64(n2) * 100 / float64(n1),nil
}
