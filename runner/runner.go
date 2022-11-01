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
	err := cmd.Start()
	if err != nil {
		cute.Check("Error: ", err)
	}
	time.Sleep(8 * time.Second)
	cute.Println("Wating the script to complete...")
	err = cmd.Wait()
	if err != nil {
		cute.Check("Error: ", err)
	}
}
