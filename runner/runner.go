package runner

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/zakaria-chahboun/cute"
)

func Run(cmd *exec.Cmd) {
	cute.SetTitleColor(cute.BrightBlue)
	cute.SetMessageColor(cute.BrightGreen)
	cute.Println("Running script, please don't kill me...")
	fmt.Println()
	err := cmd.Start()
	if err != nil {
		cute.Check("Error: ", err)
	}
	time.Sleep(15 * time.Second)
	cute.Println("Finishing...")
	fmt.Println()
	err = cmd.Wait()
	if err != nil {
		cute.Check("Error: ", err)
	}
}
