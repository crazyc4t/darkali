package runner

import (
	"os/exec"
	"time"

	"github.com/zakaria-chahboun/cute"
)

func Run(cmd *exec.Cmd) {
	cute.SetTitleColor(cute.BrightBlue)
	cute.SetMessageColor(cute.BrightGreen)
	cute.Println("Running script, please don't kill me...")
	time.Sleep(10 * time.Second)
	cute.Println("Almost finishing...")
	err := cmd.Run()
	if err != nil {
		cute.Check("Error: ", err)
	} else {
		cute.Println("Done!")
	}
}
