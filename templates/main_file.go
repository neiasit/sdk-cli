package templates

const MainFileTemplate = `package main

import (
	"github.com/go-playground/validator/v10"
	{{- range .Libraries }}
	"{{ ImportPath . }}"
	{{- end }}
	"go.uber.org/fx"
	{{- if HasLibrary "logging" }}
	"go.uber.org/fx/fxevent"
	"log/slog"
	{{- end }}
)

func main() {
	app := fx.New(

		// setting validator
		fx.Provide(func() *validator.Validate {
			return validator.New(
				validator.WithRequiredStructEnabled(),
			)
		}),
	
		{{- if HasLibrary "logging" }}
		// setting logger
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: logger,
			}
		}),
		{{- end }}
		
		// including platform libs here
		{{- if HasLibrary "logging" }}
		logging_library.Module,
		{{- end }}
		{{- if and (HasLibrary "grpc") (not (HasLibrary "auth")) }}
		grpc_library.Module,
		{{- else if and (HasLibrary "grpc") (HasLibrary "auth") }}
		grpc_library.ModuleWithAuth,
		{{- end }}
		{{- if HasLibrary "http-support" }}
		http_support_library.Module,
		{{- end }}
		{{- if HasLibrary "auth" }}
		auth_library.AuthKeycloakModule,
		{{- end }}

		// including app modules here
	
	)

	app.Run()
}

`
