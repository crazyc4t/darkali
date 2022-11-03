package cmd

import (
	"os/exec"

	"github.com/crazyc4t/darkali/header"
	"github.com/crazyc4t/darkali/privilages"
	"github.com/crazyc4t/darkali/runner"
	"github.com/fatih/color"
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
	color.Blue(header.Header)
	if privilages.Privilaged() {
		autoClean := exec.Command("apt", "autoclean", "-yy")
		autoRemove := exec.Command("apt", "autoremove", "-yy")
		aptClean := exec.Command("apt", "clean", "-yy")
		cute.Println("Updating...")
		runner.Run(autoClean)
		runner.Run(autoRemove)
		runner.Run(aptClean)
		cute.Println("Done!")
	} else {
		autoClean := exec.Command("sudo", "apt", "autoclean", "-yy")
		autoRemove := exec.Command("sudo", "apt", "autoremove", "-yy")
		aptClean := exec.Command("sudo", "apt", "clean", "-yy")
		cute.Println("Updating...")
		runner.Run(autoClean)
		runner.Run(autoRemove)
		runner.Run(aptClean)
		cute.Println("Done!")
	}
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
