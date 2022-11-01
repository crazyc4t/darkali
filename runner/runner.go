package runner

import (
	"os/exec"

	"github.com/zakaria-chahboun/cute"
)

func Run(cmd *exec.Cmd) {
	cute.SetTitleColor(cute.BrightBlue)
	cute.SetMessageColor(cute.BrightGreen)
	cute.Println("Running script, please don't kill me...")
	err := cmd.Run()
	if err != nil {
		cute.Check("Err: ", err)
	} else {
		cute.Println("Done!")
	}
}
