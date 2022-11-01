package cmd

import (
	"os/exec"
	"strconv"

	"github.com/crazyc4t/darkali/runner"
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
		autoUpdate := exec.Command("apt", "update")
		autoUpgrade := exec.Command("apt", "upgrade")
		aptDist := exec.Command("apt", "dist-upgrade")
		runner.Run(autoUpdate)
		runner.Run(autoUpgrade)
		runner.Run(aptDist)
	} else {
		autoUpdate := exec.Command("sudo", "apt", "update")
		autoUpgrade := exec.Command("sudo", "apt", "upgrade")
		aptDist := exec.Command("sudo", "apt", "dist-upgrade")
		runner.Run(autoUpdate)
		runner.Run(autoUpgrade)
		runner.Run(aptDist)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
