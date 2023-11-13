package main

import (
	"github.com/mrexmelle/connect-auth/cmd/opts"
)

func main() {
	opts.RootCmd.CompletionOptions.DisableDefaultCmd = true
	opts.RootCmd.AddCommand(opts.ServeCmd)
	opts.RootCmd.Execute()
}
