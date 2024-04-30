package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var Read = &cobra.Command{
	Use:   "read",
	Short: "Read a note",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("Could not read file: %s", err)
		}
		fmt.Println(string(data))
	},
}
