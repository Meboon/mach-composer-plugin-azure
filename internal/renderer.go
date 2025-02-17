package internal

import (
	"embed"

	"github.com/flosch/pongo2/v5"
	"github.com/mach-composer/mach-composer-plugin-helpers/helpers"
)

//go:embed templates/*
var templates embed.FS

func renderResources(site, env string, cfg *SiteConfig, endpoints []EndpointConfig) (string, error) {
	templateSet := pongo2.NewSet("", &helpers.EmbedLoader{Content: templates})
	template := pongo2.Must(templateSet.FromFile("main.tf"))

	return template.Execute(pongo2.Context{
		"azure":     cfg,
		"siteName":  site,
		"envName":   env,
		"endpoints": endpoints,
	})
}
