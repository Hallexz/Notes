package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rewrite bool

var CreateCMD = &cobra.Command{
	Use:   "create [filename]",
	Short: "Create a new file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		if !rewrite {
			if _, err := os.Stat(filename); err == nil {
				fmt.Println("File already exists:", filename)
				os.Exit(1)
			} else if !os.IsNotExist(err) {
				fmt.Println("Error checking file:", err)
				os.Exit(1)
			}
		}

		dir := filepath.Dir(filename)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}

		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		file.Close()

		fmt.Println("File created successfully:", filename)
	},
}

func init() {
	CreateCMD.Flags().BoolVarP(&rewrite, "rewrite", "r", false, "Rewrite the file if it already exists")
}
