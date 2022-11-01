package cmd

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/crazyc4t/darkali/runner"
	"github.com/spf13/cobra"
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
		log.Fatal(err)
	}
	i, err := strconv.Atoi(string(permOut[:len(permOut)-1]))
	if err != nil {
		log.Fatal(err)
	}
	if i == 0 {
		log.Fatal("Try again without sudo privilages, exiting...")
	} else {
		install := exec.Command("sudo", "./cmd/configs/system.sh")
		pip := exec.Command("pip", "install", "cairocffi xcffib psutil")
		runner.Run(install)
		runner.Run(pip)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
}
