package cmd

import (
	"fmt"
	"os"

	"github.com/af-go/basic-app/cmd/agent"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "basic-app",
	Short: "",
}

func init() {
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(agent.AgentCmd)
}

// Exec excute root command
func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
