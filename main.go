package main

import (
	"fmt"
	"os"

	"Notes/cmd"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "App is a CLI for managing files",
	}

	rootCmd.AddCommand(cmd.CreateCMD)
	rootCmd.AddCommand(cmd.Read)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
