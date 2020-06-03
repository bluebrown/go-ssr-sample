package main

import (
	"io"
	"log"
	"text/template"
)

// View respresents the view
type View struct {
	pages map[string]*template.Template
}

// Parse parses the given file and register it with the name. Base templates are used  before.
func (v *View) Parse(name, glob string) {
	t, err := newT(glob)
	if err != nil {
		log.Fatalln(err)
	}
	v.pages[name] = t
}

// Exec executes template view data register under given name.
func (v *View) Exec(name string, w io.Writer, data interface{}) error {
	return v.pages[name].Execute(w, data)
}

// NewView returns a new view
func NewView() *View {
	var v View
	v.pages = make(map[string]*template.Template)
	return &v
}

func newT(glob string) (*template.Template, error) {
	t, err := template.ParseFiles("view/base.html")
	if err != nil {
		return nil, err
	}
	t, err = t.ParseGlob("view/partials/*.html")
	if err != nil {
		return nil, err
	}
	return t.ParseGlob(glob)
}
