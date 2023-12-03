package main

import (
	"github.com/mrexmelle/connect-auth/cmd/opts"
	_ "github.com/mrexmelle/connect-auth/docs"
)

// @title           Connect Auth API
// @version         1.0
// @description     Auth API for Connect.

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
