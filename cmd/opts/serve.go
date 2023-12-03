package opts

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mrexmelle/connect-auth/internal/config"
	"github.com/mrexmelle/connect-auth/internal/credential"
	"github.com/mrexmelle/connect-auth/internal/profile"
	"github.com/mrexmelle/connect-auth/internal/session"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/dig"
)

func NewConfig() *config.Config {
	cfg, err := config.New(
		"application", "yaml",
		[]string{
			"/etc/conf",
			"./config",
		},
	)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Serve(cmd *cobra.Command, args []string) {
	container := dig.New()
	container.Provide(NewConfig)

	container.Provide(credential.NewRepository)
	container.Provide(profile.NewRepository)

	container.Provide(credential.NewService)
	container.Provide(profile.NewService)
	container.Provide(session.NewService)

	container.Provide(credential.NewController)
	container.Provide(profile.NewController)
	container.Provide(session.NewController)

	process := func(
		credentialController *credential.Controller,
		profileController *profile.Controller,
		sessionController *session.Controller,
		config *config.Config,
	) {
		r := chi.NewRouter()

		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://localhost:3000"},
			AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		if config.Profile == "local" {
			r.Mount("/swagger", httpSwagger.WrapHandler)
		}

		r.Route("/sessions", func(r chi.Router) {
			r.Post("/", sessionController.Post)
		})

		r.Route("/credentials", func(r chi.Router) {
			r.Post("/", credentialController.Post)
			r.Delete("/{eid}", credentialController.Delete)
			r.Patch("/{eid}/password", credentialController.PatchPassword)
			r.Delete("/{eid}/password", credentialController.DeletePassword)
		})

		r.Route("/profiles", func(r chi.Router) {
			r.Get("/{ehid}", profileController.Get)
			r.Patch("/{ehid}", profileController.Patch)
			r.Delete("/{ehid}", profileController.Delete)
		})

		err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)

		if err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(process); err != nil {
		panic(err)
	}
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Connect Auth server",
	Run:   Serve,
}
