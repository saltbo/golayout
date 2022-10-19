package serve

import (
	"fmt"

	"golayout/internal/app/server"
	"golayout/pkg/config"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run a new server",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.New()
		if err != nil {
			return err
		}

		fmt.Println(cfg)

		return server.NewServer(cfg).Run(cmd.Context())
	},
}

func New() *cobra.Command {
	return serveCmd
}
