package cmd

import (
	"github.com/slalom/go-slalom/pkg/api"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the web service",
	Long: `Performs the following actions:

- Starts web service`,
	Run: start,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	// create HTTP server
	api.NewServer().ListenAndServe()
}
