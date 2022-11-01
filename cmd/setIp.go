/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/zakaria-chahboun/cute"
)

// setIpCmd represents the setIp command
var setIpCmd = &cobra.Command{
	Use:   "setIp",
	Short: "Sets your target IP in the qtile bar",
	Long:  `This is useful to remember your target IP when pentesting for example a HackTheBox machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip := `"` + args[0] + `"`
		setIp := exec.Command("qtile", "cmd-obj", "widget", "textbox", "-f", "update", "-a", ip)
		// qtile cmd-obj -o widget textbox -f update -a "New text"
		err := setIp.Run()
		if err != nil {
			cute.Check("Error happened:", err)
		}
		cute.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(setIpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setIpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setIpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
