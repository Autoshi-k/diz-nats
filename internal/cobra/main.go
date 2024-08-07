package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func FirstLog() {
	var rootCmd = &cobra.Command{
		Use:   "mytool",
		Short: "MyTool is a CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, MyTool!")
		},
	}

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
