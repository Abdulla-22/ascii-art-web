package Web_Art

import (
	web_art "Web_Art/Functions"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Struct to collect the data from the user
type PageData struct {
	UserInput     string
	Banner        string
	ModifiedInput string
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {

	// If it is a git method just show the main page
	if r.Method == "GET" {

		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "../templates/PageNotFound.html")
			fmt.Println("Not Found (404)")
			fmt.Println()

		} else if r.URL.Path == "/" {
			// To make all the filled look empty
			userInput := ""
			BannerName := ""
			modifiedInput := ""

			data := PageData{
				UserInput:     userInput,
				Banner:        BannerName,
				ModifiedInput: modifiedInput,
			}

			if data.UserInput == "" {
				tmpl := template.Must(template.ParseFiles("../templates/art.html"))
				tmpl.Execute(w, data)
			}

		}

		// In the post method
	} else if r.Method == "POST" {

		if r.URL.Path != "/submit" {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "../templates/PageNotFound.html")
			fmt.Println("Not Found (404)")
			fmt.Println()

		} else {

			// Take the input from the user
			userInput := r.FormValue("userInput")
			BannerName := r.FormValue("asciiOption")

			// TODO: move this part to the art functions

			// Check the input if they are acspted or not
			for _, ascii := range userInput {
				if ((ascii < 32) || (ascii > 126)) && (ascii != 10) && (ascii != 13) {

					http.ServeFile(w, r, "../templates/BadReq.html")
					fmt.Println("Bad Request (400)")
					fmt.Println("USE ONLY ENGLISH")
					fmt.Println()
					break

				} else {
					continue
				}
			}

			// Do the modification to the string
			ok := strings.Split(userInput, "\r")

			for i, l := range ok {
				if l == "\r" {
					ok[i] = "\n"
				}
			}
			f := strings.Join(ok, " ")
			modifiedInput, err1 := web_art.PrintArt(f, BannerName)

			// Make sure the file is there or not.
			if err1 != "" {
				w.WriteHeader(http.StatusInternalServerError)
				http.ServeFile(w, r, "../templates/InternalServerError.html")
				fmt.Println("Internal Server Error (500)")
				fmt.Println()
				return
			}

			data := PageData{
				UserInput:     userInput,
				Banner:        BannerName,
				ModifiedInput: modifiedInput,
			}

			tpl, err := template.ParseFiles("../templates/art.html")

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				http.ServeFile(w, r, "../templates/InternalServerError.html")
				fmt.Println("Internal Server Error (500)")
				fmt.Println()
				return
			}

			// Write the data to the website
			err = tpl.ExecuteTemplate(w, "art.html", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				http.ServeFile(w, r, "../templates/InternalServerError.html")
				fmt.Println("Internal Server Error (500)")
				fmt.Println()
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Println("OK (200)")
			fmt.Println()

		}

	} else {
		//Do the bad requst option here if any thing go wronge.
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "../templates/InternalServerError.html")
		fmt.Println("Internal Server Error (500)")
		fmt.Println()

	}
}
