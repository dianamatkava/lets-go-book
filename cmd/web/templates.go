package main

import "github.com/dianamatkava/snippetbox/cmd/internal/models"


type SnippetTemplate struct {
	Snippet models.Snippet
}

type SnippetsTemplate struct {
	Snippets []models.Snippet
}