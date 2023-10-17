package main

import "github.com/TiagoSMarques/Snippetbox/internal/models"

// define a templateData type to act has the struct that holds all the dynamic data to pass to the html templates
type templateData struct {
	Snippet *models.Snippet
}
