package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	models "github.com/TiagoSMarques/Snippetbox"
)

// Define a home handler funtion which writes a byte slice containing
// the string as the response body
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it
	// doesn't, use the http.NotFound() function to send a 404 response to the clien
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// Use template.ParseFiles() to read the template file into a template set.
	// If there's an error, we log the detailed error message and use
	// http.Error() to send a generic 500 Internal Server Error response to the user.
	files := []string{
		"./ui/html/pages/home.html",
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll leave as nil

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function.
	//If it can't be converted to an integer, or the value is less than 1, we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	//If its not send a method not allowed
	if r.Method != http.MethodPost {
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	//pass the data to snippetModel.Insert() method, recieving the id of the record back
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//redirect the user to the relevant page for the snippetCreate
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
