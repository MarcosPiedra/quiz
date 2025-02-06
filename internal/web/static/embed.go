package static

import (
	"embed"
)

//go:embed index.html
var SwaggerIndex embed.FS

//go:embed swagger-ui/*
var SwaggerUi embed.FS
