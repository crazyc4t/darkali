package cmd

import (
	"os/exec"
	"strconv"

	"github.com/crazyc4t/darkali/runner"
	"github.com/spf13/cobra"
	"github.com/zakaria-chahboun/cute"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "DarKali will clean your system",
	Long:  `DarKali will autoremove packages and clean your cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		clean()
	},
}

func clean() {
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
		autoClean := exec.Command("apt", "autoclean")
		autoRemove := exec.Command("apt", "autoremove")
		aptClean := exec.Command("apt", "clean")
		runner.Run(autoClean)
		runner.Run(autoRemove)
		runner.Run(aptClean)
	} else {
		autoClean := exec.Command("sudo", "apt", "autoclean")
		autoRemove := exec.Command("sudo", "apt", "autoremove")
		aptClean := exec.Command("sudo", "apt", "clean")
		runner.Run(autoClean)
		runner.Run(autoRemove)
		runner.Run(aptClean)
	}
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
