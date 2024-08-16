package application

import (
	"diz-nats/internal/docker"
	"diz-nats/internal/memory"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type DizNatsApp struct {
	memory memory.Organizer
	docker docker.Docker
}

func NewApp(memory memory.Organizer, docker docker.Docker) DizNatsApp {
	app := DizNatsApp{
		memory: memory,
		docker: docker,
	}

	// temp solution...
	var rootCmd = &cobra.Command{
		Use:   "diz-nats",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			app.NewServer("shani", "latest", "")
		},
	}

	// took that from init function for both root and test commands
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(testCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	// why?...
	return app
}
