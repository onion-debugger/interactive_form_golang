package main

import (
	"html/template"
	"fmt"
	"net/http"
	"log"
	"os"
)

// creating a struct to handle form validation
type ContactInfo struct {
	Name string
	Email string
	PhoneNumber string
	Message string
	Errors  map[string]string
}


// routing function for handling form submission
func confirmation(w http.ResponseWriter, r *http.Request) {
	routeSecurity(w, r)
	contactInfo := &ContactInfo{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		PhoneNumber: r.FormValue("phone"),
		Message: r.FormValue("message"),
	}

	// handles empty name
	if contactInfo.Name == "" {
		fmt.Fprintln(w, "<p>Please go and fill out the form</p>")
		return
	}
	render(w, "./static/confirmation.html", contactInfo)
}

// Handling route Security
func routeSecurity(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/confirmation" {
		http.Error(writer, "<p>This Page doesn't exist</p>", http.StatusNotFound)
		return
	}
}

// Listening to port from heroku
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
} 


func main() {
	// loading static file
	homePage := http.FileServer(http.Dir("./static"))
	http.Handle("/", homePage)

	http.HandleFunc("/confirmation", confirmation)

	fmt.Println("Server started at port 8080")

	// Handling error and starting server
	if err := http.ListenAndServe(getPort(), nil); err != nil {
		log.Fatal(err)
	}
}