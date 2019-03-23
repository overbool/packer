package main

import (
	"github.com/overbool/packer"
	"github.com/spf13/cobra"
)

var all bool

func init() {
	versionCMD.Flags().BoolVar(&all, "all", false, "show all version info")
}

var versionCMD = &cobra.Command{
	Use:   "version [flags]",
	Short: "Show version about app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("packer version: %s-%s\n", packer.CurrentVersion, packer.CurrentCommit)
		if all {
			cmd.Printf("App build date: %s\n", packer.BuildDate)
			cmd.Printf("System version: %s\n", packer.Platform)
			cmd.Printf("Golang version: %s\n", packer.GoVersion)
		}
	},
}
