package privilages

import (
	"os/exec"
	"strconv"

	"github.com/zakaria-chahboun/cute"
)

func Privilaged() bool {
	perm := exec.Command("id", "-u")
	permOut, err := perm.Output()
	if err != nil {
		cute.Check("Error: ", err)
	}
	i, err := strconv.Atoi(string(permOut[:len(permOut)-1]))
	if err != nil {
		cute.Check("Error: ", err)
	}
	if i == 0 {
		return true
	} else {
		return false
	}
}
