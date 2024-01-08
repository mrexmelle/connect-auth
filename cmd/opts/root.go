package opts

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "connect-authx",
	Short:   "Connect Authx",
	Long:    "Connect Authx - An authentication service for Connect",
	Version: "0.2.0",
}
