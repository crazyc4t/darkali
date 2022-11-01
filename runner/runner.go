package runner

import (
	"os/exec"

	"github.com/zakaria-chahboun/cute"
)

func Run(cmd *exec.Cmd) {
	cute.SetTitleColor(cute.BrightRed)
	cute.SetMessageColor(cute.BrightBlue)
	cute.Println("Running script...")
	err := cmd.Run()
	if err != nil {
		cute.Check("err: %v", err)
	} else {
		cute.Println("Done!")
	}
}
