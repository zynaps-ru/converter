package main

import (
	"html/template"
)

type Templates struct {
	index   *template.Template
	convert *template.Template
	enter   *template.Template
}

func NewTemplates() Templates {
	p := Templates{
		template.Must(template.ParseFiles("index.html")),
		template.Must(template.ParseFiles("convert.html")),
		template.Must(template.ParseFiles("enterData.html")),
	}

	return p
}
