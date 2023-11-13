package opts

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "connect-auth",
	Short:   "Connect Auth",
	Long:    "Connect Auth - An authentication and authorization service for Connect",
	Version: "0.1.0",
}
