package serve

import (
	"golayout/internal/app/server"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run a new server",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := server.NewServer()
		if err != nil {
			return err
		}

		return s.Run(cmd.Context())
	},
}

func New() *cobra.Command {
	return serveCmd
}
