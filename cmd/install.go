package cmd

import (
	"os/exec"
	"strconv"

	"github.com/crazyc4t/darkali/runner"
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
	- Firefox-esr
	`,
	Run: func(cmd *cobra.Command, args []string) {
		install()
	},
}

func install() {
	perm := exec.Command("id", "-u")
	permOut, err := perm.Output()
	if err != nil {
		cute.Check("Error:", err)
	}
	i, err := strconv.Atoi(string(permOut[:len(permOut)-1]))
	if err != nil {
		cute.Check("Error:", err)
	}
	if i == 0 {
		cute.Println("Try again without sudo privilages, exiting...")
	} else {
		installSys := exec.Command("sudo", "./cmd/configs/system.sh")
		installNon := exec.Command("./cmd/configs/nonroot.sh")
		runner.Run(installSys)
		runner.Run(installNon)
		cute.Println("Rebooting is necessary for applying changes, DarKali installed.")
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
}
