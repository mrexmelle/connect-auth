package main

import (
	"github.com/mrexmelle/connect-authx/cmd/opts"
	_ "github.com/mrexmelle/connect-authx/docs"
)

// @title           Connect Authx API
// @version         0.1.1
// @description     Authx API for Connect.

// @contact.email  mrexmelle@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	opts.RootCmd.CompletionOptions.DisableDefaultCmd = true
	opts.RootCmd.AddCommand(opts.ServeCmd)
	opts.RootCmd.Execute()
}
