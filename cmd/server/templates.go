package main

import "snippetbox.austinsofaer.dev/internal/models"

// dynamic data type to pass into html templates
type templateData struct {
	Snippet models.Snippet
}
