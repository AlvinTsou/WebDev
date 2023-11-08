package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	UseCache       bool
	TemplatesCache map[string]*template.Template
	InfoLog        *log.Logger
}
