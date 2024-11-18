package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.austinsofaer.dev/internal/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, s := range snippets {
		fmt.Fprintf(w, "%+v\n\n", s)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/components/nav.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// }
	//
	// // Parse html template into template set
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
	//
	// // write content of "base" template as response body
	// ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// initialise slice containing paths to view.tmpl.html file
	// and base layout + nav component
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/components/nav.tmpl.html",
		"./ui/html/pages/view.tmpl.html",
	}

	// parse template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// create templateData truct
	data := templateData{
		Snippet: s,
	}

	// execute template with snippet passed in
	if err := ts.ExecuteTemplate(w, "base", data); err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet... "))
}

func (app *Application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// dummy data
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// insert values as new snippet
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// redirect user to relevant page for snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
