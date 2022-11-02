package runner

import (
	"os/exec"

	"github.com/zakaria-chahboun/cute"
)

func Run(cmd *exec.Cmd) {
	cute.SetTitleColor(cute.BrightBlue)
	cute.SetMessageColor(cute.BrightGreen)
	err := cmd.Start()
	if err != nil {
		cute.Check("Error: ", err)
	}
	err = cmd.Wait()
	if err != nil {
		cute.Check("Error: ", err)
	}
}
