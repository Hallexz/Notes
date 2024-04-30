package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var Edit = &cobra.Command{
	Use:   "edit",
	Short: "Edit a note",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		file.Close()

		file, err = os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			os.Exit(1)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		for _, line := range lines {
			newLine := strings.ToUpper(line)
			fmt.Fprintln(writer, newLine)
		}

		writer.Flush()
	},
}
