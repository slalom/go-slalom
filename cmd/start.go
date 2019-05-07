package cmd

import (
	"github.com/spf13/cobra"
	"go-slalom/pkg/api"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the web service",
	Long: `Performs the following actions:

- Starts web service`,
	Run:    start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	// create HTTP server
	api.NewServer().ListenAndServe()
}