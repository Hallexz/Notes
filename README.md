# CLI Notes

CLI application for creating and editing notes
Created with Cobra


```
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
	rootCmd.AddCommand(cmd.Edit)
	rootCmd.AddCommand(cmd.Remove)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

### Strat off
- Download git rrepository
  
- Go to the installed directory

```
Cd /your/install/project/Notes
```

- Launching the application
  
```
go run main.go

```
