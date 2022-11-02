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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "DarKali updates your system with apt",
	Long:  `DarKali will update your system and give you the latest and greatest packages.`,
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func update() {
	if privilages.Privilaged() {
		color.Blue(header.Header)
		cute.Println("Updating...")
		autoUpdate := exec.Command("apt", "update", "yy")
		autoUpgrade := exec.Command("apt", "upgrade", "-yy")
		aptDist := exec.Command("apt", "dist-upgrade", "-yy")
		runner.Run(autoUpdate)
		runner.Run(autoUpgrade)
		runner.Run(aptDist)
	} else {
		color.Blue(header.Header)
		cute.Println("Updating...")
		autoUpdate := exec.Command("sudo", "apt", "update", "-yy")
		autoUpgrade := exec.Command("sudo", "apt", "upgrade", "-yy")
		aptDist := exec.Command("sudo", "apt", "dist-upgrade", "-yy")
		runner.Run(autoUpdate)
		runner.Run(autoUpgrade)
		runner.Run(aptDist)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
