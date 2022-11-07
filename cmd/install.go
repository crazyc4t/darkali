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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "DarKali install will set up your minimalist kali linux",
	Long: `DarKali will install:
	- Qtile window manager
	- Kitty terminal
	- Neovim config
	- Librewolf

	You need to run it inside the darkali directory.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}

func install() {
	if privilages.Privilaged() {
		cute.Println("Try again without sudo privilages, exiting...")
	} else {
		cute.SetTitleColor(cute.BrightBlue)
		cute.SetMessageColor(cute.BrightGreen)
		installSys := exec.Command("sudo", "./cmd/configs/system.sh")
		installNon := exec.Command("./cmd/configs/nonroot.sh")
		color.Blue(header.Header)
		cute.Println("Installing DarKali, Please don't kill the program...")
		runner.Run(installSys)
		cute.Println("Installing configs...")
		runner.Run(installNon)
		cute.Println("Rebooting is necessary for applying changes, DarKali installed.")
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
}
