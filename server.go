package main

import (
	"fmt"
	"net/http"
	"log"
)

const confirmationMessage = `<h1>Confirmation</h1>
<p>Your message has been sent!</p>`

// routing function contactHandler() for handling form submission
func contactHandler(writer http.ResponseWriter, request *http.Request) {
	// handling route security
	routeSecurity(writer, request)

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintln(writer, confirmationMessage)
	name := request.FormValue("name")
	email := request.FormValue("email")
	phoneNumber := request.FormValue("phone")
	message := request.FormValue("message")

	fmt.Fprintf(writer, "Name: %s\n", name)
	fmt.Fprintf(writer, "Email: %s\n", email)
	fmt.Fprintf(writer, "Phone Number: %s\n", phoneNumber)
	fmt.Fprintf(writer, "Message: %s\n", message)
	
}

// Handling route Security
func routeSecurity(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, "<p>This Page doesn't exist</p>", http.StatusNotFound)
		return
	}

	if request.Method != "POST" {
		http.Error(writer, "We cannot process this request", http.StatusNotFound)
		return
	}
}


func main() {
	// loading static file
	homePage := http.FileServer(http.Dir("./static"))
	http.Handle("/", homePage)

	http.HandleFunc("/submitContact", contactHandler)

	fmt.Println("Server started at port 3000")

	// Handling error and starting server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}